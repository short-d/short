package app

import (
	"database/sql"
	"short/app/graphql"
	"short/app/routing"
	"short/fw"
)

type WwwRoot string

func NewGraphQlApi(logger fw.Logger, tracer fw.Tracer, db *sql.DB) fw.GraphQlApi {
	return graphql.NewShort(logger, tracer, db)
}

func NewRoutes(logger fw.Logger, tracer fw.Tracer, db *sql.DB, wwwRoot WwwRoot) []fw.Route {
	return routing.NewShort(logger, tracer, string(wwwRoot), db)
}
