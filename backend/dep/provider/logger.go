package provider

import (
	"github.com/short-d/app/fw/io"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/runtime"
	"github.com/short-d/app/fw/timer"
)

// LogPrefix represents the prefix of a log message.
// "[Short]" is the prefix part of the following log message:
// [Short] [Info] 2020-01-07 04:33:22 line 25 at service.go GraphQL API started
type LogPrefix string

// NewLogger creates logger with LogPrefix to uniquely identify log prefix
// during dependency injection.
func NewLogger(
	prefix LogPrefix,
	level logger.LogLevel,
	timer timer.Timer,
	programRuntime runtime.Program,
	entryRepo logger.EntryRepository,
) logger.Logger {
	return logger.NewLogger(string(prefix), level, timer, programRuntime, entryRepo)
}

// NewLocalEntryRepo create LocalEntryRepo with line number disabled in logs.
func NewLocalEntryRepo(output io.Output) logger.Local {
	return logger.NewLocal(output, false)
}
