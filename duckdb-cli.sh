#!/bin/zsh

# First argument ($1): main script to run
# Options:
# --deps /path/to/file1 /path/to/file2 => list of dependecy files (files used by the main script)

# ==============================================================================================================
# VARIABLES
# ==============================================================================================================

# location (directory) of this running script
this_script_dir=$(dirname "$0")
# absolute path of location this script (project root dir)
project_root_abs_path=$(realpath "$this_script_dir")

# relative path of duckdb input target dir (duckdb volume for scripts)
input_target_dir_path=$project_root_abs_path/target/duckdb-input
# path of final input script in duckdb volume
std_input_file_path=$input_target_dir_path/input.sql

# ==============================================================================================================
# FUNCTIONS
# ==============================================================================================================

verify_file () {
    if [ ! -f "$1" ]; then
        echo "File $1 does not exist"
        return 1
    fi
}

# ==============================================================================================================
# SCRIPT BODY
# ==============================================================================================================

# checks if the target input path exists in the project
if [ ! -d "$input_target_dir_path" ]; then
    # if it does not, create it
    mkdir "$input_target_dir_path"
elseif [ -f "$std_input_file_path" ]
    # if it exists, try to remove any previous copies of an existent input file
    rm "$std_input_file_path"
fi

# checks the usage of the first argument
if [ -n "$1" ]; then

    # checks if it is a file
    verify_file "$1"
    if [ ! $? -eq 1 ]; then
        # copies the file to the duckdb input target dir
        echo "Copying main file $1 to DuckDB input"
        cp "$1" "$std_input_file_path"
    else
        exit 1
    fi
else
    echo "A input script file must be informed as a parameter"
    exit 1
fi

dep_file_paths=()

# move positional parameter 1 position to work with options
shift

# parse command-line arguments after main script
while [[ $# -gt 0 ]]; do

    # processing each options argument for known ones
    case $1 in

        # files argument
        --deps)
            # move positional parameters inside dependencies options (--deps is $0)
            shift
            # loop through all the dependency files parameters until finding a new "--" option
            # stores them all in the array
            while [[ $# -gt 0 && ! $1 =~ ^-- ]]; do

                # checks if it is a file
                verify_file "$1"
                # if it is copy it with the same name
                if [ ! $? -eq 1 ]; then
                    echo "Copying dependency file $1 to DuckDB input"
                    dest_file_path=$input_target_dir_path/$(basename "$1")
                    cp "$1" "$dest_file_path"
                    dep_file_paths+=("$dest_file_path")
                else
                    exit 1
                fi

                # move to the next positional parameter
                shift
            done
            ;;

        # Postgres port argument
        --pg-port)
            # move positional parameters inside dependencies options (--deps is $0)
            shift

            # checks if the argument is a number
            if [[ $1 =~ ^[0-9]+$ ]]; then
                pgport="$1"
            else
                echo "Postgres port must be a number"
                exit 1
            fi

            # move to the next positional parameter
            shift
            ;;

        # default verification
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
done

echo "Executing duckdb script $1"
export PGPORT=$pgport
docker compose -f "$project_root_abs_path"/src/main/docker/duckdb/docker-compose-duckdb.yml run --rm duckdb-cli

echo "Cleaning input files"
rm "$std_input_file_path"
for file_path in "${dep_file_paths[@]}"; do
    rm "$file_path"
done

exit 0
