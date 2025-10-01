#!/bin/bash
set -e

PROVIDER=$1
RELEASE_TYPE=$2
VERSION_OVERRIDE=$3

if [ -z "$RELEASE_TYPE" ]; then
  echo "Usage: $0 <provider> <release_type> [version_override]"
  echo "Note: <provider> can be an empty string \"\" for major releases to scan all providers."
  exit 1
fi

ALL_RELEASES=""

if [ -z "$PROVIDER" ]; then
  # Scan all provider directories to find the absolute latest release
  if [ "$RELEASE_TYPE" != "major" ] && [ "$RELEASE_TYPE" != "minor" ]; then
    echo "Error: An empty provider is only allowed for 'major' or 'minor' release types."
    exit 1
  fi
  PROVIDER_DIRS=$(ls -d */ | grep -v -E '^(announcements|app|helm|sdk|tools|archived|kvm)$' | tr '\n' ' ')
  ALL_RELEASES=$(find $PROVIDER_DIRS -maxdepth 1 -type d -name 'v[0-9]*' -not -path '*/archived/*' -exec basename {} \; | sort -V)
else
  PROVIDER_DIR="./${PROVIDER}"
  if [ "$PROVIDER" == "aws" ]; then
    PROVIDER_DIR="./capa"
  fi

  if [ ! -d "$PROVIDER_DIR" ]; then
    echo "Provider directory not found: $PROVIDER_DIR"
    exit 1
  fi
  # Get a sorted list of all releases for the provider, excluding archived releases.
  ALL_RELEASES=$(find "$PROVIDER_DIR" -maxdepth 1 -type d -name "v[0-9]*" -not -path "*/archived/*" -exec basename {} \; | sort -V)
fi


if [ -z "$ALL_RELEASES" ]; then
  echo "No releases found."
  exit 1
fi

if [ -n "$VERSION_OVERRIDE" ]; then
  # Version is specified, use it and determine the base release.
  NEXT_RELEASE=$VERSION_OVERRIDE
  # Remove 'v' prefix for version manipulation
  NEXT_VERSION_NO_V=${NEXT_RELEASE#v}
  IFS='.' read -r -a VERSION_PARTS <<< "$NEXT_VERSION_NO_V"
  MAJOR=${VERSION_PARTS[0]}
  MINOR=${VERSION_PARTS[1]}
  PATCH=${VERSION_PARTS[2]}

  if [ "$PATCH" -gt 0 ]; then
    # This is a patch release. The base is the previous patch.
    PREVIOUS_PATCH=$((PATCH - 1))
    LATEST_RELEASE="v${MAJOR}.${MINOR}.${PREVIOUS_PATCH}"
    # Verify that the base release exists
    if [ ! -d "${PROVIDER_DIR}/${LATEST_RELEASE}" ]; then
      echo "Error: Base release ${LATEST_RELEASE} for backport not found."
      exit 1
    fi
  else
    # This is a new minor or major release. Find the latest release from the previous minor/major series.
    PREVIOUS_MINOR=$((MINOR - 1))
    if [ $PREVIOUS_MINOR -lt 0 ]; then
        # This is a major version, find latest release from previous major
        PREVIOUS_MAJOR=$((MAJOR - 1))
        LATEST_RELEASE=$(echo "$ALL_RELEASES" | grep "^v${PREVIOUS_MAJOR}\." | tail -n 1)
    else
        # This is a minor version, find latest release from previous minor
        LATEST_RELEASE=$(echo "$ALL_RELEASES" | grep "^v${MAJOR}\.${PREVIOUS_MINOR}\." | tail -n 1)
    fi

    if [ -z "$LATEST_RELEASE" ]; then
        echo "Error: Could not determine a base release for ${NEXT_RELEASE}."
        exit 1
    fi
  fi
else
  # No version override, calculate the next version from the latest release.
  LATEST_RELEASE=$(echo "$ALL_RELEASES" | tail -n 1)
  LATEST_VERSION=${LATEST_RELEASE#v}
  IFS='.' read -r -a VERSION_PARTS <<< "$LATEST_VERSION"
  MAJOR=${VERSION_PARTS[0]}
  MINOR=${VERSION_PARTS[1]}
  PATCH=${VERSION_PARTS[2]}

  case $RELEASE_TYPE in
    major)
      NEXT_MAJOR=$((MAJOR + 1))
      NEXT_MINOR=0
      NEXT_PATCH=0
      ;;
    minor)
      NEXT_MAJOR=$MAJOR
      NEXT_MINOR=$((MINOR + 1))
      NEXT_PATCH=0
      ;;
    patch)
      NEXT_MAJOR=$MAJOR
      NEXT_MINOR=$MINOR
      NEXT_PATCH=$((PATCH + 1))
      ;;
    *)
      echo "Invalid release type: $RELEASE_TYPE"
      exit 1
      ;;
  esac
  NEXT_RELEASE="v${NEXT_MAJOR}.${NEXT_MINOR}.${NEXT_PATCH}"
fi

echo "Base release: $LATEST_RELEASE"
echo "New release: $NEXT_RELEASE"

# Set provider acronym for PR titles
PROVIDER_ACRONYM=""
if [ -n "$PROVIDER" ]; then
  case "$PROVIDER" in
    "aws")
      PROVIDER_ACRONYM="CAPA"
      ;;
    "azure")
      PROVIDER_ACRONYM="CAPZ"
      ;;
    "vsphere")
      PROVIDER_ACRONYM="CAPV"
      ;;
    "cloud-director")
      PROVIDER_ACRONYM="CAPVCD"
      ;;
    *)
      PROVIDER_ACRONYM=$(echo "$PROVIDER" | tr 'a-z' 'A-Z')
      ;;
  esac
else
  PROVIDER_ACRONYM="CAPI"
fi

if [ -n "$GITHUB_OUTPUT" ]; then
  echo "LATEST_RELEASE=${LATEST_RELEASE}" >> $GITHUB_OUTPUT
  echo "NEXT_RELEASE=${NEXT_RELEASE}" >> $GITHUB_OUTPUT
fi

# Set output for GitHub Actions
if [ -n "${GITHUB_OUTPUT-}" ]; then
    echo "version=${NEXT_RELEASE}" >> "$GITHUB_OUTPUT"
    echo "base=${LATEST_RELEASE}" >> "$GITHUB_OUTPUT"
    echo "provider_acronym=${PROVIDER_ACRONYM}" >> "$GITHUB_OUTPUT"
fi

# Print for local use and piping
echo "version=${NEXT_RELEASE}"
echo "base=${LATEST_RELEASE}"
echo "provider_acronym=${PROVIDER_ACRONYM}"
