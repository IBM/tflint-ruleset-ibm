#!/usr/bin/env python3
# Copyright IBM Corp. 2026 All Rights Reserved.
# Licensed under the Mozilla Public License v2.0
"""Download IBM Cloud OpenAPI documents into a private per-user cache (disposable).

The manifest records a SHA-256 for every saved document; openapi-overlay.py
refuses documents that no longer match it.
"""

from __future__ import annotations

import json
import os
import re
import sys
from concurrent.futures import ThreadPoolExecutor, as_completed
from pathlib import Path
from urllib.error import HTTPError, URLError
from urllib.request import Request, urlopen

from apidocs_cache import raw_dir, sha256_file

RAW = raw_dir()
SCRIPT_DIR = Path(__file__).resolve().parent
CATALOG = Path(os.environ.get("IBM_APIDOCS_CATALOG", str(SCRIPT_DIR / "api-catalog.json")))

UA = "tflint-ruleset-ibm-apidoc-fetch/1.0"

# Terraform package → IBM Cloud apidocs hrefs (from the public API catalog).
PACKAGE_HREFS: dict[str, list[str]] = {
    "accountmanagement": [
        "/docs/apis/bss-account-management",
        "/docs/apis/user-management",
        "/docs/apis/support-center/case-management",
        "/docs/apis/support-center/case-management-v2",
    ],
    "appconfiguration": ["/docs/apis/app-configuration"],
    "appid": [
        "/docs/apis/app-id/auth",
        "/docs/apis/app-id/management",
        "/docs/apis/app-id/profiles",
    ],
    "atracker": ["/docs/apis/atracker"],
    "backuprecovery": [
        "/docs/apis/baas/backup-recovery",
        "/docs/apis/baas/ibm_reporting_api",
    ],
    "catalogmanagement": [
        "/docs/apis/resource-catalog/private-catalog",
        "/docs/apis/resource-catalog/global-catalog",
        "/docs/apis/search",
    ],
    "cdtektonpipeline": ["/docs/apis/tekton-pipeline"],
    "cdtoolchain": ["/docs/apis/toolchain"],
    "cis": ["/docs/apis/cis"],
    "classicinfrastructure": [],
    "cloudant": ["/docs/apis/cloudant/cloudant-gen2"],
    "cloudfoundry": [],
    "cloudshell": ["/docs/apis/cloudshell"],
    "codeengine": ["/docs/apis/codeengine/v2"],
    "configurationaggregator": ["/docs/apis/app-configuration"],
    "contextbasedrestrictions": ["/docs/apis/context-based-restrictions"],
    "cos": [
        "/docs/apis/cos/cos-configuration",
        "/docs/apis/cos/cos-compatibility",
        "/docs/apis/cos/cos-configuration-f2019",
        "/docs/apis/cos/cos-compatibility-f2019",
    ],
    "database": [
        "/docs/apis/cloud-databases-api/cloud-databases-api-v5",
        "/docs/apis/cloud-databases-api/cloud-databases-api-v4",
    ],
    "db2": [
        "/docs/apis/db2-on-cloud/db2-on-cloud-v4",
        "/docs/apis/db2-on-cloud/db2-on-cloud-byoc",
        "/docs/apis/db2-on-cloud/db2-saas-perf-v4",
        "/docs/apis/db2-warehouse-on-cloud/db2-warehouse-on-cloud-v5",
        "/docs/apis/db2-warehouse-on-cloud/db2-warehouse-on-cloud-v4",
        "/docs/apis/db2-warehouse-on-cloud/db2-warehouse-on-cloud-byoc",
    ],
    "directlink": ["/docs/apis/direct_link", "/docs/apis/direct_link_provider_api"],
    "dnsservices": ["/docs/apis/dns-svcs"],
    "drautomationservice": ["/docs/apis/dr-automation-powervs"],
    "enterprise": [
        "/docs/apis/enterprise-apis/enterprise",
        "/docs/apis/enterprise-apis/billing-unit",
        "/docs/apis/enterprise-apis/resource-usage-reports",
    ],
    "eventnotification": ["/docs/apis/event-notifications"],
    "eventstreams": [
        "/docs/apis/event-streams/adminrest",
        "/docs/apis/event-streams/schemaregistry",
        "/docs/apis/event-streams/restproducer",
        "/docs/apis/event-streams/restproducer_v2",
    ],
    "functions": [],
    "globaltagging": ["/docs/apis/tagging"],
    "hpcs": ["/docs/apis/hs-crypto", "/docs/apis/uko"],
    "iamaccessgroup": ["/docs/apis/iam-access-groups"],
    "iamidentity": ["/docs/apis/iam-identity-token-api"],
    "iampolicy": ["/docs/apis/iam-policy-management"],
    "kms": ["/docs/apis/key-protect", "/docs/apis/hs-crypto", "/docs/apis/uko"],
    "kubernetes": ["/docs/apis/kubernetes/containers-v1-v2"],
    "logs": ["/docs/apis/logs-service-api"],
    "logsrouter": ["/docs/apis/logs-router-service-api/logs-router-v1"],
    "logsrouting": ["/docs/apis/logs-router-service-api/logs-router-v1"],
    "metricsrouter": ["/docs/apis/metrics-router/metrics-router-v3"],
    "mqcloud": ["/docs/apis/mq-on-cloud"],
    "pag": ["/docs/apis/privileged-access-gateway"],
    "partnercentersell": [],
    "platformnotifications": ["/docs/apis/platform-notifications"],
    "power": ["/docs/apis/power-cloud", "/docs/apis/power-cloud-metadata-identity"],
    "powerhaautomationservice": [],
    "project": ["/docs/apis/projects"],
    "pushnotification": [],
    "registry": ["/docs/apis/container-registry", "/docs/apis/vulnerability-advisor"],
    "resourcecontroller": [
        "/docs/apis/resource-controller/resource-controller",
        "/docs/apis/resource-controller/ibm-cloud-osb-api",
    ],
    "resourcemanager": ["/docs/apis/resource-controller/resource-manager"],
    "satellite": [],
    "scc": ["/docs/apis/security-compliance", "/docs/apis/workload-protection"],
    "schematics": ["/docs/apis/schematics/schematics"],
    "sdsaas": ["/docs/apis/block-storage", "/docs/apis/object-storage"],
    "secretsmanager": ["/docs/apis/secrets-manager/secrets-manager-v2"],
    "secretsmanagerinstancemanagement": [
        "/docs/apis/secrets-manager/secrets-manager-instance-management-v2"
    ],
    "transitgateway": ["/docs/apis/transit-gateway"],
    "usagereports": ["/docs/apis/metering-reporting", "/docs/apis/usage-metering"],
    "vmware": ["/docs/apis/vmware-service", "/docs/apis/vmware-solutions"],
    "vpc": [
        "/docs/apis/vpc/latest",
        "/docs/apis/vpc-identity/latest",
        "/docs/apis/vpc-metadata/latest",
    ],
}


def hrefs_from_catalog() -> list[str]:
    if not CATALOG.is_file():
        return []
    try:
        data = json.loads(CATALOG.read_text())
    except json.JSONDecodeError:
        return []
    out: list[str] = []
    for cat in data.get("categories") or []:
        for svc in cat.get("services") or []:
            href = svc.get("href")
            if href:
                out.append(href)
    return out


def slug_of(href: str) -> str:
    href = href.strip("/")
    href = re.sub(r"^docs/apis/", "", href)
    return href.replace("/", "-")


def urls_for(href: str) -> list[str]:
    path = href.strip("/")
    if path.startswith("docs/apis/"):
        path = path[len("docs/apis/") :]
    return [
        f"https://cloud.ibm.com/apidocs/{path}.json",
        f"https://cloud.ibm.com/docs/apis/{path}.json",
    ]


def fetch(url: str) -> dict | None:
    req = Request(url, headers={"User-Agent": UA, "Accept": "application/json"})
    try:
        with urlopen(req, timeout=45) as resp:
            raw = resp.read()
    except (HTTPError, URLError, TimeoutError, OSError):
        return None
    try:
        data = json.loads(raw)
    except json.JSONDecodeError:
        return None
    if not isinstance(data, dict):
        return None
    if not (data.get("openapi") or data.get("swagger") or data.get("paths")):
        return None
    return data


def save_spec(slug: str, data: dict) -> Path:
    path = RAW / f"{slug}.json"
    path.write_text(json.dumps(data, indent=2) + "\n")
    return path


def all_hrefs() -> list[str]:
    seen: set[str] = set()
    out: list[str] = []
    for href in hrefs_from_catalog():
        if href not in seen:
            seen.add(href)
            out.append(href)
    for hrefs in PACKAGE_HREFS.values():
        for href in hrefs:
            if href not in seen:
                seen.add(href)
                out.append(href)
    return out


def download_one(href: str) -> tuple[str, str, bool, str, str]:
    slug = slug_of(href)
    last_err = "no url"
    for url in urls_for(href):
        data = fetch(url)
        if data:
            path = save_spec(slug, data)
            title = (data.get("info") or {}).get("title") or slug
            return slug, title, True, url, sha256_file(path)
        last_err = url
    return slug, "", False, last_err, ""


def write_manifest(results: list[tuple[str, str, bool, str, str]]) -> None:
    mapping: dict[str, list[str]] = {}
    for pkg, hrefs in PACKAGE_HREFS.items():
        mapping[pkg] = [slug_of(h) for h in hrefs]
    manifest = {
        "packages": mapping,
        "downloaded": [
            {"slug": s, "title": t, "ok": ok, "url": u, "sha256": h}
            for s, t, ok, u, h in sorted(results)
        ],
    }
    (RAW / "manifest.json").write_text(json.dumps(manifest, indent=2) + "\n")


def main() -> int:
    hrefs = all_hrefs()
    print(f"fetching {len(hrefs)} IBM Cloud API specs into {RAW}", flush=True)
    results: list[tuple[str, str, bool, str, str]] = []
    ok_n = 0
    with ThreadPoolExecutor(max_workers=12) as pool:
        futs = {pool.submit(download_one, h): h for h in hrefs}
        for fut in as_completed(futs):
            slug, title, ok, url, digest = fut.result()
            results.append((slug, title, ok, url, digest))
            if ok:
                ok_n += 1
                print(f"  ok  {slug:50} {title[:50]}", flush=True)
            else:
                print(f"  miss {slug:50}", flush=True)
    write_manifest(results)
    print(f"done: {ok_n}/{len(hrefs)} specs at {RAW}", flush=True)
    return 0 if ok_n else 1


if __name__ == "__main__":
    sys.exit(main())
