package modern

import (
	"log"
	"short/fw"
)

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

func NewLocalLogger() fw.Logger {
	return LocalLogger{}
}
