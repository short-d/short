package graphql

import (
	"database/sql"
	"short/app/graphql/resolver"
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

func NewShort(logger fw.Logger, tracer fw.Tracer, db *sql.DB) fw.GraphQlApi {
	r := resolver.NewResolver(logger, tracer, db)
	return &Short{
		resolver: &r,
	}
}
