#!/bin/zsh

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")

cd "$project_root"/src/main/docker || exit
docker compose down
docker compose -f docker-compose-dev.yml down
