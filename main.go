package main

import (
	l "github.com/yanc0/go-packages/logs"
)

func main() {
	l.ConfigSeverity("trace")
	l.ConfigFormat("json")

	l.Trace("ceci est une trace")
	l.Debug("un debug?")
	trace := l.CraftTrace()
	trace.With("music", "despacito")
	trace.With("listeners", 0)
	trace.Dump()

	l.Panic("Oh my god !")
}
