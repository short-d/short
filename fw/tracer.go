package fw

type Trace interface {
	End()
	Next(name string) Trace
}

type Tracer interface {
	BeginTrace(name string) Trace
}
