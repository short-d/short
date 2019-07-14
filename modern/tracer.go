package modern

import (
	"log"
	"time"
	"tinyURL/fw"
)

type LocalTracer struct{}

func (LocalTracer) Begin() func(string) {
	start := time.Now()

	return func(name string) {
		end := time.Now()
		diff := end.Sub(start)
		log.Printf("%s %v", name, diff)
	}
}

func NewLocalTracer() fw.Tracer {
	return LocalTracer{}
}
