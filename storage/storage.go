package storage

import (
	"database/sql"
	"github.com/adrg/xdg"
	"github.com/go-gorp/gorp"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"
	"sts2stats/spool"
	"sync"
)

var (
	Cache, CacheErr = xdg.CacheFile("sts2stats/db." + driver)
	Db              = pflag.String("database", Cache, "file path for database")
	Reset           = pflag.BoolP("reindex", "r", true, "reindex all runs")
)

var conn *sql.DB
var dbmap gorp.DbMap
var lock = &sync.Mutex{}

func Init(items ...any) error {
	if *Reset {
		spool.Warn("reindex expected, removing database\n")
		err := os.Remove(*Db)
		if err != nil {
			return err
		}
	}

	if CacheErr != nil {
		return CacheErr
	}
	spool.Info("opening %v database in %s\n", driver, *Db)
	dir, _ := filepath.Split(Cache)
	_ = os.MkdirAll(dir, 0600)
	db, err := sql.Open(driver, *Db)
	if err != nil {
		return err
	}
	conn = db

	dbmap = gorp.DbMap{
		Db:      conn,
		Dialect: gorp.SqliteDialect{},
	}

	err = register(items...)
	if err != nil {
		return err
	}

	return dbmap.CreateTablesIfNotExists()
}

func register(item ...any) error {
	for _, t := range item {
		dbmap.AddTable(t)
	}
	return dbmap.CreateTablesIfNotExists()
}

func SaveNow(item ...any) error {
	lock.Lock()
	err := dbmap.Insert(item...)
	lock.Unlock()
	return err
}

func Close() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			spool.Warn("%v\n", err)
		}
	}
}

func UI() error {
	_, err := dbmap.Db.Exec("SET ui_local_port = 4213; CALL start_ui_server();")
	return err
}

func Query(query string, args ...any) (*sql.Rows, error) {
	return dbmap.Query(query, args...)
}

func Entities[T any](query string, args ...any) ([]T, error) {
	thing := new(T)
	res, err := dbmap.Select(thing, query, args...)
	var result []T
	for _, r := range res {
		result = append(result, r.(T))
	}
	return result, err
}
