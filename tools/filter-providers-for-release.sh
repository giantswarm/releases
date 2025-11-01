#!/bin/bash

set -euo pipefail

# This script filters providers based on whether they can be included in a consolidated release
# for a given target version. It excludes providers that are more than 1 minor version behind
# to prevent accidental major version jumps.

if [ "$#" -lt 2 ]; then
    echo "Usage: $0 <target_version> <provider1> <provider2> ..."
    echo "Example: $0 v33.1.0 aws azure cloud-director eks vsphere"
    echo ""
    echo "Filters providers that can be safely included in a release for the target version."
    echo "Providers more than 1 minor version behind are excluded to prevent major version jumps."
    exit 1
fi

TARGET_VERSION="$1"
shift
PROVIDERS=("$@")

# Extract target major and minor from target version (e.g., v33.1.0 -> major=33, minor=1)
TARGET_MAJOR=$(echo "$TARGET_VERSION" | sed -E 's/^v([0-9]+)\..*/\1/')
TARGET_MINOR=$(echo "$TARGET_VERSION" | sed -E 's/^v[0-9]+\.([0-9]+)\..*/\1/')

echo "Target version: $TARGET_VERSION (Major: $TARGET_MAJOR, Minor: $TARGET_MINOR)"
echo "Filtering providers: ${PROVIDERS[*]}"
echo ""

INCLUDED_PROVIDERS=()
EXCLUDED_PROVIDERS=()

for provider in "${PROVIDERS[@]}"; do
    # Handle aws -> capa directory mapping for modern releases
    search_dir=$provider
    if [ "$provider" == "aws" ]; then
        search_dir="capa"
    fi

    if [ ! -d "$search_dir" ]; then
        echo "WARNING: Directory for provider '$provider' ('$search_dir') does not exist. Excluding."
        EXCLUDED_PROVIDERS+=("$provider")
        continue
    fi

    # Find the latest release version in the provider's directory, excluding archived releases.
    LATEST_RELEASE=$(find "$search_dir" -name "v[0-9]*" -type d -not -path "*/archived/*" -exec basename {} \; | sort -V | tail -n 1)

    if [ -z "$LATEST_RELEASE" ]; then
        echo "WARNING: No active releases found for provider '$provider'. Excluding."
        EXCLUDED_PROVIDERS+=("$provider")
        continue
    fi

    # Extract current major and minor version
    CURRENT_MAJOR=$(echo "$LATEST_RELEASE" | sed -E 's/^v([0-9]+)\..*/\1/')
    CURRENT_MINOR=$(echo "$LATEST_RELEASE" | sed -E 's/^v[0-9]+\.([0-9]+)\..*/\1/')

    echo "Provider: $provider"
    echo "  Current: $LATEST_RELEASE (Major: $CURRENT_MAJOR, Minor: $CURRENT_MINOR)"
    
    # Decision logic:
    # 1. If provider is at the same major version as target -> INCLUDE
    # 2. If provider is 1 major version behind -> EXCLUDE (would be a major jump)
    # 3. If provider is 2+ major versions behind -> EXCLUDE
    
    if [ "$CURRENT_MAJOR" -eq "$TARGET_MAJOR" ]; then
        # Same major version - safe to include
        echo "  ✓ INCLUDE (same major version)"
        INCLUDED_PROVIDERS+=("$provider")
    elif [ "$CURRENT_MAJOR" -eq $((TARGET_MAJOR - 1)) ]; then
        # One major version behind - exclude to prevent major jump
        echo "  ✗ EXCLUDE (would jump from major $CURRENT_MAJOR to $TARGET_MAJOR)"
        EXCLUDED_PROVIDERS+=("$provider")
    else
        # More than one major version behind - exclude
        echo "  ✗ EXCLUDE (too far behind: major $CURRENT_MAJOR vs target $TARGET_MAJOR)"
        EXCLUDED_PROVIDERS+=("$provider")
    fi
    echo ""
done

# Output results
echo "========================================"
echo "INCLUDED PROVIDERS: ${INCLUDED_PROVIDERS[*]:-none}"
echo "EXCLUDED PROVIDERS: ${EXCLUDED_PROVIDERS[*]:-none}"
echo "========================================"

# Format for GitHub Actions (space-separated list)
if [ ${#INCLUDED_PROVIDERS[@]} -gt 0 ]; then
    INCLUDED_LIST=$(printf "%s " "${INCLUDED_PROVIDERS[@]}")
    INCLUDED_LIST=${INCLUDED_LIST% } # Remove trailing space
else
    INCLUDED_LIST=""
fi

# Set output for GitHub Actions
if [ -n "${GITHUB_OUTPUT-}" ]; then
    echo "included_providers=$INCLUDED_LIST" >> "$GITHUB_OUTPUT"
    echo "excluded_providers=${EXCLUDED_PROVIDERS[*]:-}" >> "$GITHUB_OUTPUT"
    echo "has_included_providers=$([ ${#INCLUDED_PROVIDERS[@]} -gt 0 ] && echo 'true' || echo 'false')" >> "$GITHUB_OUTPUT"
fi

# Also output to stdout for easy parsing
echo "included_providers=$INCLUDED_LIST"
echo "excluded_providers=${EXCLUDED_PROVIDERS[*]:-}"

# Exit with error if no providers are included
if [ ${#INCLUDED_PROVIDERS[@]} -eq 0 ]; then
    echo ""
    echo "ERROR: No providers can be included in this release!"
    exit 1
fi

