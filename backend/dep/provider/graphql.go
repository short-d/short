package provider

import (
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/service"
)

// GraphQLPath represents the path for GraphQL APIs.
type GraphQLPath string

// NewGraphQLService creates GraphQL service with GraphQLPath to uniquely
// identify graphQLPath during dependency injection.
func NewGraphQLService(
	gqlPath GraphQLPath,
	handler graphql.Handler,
	webUI graphql.WebUI,
	logger logger.Logger,
) service.GraphQL {
	return service.NewGraphQL(logger, string(gqlPath), handler, webUI)
}

// GraphiQLDefaultQuery represents the default GraphQL query showing up in
// GraphiQL editor when it first loads.
type GraphiQLDefaultQuery string

// NewGraphiQL creates GraphiQL with GraphiQLDefaultQuery to uniquely identify
// constructor parameters during dependency injection.
func NewGraphiQL(gqlPath GraphQLPath, defaultQuery GraphiQLDefaultQuery) graphql.GraphiQL {
	return graphql.NewGraphiQL(string(gqlPath), string(defaultQuery))
}
