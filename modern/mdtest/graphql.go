package mdtest

import (
	"github.com/graph-gophers/graphql-go"
	"tinyURL/fw"
)

func IsGraphQlApiValid(api fw.GraphQlApi) bool {
	_, err := graphql.ParseSchema(api.GetSchema(), api.GetResolver())
	return err == nil
}
