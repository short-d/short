//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/account"
	"short/app/adapter/graphql"
	"short/app/adapter/repo"
	"short/app/usecase/keygen"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"short/dep/inject"
	"short/modern/mdhttp"
	"short/modern/mdlogger"
	"short/modern/mdrequest"
	"short/modern/mdrouting"
	"short/modern/mdservice"
	"short/modern/mdtimer"
	"short/modern/mdtracer"

	"github.com/google/wire"
)

func InitGraphQlService(
	name string,
	db *sql.DB,
	graphqlPath inject.GraphQlPath,
	secret inject.ReCaptchaSecret,
) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		inject.GraphGophers,
		mdhttp.NewClient,
		mdrequest.NewHttp,

		repo.NewUrlSql,
		keygen.NewInMemory,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		inject.ReCaptchaService,
		requester.NewVerifier,
		graphql.NewShort,
	)
	return mdservice.Service{}
}

func InitRoutingService(
	name string,
	db *sql.DB,
	wwwRoot inject.WwwRoot,
	githubClientId inject.GithubClientId,
	githubClientSecret inject.GithubClientSecret,
	jwtSecret inject.JwtSecret,
) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		mdrouting.NewBuiltIn,
		mdhttp.NewClient,
		mdrequest.NewHttp,
		mdrequest.NewGraphQl,
		mdtimer.NewTimer,
		inject.JwtGo,

		repo.NewUrlSql,
		url.NewRetrieverPersist,
		inject.GithubOAuth,
		account.NewGithub,
		inject.Authenticator,
		inject.ShortRoutes,
	)
	return mdservice.Service{}
}
