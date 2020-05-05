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
func NewGraphQLService(gqlPath GraphQLPath, handler graphql.Handler, logger logger.Logger) service.GraphQL {
	return service.NewGraphQL(logger, string(gqlPath), handler)
}
