//go:build sqlite

package storage

import (
	// cgo driver
	_ "github.com/mattn/go-sqlite3"
)

var driver = "sqlite3"
