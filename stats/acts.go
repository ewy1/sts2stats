package stats

import (
	"slices"
	"strings"
	"sts2stats/model"
)

type Act struct {
	Index int
	Label string
	Key   string `db:"Key,primarykey"`
}

var actKeys []string

func EnrichActs(run model.Run, stat RunStat) (result []any, err error) {
	for i, a := range run.Acts {
		if slices.Contains(actKeys, a) {
			continue
		}
		actKeys = append(actKeys, a)

		act := Act{
			Index: i,
			Key:   a,
			Label: strings.SplitN(a, ".", 2)[0],
		}

		result = append(result, &act)
	}
	return result, nil
}
