#!/usr/bin/env bash
set -euo pipefail
TMP="$(mktemp -d)"
go build -tags duckdb "$@" -o "$TMP/sts2s" .
export XDG_CACHE_HOME="$TMP"
"$TMP/sts2s" --profile "$PROFILE" --database "$TMP/db"