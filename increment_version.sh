#!/bin/bash

# Get the latest tag
LATEST_TAG=$(git describe --tags $(git rev-list --tags --max-count=1) 2>/dev/null)

# Check if a tag was found
if [ -z "$LATEST_TAG" ]; then
  # No tags found, create the initial tag
  NEW_TAG="v0.1.0"
else
  # Extract the version numbers
  IFS='.' read -r -a VERSION_PARTS <<< "${LATEST_TAG#v}"

  # Increment the patch version
  PATCH_VERSION=$((VERSION_PARTS[2]+1))

  # Create the new tag
  NEW_TAG="v${VERSION_PARTS[0]}.${VERSION_PARTS[1]}.$PATCH_VERSION"
fi

# Output the new tag
echo "New tag: $NEW_TAG"

# Create and push the new tag
git tag "$NEW_TAG"
git push origin "$NEW_TAG"
