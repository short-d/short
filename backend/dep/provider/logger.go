package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdlogger"
)

type LogPrefix string

func NewLocalLogger(
	prefix LogPrefix,
	level fw.LogLevel,
	stdout fw.StdOut,
	timer fw.Timer,
	programRuntime fw.ProgramRuntime,
) mdlogger.Local {
	return mdlogger.NewLocal(string(prefix), level, stdout, timer, programRuntime)
}
