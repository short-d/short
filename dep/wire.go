//+build wireinject

package dep

import (
	"tinyURL/app/graphql"
	"tinyURL/fw"

	"github.com/google/wire"
)

func InitializeApp() fw.App {
	wire.Build(fw.NewApp, fw.NewGraphGophers, graphql.NewTinyUrlGraphQl)
	return fw.App{}
}
