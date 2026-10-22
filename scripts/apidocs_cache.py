# Copyright IBM Corp. 2026 All Rights Reserved.
# Licensed under the Mozilla Public License v2.0
"""Private, per-user cache directory for downloaded IBM Cloud API documents."""

from __future__ import annotations

import hashlib
import os
from pathlib import Path


def raw_dir() -> Path:
    """Return the spec cache, refusing locations other users can tamper with.

    Defaults to $XDG_CACHE_HOME/tflint-ruleset-ibm/apidocs (or
    ~/.cache/...). Set IBM_APIDOCS_DIR to override.
    """
    override = os.environ.get("IBM_APIDOCS_DIR")
    if override:
        path = Path(override)
    else:
        cache = os.environ.get("XDG_CACHE_HOME") or str(Path.home() / ".cache")
        path = Path(cache) / "tflint-ruleset-ibm" / "apidocs"
    path.mkdir(parents=True, exist_ok=True, mode=0o700)
    if path.is_symlink():
        raise SystemExit(f"refusing symlinked spec cache {path}")
    st = path.stat()
    if hasattr(os, "getuid") and st.st_uid != os.getuid():
        raise SystemExit(f"spec cache {path} is not owned by the current user")
    if st.st_mode & 0o022:
        raise SystemExit(f"spec cache {path} is writable by other users")
    return path


def sha256_file(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()
