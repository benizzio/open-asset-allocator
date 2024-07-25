#!/bin/zsh

script_dir=$(dirname "$0")
project_root=$(realpath "$script_dir")

target_relative_path=target/duckdb-input
std_input_file=$project_root/$target_relative_path/input.sql

if [ ! -d $project_root/$target_relative_path ]; then
    mkdir $project_root/$target_relative_path
elseif [ -f $std_input_file ]
    rm $std_input_file
fi


if [ -n "$1" ]; then

    if [ ! -f $1 ]; then
        echo "File $1 does not exist"
        exit 1
    else
        cp $1 $std_input_file
    fi

else
    echo "A file must be informed as a parameter"
    exit 1
fi

echo "Executing duckdb script $1"
docker compose -f $project_root/src/main/docker/duckdb/docker-compose-duckdb.yml run --rm duckdb-cli
rm $std_input_file

exit 0