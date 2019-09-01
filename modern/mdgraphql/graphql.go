package mdgraphql

import (
	"short/fw"
	"short/modern/mdhttp"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type GraphGophers struct {
	logger fw.Logger
	server fw.Server
}

func (g GraphGophers) Shutdown() error {
	return g.server.Shutdown()
}

func (g GraphGophers) ListenAndServe(port int) error {
	return g.server.ListenAndServe(port)
}

func NewGraphGophers(graphqlPath string, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlAPI) fw.Server {
	schema := graphql.MustParseSchema(g.GetSchema(), g.GetResolver())

	relayHandler := relay.Handler{
		Schema: schema,
	}

	server := mdhttp.NewServer(logger, tracer)
	server.HandleFunc(graphqlPath, &relayHandler)

	return GraphGophers{
		logger: logger,
		server: &server,
	}
}
