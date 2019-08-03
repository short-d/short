package modern

import (
	"log"
	"time"
	"tinyURL/fw"

	uuid "github.com/satori/go.uuid"
)

type LocalTracer struct{}

type LocalTrace struct {
	id    string
	name  string
	start time.Time
}

func (t LocalTrace) Next(name string) fw.Trace {
	start := time.Now()
	log.Printf("[Trace Start id=%s name=%s startAt=%v]", t.id, name, start)
	return LocalTrace{
		id:    t.id,
		name:  name,
		start: start,
	}
}

func (t LocalTrace) End() {
	end := time.Now()
	diff := end.Sub(t.start)
	log.Printf("[Trace End   id=%s name=%s endAt=%v duration=%v]", t.id, t.name, end, diff)
}

func (LocalTracer) BeginTrace(name string) fw.Trace {
	id := uuid.NewV4().String()
	start := time.Now()

	log.Printf("[Trace Start id=%s name=%s startAt=%v]", id, name, start)
	return LocalTrace{
		id:    id,
		name:  name,
		start: start,
	}
}

func NewLocalTracer() fw.Tracer {
	return LocalTracer{}
}
