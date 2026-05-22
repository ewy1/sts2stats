//go:build.sh pprof

package main

import (
	"net/http"
	"runtime"
)

func init() {
	runtime.SetBlockProfileRate(1)
	runtime.SetMutexProfileFraction(1)
	go http.ListenAndServe("localhost:6060", nil)
}
