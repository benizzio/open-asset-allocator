#!/bin/zsh

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")

docker compose -f $project_root/src/main/docker/flyway/docker-compose-flyway.yml run --rm flyway
