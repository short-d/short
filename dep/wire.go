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
		mdrequest.NewHTTP,

		inject.URLRepoSQL,
		keygen.NewInMemory,
		inject.URLRetrieverPersist,
		inject.URLCreatorPersist,
		inject.ReCaptchaService,
		requester.NewVerifier,
		inject.ShortGraphQlAPI,
	)
	return mdservice.Service{}
}

func InitRoutingService(
	name string,
	db *sql.DB,
	wwwRoot inject.WwwRoot,
	githubClientID inject.GithubClientID,
	githubClientSecret inject.GithubClientSecret,
	jwtSecret inject.JwtSecret,
) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		mdrouting.NewBuiltIn,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdrequest.NewGraphQl,
		mdtimer.NewTimer,
		inject.JwtGo,

		inject.URLRepoSQL,
		inject.URLRetrieverPersist,
		inject.GithubOAuth,
		account.NewGithub,
		inject.Authenticator,
		inject.ShortRoutes,
	)
	return mdservice.Service{}
}
