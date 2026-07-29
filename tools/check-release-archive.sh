#!/bin/bash
set -e

GRAFANA_API_KEY=$1

if [[ -z "$GRAFANA_API_KEY" ]]; then
  echo "Error: GRAFANA_API_KEY is required"
  echo "Usage: $0 <GRAFANA_API_KEY>"
  exit 1
fi

# Check for any renamed files in the last commit
DEFAULT_BRANCH=$(git remote show origin | grep 'HEAD branch' | cut -d' ' -f5)
echo "Using default branch: $DEFAULT_BRANCH"
git fetch origin $DEFAULT_BRANCH

# Get renamed files
renamed_files=$(git diff --name-status origin/$DEFAULT_BRANCH...HEAD | grep "^R" || echo "")

if [[ -z "$renamed_files" ]]; then
  echo "No renamed files found in this PR. Nothing to check."
  exit 0
fi

# Array to store active releases (still in use)
declare -a active_releases
# Releases blocked because a smaller active release still exists
declare -a below_floor_releases
# Releases blocked because they were never deprecated
declare -a still_active_releases

# Caches, kept as "key|value" records so this stays compatible with bash 3.2 (macOS)
# which has no associative arrays.
in_use_records=""    # "provider|version" per line, versions currently in use
queried_providers="" # "provider" per line, providers already queried
floor_records=""     # "provider|floor" per line, empty floor means none applies

# Compare two semver versions, true if $1 > $2
version_gt() {
  [[ "$1" != "$2" ]] && [[ "$(printf '%s\n%s\n' "$1" "$2" | sort -V | head -n1)" == "$2" ]]
}

# Resolve the archive floor for a provider into ARCHIVE_FLOOR, determining it only
# once per provider. Sets an empty value when no floor applies. This assigns to a
# global instead of printing, so that the cache survives across calls.
ARCHIVE_FLOOR=""
resolve_archive_floor() {
  local provider=$1

  if ! printf '%s\n' "$floor_records" | grep -q "^${provider}|"; then
    local floor
    floor=$("$(dirname "$0")/lowest-active-release.sh" "$provider")
    floor_records="${floor_records}${provider}|${floor}"$'\n'

    if [[ -n "$floor" ]]; then
      echo "Archive floor for $provider is $floor (lowest active release)"
    else
      echo "Provider $provider has no active releases, no archive floor applies"
    fi
  fi

  ARCHIVE_FLOOR=$(printf '%s\n' "$floor_records" | grep "^${provider}|" | head -n1 | cut -d'|' -f2)
}

# Check whether a version sits above the provider's archive floor, i.e. whether a
# smaller active release still exists. Such releases must stay available even when
# deprecated, because clusters on the lower release may still need them as an
# intermediate upgrade step. See https://github.com/giantswarm/roadmap/issues/4325
is_above_archive_floor() {
  local provider=$1
  local version=$2

  resolve_archive_floor "$provider"

  [[ -n "$ARCHIVE_FLOOR" ]] && version_gt "$version" "$ARCHIVE_FLOOR"
}

# Get all active versions for a provider
get_active_versions() {
  local provider=$1
  
  # Rename provider for Azure for Grafana query
  api_provider=$provider
  if [[ "$provider" == "azure" ]]; then
    api_provider="capz"
  fi
  
  echo "Fetching currently used versions for provider $provider..."
  
  # Grafana DataSource Query with error handling.
  # The 7d lookback keeps a release from looking unused just because its metrics were
  # briefly missing: a scrape gap, an unreachable management cluster, or a dev cluster
  # that is shut down over the weekend.
  response=$(curl --silent --fail --show-error --location --request POST 'https://giantswarm.grafana.net/api/ds/query' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H "Authorization: Bearer $GRAFANA_API_KEY" \
    -d "{\"from\":\"now-5m\",\"to\":\"now\",\"queries\":[{\"refId\":\"A\",\"expr\":\"sum(max_over_time(aggregation:giantswarm:cluster_release_version{provider=\\\"$api_provider\\\", release_version=~\\\".*\\\", installation=~\\\".*\\\", cluster_type=~\\\".*\\\", customer=~\\\".*\\\"}[7d])) by (release_version)\",\"datasource\":{\"uid\":\"000000006\",\"type\":\"prometheus\"}}]}" 2>&1)
  
  # Check for curl failures
  if [[ $? -ne 0 ]]; then
    echo "::error:: Failed to connect to Grafana API: $response"
    echo "::error:: Cannot verify if releases are safe to archive. Blocking archive operation."
    return 1
  fi
  
  # Validate API response
  if ! echo "$response" | jq -e '.results.A.frames' >/dev/null 2>&1; then
    echo "::error:: Invalid or unexpected response from Grafana API:"
    echo "$response" | head -n 20
    echo "::error:: Cannot verify if releases are safe to archive. Blocking archive operation."
    return 1
  fi
  
  # Extract active release versions for the current provider
  used_versions=$(echo "$response" | jq -r '.results.A.frames[].schema.fields[] | select(.labels != null) | .labels.release_version')
  
  # Check if we got any versions (empty result could indicate a misconfigured query)
  if [[ -z "$used_versions" ]]; then
    echo "::warning:: No active release versions found for provider $provider."
    echo "::warning:: This could mean the query is misconfigured or no clusters exist with this provider."
  fi
  
  echo "Currently used versions for $provider:"
  echo "$used_versions"
  
  # Store the result
  while IFS= read -r used_version; do
    [[ -z "$used_version" ]] && continue
    in_use_records="${in_use_records}${provider}|${used_version}"$'\n'
  done <<< "$used_versions"
  queried_providers="${queried_providers}${provider}"$'\n'

  return 0
}

# Check if a specific version is in use
is_version_in_use() {
  local provider=$1
  local version=$2
  
  # If we haven't queried this provider yet, do it now
  if ! printf '%s\n' "$queried_providers" | grep -qxF "$provider"; then
    if ! get_active_versions "$provider"; then
      # If API call failed, assume version is in use to be safe
      return 0
    fi
  fi

  # Check if the version is in the list of used versions
  printf '%s\n' "$in_use_records" | grep -qxF "${provider}|${version}"
}

# Process renamed files
found_archive_move=false
declare -a versions_to_check

while IFS= read -r line; do
  # Skip empty lines
  [ -z "$line" ] && continue
  
  # Split the line into parts
  change_type=$(echo "$line" | cut -f1)
  source_path=$(echo "$line" | cut -f2)
  target_path=$(echo "$line" | cut -f3)
  
  # Extract provider and version from paths
  if [[ $source_path =~ ^([^/]+)/v([0-9]+\.[0-9]+\.[0-9]+.*)/[^/]+$ ]]; then
    provider=${BASH_REMATCH[1]}
    version=${BASH_REMATCH[2]}
    
    # Check target path is in archived directory
    if [[ $target_path =~ ^([^/]+)/archived/v([0-9]+\.[0-9]+\.[0-9]+.*)/[^/]+$ ]]; then
      found_archive_move=true
      
      # Track unique provider/version combinations to check
      version_entry="$provider/$version"
      if ! echo "${versions_to_check[@]}" | grep -q "$version_entry"; then
        versions_to_check+=("$version_entry")
      fi
    fi
  fi
done <<< "$renamed_files"

if [ "$found_archive_move" = false ]; then
  echo "No release versions being moved to archive were detected."
  exit 0
fi

# Check all versions against their respective providers
for version_entry in "${versions_to_check[@]}"; do
  provider=$(echo "$version_entry" | cut -d'/' -f1)
  version=$(echo "$version_entry" | cut -d'/' -f2)

  # Only deprecated releases may be archived. The archive floor below is derived from
  # the active releases, so archiving one straight from active would move the floor
  # instead of being blocked by it.
  archived_release_yaml="$provider/archived/v$version/release.yaml"
  if [[ -f "$archived_release_yaml" ]] && grep -q -E "^\s*state:\s*active\s*$" "$archived_release_yaml"; then
    still_active_releases+=("$provider/$version")
    continue
  fi

  if is_version_in_use "$provider" "$version"; then
    active_releases+=("$provider/$version")
    continue
  fi

  if is_above_archive_floor "$provider" "$version"; then
    below_floor_releases+=("$provider/$version")
  fi
done

blocked=false

if [ ${#still_active_releases[@]} -gt 0 ]; then
  blocked=true
  echo "::error:: The following releases cannot be archived because they are still active."
  echo "::error:: A release has to be deprecated before it can be archived."
  for release in "${still_active_releases[@]}"; do
    echo "  - $release"
  done
fi

if [ ${#active_releases[@]} -gt 0 ]; then
  blocked=true
  echo "::error:: The following releases cannot be archived because they are still in use:"
  for release in "${active_releases[@]}"; do
    echo "  - $release"
  done
fi

if [ ${#below_floor_releases[@]} -gt 0 ]; then
  blocked=true
  echo "::error:: The following releases cannot be archived because a smaller active release still exists."
  echo "::error:: Clusters on the lower release may still need them as an intermediate upgrade step."
  echo "::error:: They can stay deprecated, but must remain available until the lower release is archived."
  for release in "${below_floor_releases[@]}"; do
    provider=$(echo "$release" | cut -d'/' -f1)
    resolve_archive_floor "$provider"
    echo "  - $release (lowest active release for $provider is $ARCHIVE_FLOOR)"
  done
fi

if [ "$blocked" = true ]; then
  exit 1
fi

echo "All releases are safe to archive."
exit 0
