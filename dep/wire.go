//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/account"
	"short/app/usecase/keygen"
	"short/app/usecase/requester"
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

		inject.UrlRepoSql,
		keygen.NewInMemory,
		inject.UrlRetrieverPersist,
		inject.UrlCreatorPersist,
		inject.ReCaptchaService,
		requester.NewVerifier,
		inject.ShortGraphQlApi,
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

		inject.UrlRepoSql,
		inject.UrlRetrieverPersist,
		inject.GithubOAuth,
		account.NewGithub,
		inject.Authenticator,
		inject.ShortRoutes,
	)
	return mdservice.Service{}
}
