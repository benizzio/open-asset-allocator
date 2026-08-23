#!/bin/zsh
# Starts the development stack and keeps Parcel logs in the foreground.
# Co-authored by: OpenCode and Igor Benicio de Mesquita

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")
dev_db_dir="$project_root"/target/postgres-dev-data

[ -d "$dev_db_dir" ] || mkdir "$dev_db_dir"
export POSTGRES_DEV_DATA_DIR="$dev_db_dir"

cd "$project_root"/src/main/docker || exit
if ! docker compose -f docker-compose-dev.yml up --attach frontend; then
  echo "Failed to start development Docker Compose stack"
  exit 1
fi
