//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/graphql"
	"short/app/adapter/repo"
	"short/app/adapter/request"
	"short/app/adapter/routing"
	"short/app/adapter/service"
	"short/app/usecase/keygen"
	"short/app/usecase/recaptcha"
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
	secret service.ReCaptchaSecret,
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
		request.NewHttp,
		service.NewReCaptcha,
		recaptcha.NewVerifier,
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
