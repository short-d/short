package graphql

import (
	"short/app/adapter/graphql/resolver"
	"short/app/usecase/recaptcha"
	"short/app/usecase/url"
	"short/fw"
)

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
	captchaVerifier recaptcha.Verifier,
) fw.GraphQlApi {
	r := resolver.NewResolver(
		logger,
		tracer,
		urlRetriever,
		urlCreator,
		captchaVerifier,
	)
	return &Short{
		resolver: &r,
	}
}
