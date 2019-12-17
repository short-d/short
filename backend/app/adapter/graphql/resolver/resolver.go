package resolver

import (
	"short/app/usecase/auth"
	"short/app/usecase/requester"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

// Resolver represents GraphQL request handlers.
type Resolver struct {
	Query
	Mutation
}

// NewResolver creates a new GraphQL resolver.
func NewResolver(
	logger fw.Logger,
	tracer fw.Tracer,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) Resolver {
	return Resolver{
		Query: newQuery(logger, tracer, authenticator, urlRetriever),
		Mutation: newMutation(
			logger,
			tracer,
			urlCreator,
			requesterVerifier,
			authenticator,
		),
	}
}
