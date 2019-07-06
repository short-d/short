//+build wireinject

package dep

import (
	"tinyURL/app"
	"tinyURL/fw"
	"tinyURL/modern"

	"github.com/google/wire"
)

func InitGraphQlService(name string) fw.Service {
	wire.Build(
		fw.NewService,
		fw.NewLocalLogger,
		modern.NewGraphGophers,
		app.NewGraphQlApi)
	return fw.Service{}
}
