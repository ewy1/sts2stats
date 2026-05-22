//go:build sqlite_wasm

package storage

import (
	// wasm driver
	_ "github.com/ncruces/go-sqlite3/driver"
)

var driver = "sqlite3"
