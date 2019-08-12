//+build wireinject

package dep

import (
	"database/sql"
	"github.com/google/wire"
	"short/app/adapter/graphql"
	"short/app/adapter/repo"
	"short/app/adapter/routing"
	"short/app/usecase/keygen"
	"short/app/usecase/url"
	"short/modern/mdgraphql"
	"short/modern/mdlogger"
	"short/modern/mdrouting"
	"short/modern/mdservice"
	"short/modern/mdtracer"
)

func InitGraphQlService(name string, db *sql.DB, graphqlPath mdgraphql.Path) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		mdgraphql.NewGraphGophers,

		repo.NewUrlSql,
		keygen.NewInMemory,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		graphql.NewShort,
	)
	return mdservice.Service{}
}

func InitRoutingService(name string, db *sql.DB, wwwRoot routing.WwwRoot) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		mdrouting.NewBuiltIn,

		repo.NewUrlSql,
		routing.NewShort,
	)
	return mdservice.Service{}
}
