package storage

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"sts2stats/spool"
)

//go:embed views
var data embed.FS

func SetupViews() error {
	err := fs.WalkDir(data, "views", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		name := strings.TrimSuffix(strings.Join(strings.Split(path, "/")[1:], "_"), ext)
		content, err := fs.ReadFile(data, path)
		if err != nil {
			spool.Warn("view %s: %s\n", name, err)
			return nil
		}
		err = AddView(name, string(content))
		if err != nil {
			spool.Warn("view %s: %s\n", name, err)
			return nil
		}
		return nil
	})
	return err
}

func AddView(name string, selectQuery string) error {
	_, err := conn.Exec(fmt.Sprintf("CREATE OR REPLACE VIEW _%s AS %s;", name, selectQuery))
	return err
}
