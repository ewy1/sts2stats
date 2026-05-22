package spool

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"runtime/debug"
	"time"
)

var (
	DebugFlag = pflag.BoolP("debug", "d", false, "show debug output")
)

func Debug(format string, args ...any) {
	if !*DebugFlag {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}

func Info(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stdout, timestamp(format), args...)
}

func Warn(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, timestamp(format), args...)
}

func Panic(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, timestamp(format), args...)
	debug.PrintStack()
	os.Exit(1)
}

func timestamp(in string) string {
	n := time.Now()
	return fmt.Sprintf("[%v] %v", n.Format(time.TimeOnly), in)
}
