#!/bin/zsh

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")
application_dir=~/.open-asset-allocator
db_dir=$application_dir/postgres-data

[ -d "$application_dir" ] || mkdir "$application_dir"
[ -d "$db_dir" ] || mkdir "$db_dir"
export POSTGRES_DATA_DIR="$db_dir"

cd "$project_root"/src/main/docker || exit
docker compose up -d