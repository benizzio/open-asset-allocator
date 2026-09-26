#!/usr/bin/env bash
# Exercises native scoped extraction, single-scope updates, and deleted-file pruning.
# Run bash src/test/graphify-integration.sh with Graphify 0.9.64 installed.
# Authored by: GPT-6 Sol

set -euo pipefail

repository_root=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/../.." && pwd)
fixture=$(mktemp -d /tmp/opencode/graphify-integration.XXXXXX)
trap 'rm -rf -- "$fixture"' EXIT

mkdir -p "$fixture/graphify-out" "$fixture/docs" "$fixture/src/main/web-static/websrc" "$fixture/src/main/go/domain" "$fixture/src/main/go/docs" "$fixture/src/test/e2e"
mkdir -p "$fixture/src/main/web-static/graphify-out" "$fixture/src/main/go/graphify-out"
cp "$repository_root/graphify-out/.graphify_build.json" "$fixture/graphify-out/.graphify_build.json"
cp "$repository_root/src/main/web-static/graphify-out/.graphify_build.json" "$fixture/src/main/web-static/graphify-out/.graphify_build.json"
cp "$repository_root/src/main/go/graphify-out/.graphify_build.json" "$fixture/src/main/go/graphify-out/.graphify_build.json"
printf 'export function primary(): number { return 1; }\n' > "$fixture/src/main/web-static/websrc/primary.ts"
printf 'export function removable(): number { return 2; }\n' > "$fixture/src/main/web-static/websrc/removable.ts"
printf 'export function excluded(): number { return 4; }\n' > "$fixture/src/main/web-static/websrc/excluded.ts"
printf '<main>Template marker</main>\n' > "$fixture/src/main/web-static/websrc/primary.html"
printf '# Backend semantic document\n' > "$fixture/src/main/go/docs/semantic.md"
printf '# Remainder semantic document\n' > "$fixture/docs/semantic.md"
printf 'package domain\ntype Portfolio struct { Name string }\n' > "$fixture/src/main/go/domain/portfolio.go"
printf 'export function e2eFixture(): boolean { return true; }\n' > "$fixture/src/test/e2e/fixture.ts"

graphify extract "$fixture/src/main/web-static" --code-only > "$fixture/frontend-build.log"
graphify extract "$fixture/src/main/go" --code-only > "$fixture/backend-build.log"
graphify extract "$fixture" --code-only > "$fixture/remainder-build.log"
backend_graph="$fixture/src/main/go/graphify-out/graph.json"
remainder_graph="$fixture/graphify-out/graph.json"
frontend_graph="$fixture/src/main/web-static/graphify-out/graph.json"
backend_hash=$(sha256sum "$backend_graph")
remainder_hash=$(sha256sum "$remainder_graph")
frontend_config_hash=$(sha256sum "$fixture/src/main/web-static/graphify-out/.graphify_build.json")
graphify_python=$(head -n 1 "$(command -v graphify)")
graphify_python=${graphify_python#'#!'}

# The host agent may read only files selected with the same tracked excludes as native extraction.
"$graphify_python" -B - "$fixture" <<'PY'
"""Check Graphify's scoped semantic candidate enumeration.

Authored by: GPT-6 Sol
"""

import json
import sys
from pathlib import Path

from graphify.detect import detect

root = Path(sys.argv[1])
for scope, directory, expected in (
    ('frontend', root / 'src/main/web-static', 'websrc/primary.html'),
    ('backend', root / 'src/main/go', 'docs/semantic.md'),
    ('remainder', root, 'docs/semantic.md'),
):
    config = json.loads((directory / 'graphify-out/.graphify_build.json').read_text())
    corpus = detect(directory, extra_excludes=config['excludes'], gitignore=config['gitignore'])
    documents = {Path(path).relative_to(directory).as_posix() for path in corpus['files']['document']}
    assert expected in documents, f'{scope}: missing own semantic document: {documents}'
    assert documents == {expected}, f'{scope}: dispatched another scope\'s document: {documents}'
PY

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
rm "$fixture/src/main/web-static/graphify-out/manifest.json"
(
  cd /tmp/opencode
  graphify update "$fixture/src/main/web-static" > "$fixture/update.log"
)

[[ $(sha256sum "$backend_graph") == "$backend_hash" ]] || { printf 'Backend graph changed during frontend update\n' >&2; exit 1; }
[[ $(sha256sum "$remainder_graph") == "$remainder_hash" ]] || { printf 'Remainder graph changed during frontend update\n' >&2; exit 1; }
[[ $(sha256sum "$fixture/src/main/web-static/graphify-out/.graphify_build.json") == "$frontend_config_hash" ]] || { printf 'Tracked Graphify build configuration changed\n' >&2; exit 1; }

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
assert 'websrc/excluded.ts' in sources['frontend'], sources['frontend']
assert 'websrc/primary.html' in sources['frontend'], 'AST update removed semantic template knowledge'
assert 'domain/portfolio.go' in sources['backend'], sources['backend']
assert 'src/test/e2e/fixture.ts' in sources['remainder'], sources['remainder']
assert not any(s.startswith(('src/main/go/', 'src/main/web-static/')) for s in sources['remainder'])
for graph in graphs.values():
    ids = {node['id'] for node in graph['nodes']}
    assert all(edge['source'] in ids and edge['target'] in ids for edge in graph['links'])
PY

# A rename and a newly excluded file must both prune stale sources without touching the other scopes.
mv "$fixture/src/main/web-static/websrc/primary.ts" "$fixture/src/main/web-static/websrc/renamed.ts"
printf 'export function renamed(): number { return 5; }\n' > "$fixture/src/main/web-static/websrc/renamed.ts"
graphify update "$fixture/src/main/web-static" > "$fixture/rename.log"
python3 - "$fixture/src/main/web-static/graphify-out/.graphify_build.json" <<'PY'
import json
import sys
from pathlib import Path

path = Path(sys.argv[1])
config = json.loads(path.read_text())
config['excludes'].append('websrc/excluded.ts')
path.write_text(json.dumps(config))
PY
graphify update "$fixture/src/main/web-static" > "$fixture/exclude.log"
python3 - "$frontend_graph" <<'PY'
import json
import sys
from pathlib import Path

graph = json.loads(Path(sys.argv[1]).read_text())
sources = {node['source_file'] for node in graph['nodes'] if node.get('source_file')}
assert 'websrc/renamed.ts' in sources, sources
assert 'websrc/primary.ts' not in sources, sources
assert 'websrc/excluded.ts' not in sources, sources
assert 'websrc/primary.html' in sources, sources
PY
[[ $(sha256sum "$backend_graph") == "$backend_hash" ]] || { printf 'Backend graph changed during frontend pruning\n' >&2; exit 1; }
[[ $(sha256sum "$remainder_graph") == "$remainder_hash" ]] || { printf 'Remainder graph changed during frontend pruning\n' >&2; exit 1; }

# An unsupported semantic backend must fail before writing an apparently refreshed graph.
frontend_hash=$(sha256sum "$frontend_graph")
if graphify extract "$fixture/src/main/web-static" --backend invalid-backend > "$fixture/failure.log" 2>&1; then
  printf 'Semantic extraction unexpectedly accepted an unsupported backend\n' >&2
  exit 1
fi
grep -q 'unknown backend' "$fixture/failure.log" || { printf 'Semantic failure was not explained\n' >&2; exit 1; }
[[ $(sha256sum "$frontend_graph") == "$frontend_hash" ]] || { printf 'Failed extraction changed the graph\n' >&2; exit 1; }

# A whole AST extractor failure must also exit nonzero and retain the last good graph.
"$graphify_python" -B - "$fixture" <<'PY'
"""Inject an AST failure into Graphify's native extraction path.

Authored by: GPT-6 Sol
"""

import hashlib
import sys
from pathlib import Path
from unittest.mock import patch

from graphify.__main__ import main

root = Path(sys.argv[1]) / 'src/main/web-static'
graph = root / 'graphify-out/graph.json'
before = hashlib.sha256(graph.read_bytes()).digest()
sys.argv = ['graphify', 'extract', str(root), '--code-only', '--force']
with patch('graphify.extract.extract', side_effect=RuntimeError('fixture AST failure')):
    try:
        main()
    except SystemExit as error:
        assert error.code == 1, f'unexpected failure status: {error.code}'
    else:
        raise AssertionError('failed AST extraction reported success')
assert hashlib.sha256(graph.read_bytes()).digest() == before, 'failed AST extraction changed the graph'
PY

# The host agent supplies semantic JSON; Graphify merges it without losing unchanged AST or other scopes.
# Normalize source locations because an agent's "L1-L3" looks like AST provenance in Graphify 0.9.64.
printf '<main>Updated semantic marker</main>\n' > "$fixture/src/main/web-static/websrc/primary.html"
"$graphify_python" -B - "$fixture" <<'PY'
"""Exercise a changed semantic template through Graphify's native merge API.

Authored by: GPT-6 Sol
"""

import hashlib
import json
import sys
from pathlib import Path

from graphify.build import build_merge
from graphify.cluster import cluster
from graphify.export import to_json

root = Path(sys.argv[1])
frontend = root / 'src/main/web-static'
graph_path = frontend / 'graphify-out/graph.json'
source = frontend / 'websrc/primary.html'
other_graphs = [root / 'graphify-out/graph.json', root / 'src/main/go/graphify-out/graph.json']
other_hashes = [hashlib.sha256(path.read_bytes()).digest() for path in other_graphs]
before = json.loads(graph_path.read_text())
assert any(node['id'] == 'template_marker' for node in before['nodes'])
chunk = {
    'nodes': [{
        'id': 'websrc_primary_updated_semantic_marker',
        'label': 'Updated semantic marker',
        'file_type': 'document',
        'source_file': str(source),
        'source_location': None,
        '_origin': 'semantic',
    }],
    'edges': [],
    'hyperedges': [],
    'input_tokens': 0,
    'output_tokens': 0,
}
assert {Path(item['source_file']) for part in ('nodes', 'edges', 'hyperedges') for item in chunk[part]} == {source}
graph = build_merge([chunk], graph_path=graph_path, root=frontend)
assert to_json(graph, cluster(graph), str(graph_path), force=True)
after = json.loads(graph_path.read_text())
node_ids = {node['id'] for node in after['nodes']}
assert 'template_marker' not in node_ids, 'old template knowledge survived the refresh'
assert 'websrc_primary_updated_semantic_marker' in node_ids
assert any(node.get('source_file') == 'websrc/renamed.ts' for node in after['nodes'])
assert [hashlib.sha256(path.read_bytes()).digest() for path in other_graphs] == other_hashes
PY

# The SQL parser workaround must reject any unexpected orphan before modifying the graph.
"$graphify_python" -B - "$fixture" "$repository_root/src/ext/graphify/prune-sql-stubs.py" <<'PY'
"""Verify narrow, fail-closed cleanup of disconnected SQL parser stubs.

Authored by: GPT-6 Sol
"""

import importlib.util
import json
import sys
from pathlib import Path

root = Path(sys.argv[1])
path = root / 'graphify-out/graph.json'
spec = importlib.util.spec_from_file_location('graphify_sql_cleanup_fixture', sys.argv[2])
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)
module.GRAPH_PATH = path
original = json.loads(path.read_text())
stub = {
    'id': 'src_main_flyway_sql_migration_1_test_sql_unowned_table',
    'label': 'unowned_table',
    'source_file': '',
    'file_type': 'code',
    '_origin': 'ast',
}
unexpected = {'id': 'unexpected_orphan', 'label': 'unexpected', 'source_file': ''}
original['nodes'].extend((stub, unexpected))
path.write_text(json.dumps(original))
unchanged = path.read_bytes()
try:
    module._main()
except RuntimeError as error:
    assert 'unexpected_orphan' in str(error)
else:
    raise AssertionError('cleanup did not fail closed on an unknown orphan')
assert path.read_bytes() == unchanged
original['nodes'].remove(unexpected)
path.write_text(json.dumps(original))
module._main()
after = json.loads(path.read_text())
assert stub['id'] not in {node['id'] for node in after['nodes']}
assert len(after['links']) == len(original['links'])
PY
printf 'Graphify scoped integration: OK\n'
