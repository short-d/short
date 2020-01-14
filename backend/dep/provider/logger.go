package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdlogger"
)

// LogPrefix represents the prefix of a log message.
// "[Short]" is the prefix part of the following log message:
// [Short] [Info] 2020-01-07 04:33:22 line 25 at service.go GraphQL API started
type LogPrefix string

// NewLocalLogger creates local logger with LogPrefix to uniquely identify log
// prefix during dependency injection.
func NewLocalLogger(
	prefix LogPrefix,
	level fw.LogLevel,
	stdout fw.StdOut,
	timer fw.Timer,
	programRuntime fw.ProgramRuntime,
) mdlogger.Local {
	return mdlogger.NewLocal(string(prefix), level, stdout, timer, programRuntime)
}
