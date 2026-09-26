"""Validate the committed monorepo Graphify graphs and their source boundaries.

Run ``python3 src/test/graphify-graphs.py`` from any directory after updating graphs.
Authored by: GPT-6 Sol
"""

import json
from pathlib import Path


REPOSITORY_ROOT = Path(__file__).resolve().parents[2]
GRAPH_PATHS = {
    "frontend": REPOSITORY_ROOT / "src/main/web-static/graphify-out/graph.json",
    "backend": REPOSITORY_ROOT / "src/main/go/graphify-out/graph.json",
    "remainder": REPOSITORY_ROOT / "graphify-out/graph.json",
}


def _validate_graph(scope: str, path: Path) -> set[str]:
    """Check the selected graph's sources and referenced endpoints."""
    graph = json.loads(path.read_text(encoding="utf-8"))
    nodes = graph["nodes"]
    ids = {node["id"] for node in nodes}
    assert len(ids) == len(nodes), f"{scope}: duplicate node IDs"
    for edge in graph["links"]:
        assert edge["source"] in ids and edge["target"] in ids, f"{scope}: dangling edge {edge}"
    for hyperedge in graph.get("graph", {}).get("hyperedges", []):
        assert set(hyperedge["nodes"]).issubset(ids), f"{scope}: dangling hyperedge {hyperedge['id']}"

    sources = {
        item["source_file"]
        for collection in (nodes, graph["links"], graph.get("graph", {}).get("hyperedges", []))
        for item in collection
        if item.get("source_file")
    }
    assert all(not Path(source).is_absolute() for source in sources), f"{scope}: absolute source path"
    assert all((path.parent.parent / source).is_file() for source in sources), f"{scope}: missing source file"
    return sources


def _main() -> None:
    """Validate all three graphs and assert representative corpus coverage."""
    expected_excludes = {
        "frontend": [],
        "backend": [],
        "remainder": ["src/main/web-static/", "src/main/go/"],
    }
    for scope, path in GRAPH_PATHS.items():
        config = json.loads((path.parent / ".graphify_build.json").read_text(encoding="utf-8"))
        assert config == {"excludes": expected_excludes[scope], "gitignore": True}, f"{scope}: wrong corpus"

    sources = {scope: _validate_graph(scope, path) for scope, path in GRAPH_PATHS.items()}
    assert any(source.endswith(".html") for source in sources["frontend"]), "frontend: missing HTMX templates"
    assert any(source.endswith(".ts") for source in sources["frontend"]), "frontend: missing TypeScript"
    assert any(source.endswith(".go") for source in sources["backend"]), "backend: missing Go"
    assert any(source.endswith(".md") for source in sources["backend"]), "backend: missing documentation"
    assert any(source.startswith("src/test/e2e/") for source in sources["remainder"]), "remainder: missing E2E"
    assert any(source.startswith("src/main/docker/") for source in sources["remainder"]), "remainder: missing Docker"
    assert not any(
        source.startswith(("src/main/web-static/", "src/main/go/")) for source in sources["remainder"]
    ), "remainder: contains an excluded module"
    print("Graphify graph integrity and scope coverage: OK")


if __name__ == "__main__":
    _main()
