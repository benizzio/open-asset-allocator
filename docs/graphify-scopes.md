# Graphify scopes

<!-- Authored by: GPT-6 Sol -->

The graph locations and routing rules are in [the root agent instructions](../AGENTS.md). The root `graphify-out/`
contains **only** the remainder of the repository. The frontend and backend have their own committed graphs and
reports. Graphify 0.9.64 is required; the expected version is stored in
`.agents/skills/graphify/.graphify_version` alongside the vendored skill.

## Maintenance

From any working directory, call the repository script by path:

```sh
./graphify.sh update frontend
./graphify.sh update backend
./graphify.sh update remainder
./graphify.sh update all
```

The `update` operation uses local AST extraction. If a graph is absent, it builds a code-only graph, then generates
its report and visualization. On a fresh checkout with committed graphs but no manifest, Graphify reconstructs the
local incremental state from the existing graph. The script recreates each scope's ignored build configuration on
every invocation; the remainder's exclusion rules cannot leak into module builds.

If HTML/HTMX templates, documents, or images changed, an AST update is insufficient. With a Graphify-supported
semantic backend configured by the operator, use `./graphify.sh refresh <scope>`; this performs native incremental
semantic extraction and regenerates the report. A missing backend causes an explicit failure rather than a misleading
successful refresh. The initial code-only build on a fresh repository without committed graphs needs a semantic
refresh to reach the committed graphs' coverage. Agent-assisted semantic extraction can run from within a module
using its module path and output directory. **Do not run the bundled skill's default full pipeline on the repository
root:** its default detection does not apply the remainder-only exclusions.

The script does not install Git hooks or CI jobs. If one is installed later, it must invoke the same scope-aware
entry point. Native Graphify hooks and direct `graphify update .` do not implement the three-scope routing.

## Initial migration and remaining coverage

On 2026-09-26 the frontend and backend code graphs were extracted independently from their scan roots. The 27
frontend and 4 backend sourced semantic nodes from the prior root-wide graph were retained with their source paths
made relative to their scan roots. Edges and hyperedges were retained only if every endpoint belonged to the scoped
graph. Source-less dependency references were produced by the new module extraction rather than indiscriminately
assigned to the remainder. The remainder was updated from the prior root-wide graph with both module trees
excluded; its existing documentation and image nodes were preserved.

This retains previously extracted template/document knowledge without assuming that an AST-only build can extract
it. The old graph's recorded source revision was `d5aa03fe8a048bd3bf26f0b14c752186d8d1674c`; migrated semantic
content should be refreshed after its source changes. Generated graph `built_at_commit` values describe the most
recent build, **not** independent freshness of every preserved semantic node.

This installation does **not** have Graphify's optional `tree_sitter_sql` parser. During the remainder update,
Graphify reported that 25 `.sql` files contributed no nodes. Flyway configuration and instructions are queryable,
but migration SQL contents are not. Install a compatible Graphify SQL extra through the project's normal dependency
management process, then rebuild the remainder and validate SQL coverage. Do not mistake an existing Flyway
documentation node for a parsed migration.

## Validation

```sh
bash src/test/graphify-scopes.sh
bash src/test/graphify-integration.sh
python3 src/test/graphify-graphs.py
```

The routing test uses a fake Graphify executable and checks clean-checkout configuration, working-directory
independence, exclusions, and failure propagation. The integration test runs real Graphify on a disposable small
repository, changes and removes a frontend source, and checks that unrelated graph files stay byte-identical.
The committed-graph check verifies scope membership, path portability, endpoints, hyperedges, and representative
frontend template/backend/E2E/Docker sources. SQL content is deliberately **not** counted as covered.

## Retrieval comparison

The following queries were run on the independently rebuilt graphs with BFS depth 2 and a 1,200-token budget.
Results are for the 2026-09-26 checkout with Graphify 0.9.64. Repeating them after a rebuild is more useful than
assuming a reduction in node counts:

| Query and chosen graph | Observed result |
| --- | --- |
| `asset management` / frontend | 237 nodes reached; the displayed nodes start with frontend `Asset`, `allocation-plan-management.ts`, and frontend asset pages. Output truncated. |
| `routing` / frontend | 194 nodes reached and output truncated. Module isolation did not reduce within-module breadth. |
| `portfolio allocation` / backend | 261 nodes reached; the displayed nodes start with backend portfolio-allocation files and types. Output truncated. |
| `flyway migrations` / remainder | 19 nodes reached, including Flyway service/instructions and E2E database references. No parsed SQL migration content. |
| `e2e fixtures` / remainder | 159 nodes reached; the displayed nodes include E2E fixtures, specs, and orchestration. Output truncated. |

For a verified cross-module boundary, `src/main/web-static/websrc/pages/asset.html` line 7 declares
`hx-get="/api/asset"`; `src/main/go/api/rest/asset_controller.go` lines 29-33 register `/api/asset` on the
backend. This conclusion comes from the source files, **not** an inferred link between separate graphs. Use
`--graph` explicitly for queries and path lookups. A merged graph is not required for this workflow.
