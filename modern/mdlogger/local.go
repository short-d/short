package mdlogger

import (
	"log"

	"short/fw"
)

type Local struct {
}

func (Local) Error(err error) {
	log.Println(err.Error())
}

func (Local) Info(info string) {
	log.Println(info)
}

func (Local) Crash(err error) {
	log.Fatal(err)
}

func NewLocal() fw.Logger {
	return Local{}
}
