package mdtest

import (
	"short/fw"

	"github.com/graph-gophers/graphql-go"
)

func IsGraphQlAPIValid(api fw.GraphQlAPI) bool {
	_, err := graphql.ParseSchema(api.GetSchema(), api.GetResolver())
	return err == nil
}
