#!/usr/bin/env zsh

# This script was generated with the assistance of AI (GitHub Copilot).
# It creates a dependency graph for a Go module dependency.
# REQUIRES: graphviz

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <dependency>"
  exit 1
fi

# Path to Go module directory
GO_MODULE_DIR="src/main/go"
# Path to output directory (relative to project root)
OUTPUT_DIR="target/dependency-graph"

# First, ensure the output directory exists (from project root)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
mkdir -p "${SCRIPT_DIR}/${OUTPUT_DIR}"

# Change to the Go module directory to run go mod commands
cd "${SCRIPT_DIR}/${GO_MODULE_DIR}" || {
  echo "Error: Cannot find Go module directory at ${SCRIPT_DIR}/${GO_MODULE_DIR}"
  exit 1
}

dep="$1"
dotfile="${dep:t}.dot"
imgfile="${dep:t}.png"
filtered_file="filtered.txt"

# Generate filtered dependency lines
go mod graph | grep "$dep" > "${SCRIPT_DIR}/${OUTPUT_DIR}/${filtered_file}"

# Create DOT file
{
  echo "digraph G {"
  awk '{print "\"" $1 "\" -> \"" $2 "\""}' "${SCRIPT_DIR}/${OUTPUT_DIR}/${filtered_file}"
  echo "}"
} > "${SCRIPT_DIR}/${OUTPUT_DIR}/${dotfile}"

# Generate PNG image
dot -Tpng "${SCRIPT_DIR}/${OUTPUT_DIR}/${dotfile}" -o "${SCRIPT_DIR}/${OUTPUT_DIR}/${imgfile}"

echo "Dependency graph generated in ${OUTPUT_DIR}:"
echo "- ${dotfile}"
echo "- ${imgfile}"
echo "- ${filtered_file}"
