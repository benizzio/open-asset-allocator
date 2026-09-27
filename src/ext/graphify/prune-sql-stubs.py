"""Remove disconnected Flyway SQL parser stubs from the remainder graph.

Graphify 0.9.64 can recreate six unowned SQL table nodes during a native root
update. This repair runs after that update, before `graphify cluster-only .`.
It refuses to rewrite the graph if any other disconnected, source-less node
exists. Run with the Python interpreter that provides Graphify:

    /path/to/graphify-python src/ext/graphify/prune-sql-stubs.py

Authored by: GPT-6 Sol
"""

import json
from pathlib import Path

from networkx.readwrite import json_graph

from graphify.cluster import cluster
from graphify.export import to_json


REPOSITORY_ROOT = Path(__file__).resolve().parents[3]
GRAPH_PATH = REPOSITORY_ROOT / "graphify-out/graph.json"
SQL_STUB_PREFIX = "src_main_flyway_sql_migration_"


def _main() -> None:
    """Remove only unowned, disconnected Flyway SQL nodes and retain graph provenance."""
    original = json.loads(GRAPH_PATH.read_text(encoding="utf-8"))
    graph = json_graph.node_link_graph(original, edges="links")
    hyperedge_nodes = {
        node for hyperedge in graph.graph.get("hyperedges", []) for node in hyperedge.get("nodes", [])
    }
    disconnected = {
        node_id
        for node_id, attributes in graph.nodes(data=True)
        if not attributes.get("source_file") and graph.degree(node_id) == 0 and node_id not in hyperedge_nodes
    }
    unexpected = {
        node_id
        for node_id in disconnected
        if not node_id.startswith(SQL_STUB_PREFIX)
        or graph.nodes[node_id].get("_origin") != "ast"
        or graph.nodes[node_id].get("source_file") != ""
    }
    if unexpected:
        raise RuntimeError(f"Unrecognized disconnected graph nodes; refusing to rewrite: {sorted(unexpected)}")
    if not disconnected:
        print("No disconnected Flyway SQL stubs to remove")
        return

    graph.remove_nodes_from(disconnected)
    if not to_json(
        graph,
        cluster(graph),
        str(GRAPH_PATH),
        force=True,
        built_at_commit=original.get("built_at_commit"),
    ):
        raise RuntimeError("Graphify refused to write the cleaned remainder graph")
    print(f"Removed {len(disconnected)} disconnected Flyway SQL stubs; regenerate the report with cluster-only")


if __name__ == "__main__":
    _main()
