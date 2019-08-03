//+build wireinject

package dep

import (
	"database/sql"
	"tinyURL/app"
	"tinyURL/modern"

	"github.com/google/wire"
)

func InitGraphQlService(name string, db *sql.DB, graphqlPath modern.GraphQlPath) modern.Service {
	wire.Build(
		modern.NewService,
		modern.NewLocalLogger,
		modern.NewLocalTracer,
		modern.NewGraphGophers,

		app.NewGraphQlApi)
	return modern.Service{}
}

func InitRoutingService(name string, db *sql.DB, wwwRoot app.WwwRoot) modern.Service {
	wire.Build(
		modern.NewService,
		modern.NewLocalLogger,
		modern.NewLocalTracer,
		modern.NewCustomRouting,
		app.NewRoutes)
	return modern.Service{}
}
