//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/graphql"
	"short/app/adapter/repo"
	"short/app/adapter/routing"
	"short/app/usecase/captcha"
	"short/app/usecase/keygen"
	"short/app/usecase/url"
	"short/modern/mdgraphql"
	"short/modern/mdhttp"
	"short/modern/mdlogger"
	"short/modern/mdrouting"
	"short/modern/mdservice"
	"short/modern/mdtracer"

	"github.com/google/wire"
)

func InitGraphQlService(
	name string,
	db *sql.DB,
	graphqlPath mdgraphql.Path,
	secret captcha.RecaptchaV3Secret,
) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		mdgraphql.NewGraphGophers,
		mdhttp.NewClient,

		repo.NewUrlSql,
		keygen.NewInMemory,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		captcha.NewRecaptchaV3Verifier,
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
