"""Validate the committed monorepo Graphify graphs and their source boundaries.

Run ``python3 src/test/graphify-graphs.py`` from any directory after updating graphs.
Authored by: GPT-6 Sol
"""

import json
import re
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

    connected = {
        endpoint for edge in graph["links"] for endpoint in (edge["source"], edge["target"])
    }
    connected.update(
        endpoint for hyperedge in graph.get("graph", {}).get("hyperedges", []) for endpoint in hyperedge["nodes"]
    )
    assert all(node.get("source_file") or node["id"] in connected for node in nodes), (
        f"{scope}: disconnected source-less reference"
    )

    report = (path.parent / "GRAPH_REPORT.md").read_text(encoding="utf-8")
    summary = re.search(r"(?m)^- (\d+) nodes · (\d+) edges", report)
    assert summary and (int(summary[1]), int(summary[2])) == (len(nodes), len(graph["links"])), (
        f"{scope}: report and graph counts differ"
    )
    revision = graph.get("built_at_commit")
    assert revision and f"Built from commit: `{revision[:8]}`" in report, f"{scope}: revision mismatch"

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
    migration_sql = {
        path.relative_to(REPOSITORY_ROOT).as_posix()
        for path in REPOSITORY_ROOT.glob("src/main/flyway/sql/*.sql")
    }
    assert migration_sql <= sources["remainder"], "remainder: missing Flyway SQL files"
    remainder = json.loads(GRAPH_PATHS["remainder"].read_text(encoding="utf-8"))
    for source in migration_sql:
        assert any(
            node.get("source_file") == source and node.get("label") != Path(source).name
            for node in remainder["nodes"]
        ), f"remainder: migration content missing from {source}"
    assert any(
        node.get("source_file", "").endswith("migration-[1]-database_creation.sql")
        and node.get("label") == "allocation_plan"
        for node in remainder["nodes"]
    ), "remainder: migration content was not indexed"
    assert not any(
        source.startswith(("src/main/web-static/", "src/main/go/")) for source in sources["remainder"]
    ), "remainder: contains an excluded module"
    print("Graphify graph integrity and scope coverage: OK")


if __name__ == "__main__":
    _main()
