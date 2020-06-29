package gqlapi

import (
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/short/backend/app/adapter/gqlapi/resolver"
	"github.com/short-d/short/backend/app/fw/filesystem"
)

// NewShort creates GraphQL API config
func NewShort(
	schemaPath string,
	fileSystem filesystem.FileSystem,
	r resolver.Resolver,
) (graphql.API, error) {
	buf, err := fileSystem.ReadFile(schemaPath)
	if err != nil {
		return graphql.API{}, err
	}
	return graphql.API{
		Schema:   string(buf),
		Resolver: &r,
	}, nil
}
