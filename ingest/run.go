package ingest

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"sts2stats/model"
	"sts2stats/spool"
	"sts2stats/stats"
	"sts2stats/storage"
	"sync"
)

type Profile struct {
	RunSaves []*model.Run
}

var SkippedExtensions = []string{".backup"}

// AddProfile adds a profile by reading from a folder.
// the path should be something like ~/.local/share/SlayTheSpire2/steam/<steamid>/profile1
func AddProfile(f fs.FS, index int) error {
	existingRows, err := storage.Query("SELECT RunId FROM RunStat WHERE InProgress = 0")
	if err != nil {
		return err
	}
	var existingRunIds []string
	for existingRows.Next() {
		var id string
		err = existingRows.Scan(&id)
		if err != nil {
			return err
		}
		existingRunIds = append(existingRunIds, id)
	}

	var runs []*model.Run
	wg := sync.WaitGroup{}
	err = fs.WalkDir(f, "saves/history", func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".run" {
			if !slices.Contains(SkippedExtensions, ext) {
				spool.Warn("skipping file with unexpected file extension: %s\n", path)
			}
			return nil
		}
		content, err := os.ReadFile(path)
		save, err := model.NewRun(content)
		if slices.Contains(existingRunIds, save.RunId) {
			return nil
		}
		if err != nil {
			return err
		}
		wg.Go(func() {
			err := stats.Enrich(save)
			if err != nil {
				spool.Warn("%v\n", err)
			}
		})
		runs = append(runs, &save)
		return nil
	})
	if err != nil {
		return err
	}
	wg.Wait()
	spool.Warn("indexed all runs")
	return nil
}
