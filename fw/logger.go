package fw

import "log"

type Logger interface {
	Info(info string)
	Error(err error)
	Crash(err error)
}

type LocalLogger struct {
}

func (LocalLogger) Error(err error) {
	log.Println(err.Error())
}

func (LocalLogger) Info(info string) {
	log.Println(info)
}

func (LocalLogger) Crash(err error) {
	log.Fatal(err)
}

func NewLocalLogger() Logger {
	return LocalLogger{}
}
