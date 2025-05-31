#!/usr/bin/env zsh

# This script was generated with the assistance of AI (GitHub Copilot).
# It creates a dependency graph for a Go module dependency.
# REQUIRES: graphviz

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <dependency>"
  exit 1
fi

# Get the root project directory (where this script is located)
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Define output directory
OUTPUT_DIR="${PROJECT_ROOT}/target/dependency-graph"

# Create output directory if it doesn't exist
mkdir -p "${OUTPUT_DIR}"

# Hardcoded Go module directory path - simpler and more predictable
GO_MODULE_DIR="${PROJECT_ROOT}/src/main/go"

# Check if go.mod exists in the expected location
if [[ ! -f "${GO_MODULE_DIR}/go.mod" ]]; then
  echo "Error: go.mod file not found at ${GO_MODULE_DIR}"
  exit 1
fi

echo "Using Go module at: ${GO_MODULE_DIR}"

# Change to the Go module directory to run go mod commands
cd "${GO_MODULE_DIR}" || {
  echo "Error: Cannot change to Go module directory at ${GO_MODULE_DIR}"
  exit 1
}

dep="$1"
dotfile="${dep:t}.dot"
imgfile="${dep:t}.png"
filtered_file="filtered.txt"

# Generate filtered dependency lines
go mod graph | grep -F "$dep" > "${OUTPUT_DIR}/${filtered_file}"

# Create DOT file
{
  echo "digraph G {"
  awk '{print "\"" $1 "\" -> \"" $2 "\""}' "${OUTPUT_DIR}/${filtered_file}"
  echo "}"
} > "${OUTPUT_DIR}/${dotfile}"

# Generate PNG image
dot -Tpng "${OUTPUT_DIR}/${dotfile}" -o "${OUTPUT_DIR}/${imgfile}"

echo "Dependency graph generated in target/dependency-graph:"
echo "- ${dotfile}"
echo "- ${imgfile}"
echo "- ${filtered_file}"
