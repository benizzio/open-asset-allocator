#!/usr/bin/env zsh
# Builds the standalone backend and Parcel development images.
# Co-authored by: OpenCode and Igor Benicio de Mesquita

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")
dev_db_dir="$project_root"/target/postgres-dev-data

if ! mkdir -p "$dev_db_dir"; then
  echo "Failed to create development PostgreSQL data directory"
  exit 1
fi
export POSTGRES_DEV_DATA_DIR="$dev_db_dir"

cd "$project_root"/src/main/docker || exit
if ! docker compose -f docker-compose-dev.yml build backend frontend; then
  echo "Failed to build development Docker Compose services"
  exit 1
fi
