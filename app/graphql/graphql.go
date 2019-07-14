package graphql

import (
	"tinyURL/app/entity"
	"tinyURL/app/graphql/resolver"
	"tinyURL/app/repo"
	"tinyURL/fw"
)

type TinyUrl struct {
	resolver *resolver.Resolver
}

func (t TinyUrl) GetSchema() string {
	return schema
}

func (t TinyUrl) GetResolver() interface{} {
	return t.resolver
}

func NewTinyUrl(logger fw.Logger, tracer fw.Tracer) fw.GraphQlApi {
	urlRepo := repo.NewUrlFake(map[string]entity.Url{
		"220uFicCJj": {
			Alias:       "220uFicCJj",
			OriginalUrl: "http://www.google.com",
		},
		"yDOBcj5HIPbUAsw": {
			Alias:       "yDOBcj5HIPbUAsw",
			OriginalUrl: "http://www.facebook.com",
		},
	})

	r := resolver.NewResolver(logger, tracer, urlRepo)
	return &TinyUrl{
		resolver: &r,
	}
}
