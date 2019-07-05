package fw

type Server interface {
	ListenAndServe(port int)
}
