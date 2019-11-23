package graphql

import (
	"short/app/adapter/graphql/resolver"
	"short/app/usecase/auth"
	"short/app/usecase/requester"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

var _ fw.GraphQlAPI = (*Short)(nil)

// Short represents GraphQL API config
type Short struct {
	resolver *resolver.Resolver
}

// GetSchema retrieves GraphQL schema
func (t Short) GetSchema() string {
	return schema
}

// GetResolver retrieves GraphQL resolver
func (t Short) GetResolver() interface{} {
	return t.resolver
}

// NewShort creates GraphQL API config
func NewShort(
	logger fw.Logger,
	tracer fw.Tracer,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) Short {
	r := resolver.NewResolver(
		logger,
		tracer,
		urlRetriever,
		urlCreator,
		requesterVerifier,
		authenticator,
	)
	return Short{
		resolver: &r,
	}
}
