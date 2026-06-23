# DuckDB CLI container image.
# Co-authored by: OpenCode and benizzio
FROM ubuntu:24.04

WORKDIR /duckdb

RUN apt-get update \
    && apt-get install -y --no-install-recommends unzip ca-certificates \
    && rm -rf /var/lib/apt/lists/*
ADD https://github.com/duckdb/duckdb/releases/download/v1.2.2/duckdb_cli-linux-amd64.zip /duckdb/duckdb.zip
RUN unzip duckdb.zip && rm duckdb.zip

CMD ["sh", "-c", "./duckdb -unsigned < ./input/$INPUT_FILE"]
