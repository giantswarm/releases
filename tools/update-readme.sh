#!/bin/bash
set -euo pipefail

if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <provider> <release_version>"
  exit 1
fi

PROVIDER=$1
RELEASE_VERSION=$2 # e.g., v31.2.0

README_FILE="README.md"
TMP_README_FILE=$(mktemp)
PROVIDER_DIR_LINK=$PROVIDER
HEADING_SEARCH_STRING=""

# Handle provider aliases and specific heading names
case $PROVIDER in
  aws)
    PROVIDER_DIR_LINK="capa"
    HEADING_SEARCH_STRING="## AWS"
    ;;
  azure)
    HEADING_SEARCH_STRING="## Azure"
    ;;
  capa)
    HEADING_SEARCH_STRING="## AWS"
    ;;
  vsphere)
    HEADING_SEARCH_STRING="## vSphere"
    ;;
  "cloud-director")
    HEADING_SEARCH_STRING="## VMware Cloud Director"
    ;;
  "eks")
    HEADING_SEARCH_STRING="## EKS"
    ;;
  *)
    echo "Error: Unknown provider '$PROVIDER'. Please add it to the case statement in tools/update-readme.sh" >&2
    exit 1
    ;;
esac

MAJOR_V_NUM=$(echo "$RELEASE_VERSION" | cut -d. -f1 | tr -d 'v')
MINOR_V_NUM=$(echo "$RELEASE_VERSION" | cut -d. -f2)

MAJOR_V_HEADER="- v${MAJOR_V_NUM}"
MINOR_V_HEADER="  - v${MAJOR_V_NUM}.${MINOR_V_NUM}"
PATCH_LINE="    - [${RELEASE_VERSION}](https://github.com/giantswarm/releases/tree/master/${PROVIDER_DIR_LINK}/${RELEASE_VERSION})"

# --- State machine variables ---
in_provider_section=false
in_major_section=false
in_minor_section=false
version_inserted=false
major_found=false
minor_found=false

while IFS= read -r line || [ -n "$line" ]; do

  # End of a section logic
  if $in_provider_section && [[ "$line" =~ ^##\  && "$line" != "$HEADING_SEARCH_STRING" ]]; then
    if ! $version_inserted; then
        if ! $major_found; then
            echo "$MAJOR_V_HEADER" >> "$TMP_README_FILE"
            major_found=true
        fi
        if ! $minor_found; then
            echo "$MINOR_V_HEADER" >> "$TMP_README_FILE"
            minor_found=true
        fi
        echo "$PATCH_LINE" >> "$TMP_README_FILE"
        version_inserted=true
    fi
    in_provider_section=false
    in_major_section=false
    in_minor_section=false
  fi
  if $in_major_section && [[ "$line" =~ ^-\ v && "$line" != "$MAJOR_V_HEADER" ]]; then
    if ! $version_inserted; then
        if ! $minor_found; then
            echo "$MINOR_V_HEADER" >> "$TMP_README_FILE"
            minor_found=true
        fi
        echo "$PATCH_LINE" >> "$TMP_README_FILE"
        version_inserted=true
    fi
    in_major_section=false
    in_minor_section=false
  fi
  if $in_minor_section && [[ "$line" =~ ^\ \ -\ v && "$line" != "$MINOR_V_HEADER" ]]; then
     if ! $version_inserted; then
        echo "$PATCH_LINE" >> "$TMP_README_FILE"
        version_inserted=true
    fi
    in_minor_section=false
  fi

  # Find provider section
  if [[ "$line" == "$HEADING_SEARCH_STRING" ]]; then
    in_provider_section=true
  fi

  # Find major version section
  if $in_provider_section && ! $major_found && [[ "$line" =~ ^-\ v([0-9]+)$ ]]; then
    EXISTING_MAJOR_V_NUM=${BASH_REMATCH[1]}
    if [[ $MAJOR_V_NUM -gt $EXISTING_MAJOR_V_NUM ]] && ! $version_inserted; then
      echo "$MAJOR_V_HEADER" >> "$TMP_README_FILE"
      echo "$MINOR_V_HEADER" >> "$TMP_README_FILE"
      echo "$PATCH_LINE" >> "$TMP_README_FILE"
      echo "" >> "$TMP_README_FILE" # Add a space AFTER a new major version
      version_inserted=true
    fi
  fi
  if $in_provider_section && [[ "$line" == "$MAJOR_V_HEADER" ]]; then
    in_major_section=true
    major_found=true
  fi

  # Find minor version section
  if $in_major_section && ! $minor_found && [[ "$line" =~ ^\ \ -\ v[0-9]+\.([0-9]+)$ ]]; then
    EXISTING_MINOR_V_NUM=${BASH_REMATCH[1]}
    if [[ $MINOR_V_NUM -gt $EXISTING_MINOR_V_NUM ]] && ! $version_inserted; then
      echo "$MINOR_V_HEADER" >> "$TMP_README_FILE"
      echo "$PATCH_LINE" >> "$TMP_README_FILE"
      version_inserted=true
    fi
  fi
  if $in_major_section && [[ "$line" == "$MINOR_V_HEADER" ]]; then
    in_minor_section=true
    minor_found=true
  fi

  # Find patch insertion point
  if $in_minor_section && ! $version_inserted && [[ "$line" =~ ^\ \ \ \ -\ \[(v[0-9]+\.[0-9]+\.[0-9]+)\]\(.*$ ]]; then
    EXISTING_VERSION=${BASH_REMATCH[1]}
    if dpkg --compare-versions "${RELEASE_VERSION#v}" "gt" "${EXISTING_VERSION#v}"; then
      echo "$PATCH_LINE" >> "$TMP_README_FILE"
      version_inserted=true
    fi
  fi

  echo "$line" >> "$TMP_README_FILE"

done < "$README_FILE"

# Handle case where the new release is the very last one in the file
if $in_provider_section && ! $version_inserted; then
    if ! $major_found; then echo "$MAJOR_V_HEADER" >> "$TMP_README_FILE"; fi
    if ! $minor_found; then echo "$MINOR_V_HEADER" >> "$TMP_README_FILE"; fi
    echo "$PATCH_LINE" >> "$TMP_README_FILE"
    version_inserted=true
fi


mv "$TMP_README_FILE" "$README_FILE"

if $version_inserted; then
    echo "Successfully updated README.md with ${RELEASE_VERSION} for ${PROVIDER}"
else
    echo "Warning: Could not automatically insert ${RELEASE_VERSION} into README.md. Manual update required." >&2
fi


