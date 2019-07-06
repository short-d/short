package fw

type Server interface {
	ListenAndServe(port int) error
	Shutdown() error
}
