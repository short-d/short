package mdtest

type tracer struct {
}

func (tracer) Begin() func(string) {
	return func(s string) {}
}

var FakeTracer = tracer{}
