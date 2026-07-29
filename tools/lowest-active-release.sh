#!/bin/bash
# Prints the lowest active release version (without the "v" prefix) for a provider.
#
# This version is the "archive floor": clusters sitting on it may still need any
# higher release as an intermediate step of an upgrade path that was tested earlier,
# so nothing above the floor may be archived even once it is deprecated.
# See https://github.com/giantswarm/roadmap/issues/4325
#
# Only releases directly under the provider directory are considered; anything
# already under <provider>/archived/ is ignored. Prints nothing if the provider has
# no active releases.
#
# Usage: lowest-active-release.sh <provider>

set -eo pipefail

PROVIDER=$1

if [[ -z "$PROVIDER" ]]; then
  echo "Error: provider is required" >&2
  echo "Usage: $0 <provider>" >&2
  exit 1
fi

if [[ ! -d "$PROVIDER" ]]; then
  echo "Error: provider directory $PROVIDER does not exist" >&2
  exit 1
fi

active_versions=""

while IFS= read -r -d $'\0' release_dir; do
  release_yaml="$release_dir/release.yaml"
  [[ -f "$release_yaml" ]] || continue

  if grep -q -E "^\s*state:\s*active\s*$" "$release_yaml"; then
    active_versions+="$(basename "$release_dir" | sed 's/^v//')"$'\n'
  fi
done < <(find "$PROVIDER" -maxdepth 1 -mindepth 1 -type d -name 'v*.*.*' -print0)

if [[ -z "${active_versions//[[:space:]]/}" ]]; then
  exit 0
fi

printf '%s' "$active_versions" | grep -v '^$' | sort -V | head -n1
