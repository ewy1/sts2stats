package stats

import (
	"github.com/spf13/pflag"
	"sts2stats/model"
	"sts2stats/spool"
	"sts2stats/storage"
	"sync"
	"time"
)

type LoadFunc = func() error

// Enricher reads data from the RunSave and fills dictionaries and data structures
type Enricher interface {
	Enrich(run model.Run, stat RunStat) ([]any, error)
}

type EnrichFunc func(run model.Run, stat RunStat) ([]any, error)

type EnrichFuncWrapper struct {
	f EnrichFunc
}

func EnrichWrap(f EnrichFunc) Enricher {
	return EnrichFuncWrapper{f: f}
}

func (e EnrichFuncWrapper) Enrich(run model.Run, stat RunStat) ([]any, error) {
	return e.f(run, stat)
}

var Enrichers = map[string]Enricher{
	"act":            EnrichWrap(EnrichActs),
	"version":        EnrichWrap(EnrichGameVersion),
	"ancient choice": EnrichWrap(EnrichAncients),
	"card choice":    EnrichWrap(EnrichCardChoice),
}

var SteamId = pflag.IntP("steamid", "s", 0, "steamid to match players to")

func Enrich(run model.Run) error {
	startTime := time.Now()
	id := run.RunId[:4]
	wg := sync.WaitGroup{}
	st := NewRunStat(run, *SteamId)
	for k, e := range Enrichers {
		wg.Go(func() {
			spool.Debug("[%v] Starting %v enrichment\n", id, k)
			res, err := e.Enrich(run, st)
			if err != nil {
				spool.Panic("%v\n", err)
			}

			if len(res) == 0 {
				spool.Debug("[%v] Finished %v enrichment\n", id, k)
				return
			}

			spool.Debug("[%v] Collected %v entities (%v)\n", id, k, len(res))
			err = storage.SaveNow(res...)
			if err != nil {
				spool.Panic("during %v: %v\n", k, err)
			}
			spool.Debug("[%v] Saved %v entities (%v)\n", id, k, len(res))
		})
	}
	wg.Wait()

	endTime := time.Now()
	spool.Info("[%v] digested run\n", id)
	spool.Debug("[%v] took %.2fs", endTime.Sub(startTime).Seconds())
	return nil
}
