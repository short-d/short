package app

import (
	"tinyURL/app/graphql"
	"tinyURL/fw"
)

func NewGraphQlApi() fw.GraphQlApi {
	return graphql.NewTinyUrl()
}
