package fw

type Tracer interface {
	Begin() func(string)
}
