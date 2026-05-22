package stats

import (
	"sts2stats/model"
	"time"
)

type RunStat struct {
	RunId         string
	StartTime     time.Time
	Ascension     int
	Version       string
	Win           bool
	FloorsClimbed int
	Abandoned     bool
	InProgress    bool
	Character     string
}

type RoomStat struct {
	Act   int
	Floor int
}

func NewRunStat(run model.Run, steamid int) RunStat {
	var player *model.Player
	for _, p := range run.Players {
		if p.ID == steamid {
			player = &p
		}
	}
	if player == nil {
		player = &run.Players[0]
	}
	st := RunStat{
		Version:       run.BuildID,
		StartTime:     time.Unix(int64(run.StartTime), 0),
		Ascension:     run.Ascension,
		Win:           run.Win,
		RunId:         run.RunId,
		Abandoned:     run.WasAbandoned,
		FloorsClimbed: runLen(run),
		InProgress:    run.KilledByEncounter != "" || run.KilledByEvent != "" || run.Win != true,
		Character:     player.Character,
	}
	return st
}

func runLen(run model.Run) int {
	var res int
	for _, a := range run.MapPointHistory {
		res += len(a)
	}
	return res
}
