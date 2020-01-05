package provider

import (
	"github.com/short-d/app/modern/mdgraphql"

	"github.com/short-d/app/fw"
)

// GraphQlPath represents the path for GraphQL APIs.
type GraphQlPath string

// NewGraphGophers creates GraphGopher GraphQL server with GraphQlPath to uniquely identify graphqlPath during dependency injection.
func NewGraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQLAPI) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}
