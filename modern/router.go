package modern

import (
	"net/http"
	"tinyURL/fw"

	"github.com/gorilla/mux"
)

type GorillaMux struct {
	logger fw.Logger
	server fw.Server
}

func (g GorillaMux) Shutdown() error {
	return g.server.Shutdown()
}

func (g GorillaMux) ListenAndServe(port int) error {
	return g.server.ListenAndServe(port)
}

func NewGorillaMux(logger fw.Logger, tracer fw.Tracer, routes fw.Routes) fw.Server {
	router := mux.NewRouter()

	for path, handler := range routes {
		router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			params := mux.Vars(r)
			handler(w, r, params)
		})
	}

	server := NewHttpServer(logger, tracer)
	server.HandleFunc("/", router)

	return GorillaMux{
		logger: logger,
		server: &server,
	}
}
