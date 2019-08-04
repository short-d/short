package mdtest

import (
	"short/fw"

	"github.com/graph-gophers/graphql-go"
)

func IsGraphQlApiValid(api fw.GraphQlApi) bool {
	_, err := graphql.ParseSchema(api.GetSchema(), api.GetResolver())
	return err == nil
}
