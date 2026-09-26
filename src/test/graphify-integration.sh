#!/usr/bin/env bash
# Exercises real scoped extraction, single-scope updates, and deleted-file pruning.
# Run bash src/test/graphify-integration.sh with Graphify 0.9.64 installed.
# Authored by: GPT-6 Sol

set -euo pipefail

repository_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd)
fixture=$(mktemp -d /tmp/opencode/graphify-integration.XXXXXX)
trap 'rm -rf -- "$fixture"' EXIT

mkdir -p "$fixture/.agents/skills/graphify" "$fixture/src/main/web-static/websrc" "$fixture/src/main/go/domain" "$fixture/src/test/e2e"
cp "$repository_root/graphify.sh" "$fixture/graphify.sh"
cp "$repository_root/.agents/skills/graphify/.graphify_version" "$fixture/.agents/skills/graphify/.graphify_version"
printf 'export function primary(): number { return 1; }\n' > "$fixture/src/main/web-static/websrc/primary.ts"
printf 'export function removable(): number { return 2; }\n' > "$fixture/src/main/web-static/websrc/removable.ts"
printf '<main>Template marker</main>\n' > "$fixture/src/main/web-static/websrc/primary.html"
printf 'package domain\ntype Portfolio struct { Name string }\n' > "$fixture/src/main/go/domain/portfolio.go"
printf 'export function e2eFixture(): boolean { return true; }\n' > "$fixture/src/test/e2e/fixture.ts"

"$fixture/graphify.sh" update all > "$fixture/build.log"
backend_graph="$fixture/src/main/go/graphify-out/graph.json"
remainder_graph="$fixture/graphify-out/graph.json"
frontend_graph="$fixture/src/main/web-static/graphify-out/graph.json"
backend_hash=$(sha256sum "$backend_graph")
remainder_hash=$(sha256sum "$remainder_graph")

# Simulates a previously extracted semantic node without calling an external LLM.
python3 - "$frontend_graph" <<'PY'
import json
import sys
from pathlib import Path

path = Path(sys.argv[1])
graph = json.loads(path.read_text())
graph['nodes'].append({
    'id': 'template_marker',
    'label': 'Template marker',
    '_origin': 'semantic',
    'file_type': 'document',
    'source_file': 'websrc/primary.html',
})
path.write_text(json.dumps(graph))
PY

printf 'export function primary(): number { return 3; }\n' > "$fixture/src/main/web-static/websrc/primary.ts"
rm "$fixture/src/main/web-static/websrc/removable.ts"
(
  cd /tmp/opencode
  "$fixture/graphify.sh" update frontend > "$fixture/update.log"
)

[[ $(sha256sum "$backend_graph") == "$backend_hash" ]] || { printf 'Backend graph changed during frontend update\n' >&2; exit 1; }
[[ $(sha256sum "$remainder_graph") == "$remainder_hash" ]] || { printf 'Remainder graph changed during frontend update\n' >&2; exit 1; }

# Verifies that the source membership and deleted-file behavior match the scope.
python3 - "$fixture" <<'PY'
import json
import sys
from pathlib import Path

root = Path(sys.argv[1])
graphs = {
    name: json.loads((root / path / 'graphify-out/graph.json').read_text())
    for name, path in (('frontend', 'src/main/web-static'), ('backend', 'src/main/go'), ('remainder', '.'))
}
sources = {
    name: {node['source_file'] for node in graph['nodes'] if node.get('source_file')}
    for name, graph in graphs.items()
}
assert 'websrc/primary.ts' in sources['frontend'], sources['frontend']
assert 'websrc/removable.ts' not in sources['frontend'], sources['frontend']
assert 'websrc/primary.html' in sources['frontend'], 'AST update removed semantic template knowledge'
assert 'domain/portfolio.go' in sources['backend'], sources['backend']
assert 'src/test/e2e/fixture.ts' in sources['remainder'], sources['remainder']
assert not any(s.startswith(('src/main/go/', 'src/main/web-static/')) for s in sources['remainder'])
for graph in graphs.values():
    ids = {node['id'] for node in graph['nodes']}
    assert all(edge['source'] in ids and edge['target'] in ids for edge in graph['links'])
PY
printf 'Graphify scoped integration: OK\n'
