package gqlapi

import (
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/adapter/gqlapi/resolver"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/requester"
	"github.com/short-d/short/backend/app/usecase/url"
)

var _ graphql.API = (*Short)(nil)

// Short represents GraphQL API config
type Short struct {
	resolver *resolver.Resolver
}

// GetSchema retrieves GraphQL schema
func (t Short) GetSchema() string {
	return schema
}

// GetResolver retrieves GraphQL resolver
func (t Short) GetResolver() graphql.Resolver {
	return t.resolver
}

// NewShort creates GraphQL API config
func NewShort(
	logger logger.Logger,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	changeLog changelog.ChangeLog,
	requesterVerifier requester.Verifier,
	authenticator authenticator.Authenticator,
) Short {
	r := resolver.NewResolver(
		logger,
		changeLog,
		urlRetriever,
		urlCreator,
		requesterVerifier,
		authenticator,
	)
	return Short{
		resolver: &r,
	}
}
