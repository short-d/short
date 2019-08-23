package graphql

import (
	"short/app/adapter/graphql/resolver"
	"short/app/usecase/auth"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"short/fw"
)

var _ fw.GraphQlAPI = (*Short)(nil)

type Short struct {
	resolver *resolver.Resolver
}

func (t Short) GetSchema() string {
	return schema
}

func (t Short) GetResolver() interface{} {
	return t.resolver
}

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
