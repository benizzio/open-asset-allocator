#!/bin/zsh

cd src/main/docker/duckdb

if [ -n "$1" ]; then
    echo "First parameter is: $1"
    docker compose -f docker-compose-duckdb.yml run --rm --env INPUT_FILE="$1" duckdb-cli
else
    docker compose -f docker-compose-duckdb.yml run --rm duckdb-cli
fi