#!/bin/bash

set -euo pipefail

if [ "$#" -eq 0 ]; then
    echo "Usage: $0 <provider1> <provider2> ..."
    echo "Checks if there are existing open PRs for minor releases."
    exit 1
fi

PROVIDERS=("$@")
EXISTING_PRS=()

echo "Checking for existing minor release PRs for providers: ${PROVIDERS[*]}"

# Get the latest MERGED major version to determine what minor version we should check for
# We need to check if there are any open PRs for the next major version to determine
# if the current major is still the latest merged one
LATEST_MAJOR=""
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
    if [ -z "$LATEST_MAJOR" ] || [ "$MAJOR_VERSION" -gt "$LATEST_MAJOR" ]; then
        LATEST_MAJOR=$MAJOR_VERSION
    fi
done

# Now check if there are any open PRs for the next major version
# If there are, it means the current major is still the latest merged one
# If there aren't, it means a new major has been merged and we should use that
NEXT_MAJOR=$((LATEST_MAJOR + 1))
echo "Checking if v$NEXT_MAJOR.0.0 has been merged..."

# Check for any open PRs for the next major version
HAS_OPEN_NEXT_MAJOR=false
for provider in "${PROVIDERS[@]}"; do
    PR=$(gh pr list --repo "$GITHUB_REPOSITORY" --head "release-v$NEXT_MAJOR.0.0-$provider" --state open --json number --jq '.[0].number' 2>/dev/null || echo "")
    if [ -n "$PR" ]; then
        HAS_OPEN_NEXT_MAJOR=true
        break
    fi
done

# Also check for consolidated PR
if [ "$HAS_OPEN_NEXT_MAJOR" == "false" ]; then
    PR=$(gh pr list --repo "$GITHUB_REPOSITORY" --head "release-v$NEXT_MAJOR.0.0" --state open --json number --jq '.[0].number' 2>/dev/null || echo "")
    if [ -n "$PR" ]; then
        HAS_OPEN_NEXT_MAJOR=true
    fi
fi

# If there are open PRs for the next major, use current major for minor releases
# If there are no open PRs for the next major, it means the next major has been merged
if [ "$HAS_OPEN_NEXT_MAJOR" == "true" ]; then
    echo "Found open PRs for v$NEXT_MAJOR.0.0. Using v$LATEST_MAJOR for minor releases."
    TARGET_MAJOR=$LATEST_MAJOR
else
    echo "No open PRs for v$NEXT_MAJOR.0.0 found. Next major has been merged. Using v$NEXT_MAJOR for minor releases."
    TARGET_MAJOR=$NEXT_MAJOR
fi

if [ -z "$LATEST_MAJOR" ]; then
    echo "Error: No valid providers found."
    exit 1
fi

echo "Latest major version across all providers: $LATEST_MAJOR"
echo "Target major version for minor releases: $TARGET_MAJOR"

# Check for existing minor release PRs for the target major version
echo "Checking for existing minor release PRs for v$TARGET_MAJOR.x.x"

for provider in "${PROVIDERS[@]}"; do
    # Check for any OPEN minor release PRs (v32.1.0, v32.2.0, etc.)
    # Only check for open PRs, not closed/merged ones
    PRS=$(gh pr list --repo "$GITHUB_REPOSITORY" --head "release-v$TARGET_MAJOR.*-$provider" --state open --json number,title,headRefName --jq '.[] | .number + ":" + .headRefName' 2>/dev/null || echo "")
    if [ -n "$PRS" ]; then
        while IFS= read -r pr_info; do
            if [ -n "$pr_info" ]; then
                EXISTING_PRS+=("$provider:$pr_info")
            fi
        done <<< "$PRS"
    fi
done

if [ ${#EXISTING_PRS[@]} -gt 0 ]; then
    echo "Found existing minor release PRs:"
    for pr_info in "${EXISTING_PRS[@]}"; do
        provider=$(echo "$pr_info" | cut -d':' -f1)
        pr_number=$(echo "$pr_info" | cut -d':' -f2 | cut -d':' -f1)
        branch=$(echo "$pr_info" | cut -d':' -f2 | cut -d':' -f2)
        echo "  - $provider: #$pr_number ($branch)"
    done
    echo "prs_exist=true"
    if [ -n "${GITHUB_OUTPUT-}" ]; then
        echo "prs_exist=true" >> "$GITHUB_OUTPUT"
    fi
else
    echo "No existing minor release PRs found for v$TARGET_MAJOR.x.x"
    echo "prs_exist=false"
    if [ -n "${GITHUB_OUTPUT-}" ]; then
        echo "prs_exist=false" >> "$GITHUB_OUTPUT"
    fi
fi
