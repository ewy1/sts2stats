package main

import (
	"github.com/spf13/pflag"
	_ "net/http/pprof"
	"os"
	"os/exec"
	"os/signal"
	"sts2stats/api"
	"sts2stats/ingest"
	"sts2stats/spool"
	"sts2stats/stats"
	"sts2stats/storage"
)

var Types = []any{
	stats.Act{},
	stats.GameVersion{},
	stats.AncientChoice{},
	stats.CardChoice{},
	stats.RunStat{},
}

var Loaders = map[string]stats.LoadFunc{
	//"card": stats.LoadCards,
}

func main() {
	pflag.Parse()
	go api.Init()
	err := storage.Init(Types...)
	for i, l := range Loaders {
		err := l()
		if err != nil {
			spool.Panic("during %v: %v\n", i, err)
		}
	}
	defer storage.Close()
	if err != nil {
		spool.Panic("%v\n", err)
	}
	root := must(os.OpenRoot(*Profile))
	defer root.Close()
	go func() { must[any](nil, ingest.AddProfile(root.FS(), 0)) }()
	if err != nil {

	}
	if !*storage.Headless {
		err = storage.UI()
		if err != nil {
			spool.Warn("ui: %v\n", err)
		}
		err = exec.Command(opener, "http://localhost:4213/").Run()
		if err != nil {
			spool.Warn("ui: %v\n", err)
		}
	}

	c := make(chan os.Signal, 1)
	go signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func must[O any](out O, err error) O {
	if err != nil {
		spool.Panic("%v\n", err)
	}
	return out
}
