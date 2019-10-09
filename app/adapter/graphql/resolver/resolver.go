package resolver

import (
	"short/app/usecase/auth"
	"short/app/usecase/requester"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

type Resolver struct {
	Query
	Mutation
}

func NewResolver(
	logger fw.Logger,
	tracer fw.Tracer,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) Resolver {
	return Resolver{
		Query: NewQuery(logger, tracer, urlRetriever, authenticator),
		Mutation: NewMutation(
			logger,
			tracer,
			urlCreator,
			requesterVerifier,
			authenticator,
		),
	}
}
