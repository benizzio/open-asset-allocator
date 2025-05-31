#!/usr/bin/env zsh

# This script was generated with the assistance of AI (GitHub Copilot).
# It creates a dependency graph for a Go module dependency.
# REQUIRES: graphviz

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <dependency>"
  exit 1
fi

# Locate Go module directory by finding go.mod
GO_MODULE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
while [[ ! -f "${GO_MODULE_DIR}/go.mod" && "${GO_MODULE_DIR}" != "/" ]]; do
  GO_MODULE_DIR="$(dirname "${GO_MODULE_DIR}")"
done
if [[ ! -f "${GO_MODULE_DIR}/go.mod" ]]; then
  echo "Error: go.mod file not found. Ensure this script is run within a Go module."
  exit 1
fi
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
go mod graph | grep -F "$dep" > "${SCRIPT_DIR}/${OUTPUT_DIR}/${filtered_file}"

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
