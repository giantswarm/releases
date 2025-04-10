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
# Array to store provider versions 
declare -a provider_versions
# Track providers which are already queried
declare -a queried_providers

# Get all active versions for a provider
get_active_versions() {
  local provider=$1
  
  # Rename provider for Azure for Grafana query
  api_provider=$provider
  if [[ "$provider" == "azure" ]]; then
    api_provider="capz"
  fi
  
  echo "Fetching currently used versions for provider $provider..."
  
  # Grafana DataSource Query with error handling
  response=$(curl --silent --fail --show-error --location --request POST 'https://giantswarm.grafana.net/api/ds/query' \
    -H 'Content-Type: application/json' \
    -H 'Accept: application/json' \
    -H "Authorization: Bearer $GRAFANA_API_KEY" \
    -d "{\"from\":\"now-5m\",\"to\":\"now\",\"queries\":[{\"refId\":\"A\",\"expr\":\"sum(aggregation:giantswarm:cluster_release_version{provider=\\\"$api_provider\\\", release_version=~\\\".*\\\", installation=~\\\".*\\\", cluster_type=~\\\".*\\\", customer=~\\\".*\\\"}) by (release_version)\",\"datasource\":{\"uid\":\"000000006\",\"type\":\"prometheus\"}}]}" 2>&1)
  
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
  provider_versions["$provider"]="$used_versions"
  queried_providers+=("$provider")
  
  return 0
}

# Check if a specific version is in use
is_version_in_use() {
  local provider=$1
  local version=$2
  
  # If we haven't queried this provider yet, do it now
  if ! echo "${queried_providers[@]}" | grep -q "$provider"; then
    if ! get_active_versions "$provider"; then
      # If API call failed, assume version is in use to be safe
      return 0
    fi
  fi
  
  # Check if the version is in the list of used versions
  if echo "${provider_versions[$provider]}" | grep -q "^$version$"; then
    return 0
  else
    return 1
  fi
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
  
  if is_version_in_use "$provider" "$version"; then
    active_releases+=("$provider/$version")
  fi
done

if [ ${#active_releases[@]} -gt 0 ]; then
  echo "::error:: The following releases cannot be archived because they are still in use:"
  for release in "${active_releases[@]}"; do
    echo "  - $release"
  done
  exit 1
fi

echo "All releases are safe to archive."
exit 0
