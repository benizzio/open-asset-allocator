#!/bin/zsh

docker compose -f docker-compose-duckdb.yml run --rm --build duckdb-cli
