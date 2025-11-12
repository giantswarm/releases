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
    # 1. Same major version:
    #    a. Same minor or 1 minor behind -> INCLUDE
    #    b. 2+ minors behind -> EXCLUDE (can't skip minor versions)
    # 2. 1 major behind AND target is x.0.0 -> INCLUDE (major version bump)
    # 3. 1 major behind AND target is x.y.z (y>0) -> EXCLUDE (must catch up first)
    # 4. 2+ major versions behind -> EXCLUDE
    
    if [ "$CURRENT_MAJOR" -eq "$TARGET_MAJOR" ]; then
        # Same major version - check minor version difference
        MINOR_DIFF=$((TARGET_MINOR - CURRENT_MINOR))
        
        if [ "$MINOR_DIFF" -le 1 ] && [ "$MINOR_DIFF" -ge 0 ]; then
            # Same minor or 1 minor behind - safe to include
            if [ "$MINOR_DIFF" -eq 0 ]; then
                echo "  ✓ INCLUDE (same major and minor version)"
            else
                echo "  ✓ INCLUDE (minor version bump from v$CURRENT_MAJOR.$CURRENT_MINOR to v$TARGET_MAJOR.$TARGET_MINOR)"
            fi
            INCLUDED_PROVIDERS+=("$provider")
        elif [ "$MINOR_DIFF" -lt 0 ]; then
            # Provider is ahead - this shouldn't happen but include anyway
            echo "  ✓ INCLUDE (provider ahead of target)"
            INCLUDED_PROVIDERS+=("$provider")
        else
            # 2+ minors behind - exclude
            echo "  ✗ EXCLUDE (too far behind: v$CURRENT_MAJOR.$CURRENT_MINOR vs target v$TARGET_MAJOR.$TARGET_MINOR, must catch up to v$TARGET_MAJOR.$((TARGET_MINOR - 1)).x first)"
            EXCLUDED_PROVIDERS+=("$provider")
        fi
    elif [ "$CURRENT_MAJOR" -gt "$TARGET_MAJOR" ]; then
        # Provider is ahead - this is valid for creating minor releases of older majors
        echo "  ✓ INCLUDE (creating minor release for older major v$TARGET_MAJOR while v$CURRENT_MAJOR exists)"
        INCLUDED_PROVIDERS+=("$provider")
    elif [ "$CURRENT_MAJOR" -eq $((TARGET_MAJOR - 1)) ]; then
        # One major version behind - only include if target is a major release (x.0.0)
        if [ "$TARGET_MINOR" -eq 0 ]; then
            echo "  ✓ INCLUDE (major version bump from v$CURRENT_MAJOR to v$TARGET_MAJOR)"
            INCLUDED_PROVIDERS+=("$provider")
        else
            echo "  ✗ EXCLUDE (must catch up to v$TARGET_MAJOR.0.0 before joining v$TARGET_MAJOR.$TARGET_MINOR releases)"
            EXCLUDED_PROVIDERS+=("$provider")
        fi
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

