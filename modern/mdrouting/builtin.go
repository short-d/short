package mdrouting

import (
	"net/http"

	"short/fw"
	"short/modern/mdhttp"
	"short/modern/mdrouter"
)

type BuiltIn struct {
	logger fw.Logger
	server fw.Server
}

func (g BuiltIn) Shutdown() error {
	return g.server.Shutdown()
}

func (g BuiltIn) ListenAndServe(port int) error {
	return g.server.ListenAndServe(port)
}

func NewBuiltIn(logger fw.Logger, tracer fw.Tracer, routes []fw.Route) fw.Server {
	httpRouter := mdrouter.NewHTTPHandler()

	for _, route := range routes {
		route := route
		err := httpRouter.AddRoute(route.Method, route.MatchPrefix, route.Path, func(w http.ResponseWriter, r *http.Request, params mdrouter.Params) {
			route.Handle(w, r, params)
		})
		if err != nil {
			panic(err)
		}
	}

	server := mdhttp.NewServer(logger, tracer)
	server.HandleFunc("/", &httpRouter)

	return BuiltIn{
		logger: logger,
		server: &server,
	}
}
