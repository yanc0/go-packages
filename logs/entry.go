package logs

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var severityCodes = map[string]int{
	"trace":    0,
	"debug":    1,
	"info":     2,
	"warn":     3,
	"error":    4,
	"critical": 5,
	"panic":    6,
}

// Entry log
type Entry struct {
	mu       *sync.Mutex
	panic    bool
	format   string
	severity string
	keys     []string
	fields   map[string]interface{}
}

// Trace and Dump Entry
func Trace(message ...string) {
	newEntry("trace", message...).Dump()
}

// CraftTrace entry
func CraftTrace(message ...string) *Entry{
	return newEntry("trace", message...)
}

// Debug entry
func Debug(message ...string) {
	newEntry("debug", message...).Dump()
}

// CraftDebug entry
func CraftDebug(message ...string) *Entry{
	return newEntry("debug", message...)
}

// Info entry
func Info(message ...string) {
	newEntry("info", message...).Dump()
}

// CraftInfo entry
func CraftInfo(message ...string) *Entry{
	return newEntry("info", message...)
}

// Warn entry
func Warn(message ...string) {
	newEntry("warn", message...).Dump()
}

// CraftWarn entry
func CraftWarn(message ...string) *Entry{
	return newEntry("warn", message...)
}

// Error entry
func Error(message ...string) {
	newEntry("error", message...).Dump()
}

// CraftError entry
func CraftError(message ...string) *Entry{
	return newEntry("error", message...)
}

// Critical entry
func Critical(message ...string) {
	newEntry("critical", message...).Dump()
}

// CraftCritical entry
func CraftCritical(message ...string) *Entry{
	return newEntry("critical", message...)
}

// Panic entry
func Panic(message ...string) {
	e := newEntry("panic", message...)
	e.panic = true
	e.Dump()
}

// CraftPanic entry
func CraftPanic(message ...string) *Entry{
	e := newEntry("panic", message...)
	e.panic = true
	return e
}

// With additionnal fields
func (e *Entry) With(key string, value interface{}) *Entry {
	e.setField(key, value)
	return e
}

// Dump entry
func (e *Entry) Dump() {
	if severityCodes[configSeverity] > severityCodes[e.severity] {
		return
	}

	if configActivateTimestamp {
		e.setField("time", time.Now())
	}

	text := ""
	std := os.Stdout
	switch configFormat {
	case "json":
		text = e.dumpJSON()
	case "text":
		text = e.dumpText()
	default:
		text = e.dumpText()
	}
	if severityCodes[e.severity] > 3 {
		std = os.Stderr
	}
	fmt.Fprint(std, text)
	if e.panic {
		panic("from logging")
	}
}

func newEntry(severity string, messages ...string) *Entry {
	e := &Entry{}
	e.mu = &sync.Mutex{}
	e.keys = make([]string, 0)
	e.fields = make(map[string]interface{})
	e.severity = severity
	msg := strings.Join(messages, "")
	if msg != "" {
		e.setField("msg", msg)
	}
	e.setField("level", severity)
	for k, v := range configDefaultFields {
		e.setField(k, v)
	}
	return e
}

func (e *Entry) setField(key string, value interface{}) {
	e.mu.Lock()
	e.keys = append(e.keys, key)
	e.fields[key] = value
	e.mu.Unlock()
}
