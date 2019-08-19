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
	"short/dep/new"
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
	graphqlPath new.GraphQlPath,
	secret new.ReCaptchaSecret,
) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		new.GraphGophers,
		mdhttp.NewClient,
		mdrequest.NewHttp,

		repo.NewUrlSql,
		keygen.NewInMemory,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		new.ReCaptchaService,
		requester.NewVerifier,
		graphql.NewShort,
	)
	return mdservice.Service{}
}

func InitRoutingService(
	name string,
	db *sql.DB,
	wwwRoot new.WwwRoot,
	githubClientId new.GithubClientId,
	githubClientSecret new.GithubClientSecret,
	jwtSecret new.JwtSecret,
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
		new.JwtGo,

		repo.NewUrlSql,
		url.NewRetrieverPersist,
		new.GithubOAuth,
		account.NewGithub,
		new.Authenticator,
		new.ShortRoutes,
	)
	return mdservice.Service{}
}
