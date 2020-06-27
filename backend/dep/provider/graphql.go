package provider

import (
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/service"
)

// GraphQLPath represents the path for GraphQL APIs.
type GraphQLPath string

// NewGraphQLService creates GraphQL service with GraphQlPath to uniquely
// identify graphqlPath during dependency injection.
func NewGraphQLService(
	gqlPath GraphQLPath,
	handler graphql.Handler,
	webUI graphql.WebUI,
	logger logger.Logger,
) service.GraphQL {
	return service.NewGraphQL(logger, string(gqlPath), handler, webUI)
}

type GraphiQLDefaultQuery string

func NewGraphiQL(gqlPath GraphQLPath, defaultQuery GraphiQLDefaultQuery) graphql.GraphiQL {
	return graphql.NewGraphiQL(string(gqlPath), string(defaultQuery))
}
