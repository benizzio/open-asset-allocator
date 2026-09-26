# Graphify scopes

<!-- Authored by: GPT-6 Sol -->

The graph locations and routing rules are in [the root agent instructions](../AGENTS.md). The root `graphify-out/`
contains **only** the remainder of the repository. The frontend and backend have their own committed graphs and
reports. The graphs were built with Graphify 0.9.64; the vendored skill records its version in
`.agents/skills/graphify/.graphify_version`.

## Maintenance

From the repository root, run the native Graphify CLI for each affected scope:

```sh
graphify update "$PWD/src/main/web-static"
graphify update "$PWD/src/main/go"
graphify update .
```

These operations use local AST extraction. The three committed `.graphify_build.json` files are Graphify's own
persisted build configurations; the root configuration excludes both module trees. Do not put these exclusions in a
root `.graphifyignore`, which would also affect module scans. On a fresh checkout with committed graphs but no
manifest, Graphify reconstructs local incremental state from the graphs. A native update of the root does **not**
update the frontend or backend. For Graphify 0.9.64, pass **absolute** module scan paths: a repository-relative
module argument was observed to prefix the graph's source paths, and a later update pruned semantic nodes. The
committed graphs use module-relative source paths; the integration test uses absolute scan paths.

Graphify 0.9.64 with its SQL parser installed can add six disconnected Flyway SQL stubs after **each** remainder
update. It has no native cleanup command. For a root update, run the narrowly scoped repair with Graphify's own
Python interpreter, regenerate the report, then validate:

```sh
graphify_python=$(head -n 1 "$(command -v graphify)")
graphify_python=${graphify_python#'#!'}
"$graphify_python" -B src/ext/graphify/prune-sql-stubs.py
graphify cluster-only "$PWD" --no-label
python3 src/test/graphify-graphs.py
```

The repair refuses to write if it finds any disconnected source-less node other than an unowned Flyway SQL AST
stub. This is a parser workaround, not a general scope/update wrapper. The current graph was generated with
Graphify 0.9.64; do not apply the workaround to an upgraded parser without checking whether it is still needed.

If a graph is absent, build its code corpus with `graphify extract <absolute-scope-path> --code-only`, then run
`graphify cluster-only <absolute-scope-path> --no-label` to generate its report and visualization. Use `$PWD` for the
remainder path, or `$PWD/src/main/web-static` and `$PWD/src/main/go` for the module paths. Without the committed graph,
a code-only build has no pre-existing semantic layer; refresh semantic sources separately.

If HTML/HTMX templates, documents, or images changed, an AST update is insufficient. With a Graphify-supported
semantic backend configured by the operator, run `graphify extract <absolute-scope-path>`, then
`graphify cluster-only <absolute-scope-path> --no-label`. Graphify uses the tracked build configuration for that scope and
fails explicitly if semantic work requires a missing backend. OpenCode can instead act as the host agent for the
generated skill's semantic extraction. Start that skill inside a frontend or backend module so its paths and output
stay in the module. For the remainder, enumerate candidate files with `graphify.detect.detect(root,
extra_excludes=config["excludes"], gitignore=config["gitignore"])` using the tracked root configuration, and
check that no module file is sent to an agent. Review each chunk against the skill's extraction schema, make its
`_origin` explicitly `semantic` and its `source_location` null, then merge it with Graphify's `build_merge(...,
root=<absolute-scope-path>)` and regenerate outputs with `graphify cluster-only <absolute-scope-path> --no-label`.
In Graphify 0.9.64, an agent's `L1-L3` location without an origin is mistaken for an AST node, so stale semantic
nodes survive a refresh. The disposable fixture exercised both the failure and the normalized replacement across
all three scopes. **Do not run the bundled skill's default full pipeline on the repository root:** its detection
does not apply the remainder-only exclusions from `.graphify_build.json`. The native CLI has no OpenCode backend;
an OpenCode-assisted merge uses the Graphify library, not `graphify extract --backend opencode`.

No Git hooks or CI jobs are installed for Graphify. A native root hook would only refresh the remainder graph; if
automation is added later, it must run updates for the affected module paths too.

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

The first scoped remainder build lacked Graphify's optional `tree_sitter_sql` parser, so 25 `.sql` files contributed
no nodes. On 2026-09-26 the `graphifyy[sql]==0.9.64` extra was installed locally and the remainder was rebuilt with
`graphify extract "$PWD" --code-only --force --no-label`. All 25 SQL files now have attributed file nodes. For
example, `migration-[1]-database_creation.sql` contributes the `allocation_plan` table at line 102 and related
tables/edges. The SQL AST parser left 7 of 14 Flyway migrations with only a file node. OpenCode-assisted extraction
then added reviewed, source-attributed concepts and relationships for those seven migrations. Every Flyway migration
now has at least one content node, but this is **partial statement coverage**, not a claim that every statement was
parsed or that host-agent extraction is a SQL parser. A fresh checkout can query the committed SQL nodes but needs a
compatible SQL extra to rebuild or refresh AST SQL content. An AST rebuild preserved existing semantic nodes; the 75
unconnected, source-less references it also retained/generated (including six SQL stubs) were removed as a
graph cleanup. Six SQL stubs recur on later root updates; the narrow repair above removes them. The integrity
check rejects all disconnected source-less nodes, including new ones not covered by the repair.

## Validation

```sh
bash src/test/graphify-integration.sh
python3 src/test/graphify-graphs.py
```

The integration test runs native Graphify on a disposable small repository with copies of the tracked build
configurations. It changes, removes, renames, and excludes frontend sources; checks a visible native failure; and
checks that a template's old semantic node can be replaced without changing unrelated graphs. A separate
OpenCode host-agent fixture extracted and refreshed HTML and Markdown in all three scopes. Its first backend
chunk used AST-shaped source locations and failed to replace the old nodes; normalizing provenance and repeating
the merge resolved that. The committed-graph check verifies scope membership, path portability, endpoints,
hyperedges, report metadata, connected source-less references, and representative template/Go/E2E/Docker and
Flyway SQL coverage. A migration file node alone is **not** counted as content coverage; supplemental host-agent
concepts are labeled as semantic rather than parsed SQL schema.

OpenCode-assisted extraction uses host-agent tokens, but the generated chunks do not report reliable token usage.
Graphify's report records zero CLI extraction tokens for that work; it is not a measurement of host-agent cost.

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

The expanded comparison on the same date uses Graphify 0.9.64, base source revision `1a2fda2d`, BFS depth 2, and
the same 1,200-token budget. The remainder rows below were repeated after the SQL and host-agent semantic refresh.
The revision stamp is the Git commit at extraction time, not proof that uncommitted host-agent changes were absent:

| Exact query and graph | Observed result, expected evidence, and follow-up |
| --- | --- |
| `asset management` / frontend | 237 reached; 41 displayed, including frontend `Asset` and asset pages; truncated. No backend/remainder sources displayed. |
| `routing` / frontend | 194 reached; 35 displayed, including `routing-navigo.ts` and routing bindings; truncated. No unrelated-module sources displayed. |
| `portfolio allocation` / backend | 261 reached; 38 displayed, including `portfolio_allocation_controller.go`; truncated. No frontend/remainder sources displayed. |
| `PortfolioAllocationRESTController` / backend | 62 reached; 35 displayed, including the expected controller at `api/rest/portfolio_allocation_controller.go:L20`; truncated. A path-qualified `explain` resolved its ten direct connections. |
| `flyway migrations` / remainder | 27 reached; all displayed, including Flyway instructions and semantic coverage guidance. No specific SQL table is displayed. Output exceeds the requested budget because Graphify keeps all edges when all nodes fit. Follow up with a migration symbol/path. |
| `migration allocation_plan` / remainder | 33 reached; 19 displayed, including migration 5/6 constraint changes and migration 12's rename, alongside unrelated E2E/UI documentation; truncated before migration 1's `allocation_plan`. A path-qualified `explain` resolves that table at line 102 and its references. |
| `e2e fixtures` / remainder | 159 reached; 41 displayed, including E2E fixtures and specs; truncated. No frontend/backend sources displayed. |
| `graphify build configuration` / remainder | 30 reached; all displayed, including scoped-configuration and CodeRabbit inclusion concepts, but **no** `.graphify_build.json` file node. Output exceeds the budget. Read the three tracked dotfiles directly; the graph query does not show their actual contents. |
| `asset.external_data` / remainder | 134 reached; 30 displayed, starting with the host-agent concept `asset.external_data JSONB column added` from migration 14; truncated with unrelated E2E nodes. A path-qualified lookup is still needed for precise evidence. |

The three scoped graphs prevent attributed cross-module sources from appearing in the displayed results. They do
not guarantee that a broad query fits the budget or finds every relevant file. The configuration query remains a
file-level retrieval miss, even though the semantic guide explains where to look. Seven Flyway migrations require
the host-agent supplement because the SQL AST extractor contributes only file nodes for them.

For a verified cross-module boundary, `src/main/web-static/websrc/pages/asset.html` line 7 declares
`hx-get="/api/asset"`; `src/main/go/api/rest/asset_controller.go` lines 29-33 register `/api/asset` on the
backend. This conclusion comes from the source files, **not** an inferred link between separate graphs. Use
`--graph` explicitly for queries and path lookups. A merged graph is not required for this workflow.
