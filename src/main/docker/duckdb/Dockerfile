FROM ubuntu:24.10

WORKDIR /duckdb

RUN apt update -y && apt upgrade -y && apt install -y wget unzip && apt clean
#RUN wget -O duckdb.zip https://github.com/duckdb/duckdb/releases/download/v1.1.3/duckdb_cli-linux-amd64.zip && unzip duckdb.zip && rm duckdb.zip
ADD https://github.com/duckdb/duckdb/releases/download/v1.2.2/duckdb_cli-linux-amd64.zip /duckdb/duckdb.zip
RUN unzip duckdb.zip && rm duckdb.zip

CMD ./duckdb -unsigned < ./input/"$INPUT_FILE"
