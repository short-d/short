package provider

import (
	"github.com/byliuyang/app/modern/mdgraphql"

	"github.com/byliuyang/app/fw"
)

// GraphQlPath GraphQl path.
type GraphQlPath string

// GraphGophers initializes a new GraphQl server.
func GraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlAPI) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}
