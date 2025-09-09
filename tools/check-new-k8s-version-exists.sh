#!/bin/bash
set -e

NEXT_RELEASE_VERSION=$1
# The GITHUB_TOKEN is not passed as an argument anymore,
# as 'gh' will automatically use it from the environment.

if [ -z "$NEXT_RELEASE_VERSION" ]; then
  echo "Usage: $0 <next_release_version>"
  exit 1
fi

# Extract major version from our release, e.g., v32.0.0 -> 32
RELEASE_MAJOR=$(echo "$NEXT_RELEASE_VERSION" | cut -d'v' -f2 | cut -d'.' -f1)

# This maps our release major to the k8s minor, e.g., 32 -> 1.32
K8S_TARGET_MINOR="1.${RELEASE_MAJOR}"

echo "Next release is ${NEXT_RELEASE_VERSION}, which requires a STABLE Kubernetes minor version ${K8S_TARGET_MINOR}"

# Use GitHub CLI to check for tags matching the target k8s minor version.
API_ENDPOINT="repos/kubernetes/kubernetes/git/matching-refs/tags/v${K8S_TARGET_MINOR}."

echo "Querying GitHub API for matching tags..."

RESPONSE=$(gh api "$API_ENDPOINT")

# We parse the JSON response to get the tag references (`ref`).
# Then, we use grep to filter out any pre-release tags (alpha, beta, rc).
# A stable tag will not contain these suffixes.
STABLE_RELEASES=$(echo "$RESPONSE" | jq -r '.[].ref' | grep -v -- '-rc' | grep -v -- '-beta' | grep -v -- '-alpha' || true)

if [[ -z "$STABLE_RELEASES" ]]; then
  echo "No STABLE releases found for Kubernetes minor ${K8S_TARGET_MINOR} yet. Skipping release creation."
  if [ -n "$GITHUB_OUTPUT" ]; then
    echo "VERSION_EXISTS=false" >> $GITHUB_OUTPUT
  fi
else
  echo "Found new STABLE Kubernetes version(s) for minor ${K8S_TARGET_MINOR}. Proceeding with release creation."
  echo "Stable releases found:"
  echo "$STABLE_RELEASES"
  if [ -n "$GITHUB_OUTPUT" ]; then
    echo "VERSION_EXISTS=true" >> $GITHUB_OUTPUT
  fi
fi
