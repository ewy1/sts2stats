#!/usr/bin/env bash
set -euo pipefail
TMP="$(mktemp -d)"
go build -tags duckdb "$@" -o "$TMP/sts2s" .
cd "$LOC"
"$TMP/sts2s" --database "$TMP/db"