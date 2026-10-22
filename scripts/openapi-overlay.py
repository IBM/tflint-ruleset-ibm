#!/usr/bin/env python3
# Copyright IBM Corp. 2026 All Rights Reserved.
# Licensed under the Mozilla Public License v2.0
"""Match IBM Cloud OpenAPI constraints onto Terraform attributes.

Writes extractor/output/<pkg>.overlay.json (and overlay-index.json). Specs stay in
the private cache from apidocs_cache.py and are never committed; each one must
match the SHA-256 recorded by fetch-ibm-apidocs.py.
"""

from __future__ import annotations

import ast
import json
import re
import sys
from pathlib import Path

from apidocs_cache import raw_dir, sha256_file

RAW = raw_dir()
ROOT = Path(__file__).resolve().parent.parent
EXTRACT = ROOT / "extractor" / "output"
FETCH_MAP = Path(__file__).resolve().parent / "fetch-ibm-apidocs.py"

FORMAT_KIND = {
    "ipv4": "ipv4",
    "ipv6": "ipv4",
    "cidr": "cidr",
    "cidrv4": "cidr",
    "email": "email",
    "uuid": "uuid",
    "uri": "uri",
    "url": "uri",
    "hostname": "hostname",
    "date": "date",
    "date-time": "datetime",
    "byte": "",
    "binary": "",
    "int32": "",
    "int64": "",
    "float": "",
    "double": "",
    "password": "",
    "crn": "crn",
}

NOISE_SCHEMA = re.compile(
    r"(error|exception|paginated|pagination|listresponse|collectionresponse|"
    r"identitybyhref|identitybyid|identitybycrn|reference|deleted|href)$",
    re.I,
)
PREFER_SCHEMA = re.compile(
    r"(prototype|patch|request|body|create|update|template|options|params)$",
    re.I,
)


def slug_of(href: str) -> str:
    href = href.strip("/")
    href = re.sub(r"^docs/apis/", "", href)
    return href.replace("/", "-")


def read_package_hrefs() -> dict[str, list[str]]:
    """Read PACKAGE_HREFS from the fetch script as data, without running it."""
    tree = ast.parse(FETCH_MAP.read_text(), str(FETCH_MAP))
    for node in tree.body:
        target = None
        if isinstance(node, ast.AnnAssign):
            target = node.target
        elif isinstance(node, ast.Assign) and len(node.targets) == 1:
            target = node.targets[0]
        if isinstance(target, ast.Name) and target.id == "PACKAGE_HREFS" and node.value is not None:
            return ast.literal_eval(node.value)
    return {}


def load_manifest() -> dict:
    man = RAW / "manifest.json"
    if not man.is_file():
        return {}
    return json.loads(man.read_text())


MANIFEST = load_manifest()
DIGESTS = {
    d["slug"]: d.get("sha256", "")
    for d in MANIFEST.get("downloaded") or []
    if d.get("ok")
}


def load_package_hrefs() -> dict[str, list[str]]:
    mapping: dict[str, list[str]] = {}
    for pkg, hrefs in read_package_hrefs().items():
        mapping[pkg] = [slug_of(h) for h in hrefs]
    if MANIFEST:
        data = MANIFEST
        for pkg, slugs in (data.get("packages") or {}).items():
            mapping.setdefault(pkg, [])
            for s in slugs:
                if s not in mapping[pkg]:
                    mapping[pkg].append(s)
    return mapping


def load_spec(slug: str) -> dict | None:
    path = RAW / f"{slug}.json"
    if not path.is_file():
        return None
    want = DIGESTS.get(slug)
    if not want:
        print(f"skipping {slug}: not recorded in manifest.json", file=sys.stderr)
        return None
    if sha256_file(path) != want:
        raise SystemExit(f"{path} does not match its manifest SHA-256; re-run fetch-ibm-apidocs.py")
    try:
        data = json.loads(path.read_text())
    except json.JSONDecodeError:
        return None
    return data if isinstance(data, dict) else None


class SpecIndex:
    def __init__(self, spec: dict, slug: str):
        self.slug = slug
        self.spec = spec
        self.schemas: dict[str, dict] = {}
        components = spec.get("components") or {}
        raw = components.get("schemas") or spec.get("definitions") or {}
        if isinstance(raw, dict):
            self.schemas = raw
        self.title = ((spec.get("info") or {}).get("title") or slug).strip()

    def resolve(self, node: dict | None, depth: int = 0) -> dict:
        if not isinstance(node, dict) or depth > 8:
            return node or {}
        ref = node.get("$ref")
        if isinstance(ref, str):
            name = ref.rsplit("/", 1)[-1]
            target = self.schemas.get(name)
            if isinstance(target, dict):
                merged = dict(target)
                for k, v in node.items():
                    if k != "$ref" and k not in merged:
                        merged[k] = v
                return self.resolve(merged, depth + 1)
        for key in ("allOf", "oneOf", "anyOf"):
            parts = node.get(key)
            if isinstance(parts, list) and parts:
                acc: dict = {}
                props: dict = {}
                required: list[str] = []
                for p in parts:
                    r = self.resolve(p if isinstance(p, dict) else {}, depth + 1)
                    if not isinstance(r, dict):
                        continue
                    for pk, pv in (r.get("properties") or {}).items():
                        props[pk] = pv
                    for req in r.get("required") or []:
                        if req not in required:
                            required.append(req)
                    for k, v in r.items():
                        if k in ("properties", "required"):
                            continue
                        acc.setdefault(k, v)
                if props:
                    acc["properties"] = props
                if required:
                    acc["required"] = required
                for k, v in node.items():
                    if k not in ("allOf", "oneOf", "anyOf", "$ref"):
                        acc[k] = v
                return acc
        return node

    def properties_of(self, schema: dict | None) -> dict[str, dict]:
        resolved = self.resolve(schema)
        props = resolved.get("properties") if isinstance(resolved, dict) else None
        if not isinstance(props, dict):
            return {}
        out: dict[str, dict] = {}
        for name, node in props.items():
            if isinstance(node, dict):
                out[name] = self.resolve(node)
        return out

    def constraint(self, node: dict) -> dict:
        node = self.resolve(node)
        items = node.get("items") if isinstance(node.get("items"), dict) else {}
        src = node
        if (node.get("type") == "array" or items) and not node.get("enum") and not node.get("pattern"):
            if items:
                src = {**self.resolve(items), "minItems": node.get("minItems"), "maxItems": node.get("maxItems")}
        enum = src.get("enum")
        if not isinstance(enum, list):
            enum = None
        else:
            enum = [str(x) for x in enum if x is not None]
        pattern = src.get("pattern") if isinstance(src.get("pattern"), str) else None
        fmt = (src.get("format") or "").lower()
        kind = FORMAT_KIND.get(fmt, "")
        min_len = src.get("minLength")
        max_len = src.get("maxLength")
        if pattern and re.fullmatch(r"\^?\.(\+|\*)\$?|\(\.\*\??\)|\.\+", pattern):
            if min_len is None:
                min_len = 1
            pattern = None
        minimum = src.get("minimum")
        maximum = src.get("maximum")
        min_items = node.get("minItems") if node.get("minItems") is not None else src.get("minItems")
        max_items = node.get("maxItems") if node.get("maxItems") is not None else src.get("maxItems")
        desc = (src.get("description") or node.get("description") or "").strip()
        if desc:
            desc = re.sub(r"\s+", " ", desc)[:400]
        out = {
            "enum": enum,
            "pattern": pattern,
            "format": fmt or None,
            "kind": kind or None,
            "min_len": int(min_len) if isinstance(min_len, (int, float)) else None,
            "max_len": int(max_len) if isinstance(max_len, (int, float)) else None,
            "min": int(minimum) if isinstance(minimum, (int, float)) else None,
            "max": int(maximum) if isinstance(maximum, (int, float)) else None,
            "min_items": int(min_items) if isinstance(min_items, (int, float)) else None,
            "max_items": int(max_items) if isinstance(max_items, (int, float)) else None,
            "description": desc or None,
            "type": src.get("type") or node.get("type"),
        }
        return {k: v for k, v in out.items() if v is not None}

    def flatten_props(self, schema: dict, prefix: str = "", depth: int = 0) -> dict[str, dict]:
        if depth > 5:
            return {}
        out: dict[str, dict] = {}
        props = self.properties_of(schema)
        for name, node in props.items():
            path = f"{prefix}{name}"
            c = self.constraint(node)
            if name in set((self.resolve(schema) or {}).get("required") or []):
                c["required"] = True
            out[path] = c
            nested = None
            if isinstance(node.get("properties"), dict) or node.get("allOf") or node.get("$ref"):
                nested = node
            items = node.get("items")
            if isinstance(items, dict):
                items = self.resolve(items)
                if items.get("properties") or items.get("$ref") or items.get("allOf"):
                    nested = items
            if nested:
                out.update(self.flatten_props(nested, path + ".", depth + 1))
        return out

    def request_schemas(self) -> list[tuple[str, dict]]:
        found: list[tuple[str, dict]] = []
        seen: set[str] = set()
        for path, item in (self.spec.get("paths") or {}).items():
            if not isinstance(item, dict):
                continue
            for method, op in item.items():
                if method not in ("post", "put", "patch") or not isinstance(op, dict):
                    continue
                body = op.get("requestBody")
                if not isinstance(body, dict):
                    continue
                body = self.resolve(body)
                content = body.get("content") or {}
                for _ctype, media in content.items():
                    if not isinstance(media, dict):
                        continue
                    schema = media.get("schema")
                    if not isinstance(schema, dict):
                        continue
                    schema = self.resolve(schema)
                    key = schema.get("title") or (schema.get("$ref") or path + ":" + method)
                    if isinstance(key, str) and key not in seen:
                        seen.add(key)
                        found.append((str(key), schema))
                params = list(item.get("parameters") or []) + list(op.get("parameters") or [])
                param_schema = {"properties": {}, "required": []}
                for p in params:
                    if not isinstance(p, dict):
                        continue
                    p = self.resolve(p)
                    name = p.get("name")
                    if not name:
                        continue
                    ps = p.get("schema") if isinstance(p.get("schema"), dict) else {}
                    param_schema["properties"][name] = self.resolve(ps) or {"type": p.get("type")}
                    if p.get("required"):
                        param_schema["required"].append(name)
                if param_schema["properties"]:
                    found.append((f"params:{method}:{path}", param_schema))
        return found

    def named_schemas(self) -> list[tuple[str, dict]]:
        out = []
        for name, schema in self.schemas.items():
            if not isinstance(schema, dict):
                continue
            if NOISE_SCHEMA.search(name):
                continue
            out.append((name, schema))
        return out


def norm(s: str) -> str:
    s = s.lower()
    s = s.replace("ibm_", "").replace("ibm", "")
    s = re.sub(r"[^a-z0-9]+", " ", s)
    drop = {
        "prototype", "patch", "request", "response", "body", "create", "update",
        "delete", "get", "list", "options", "params", "collection", "template",
        "identity", "href", "schema", "model", "input", "output", "v1", "v2",
        "v3", "v4", "resource", "data",
    }
    return " ".join(t for t in s.split() if t and t not in drop)


def tokens(s: str) -> set[str]:
    return {t for t in norm(s).split() if len(t) > 1}


def writable_attrs(attrs: dict | None, prefix: str = "") -> dict[str, dict]:
    out: dict[str, dict] = {}
    if not isinstance(attrs, dict):
        return out
    for name, a in attrs.items():
        if not isinstance(a, dict):
            continue
        path = f"{prefix}{name}"
        req = bool(a.get("required") and not a.get("computed"))
        opt = bool(a.get("optional"))
        elem = a.get("elem") or {}
        nested = elem.get("kind") == "resource" and a.get("type") in ("list", "set")
        if nested:
            out.update(writable_attrs(elem.get("attributes") or {}, path + "."))
            continue
        if req or opt:
            out[path] = a
    return out


def jaccard(a: set[str], b: set[str]) -> float:
    if not a or not b:
        return 0.0
    inter = len(a & b)
    if not inter:
        return 0.0
    return inter / len(a | b)


def score_schema(tf_name: str, tf_keys: set[str], schema_name: str, schema_keys: set[str]) -> float:
    if not schema_keys:
        return 0.0
    prop = jaccard(tf_keys, schema_keys)
    tf_t = tokens(tf_name)
    sc_t = tokens(schema_name)
    name = jaccard(tf_t, sc_t)
    if tf_t and sc_t and (tf_t <= sc_t or sc_t <= tf_t):
        name = max(name, 0.7)
    prefer = 0.08 if PREFER_SCHEMA.search(schema_name or "") else 0.0
    return prop * 0.7 + name * 0.3 + prefer


def merge_constraint(dst: dict, src: dict) -> None:
    if not src:
        return
    if src.get("enum") and not dst.get("enum"):
        dst["enum"] = src["enum"]
    elif src.get("enum") and dst.get("enum"):
        seen = list(dst["enum"])
        for x in src["enum"]:
            if x not in seen:
                seen.append(x)
        dst["enum"] = seen
    if src.get("pattern") and not dst.get("pattern"):
        dst["pattern"] = src["pattern"]
    if src.get("kind") and not dst.get("kind"):
        dst["kind"] = src["kind"]
    if src.get("format") and not dst.get("format"):
        dst["format"] = src["format"]
    if src.get("min_len") is not None and dst.get("min_len") is None:
        dst["min_len"] = src["min_len"]
    if src.get("max_len") is not None and dst.get("max_len") is None:
        dst["max_len"] = src["max_len"]
    if src.get("min") is not None and dst.get("min") is None:
        dst["min"] = src["min"]
    if src.get("max") is not None and dst.get("max") is None:
        dst["max"] = src["max"]
    if src.get("description") and not dst.get("description"):
        dst["description"] = src["description"]
    if src.get("required"):
        dst["api_required"] = True


def overlay_package(pkg: str, slugs: list[str]) -> dict:
    extract_path = EXTRACT / f"{pkg}.json"
    if not extract_path.is_file():
        return {"package": pkg, "resources": [], "matched": 0, "specs": [], "resource_count": 0, "overlay_attrs": 0}
    data = json.loads(extract_path.read_text())
    indexes: list[SpecIndex] = []
    spec_meta = []
    for slug in slugs:
        spec = load_spec(slug)
        if not spec:
            continue
        idx = SpecIndex(spec, slug)
        indexes.append(idx)
        spec_meta.append(
            {
                "slug": slug,
                "title": idx.title,
                "schema_count": len(idx.schemas),
                "path_count": len(spec.get("paths") or {}),
            }
        )

    catalog: list[tuple[SpecIndex, str, dict, set[str], dict]] = []
    for idx in indexes:
        for name, schema in idx.request_schemas() + idx.named_schemas():
            keys = set(idx.properties_of(schema).keys())
            if len(keys) < 1:
                continue
            flat = idx.flatten_props(schema)
            catalog.append((idx, name, schema, keys, flat))

    unique_leaf: dict[str, tuple[str, str, dict]] = {}
    ambiguous: set[str] = set()
    for idx, sname, _schema, _keys, flat in catalog:
        for path, c in flat.items():
            useful = any(c.get(k) for k in ("enum", "pattern", "kind", "min_len", "max_len", "min", "max", "format"))
            if not useful:
                continue
            leaf = path.split(".")[-1]
            if leaf in ambiguous:
                continue
            prev = unique_leaf.get(leaf)
            if prev is None:
                unique_leaf[leaf] = (idx.slug, sname, c)
            elif (
                prev[2].get("enum") != c.get("enum")
                or prev[2].get("pattern") != c.get("pattern")
                or prev[2].get("kind") != c.get("kind")
                or prev[2].get("min") != c.get("min")
                or prev[2].get("max") != c.get("max")
            ):
                ambiguous.add(leaf)
                unique_leaf.pop(leaf, None)

    resources_out = []
    matched_n = 0
    for res in data.get("resources") or []:
        tf_name = res.get("name") or ""
        wattrs = writable_attrs(res.get("attributes") or {})
        tf_keys = {p.split(".")[-1] for p in wattrs}
        best = None
        best_score = 0.0
        best_flat: dict = {}
        for idx, sname, schema, keys, flat in catalog:
            sc = score_schema(tf_name, tf_keys, sname, keys)
            if sc > best_score:
                best_score = sc
                best = (idx, sname, schema)
                best_flat = flat
        attr_overlay: dict[str, dict] = {}
        schema_name = ""
        spec_slug = ""
        if best and best_score >= 0.12:
            idx, schema_name, schema = best
            spec_slug = idx.slug
            by_leaf: dict[str, dict] = {}
            for path, c in best_flat.items():
                leaf = path.split(".")[-1]
                by_leaf.setdefault(leaf, c)
            for path, _a in wattrs.items():
                leaf = path.split(".")[-1]
                src = best_flat.get(path) or by_leaf.get(leaf)
                if src:
                    dst: dict = {}
                    merge_constraint(dst, src)
                    if dst:
                        dst["source"] = spec_slug
                        dst["schema"] = schema_name
                        attr_overlay[path] = dst
            matched_n += 1
        for path, _a in wattrs.items():
            if path in attr_overlay:
                continue
            leaf = path.split(".")[-1]
            hit = unique_leaf.get(leaf)
            if not hit:
                continue
            slug, sname, src = hit
            dst: dict = {}
            merge_constraint(dst, src)
            if dst:
                dst["source"] = slug
                dst["schema"] = sname
                attr_overlay[path] = dst
        if attr_overlay and not schema_name:
            matched_n += 1
            schema_name = "leaf-unique"
            spec_slug = next(iter(attr_overlay.values())).get("source") or ""
        resources_out.append(
            {
                "name": tf_name,
                "kind": res.get("kind") or "resource",
                "schema": schema_name,
                "spec": spec_slug,
                "score": round(best_score, 4),
                "attrs": attr_overlay,
            }
        )
    return {
        "package": pkg,
        "specs": spec_meta,
        "matched": matched_n,
        "resource_count": len(resources_out),
        "overlay_attrs": sum(len(r["attrs"]) for r in resources_out),
        "resources": resources_out,
    }


def main() -> int:
    mapping = load_package_hrefs()
    pkgs = sorted(
        p.stem
        for p in EXTRACT.glob("*.json")
        if p.name not in ("index.json", "overlay-index.json") and not p.name.endswith(".overlay.json")
    )
    index = {"packages": [], "matched": 0, "overlay_attrs": 0, "specs_used": []}
    specs_used: set[str] = set()
    for pkg in pkgs:
        slugs = list(mapping.get(pkg) or [])
        if not slugs:
            compact = pkg.replace("_", "")
            for p in RAW.glob("*.json"):
                if p.name == "manifest.json":
                    continue
                stem = p.stem.replace("-", "").replace("_", "")
                if compact == stem or (len(compact) >= 8 and compact in stem):
                    slugs.append(p.stem)
        overlay = overlay_package(pkg, slugs)
        out_path = EXTRACT / f"{pkg}.overlay.json"
        out_path.write_text(json.dumps(overlay, indent=2) + "\n")
        index["packages"].append(
            {
                "package": pkg,
                "matched": overlay["matched"],
                "resources": overlay["resource_count"],
                "overlay_attrs": overlay["overlay_attrs"],
                "specs": [s.get("slug") for s in overlay.get("specs") or []],
            }
        )
        index["matched"] += overlay["matched"]
        index["overlay_attrs"] += overlay["overlay_attrs"]
        for s in overlay.get("specs") or []:
            specs_used.add(s.get("slug") or "")
        print(
            f"{pkg:35} specs={len(overlay.get('specs') or []):2} "
            f"matched={overlay['matched']:4}/{overlay['resource_count']:4} "
            f"attrs={overlay['overlay_attrs']:5}",
            flush=True,
        )
    index["specs_used"] = sorted(s for s in specs_used if s)
    (EXTRACT / "overlay-index.json").write_text(json.dumps(index, indent=2) + "\n")
    print(
        f"done: matched {index['matched']} types, {index['overlay_attrs']} attribute overlays",
        flush=True,
    )
    return 0


if __name__ == "__main__":
    sys.exit(main())
