#!/usr/bin/env zsh

# REQUIRES: graphviz

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <dependency>"
  exit 1
fi

dep="$1"
dotfile="${dep:t}.dot"
imgfile="${dep:t}.png"

# Generate filtered dependency lines
go mod graph | grep "$dep" > filtered.txt

# Create DOT file
{
  echo "digraph G {"
  awk '{print "\"" $1 "\" -> \"" $2 "\""}' filtered.txt
  echo "}"
} > "$dotfile"

# Generate PNG image
dot -Tpng "$dotfile" -o "$imgfile"

echo "Dependency graph generated: $imgfile"