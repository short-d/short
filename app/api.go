package app

import (
	"tinyURL/app/graphql"
	"tinyURL/fw"
)

func NewGraphQlApi(logger fw.Logger, tracer fw.Tracer) fw.GraphQlApi {
	return graphql.NewTinyUrl(logger, tracer)
}
