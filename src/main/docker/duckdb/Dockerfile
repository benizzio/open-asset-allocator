# DuckDB CLI container image.
# Co-authored by: OpenCode and benizzio
FROM ubuntu:24.04

WORKDIR /duckdb

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        ca-certificates=20260601~24.04.1 \
        unzip=6.0-28ubuntu4.1 \
        wget=1.21.4-1ubuntu4.4 \
    && rm -rf /var/lib/apt/lists/*
RUN wget --quiet --https-only --output-document=duckdb.zip \
        https://github.com/duckdb/duckdb/releases/download/v1.2.2/duckdb_cli-linux-amd64.zip \
    && printf '%s\n' "fc153822f59283e0a9374168cce5bc85a9985e699d9857842597882062fd2cb5  duckdb.zip" \
        > duckdb.zip.sha256 \
    && sha256sum -c duckdb.zip.sha256 \
    && unzip duckdb.zip \
    && rm duckdb.zip duckdb.zip.sha256 \
    && groupadd --gid 10001 duckdb \
    && useradd --uid 10001 --gid duckdb --home-dir /duckdb --no-create-home --shell /usr/sbin/nologin duckdb \
    && chown -R duckdb:duckdb /duckdb

USER duckdb
CMD ["sh", "-c", "./duckdb -unsigned < ./input/$INPUT_FILE"]
