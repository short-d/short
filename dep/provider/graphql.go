package provider

import (
	"github.com/byliuyang/app/modern/mdgraphql"

	"github.com/byliuyang/app/fw"
)

// GraphQlPath represents the path for GraphQl APIs.
type GraphQlPath string

// GraphGophers creates GraphGopher GraphQL server with GraphQlPath to uniquely identify graphqlPath during dependency injection.
func GraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlAPI) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}
