#!/bin/zsh

# Removes the production and development Docker Compose resources.
# Co-authored by: OpenCode and Igor Benicio de Mesquita

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")
export POSTGRES_DEV_DATA_DIR="$project_root/target/postgres-dev-data"
export POSTGRES_DATA_DIR="$HOME/.open-asset-allocator/postgres-data"

cd "$project_root"/src/main/docker || exit
docker compose down
docker compose -f docker-compose-dev.yml down
