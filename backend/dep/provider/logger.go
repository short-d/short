package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdlogger"
)

// LogPrefix represents the prefix of a log message.
// "[Short]" is the prefix part of the following log message:
// [Short] [Info] 2020-01-07 04:33:22 line 25 at service.go GraphQL API started
type LogPrefix string

// NewLogger creates logger with LogPrefix to uniquely identify log prefix
// during dependency injection.
func NewLogger(
	prefix LogPrefix,
	level fw.LogLevel,
	timer fw.Timer,
	programRuntime fw.ProgramRuntime,
	entryRepo mdlogger.EntryRepository,
) mdlogger.Logger {
	return mdlogger.NewLogger(string(prefix), level, timer, programRuntime, entryRepo)
}
