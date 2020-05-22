package gqlapi

import (
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/short/backend/app/adapter/gqlapi/resolver"
)

// NewShort creates GraphQL API config
func NewShort(
	r resolver.Resolver,
) graphql.API {
	return graphql.API{
		Schema:   schema,
		Resolver: &r,
	}
}
