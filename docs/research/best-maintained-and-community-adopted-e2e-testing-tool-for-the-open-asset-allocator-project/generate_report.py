#!/usr/bin/env python3
"""Generate a Markdown report from the structured E2E research results.

The converter follows the research outline for item ordering, supports flat and
nested result JSON, omits uncertain values, and retains explicitly named
uncertainties as separate report entries.

Run from any directory with ``python3 path/to/generate_report.py``.

Authored by: OpenCode
"""

from __future__ import annotations

import hashlib
import json
import re
from pathlib import Path
from typing import Any

import yaml


PROJECT_DIR = Path(__file__).resolve().parent
OUTLINE_PATH = PROJECT_DIR / "outline.yaml"
FIELDS_PATH = PROJECT_DIR / "fields.yaml"
REPORT_PATH = PROJECT_DIR / "report.md"
UNCERTAIN_MARKER = "[uncertain]"
INTERNAL_FIELDS = {"_source_file", "uncertain"}
TOC_FIELDS = ("latest_stable_release", "npm_downloads")
TOC_LABELS = {
    "latest_stable_release": "Latest Release",
    "npm_downloads": "npm Downloads",
}

CATEGORY_MAPPING = {
    "Basic Info": ["basic_info", "Basic Info"],
    "Technical Features": [
        "technical_features",
        "technical_characteristics",
        "Technical Features",
    ],
    "Performance Metrics": [
        "performance_metrics",
        "performance",
        "Performance Metrics",
    ],
    "Milestone Significance": [
        "milestone_significance",
        "milestones",
        "Milestone Significance",
    ],
    "Business Info": ["business_info", "commercial_info", "Business Info"],
    "Competition & Ecosystem": [
        "competition_ecosystem",
        "competition",
        "Competition & Ecosystem",
    ],
    "History": ["history", "History"],
    "Market Positioning": [
        "market_positioning",
        "market",
        "Market Positioning",
    ],
}


def _load_mapping(path: Path) -> dict[str, Any]:
    """Load a YAML mapping from the supplied path."""
    with path.open(encoding="utf-8") as source:
        data = yaml.safe_load(source)
    if not isinstance(data, dict):
        raise ValueError(f"Expected a mapping in {path}")
    return data


def _load_field_categories(path: Path) -> list[dict[str, Any]]:
    """Load supported fields.yaml structures into one category list."""
    data = _load_mapping(path)
    categories = data.get("field_categories")
    if isinstance(categories, list):
        return categories

    fields = data.get("fields")
    if isinstance(fields, dict):
        return [
            {"category": category, "fields": definitions}
            for category, definitions in fields.items()
        ]
    if isinstance(fields, list):
        return [{"category": "Fields", "fields": fields}]
    raise ValueError(f"No field definitions found in {path}")


def _slug_key(value: str) -> str:
    """Convert a category label to its common snake-case key."""
    return re.sub(r"[^a-z0-9]+", "_", value.lower()).strip("_")


def _category_aliases(category: str) -> list[str]:
    """Return bidirectional configured and derived aliases for a category."""
    aliases = [category, _slug_key(category)]
    for canonical, configured_aliases in CATEGORY_MAPPING.items():
        candidates = [canonical, *configured_aliases]
        normalized = {_slug_key(candidate) for candidate in candidates}
        if category in candidates or _slug_key(category) in normalized:
            aliases.extend(candidates)
            aliases.extend(normalized)
    return list(dict.fromkeys(alias for alias in aliases if alias))


def _find_nested_field(node: Any, field_name: str) -> tuple[bool, Any]:
    """Find the first matching field while traversing nested dictionaries."""
    if isinstance(node, dict):
        if field_name in node:
            return True, node[field_name]
        for value in node.values():
            found, result = _find_nested_field(value, field_name)
            if found:
                return True, result
    elif isinstance(node, list):
        for value in node:
            found, result = _find_nested_field(value, field_name)
            if found:
                return True, result
    return False, None


def _lookup_field(
    data: dict[str, Any], category: str, field_name: str
) -> tuple[bool, Any]:
    """Look up a field at top level, by category alias, then recursively."""
    if field_name in data:
        return True, data[field_name]
    for alias in _category_aliases(category):
        category_data = data.get(alias)
        if isinstance(category_data, dict) and field_name in category_data:
            return True, category_data[field_name]
    return _find_nested_field(data, field_name)


def _contains_uncertain(value: Any) -> bool:
    """Return whether a value contains the reserved uncertainty marker."""
    if isinstance(value, str):
        return UNCERTAIN_MARKER in value
    if isinstance(value, dict):
        return any(_contains_uncertain(item) for item in value.values())
    if isinstance(value, list):
        return any(_contains_uncertain(item) for item in value)
    return False


def _is_empty(value: Any) -> bool:
    """Return whether a value should be treated as absent."""
    return value is None or (isinstance(value, str) and not value.strip())


def _is_visible(field_name: str, value: Any, uncertain: set[str]) -> bool:
    """Return whether a field is safe and meaningful to render."""
    return (
        field_name not in uncertain
        and not _is_empty(value)
        and not _contains_uncertain(value)
    )


def _humanize(value: str) -> str:
    """Convert a machine field name into a readable heading."""
    return re.sub(r"[_-]+", " ", value).strip().title()


def _anchor(value: str) -> str:
    """Create a stable ASCII Markdown anchor component."""
    anchor = re.sub(r"[^a-z0-9]+", "-", value.lower()).strip("-")
    return anchor or "item"


def _escape_inline(value: str) -> str:
    """Escape Markdown table-of-contents separators and flatten whitespace."""
    return re.sub(r"\s+", " ", value).strip().replace("|", "\\|")


def _compact_summary(value: Any, limit: int = 180) -> str:
    """Produce a concise first-sentence summary for a TOC entry."""
    if isinstance(value, (dict, list)):
        text = json.dumps(value, ensure_ascii=False, separators=(", ", ": "))
    else:
        text = str(value)
    text = _escape_inline(text)
    first_sentence = re.split(r"(?<=[.!?])\s+(?=[A-Z\[])", text, maxsplit=1)[0]
    first_sentence = re.sub(r"\s*\[S\d+(?:[-,]\s*S?\d+)*\]\s*$", "", first_sentence)
    if len(first_sentence) > limit:
        first_sentence = first_sentence[: limit - 3].rstrip() + "..."
    return first_sentence


def _long_text(value: str) -> list[str]:
    """Format long prose as a readable Markdown block quote."""
    flattened = re.sub(r"\s+", " ", value).strip()
    with_breaks = re.sub(
        r"(?<=[.!?])\s+(?=[A-Z\[])",
        "<br>\n> ",
        flattened,
    )
    return [f"> {with_breaks}"]


def _inline_value(value: Any) -> str:
    """Format a value for one compact Markdown line."""
    if isinstance(value, bool):
        return "Yes" if value else "No"
    if value is None:
        return ""
    if isinstance(value, dict):
        parts = [
            f"{_humanize(str(key))}: {_inline_value(item)}"
            for key, item in value.items()
        ]
        return "; ".join(parts)
    if isinstance(value, list):
        return ", ".join(_inline_value(item) for item in value)
    return _escape_inline(str(value))


def _format_value(value: Any) -> list[str]:
    """Format scalar and complex values as readable Markdown lines."""
    if isinstance(value, list):
        if not value:
            return []
        if all(isinstance(item, dict) for item in value):
            return [
                "- "
                + " | ".join(
                    f"{_humanize(str(key))}: {_inline_value(item_value)}"
                    for key, item_value in item.items()
                )
                for item in value
            ]
        compact = [_inline_value(item) for item in value]
        if len(compact) <= 3 and sum(map(len, compact)) <= 180:
            return [", ".join(compact)]
        return [f"- {item}" for item in compact]

    if isinstance(value, dict):
        lines = []
        for key, item in value.items():
            if isinstance(item, (dict, list)):
                lines.append(f"- **{_humanize(str(key))}:**")
                lines.extend(f"  {line}" for line in _format_value(item))
            else:
                lines.append(
                    f"- **{_humanize(str(key))}:** {_inline_value(item)}"
                )
        return lines

    text = _inline_value(value)
    return _long_text(text) if len(text) > 100 else [text]


def _base_result_name(item_name: str) -> str:
    """Derive the authoritative base result filename from an item name."""
    slug = re.sub(r"\s+", "_", item_name)
    slug = re.sub(r"[^A-Za-z0-9_-]", "", slug).strip("_-") or "item"
    return f"{slug}.json"


def _build_output_mapping(
    items: list[dict[str, Any]], output_dir: Path
) -> dict[str, Path]:
    """Build collision-safe output paths using the deep-research rules."""
    names = [str(item["name"]) for item in items]
    if len(names) != len(set(names)):
        raise ValueError("The outline contains duplicate item names")

    digest_lengths = {name: 0 for name in names}
    while True:
        filenames = {}
        grouped: dict[str, list[str]] = {}
        for name in names:
            base = Path(_base_result_name(name)).stem
            digest_length = digest_lengths[name]
            if digest_length:
                digest = hashlib.sha256(name.encode("utf-8")).hexdigest()
                filename = f"{base}_{digest[:digest_length]}.json"
            else:
                filename = f"{base}.json"
            filenames[name] = filename
            grouped.setdefault(filename.casefold(), []).append(name)

        collisions = [group for group in grouped.values() if len(group) > 1]
        if not collisions:
            break
        for group in collisions:
            for name in group:
                digest_lengths[name] = max(8, digest_lengths[name] + 1)
                if digest_lengths[name] > 64:
                    raise ValueError(f"Could not disambiguate result for {name}")

    mapping = {name: (output_dir / filenames[name]).resolve() for name in names}
    if any(path.parent != output_dir for path in mapping.values()):
        raise ValueError("Mapped result paths must be direct output children")
    if len({str(path).casefold() for path in mapping.values()}) != len(mapping):
        raise ValueError("Mapped result paths are not unique")
    return mapping


def _contains_defined_field(node: Any, defined_fields: set[str]) -> bool:
    """Return whether a nested object contains a defined report field."""
    if not isinstance(node, dict):
        return False
    return any(
        key in defined_fields or _contains_defined_field(value, defined_fields)
        for key, value in node.items()
    )


def _collect_extra_fields(
    data: dict[str, Any],
    defined_fields: set[str],
    category_keys: set[str],
) -> dict[str, Any]:
    """Collect non-schema JSON values without duplicating category containers."""
    extras: dict[str, Any] = {}

    def visit(node: dict[str, Any], path: tuple[str, ...]) -> None:
        for key, value in node.items():
            if key in INTERNAL_FIELDS or key in defined_fields:
                continue
            if key in category_keys and isinstance(value, dict):
                visit(value, (*path, key))
                continue
            if isinstance(value, dict) and _contains_defined_field(
                value, defined_fields
            ):
                visit(value, (*path, key))
                continue
            extra_name = ".".join((*path, key)) if path else key
            extras[extra_name] = value

    visit(data, ())
    return extras


def _unique_anchors(names: list[str]) -> dict[str, str]:
    """Create unique anchors while preserving item order."""
    used: set[str] = set()
    anchors = {}
    for name in names:
        base = _anchor(name)
        candidate = base
        suffix = 2
        while candidate in used:
            candidate = f"{base}-{suffix}"
            suffix += 1
        used.add(candidate)
        anchors[name] = candidate
    return anchors


def _render_field(field_name: str, value: Any) -> list[str]:
    """Render one detailed field section."""
    lines = [f"#### {_humanize(field_name)}", ""]
    lines.extend(_format_value(value))
    lines.append("")
    return lines


def _validate_topic_dir(outline: dict[str, Any]) -> None:
    """Validate the persisted canonical directory against its parent path."""
    topic_dir = str(outline.get("topic_dir", ""))
    if not re.fullmatch(r"[a-z0-9]+(?:-[a-z0-9]+)*", topic_dir):
        raise ValueError(f"Invalid topic_dir: {topic_dir}")
    if topic_dir != PROJECT_DIR.name:
        raise ValueError("topic_dir does not match the outline parent directory")


def main() -> None:
    """Build report.md from the configured research results.

    The command reads ``outline.yaml``, ``fields.yaml``, and every mapped JSON
    result beneath the configured output directory. It writes ``report.md`` in
    the canonical research project directory.

    Example:
        ``python3 generate_report.py``

    Authored by: OpenCode
    """
    outline = _load_mapping(OUTLINE_PATH)
    _validate_topic_dir(outline)
    categories = _load_field_categories(FIELDS_PATH)
    items = outline.get("items")
    if not isinstance(items, list):
        raise ValueError("outline.yaml items must be a list")

    output_setting = Path(outline.get("execution", {}).get("output_dir", "./results"))
    output_dir = (
        output_setting.resolve()
        if output_setting.is_absolute()
        else (PROJECT_DIR / output_setting).resolve()
    )
    mapping = _build_output_mapping(items, output_dir)

    records: list[tuple[str, dict[str, Any], Path]] = []
    mapped_paths = set(mapping.values())
    for item in items:
        name = str(item["name"])
        result_path = mapping[name]
        if not result_path.is_file():
            raise FileNotFoundError(f"Missing mapped result: {result_path}")
        with result_path.open(encoding="utf-8") as source:
            records.append((name, json.load(source), result_path))

    for result_path in sorted(output_dir.glob("*.json")):
        resolved_path = result_path.resolve()
        if resolved_path in mapped_paths:
            continue
        with result_path.open(encoding="utf-8") as source:
            data = json.load(source)
        fallback_name = data.get("name")
        name = (
            str(fallback_name)
            if isinstance(fallback_name, (str, int, float))
            else result_path.stem.replace("_", " ")
        )
        records.append((name, data, resolved_path))

    field_to_category = {
        str(field["name"]): str(category.get("category", "Fields"))
        for category in categories
        for field in category.get("fields", [])
    }
    defined_fields = set(field_to_category)
    category_keys = {
        alias
        for category in categories
        for alias in _category_aliases(str(category.get("category", "Fields")))
    }
    names = [name for name, _, _ in records]
    anchors = _unique_anchors(names)

    lines = [
        f"# {outline.get('topic', 'Research Report')}",
        "",
        f"Generated from {len(records)} structured research results.",
        "",
        "## Table of Contents",
        "",
    ]

    for index, (name, data, _) in enumerate(records, 1):
        uncertain = set(map(str, data.get("uncertain", [])))
        summaries = []
        for field_name in TOC_FIELDS:
            category = field_to_category.get(field_name, "")
            found, value = _lookup_field(data, category, field_name)
            if found and _is_visible(field_name, value, uncertain):
                label = TOC_LABELS.get(field_name, _humanize(field_name))
                summaries.append(f"{label}: {_compact_summary(value)}")
        suffix = f" - {' | '.join(summaries)}" if summaries else ""
        lines.append(f"{index}. [{name}](#{anchors[name]}){suffix}")

    lines.extend(["", "## Detailed Results", ""])
    for index, (name, data, result_path) in enumerate(records, 1):
        uncertain_names = [str(value) for value in data.get("uncertain", [])]
        uncertain = set(uncertain_names)
        lines.extend(
            [
                f'<a id="{anchors[name]}"></a>',
                f"## {index}. {name}",
                "",
                f"Source result: `{result_path.name}`",
                "",
            ]
        )

        for category in categories:
            category_name = str(category.get("category", "Fields"))
            rendered_fields = []
            for field in category.get("fields", []):
                field_name = str(field["name"])
                found, value = _lookup_field(data, category_name, field_name)
                if found and _is_visible(field_name, value, uncertain):
                    rendered_fields.extend(_render_field(field_name, value))
            if rendered_fields:
                lines.extend([f"### {_humanize(category_name)}", ""])
                lines.extend(rendered_fields)

        extras = _collect_extra_fields(data, defined_fields, category_keys)
        visible_extras = {
            key: value
            for key, value in extras.items()
            if _is_visible(key, value, uncertain)
        }
        if visible_extras:
            lines.extend(["### Other Info", ""])
            for field_name, value in visible_extras.items():
                lines.extend(_render_field(field_name, value))

        if uncertain_names:
            lines.extend(["### Uncertain Fields", ""])
            lines.extend(f"- `{field_name}`" for field_name in uncertain_names)
            lines.append("")

    REPORT_PATH.write_text("\n".join(lines).rstrip() + "\n", encoding="utf-8")
    print(f"Generated {REPORT_PATH} from {len(records)} results")


if __name__ == "__main__":
    main()
