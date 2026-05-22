//go:build !sqlite && !sqlite_wasm

package storage

import (
	_ "github.com/duckdb/duckdb-go/v2"
)

var driver = "duckdb"
