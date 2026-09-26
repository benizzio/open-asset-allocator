---
description: Run the Graphify skill with the supplied arguments
---

Load and follow the `graphify` skill with exactly the argument text below. This repository has three scoped graphs,
not the skill's default single root-wide graph. Before a query, path, or explain, select the graph from `AGENTS.md`
and pass `--graph` explicitly. For an update or rebuild, follow the scope-aware commands in `AGENTS.md` instead of
the skill's default `graphify update .` or root-wide extraction. If the arguments do not identify a scope, infer it
from the requested source area; ask only if that is genuinely ambiguous. Preserve all other supplied arguments.

<graphify-arguments>
$ARGUMENTS
</graphify-arguments>

<!-- Co-authored by: GPT-6 Sol and Igor Benicio de Mesquita -->
