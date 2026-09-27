# Graphify scopes

<!-- Co-authored by: GPT-6 Sol and GPT-6 Luna -->

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
