#!/bin/zsh
# Builds the production monolith directly from its source build context.
# Co-authored by: OpenCode and Igor Benicio de Mesquita

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")

if ! docker build --file "$project_root/src/main/docker/monolith/Dockerfile" --tag open-asset-allocator-monolith:latest "$project_root/src/main"; then
  echo "Failed to build production monolith image"
  exit 1
fi
