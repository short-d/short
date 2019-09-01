package mdtest

import "short/fw"

type tracer struct {
}

type fakeTrace struct {
}

func (fakeTrace) End() {
}

func (fakeTrace) Next(name string) fw.Trace {
	return fakeTrace{}
}

func (tracer) BeginTrace(name string) fw.Trace {
	return fakeTrace{}
}

func (tracer) Begin() func(string) {
	return func(s string) {}
}

var FakeTracer fw.Tracer = tracer{}
