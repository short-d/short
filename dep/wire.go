//+build wireinject

package dep

import (
	"tinyURL/app"
	"tinyURL/modern"

	"github.com/google/wire"
)

func InitGraphQlService(name string, graphqlPath modern.GraphQlPath) modern.Service {
	wire.Build(
		modern.NewService,
		modern.NewLocalLogger,
		modern.NewLocalTracer,
		modern.NewGraphGophers,
		app.NewGraphQlApi)
	return modern.Service{}
}
