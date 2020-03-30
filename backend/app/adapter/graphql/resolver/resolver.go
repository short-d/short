package resolver

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/requester"
	"github.com/short-d/short/app/usecase/url"
)

// Resolver contains GraphQL request handlers.
type Resolver struct {
	Query
	Mutation
}

// NewResolver creates a new GraphQL resolver.
func NewResolver(
	logger fw.Logger,
	tracer fw.Tracer,
	timer fw.Timer,
	changeLog changelog.ChangeLog,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) Resolver {
	return Resolver{
		Query: newQuery(logger, tracer, timer, authenticator, changeLog, urlRetriever),
		Mutation: newMutation(
			logger,
			tracer,
			changeLog,
			urlCreator,
			requesterVerifier,
			authenticator,
		),
	}
}
