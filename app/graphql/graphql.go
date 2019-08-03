package graphql

import (
	"database/sql"
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

func NewTinyUrl(logger fw.Logger, tracer fw.Tracer, db *sql.DB) fw.GraphQlApi {
	r := resolver.NewResolver(logger, tracer, db)
	return &TinyUrl{
		resolver: &r,
	}
}
