#!/usr/bin/env zsh

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")

cd "$project_root"/src/main/docker || exit
if ! docker compose -f docker-compose-dev.yml build backend; then
  echo "Failed to build docker compose backend service"
  exit 1
fi
