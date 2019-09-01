package inject

import (
	"short/app/adapter/graphql"
	"short/app/usecase/auth"
	"short/app/usecase/requester"
	"short/app/usecase/url"

	"short/modern/mdgraphql"

	"short/fw"
)

type GraphQlPath string

func GraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlAPI) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}

func ShortGraphQlAPI(
	logger fw.Logger,
	tracer fw.Tracer,
	urlRetriever url.Retriever,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) fw.GraphQlAPI {
	return graphql.NewShort(
		logger,
		tracer,
		urlRetriever,
		urlCreator,
		requesterVerifier,
		authenticator,
	)
}
