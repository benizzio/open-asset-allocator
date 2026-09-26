#!/usr/bin/env bash
# Maintains independent Graphify corpora for the frontend, backend, and remainder.
# Run ./graphify.sh update frontend (or backend, remainder, all) from any directory.
# Authored by: GPT-6 Sol

set -euo pipefail

repository_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
graphify_command=${GRAPHIFY_BIN:-graphify}
expected_version=$(<"$repository_root/.agents/skills/graphify/.graphify_version")

# Prints the supported operations and their scope arguments.
usage() {
  printf 'Usage: %s <update|refresh> <frontend|backend|remainder|all>\n' "$0" >&2
  printf '  update: local AST rebuild; preserves previously extracted semantic nodes.\n' >&2
  printf '  refresh: semantic extraction through a configured Graphify backend.\n' >&2
}

# Reports an invalid configuration or a failed scope operation.
fail() {
  printf 'Graphify scopes: %s\n' "$1" >&2
  exit 1
}

# Selects a scan root and graph-specific exclusions without inherited ignore rules.
select_scope() {
  case $1 in
    frontend)
      scan_root="$repository_root/src/main/web-static"
      excludes=()
      ;;
    backend)
      scan_root="$repository_root/src/main/go"
      excludes=()
      ;;
    remainder)
      scan_root="$repository_root"
      excludes=('src/main/web-static/' 'src/main/go/')
      ;;
    *) fail "unknown scope: $1" ;;
  esac
  output_dir="$scan_root/graphify-out"
}

# Recreates the ignored Graphify build configuration from versioned scope rules.
prepare_scope() {
  mkdir -p "$output_dir"
  local json_excludes=""
  local pattern
  for pattern in "${excludes[@]}"; do
    [[ -z $json_excludes ]] || json_excludes+=,
    json_excludes+="\"$pattern\""
  done
  printf '{"excludes":[%s],"gitignore":true}\n' "$json_excludes" > "$output_dir/.graphify_build.json"
}

# Runs a native Graphify operation and reports which corpus failed.
run_scope() {
  local scope=$1
  select_scope "$scope"
  prepare_scope
  printf '[graphify scopes] %s: %s\n' "$scope" "$operation"

  local -a exclude_args=()
  local pattern
  for pattern in "${excludes[@]}"; do
    exclude_args+=(--exclude "$pattern")
  done

  if [[ $operation == update && -f $output_dir/graph.json ]]; then
    "$graphify_command" update "$scan_root" || fail "$scope: AST update failed"
  elif [[ $operation == update ]]; then
    "$graphify_command" extract "$scan_root" --code-only "${exclude_args[@]}" || fail "$scope: initial AST extraction failed"
    "$graphify_command" cluster-only "$scan_root" --no-label || fail "$scope: report generation failed"
  else
    "$graphify_command" extract "$scan_root" "${exclude_args[@]}" || fail "$scope: semantic refresh failed"
    "$graphify_command" cluster-only "$scan_root" --no-label || fail "$scope: report generation failed"
  fi
}

if [[ $# -ne 2 ]]; then
  usage
  exit 2
fi
operation=$1
scope=$2
[[ $operation == update || $operation == refresh ]] || { usage; exit 2; }
[[ $scope == frontend || $scope == backend || $scope == remainder || $scope == all ]] || { usage; exit 2; }

actual_version=$("$graphify_command" --version) || fail 'Graphify is not installed'
[[ $actual_version == "graphify $expected_version" ]] || fail "expected Graphify $expected_version, found $actual_version"

cd "$repository_root"
if [[ $scope == all ]]; then
  for selected_scope in frontend backend remainder; do
    run_scope "$selected_scope"
  done
else
  run_scope "$scope"
fi
