package stats

import (
	"slices"
	"sts2stats/model"
)

type GameVersion struct {
	Version string `db:"Version,primarykey"`
}

var versions []string

func EnrichGameVersion(run model.Run, stat RunStat) (result []any, err error) {
	if !slices.Contains(versions, run.BuildID) {
		versions = append(versions, run.BuildID)
		v := GameVersion{
			Version: run.BuildID,
		}
		result = append(result, &v)
	}
	return result, nil
}
