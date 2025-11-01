#!/bin/bash

set -euo pipefail

if [ "$#" -eq 0 ]; then
    echo "Usage: $0 <provider1> <provider2> ..."
    echo "Checks if the latest releases for the given providers share the same minor version."
    exit 1
fi

PROVIDERS=("$@")
MINOR_VERSIONS=()
# Providers to exclude from alignment check (but still process for releases)
EXCLUDED_PROVIDERS=("eks")

echo "Checking minor version alignment for providers: ${PROVIDERS[*]}"
echo "Providers excluded from alignment check: ${EXCLUDED_PROVIDERS[*]}"

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

    # Extract minor version (e.g., v31.2.0 -> 31.2)
    MINOR_VERSION=$(echo "$LATEST_RELEASE" | sed -E 's/^v([0-9]+\.[0-9]+)\..*/\1/')
    echo "Latest release for $provider is $LATEST_RELEASE (Minor: $MINOR_VERSION)"

    # Check if this provider should be excluded from alignment check
    IS_EXCLUDED=false
    for excluded in "${EXCLUDED_PROVIDERS[@]}"; do
        if [ "$provider" == "$excluded" ]; then
            IS_EXCLUDED=true
            echo "  -> Excluding $provider from alignment check"
            break
        fi
    done

    # Only include in alignment check if not excluded
    if [ "$IS_EXCLUDED" == "false" ]; then
        MINOR_VERSIONS+=("$MINOR_VERSION")
    fi
done

ALIGNMENT="aligned"
if [ ${#MINOR_VERSIONS[@]} -ge 2 ]; then
    FIRST_MINOR=${MINOR_VERSIONS[0]}
    for minor in "${MINOR_VERSIONS[@]}"; do
        if [ "$minor" != "$FIRST_MINOR" ]; then
            ALIGNMENT="misaligned"
            break
        fi
    done
fi

echo "Alignment status: $ALIGNMENT"

# Set output for GitHub Actions
if [ -n "${GITHUB_OUTPUT-}" ]; then
    echo "alignment=$ALIGNMENT" >> "$GITHUB_OUTPUT"
fi
