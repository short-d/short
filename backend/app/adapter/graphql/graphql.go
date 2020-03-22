package graphql

import (
	"github.com/short-d/short/app/adapter/graphql/resolver"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/requester"
	"github.com/short-d/short/app/usecase/url"

	"github.com/short-d/app/fw"
)

var _ fw.GraphQLAPI = (*Short)(nil)

// Short represents GraphQL API config
type Short struct {
	resolver *resolver.Resolver
}

// GetSchema retrieves GraphQL schema
func (t Short) GetSchema() string {
	return schema
}

// GetResolver retrieves GraphQL resolver
func (t Short) GetResolver() fw.Resolver {
	return t.resolver
}

// NewShort creates GraphQL API config
func NewShort(
	logger fw.Logger,
	tracer fw.Tracer,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	changeLogRetriever changelog.Retriever,
	changeLogCreator changelog.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) Short {
	r := resolver.NewResolver(
		logger,
		tracer,
		changeLogRetriever,
		changeLogCreator,
		urlRetriever,
		urlCreator,
		requesterVerifier,
		authenticator,
	)
	return Short{
		resolver: &r,
	}
}
