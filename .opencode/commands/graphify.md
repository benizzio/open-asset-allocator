---
description: Run the Graphify skill with the supplied arguments
---

Load the `graphify` skill, then apply the repository's `AGENTS.md` scope table before choosing a graph or scan root. Keep the arguments below intact. For a question about frontend or backend code, follow the skill's query workflow using the corresponding graph with an explicit `--graph`; never use its default root graph as an all-repository fallback. For mixed questions, query each relevant graph separately and verify boundaries in source.

For an update, extraction, or bare invocation without an identifiable source area, ask which of frontend, backend, or remainder is intended before running anything that writes graph artifacts. The generated skill's default full pipeline and `--update` on the repository root do not honor the remainder graph's tracked exclusions; use the repository's native per-scope instructions instead. Keep the generated skill untouched. Honor an explicit help request as documented by the skill.

<graphify-arguments>
$ARGUMENTS
</graphify-arguments>

<!-- Co-authored by: GPT-5.6 Sol and GPT-6 Sol -->
