#!/usr/bin/env bash
# Checks scope routing and clean-checkout configuration without invoking Graphify.
# Run bash src/test/graphify-scopes.sh from any working directory.
# Authored by: GPT-6 Sol

set -euo pipefail

repository_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd)
fixture=$(mktemp -d /tmp/opencode/graphify-scopes.XXXXXX)
trap 'rm -rf -- "$fixture"' EXIT

mkdir -p "$fixture/.agents/skills/graphify" "$fixture/src/main/web-static" "$fixture/src/main/go"
cp "$repository_root/graphify.sh" "$fixture/graphify.sh"
cp "$repository_root/.agents/skills/graphify/.graphify_version" "$fixture/.agents/skills/graphify/.graphify_version"

# Simulates the supported Graphify CLI while recording arguments and scan roots.
cat > "$fixture/graphify-mock" <<'MOCK'
#!/usr/bin/env bash
set -euo pipefail
if [[ $1 == --version ]]; then
  printf 'graphify 0.9.64\n'
  exit
fi
printf '%s\n' "$*" >> "$GRAPHIFY_TEST_LOG"
if [[ ${GRAPHIFY_TEST_FAIL:-} == "$1" ]]; then exit 1; fi
if [[ $1 == extract ]]; then
  mkdir -p "$2/graphify-out"
  printf '{"nodes":[],"links":[]}\n' > "$2/graphify-out/graph.json"
fi
MOCK
chmod +x "$fixture/graphify-mock"
export GRAPHIFY_BIN="$fixture/graphify-mock" GRAPHIFY_TEST_LOG="$fixture/operations.log"

# Fails the test with a message tied to the broken routing contract.
check() {
  "$@" || { printf 'Graphify scope test failed: %s\n' "$*" >&2; exit 1; }
}

(
  cd /tmp/opencode
  "$fixture/graphify.sh" update frontend
)
check test -f "$fixture/src/main/web-static/graphify-out/graph.json"
check test ! -e "$fixture/src/main/go/graphify-out/graph.json"
check grep -Fq "extract $fixture/src/main/web-static --code-only" "$GRAPHIFY_TEST_LOG"
check grep -Fq "cluster-only $fixture/src/main/web-static --no-label" "$GRAPHIFY_TEST_LOG"
check grep -Fq '"excludes":[]' "$fixture/src/main/web-static/graphify-out/.graphify_build.json"

"$fixture/graphify.sh" update all
check test -f "$fixture/graphify-out/graph.json"
check test -f "$fixture/src/main/go/graphify-out/graph.json"
check grep -Fq "update $fixture/src/main/web-static" "$GRAPHIFY_TEST_LOG"
check grep -Fq "extract $fixture --code-only --exclude src/main/web-static/ --exclude src/main/go/" "$GRAPHIFY_TEST_LOG"
check grep -Fq '"excludes":["src/main/web-static/","src/main/go/"]' "$fixture/graphify-out/.graphify_build.json"

"$fixture/graphify.sh" refresh backend
check grep -Fq "extract $fixture/src/main/go" "$GRAPHIFY_TEST_LOG"
"$fixture/graphify.sh" refresh remainder
check grep -Fq "extract $fixture --exclude src/main/web-static/ --exclude src/main/go/" "$GRAPHIFY_TEST_LOG"
if "$fixture/graphify.sh" update unknown >/dev/null 2>&1; then
  printf 'Graphify scope test failed: invalid scope accepted\n' >&2
  exit 1
fi
if GRAPHIFY_TEST_FAIL=extract "$fixture/graphify.sh" refresh remainder >/dev/null 2>&1; then
  printf 'Graphify scope test failed: semantic extraction error ignored\n' >&2
  exit 1
fi
printf '0.9.65\n' > "$fixture/.agents/skills/graphify/.graphify_version"
if "$fixture/graphify.sh" update frontend >/dev/null 2>&1; then
  printf 'Graphify scope test failed: mismatched installed version accepted\n' >&2
  exit 1
fi

printf 'Graphify scope routing: OK\n'
