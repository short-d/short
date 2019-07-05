package fw

type App struct {
	server Server
}

func NewApp(s Server) App {
	return App{
		server: s,
	}
}

func (a App) Start() {
	a.server.ListenAndServe(8080)
}
