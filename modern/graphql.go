package modern

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"tinyURL/fw"
)

type GraphGophers struct {
	logger fw.Logger
	server *fw.HttpServer
}

func (g GraphGophers) Shutdown() error {
	return g.server.Shutdown()
}

func (g GraphGophers) ListenAndServe(port int) error {
	return g.server.ListenAndServe(port)
}

func NewGraphGophers(logger fw.Logger, g fw.GraphQlApi) fw.Server {
	schema := graphql.MustParseSchema(g.GetSchema(), g.GetResolver())

	relayHandler := relay.Handler{
		Schema: schema,
	}

	server := fw.NewHttpServer(&relayHandler, logger)

	return GraphGophers{
		logger: logger,
		server: &server,
	}
}
