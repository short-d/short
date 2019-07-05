package graphql

import (
	"tinyURL/app/graphql/resolver"
	"tinyURL/fw"
)

type TinyUrlGraphQl struct {
	resolver *resolver.Resolver
}

func (t TinyUrlGraphQl) GetSchema() string {
	return schema
}

func (t TinyUrlGraphQl) GetResolver() interface{} {
	return t.resolver
}

func NewTinyUrlGraphQl() fw.GraphQl {
	return &TinyUrlGraphQl{
		resolver: &resolver.Resolver{},
	}
}
