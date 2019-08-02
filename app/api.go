package app

import (
	"tinyURL/app/graphql"
	"tinyURL/app/routing"
	"tinyURL/fw"
)

type WwwRoot string

func NewGraphQlApi(logger fw.Logger, tracer fw.Tracer) fw.GraphQlApi {
	return graphql.NewTinyUrl(logger, tracer)
}

func NewRoutes(logger fw.Logger, tracer fw.Tracer, wwwRoot WwwRoot) []fw.Route {
	return routing.NewTinyUrl(logger, tracer, string(wwwRoot))
}
