package graphql

import (
	"tinyURL/app/graphql/resolver"
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

func NewTinyUrl() fw.GraphQlApi {
	return &TinyUrl{
		resolver: &resolver.Resolver{},
	}
}
