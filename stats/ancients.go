package stats

import (
	"sts2stats/model"
)

const (
	AncientKey = "ancient"
)

type AncientChoice struct {
	RunStat
	AncientOption
	ActIndex int
	ActName  string
	Chosen   bool
}

type AncientOption struct {
	Key  string
	Type string
}

func EnrichAncients(run model.Run, st RunStat) (opts []any, err error) {
	for actIndex, act := range run.MapPointHistory {
		for _, floor := range act {
			if floor.MapPointType != AncientKey {
				continue
			}

			for _, stat := range floor.PlayerStats {
				for _, choice := range stat.AncientChoice {

					opts = append(opts, &AncientChoice{
						AncientOption: AncientOption{
							Key:  choice.TextKey,
							Type: choice.Title.Table,
						},
						RunStat:  st,
						ActIndex: actIndex,
						ActName:  run.Acts[actIndex],
						Chosen:   choice.WasChosen,
					})
				}
			}

		}
	}
	return opts, nil
}
