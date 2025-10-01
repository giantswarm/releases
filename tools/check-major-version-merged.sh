#!/bin/bash

set -euo pipefail

if [ "$#" -eq 0 ]; then
    echo "Usage: $0 <provider1> <provider2> ..."
    echo "Checks if the latest major version for the given providers is fully merged."
    exit 1
fi

PROVIDERS=("$@")
MAJOR_VERSIONS=()
LATEST_MAJOR=""

echo "Checking if latest major version is fully merged for providers: ${PROVIDERS[*]}"

for provider in "${PROVIDERS[@]}"; do
    # Handle aws -> capa directory mapping for modern releases
    search_dir=$provider
    if [ "$provider" == "aws" ]; then
        search_dir="capa"
    fi

    if [ ! -d "$search_dir" ]; then
        echo "Error: Directory for provider '$provider' ('$search_dir') does not exist. Skipping."
        continue
    fi

    # Find the latest release version in the provider's directory, excluding archived releases.
    LATEST_RELEASE=$(find "$search_dir" -name "v[0-9]*" -type d -not -path "*/archived/*" -exec basename {} \; | sort -V | tail -n 1)

    if [ -z "$LATEST_RELEASE" ]; then
        echo "Warning: No active releases found for provider '$provider'. Skipping."
        continue
    fi

    # Extract major version (e.g., v31.2.0 -> 31)
    MAJOR_VERSION=$(echo "$LATEST_RELEASE" | sed -E 's/^v([0-9]+)\..*/\1/')
    echo "Latest release for $provider is $LATEST_RELEASE (Major: $MAJOR_VERSION)"
    MAJOR_VERSIONS+=("$MAJOR_VERSION")
done

if [ ${#MAJOR_VERSIONS[@]} -eq 0 ]; then
    echo "Error: No valid providers found."
    exit 1
fi

# Find the highest major version
LATEST_MAJOR=$(printf '%s\n' "${MAJOR_VERSIONS[@]}" | sort -n | tail -n 1)
echo "Latest major version across all providers: $LATEST_MAJOR"

# Check if there are any open PRs for the next major version
NEXT_MAJOR=$((LATEST_MAJOR + 1))
echo "Checking for open PRs for next major version: v$NEXT_MAJOR.0.0"

# Check for consolidated PR (open or closed)
CONSOLIDATED_PR=$(gh pr list --repo "$GITHUB_REPOSITORY" --head "release-v$NEXT_MAJOR.0.0" --state all --json number,state --jq '.[0] | if . then (.number | tostring) + ":" + .state else "" end' 2>/dev/null || echo "")

# Check for individual PRs (open or closed)
INDIVIDUAL_PRS=()
for provider in "${PROVIDERS[@]}"; do
    PR=$(gh pr list --repo "$GITHUB_REPOSITORY" --head "release-v$NEXT_MAJOR.0.0-$provider" --state all --json number,state --jq '.[0] | if . then (.number | tostring) + ":" + .state else "" end' 2>/dev/null || echo "")
    if [ -n "$PR" ]; then
        INDIVIDUAL_PRS+=("$provider:$PR")
    fi
done

# Check if consolidated PR exists and is merged
if [ -n "$CONSOLIDATED_PR" ]; then
    PR_STATE=$(echo "$CONSOLIDATED_PR" | cut -d':' -f2)
    PR_NUMBER=$(echo "$CONSOLIDATED_PR" | cut -d':' -f1)
    if [[ "$PR_STATE" == "MERGED" ]]; then
        echo "Found merged consolidated PR for v$NEXT_MAJOR.0.0: #$PR_NUMBER"
        echo "Major version v$NEXT_MAJOR is now available."
        echo "merged=true"
        if [ -n "${GITHUB_OUTPUT-}" ]; then
            echo "merged=true" >> "$GITHUB_OUTPUT"
            echo "latest_major=$NEXT_MAJOR" >> "$GITHUB_OUTPUT"
            echo "next_major=$((NEXT_MAJOR + 1))" >> "$GITHUB_OUTPUT"
        fi
    else
        echo "Found $PR_STATE consolidated PR for v$NEXT_MAJOR.0.0: #$PR_NUMBER"
        echo "Major version v$LATEST_MAJOR is not fully merged yet."
        echo "merged=false"
        if [ -n "${GITHUB_OUTPUT-}" ]; then
            echo "merged=false" >> "$GITHUB_OUTPUT"
            echo "latest_major=$LATEST_MAJOR" >> "$GITHUB_OUTPUT"
            echo "next_major=$NEXT_MAJOR" >> "$GITHUB_OUTPUT"
        fi
    fi
elif [ ${#INDIVIDUAL_PRS[@]} -gt 0 ]; then
    # Check if all individual PRs are merged
    ALL_MERGED=true
    for pr_info in "${INDIVIDUAL_PRS[@]}"; do
        provider=$(echo "$pr_info" | cut -d':' -f1)
        pr_state=$(echo "$pr_info" | cut -d':' -f3)
        pr_number=$(echo "$pr_info" | cut -d':' -f2)
        echo "Found $pr_state individual PR for v$NEXT_MAJOR.0.0-$provider: #$pr_number"
        if [[ "$pr_state" != "MERGED" ]]; then
            ALL_MERGED=false
        fi
    done
    
    if [[ "$ALL_MERGED" == "true" ]]; then
        echo "All individual PRs for v$NEXT_MAJOR.0.0 are merged."
        echo "Major version v$NEXT_MAJOR is now available."
        echo "merged=true"
        if [ -n "${GITHUB_OUTPUT-}" ]; then
            echo "merged=true" >> "$GITHUB_OUTPUT"
            echo "latest_major=$NEXT_MAJOR" >> "$GITHUB_OUTPUT"
            echo "next_major=$((NEXT_MAJOR + 1))" >> "$GITHUB_OUTPUT"
        fi
    else
        echo "Not all individual PRs for v$NEXT_MAJOR.0.0 are merged yet."
        echo "Major version v$LATEST_MAJOR is not fully merged yet."
        echo "merged=false"
        if [ -n "${GITHUB_OUTPUT-}" ]; then
            echo "merged=false" >> "$GITHUB_OUTPUT"
            echo "latest_major=$LATEST_MAJOR" >> "$GITHUB_OUTPUT"
            echo "next_major=$NEXT_MAJOR" >> "$GITHUB_OUTPUT"
        fi
    fi
else
    echo "No PRs found for v$NEXT_MAJOR.0.0"
    echo "Major version v$LATEST_MAJOR is fully merged."
    echo "merged=true"
    if [ -n "${GITHUB_OUTPUT-}" ]; then
        echo "merged=true" >> "$GITHUB_OUTPUT"
        echo "latest_major=$LATEST_MAJOR" >> "$GITHUB_OUTPUT"
        echo "next_major=$NEXT_MAJOR" >> "$GITHUB_OUTPUT"
    fi
fi
