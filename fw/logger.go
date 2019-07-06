package fw

type Logger interface {
	Info(info string)
	Error(err error)
	Crash(err error)
}
