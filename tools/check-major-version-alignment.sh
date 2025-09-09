#!/bin/bash

set -euo pipefail

if [ "$#" -eq 0 ]; then
    echo "Usage: $0 <provider1> <provider2> ..."
    echo "Checks if the latest releases for the given providers share the same major version."
    exit 1
fi

PROVIDERS=("$@")
MAJOR_VERSIONS=()

echo "Checking major version alignment for providers: ${PROVIDERS[*]}"

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
    LATEST_RELEASE=$(find "$search_dir" -name "v*" -type d -not -path "*/archived/*" -print0 | xargs -0 basename | sort -V | tail -n 1)

    if [ -z "$LATEST_RELEASE" ]; then
        echo "Warning: No active releases found for provider '$provider'. Skipping."
        continue
    fi

    # Extract major version (e.g., v31.2.0 -> 31)
    MAJOR_VERSION=$(echo "$LATEST_RELEASE" | sed -E 's/^v([0-9]+)\..*/\1/')
    echo "Latest release for $provider is $LATEST_RELEASE (Major: $MAJOR_VERSION)"
    MAJOR_VERSIONS+=("$MAJOR_VERSION")
done

ALIGNMENT="aligned"
if [ ${#MAJOR_VERSIONS[@]} -ge 2 ]; then
    FIRST_MAJOR=${MAJOR_VERSIONS[0]}
    for major in "${MAJOR_VERSIONS[@]}"; do
        if [ "$major" != "$FIRST_MAJOR" ]; then
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
