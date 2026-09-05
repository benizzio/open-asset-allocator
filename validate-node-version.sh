#!/usr/bin/env bash
# Validates the repository-owned Node.js runtime references before builds and CI.
#
# Run `./validate-node-version.sh` from any working directory. The command exits
# unsuccessfully when .nvmrc and the application Node.js images diverge.
# Authored by: OpenCode

set -euo pipefail

script_directory=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
node_version_file="$script_directory/.nvmrc"
frontend_dockerfile="$script_directory/src/main/docker/frontend/Dockerfile"
monolith_dockerfile="$script_directory/src/main/docker/monolith/Dockerfile"

# Terminates validation with a consistent diagnostic.
fail() {
  printf 'Node.js version validation failed: %s\n' "$1" >&2
  exit 1
}

# Extracts one literal Node.js image tag and optional digest from a Dockerfile.
read_node_image() {
  local dockerfile=$1
  local reference

  reference=$(awk '
    $1 == "FROM" && $2 ~ /^node:/ { count += 1; image = $2 }
    END { if (count == 1) print image; else exit 1 }
  ' "$dockerfile") || fail "$dockerfile must contain exactly one literal Node.js FROM instruction"
  reference=${reference#node:}
  local tag=${reference%@*}
  local digest=""
  if [[ $reference == *@* ]]; then
    digest=${reference#*@}
    [[ $digest =~ ^sha256:[a-f0-9]{64}$ ]] || fail "$dockerfile contains an invalid Node.js image digest"
  fi

  printf '%s|%s\n' "$tag" "$digest"
}

[[ -f $node_version_file ]] || fail "$node_version_file does not exist"
node_version=$(tr -d '[:space:]' < "$node_version_file")
[[ $node_version =~ ^[0-9]+\.[0-9]+\.[0-9]+$ ]] || fail ".nvmrc must contain one exact major.minor.patch version"

expected_tag="$node_version-bookworm-slim"
IFS='|' read -r frontend_tag frontend_digest < <(read_node_image "$frontend_dockerfile")
IFS='|' read -r monolith_tag monolith_digest < <(read_node_image "$monolith_dockerfile")

[[ $frontend_tag == "$expected_tag" ]] || fail "$frontend_dockerfile uses $frontend_tag instead of $expected_tag"
[[ $monolith_tag == "$expected_tag" ]] || fail "$monolith_dockerfile uses $monolith_tag instead of $expected_tag"

[[ -n $frontend_digest && -n $monolith_digest ]] || fail "both Node.js images must be digest-pinned"
[[ $frontend_digest == "$monolith_digest" ]] || fail "the Node.js image digests must match"

printf 'Node.js application runtime references match: %s\n' "$node_version"
