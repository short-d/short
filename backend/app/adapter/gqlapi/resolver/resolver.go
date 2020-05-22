package resolver

import (
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/requester"
	"github.com/short-d/short/backend/app/usecase/url"
)

// Resolver contains GraphQL request handlers.
type Resolver struct {
	Query
	Mutation
}

// NewResolver creates a new GraphQL resolver.
func NewResolver(
	logger logger.Logger,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	changeLog changelog.ChangeLog,
	requesterVerifier requester.Verifier,
	authenticator authenticator.Authenticator,
) Resolver {
	return Resolver{
		Query: newQuery(logger, authenticator, changeLog, urlRetriever),
		Mutation: newMutation(
			logger,
			changeLog,
			urlCreator,
			requesterVerifier,
			authenticator,
		),
	}
}
