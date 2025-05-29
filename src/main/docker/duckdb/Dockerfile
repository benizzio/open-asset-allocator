FROM ubuntu:24.10

WORKDIR /duckdb

RUN apt update -y && apt upgrade -y && apt install -y wget unzip && apt clean
ADD https://github.com/duckdb/duckdb/releases/download/v1.2.2/duckdb_cli-linux-amd64.zip /duckdb/duckdb.zip
RUN unzip duckdb.zip && rm duckdb.zip

CMD ["sh", "-c", "./duckdb -unsigned < ./input/$INPUT_FILE"]
