package modern

import (
	"net/http"
	"tinyURL/fw"
	"tinyURL/modern/router"
)

type CustomRouting struct {
	logger fw.Logger
	server fw.Server
}

func (g CustomRouting) Shutdown() error {
	return g.server.Shutdown()
}

func (g CustomRouting) ListenAndServe(port int) error {
	return g.server.ListenAndServe(port)
}

func NewCustomRouting(logger fw.Logger, tracer fw.Tracer, routes []fw.Route) fw.Server {
	httpRouter := router.NewHttpHandler()

	for _, route := range routes {
		route := route
		err := httpRouter.AddRoute(route.Method, route.MatchPrefix, route.Path, func(w http.ResponseWriter, r *http.Request, params router.Params) {
			route.Handle(w, r, params)
		})
		if err != nil {
			panic(err)
		}
	}

	server := NewHttpServer(logger, tracer)
	server.HandleFunc("/", &httpRouter)

	return CustomRouting{
		logger: logger,
		server: &server,
	}
}
