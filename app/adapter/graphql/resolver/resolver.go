package resolver

import (
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"short/fw"
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
) Resolver {
	return Resolver{
		Query: NewQuery(logger, tracer, urlRetriever),
		Mutation: NewMutation(
			logger,
			tracer,
			urlCreator,
			requesterVerifier,
		),
	}
}
