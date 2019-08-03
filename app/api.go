package app

import (
	"database/sql"
	"tinyURL/app/graphql"
	"tinyURL/app/routing"
	"tinyURL/fw"
)

type WwwRoot string

func NewGraphQlApi(logger fw.Logger, tracer fw.Tracer, db *sql.DB) fw.GraphQlApi {
	return graphql.NewTinyUrl(logger, tracer, db)
}

func NewRoutes(logger fw.Logger, tracer fw.Tracer, db *sql.DB, wwwRoot WwwRoot) []fw.Route {
	return routing.NewTinyUrl(logger, tracer, string(wwwRoot), db)
}
