package stats

import (
	"sts2stats/model"
)

type CardChoice struct {
	RunStat
	model.PlayerStat
	RoomStat
	Card    string
	Upgrade int
	Picked  bool
}

func EnrichCardChoice(run model.Run, st RunStat) (result []any, err error) {
	for ia, act := range run.MapPointHistory {
		for i, floor := range act {
			for _, stat := range floor.PlayerStats {
				for _, choice := range stat.CardChoices {
					result = append(result, &CardChoice{
						RunStat:    st,
						Card:       choice.Card.Id,
						Upgrade:    choice.Card.CurrentUpgradeLevel,
						PlayerStat: stat.PlayerStat,
						RoomStat: RoomStat{
							Floor: i,
							Act:   ia + 1,
						},
						Picked: false,
					})
				}
			}
		}
	}
	return
}
