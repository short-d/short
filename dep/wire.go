//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/graphql"
	"short/app/adapter/recaptcha"
	"short/app/adapter/repo"
	"short/app/adapter/request"
	"short/app/adapter/routing"
	"short/app/usecase/keygen"
	"short/app/usecase/requester"
	"short/app/usecase/service"
	"short/app/usecase/url"
	"short/fw"
	"short/modern/mddb"
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
	graphqlPath GraphQlPath,
	secret ReCaptchaSecret,
) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		NewGraphGophers,
		mdhttp.NewClient,

		repo.NewUrlSql,
		keygen.NewInMemory,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		request.NewHttp,
		NewReCaptchaService,
		requester.NewVerifier,
		graphql.NewShort,
	)
	return mdservice.Service{}
}

func InitRoutingService(
	name string,
	db *sql.DB,
	wwwRoot WwwRoot,
	githubClientId GithubClientId,
	githubClientSecret GithubClientSecret,
) mdservice.Service {
	wire.Build(
		mdservice.New,
		mdlogger.NewLocal,
		mdtracer.NewLocal,
		mdrouting.NewBuiltIn,
		mdhttp.NewClient,

		repo.NewUrlSql,
		url.NewRetrieverPersist,
		request.NewHttp,
		NewShortRoutes,
	)
	return mdservice.Service{}
}

type GraphQlPath string

func NewGraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlApi) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}

type ReCaptchaSecret string

func NewReCaptchaService(req request.Http, secret ReCaptchaSecret) service.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

type WwwRoot string
type GithubClientId string
type GithubClientSecret string

func NewShortRoutes(
	logger fw.Logger,
	tracer fw.Tracer,
	wwwRoot WwwRoot,
	urlRetriever url.Retriever,
	req request.Http,
	githubClientId GithubClientId,
	githubClientSecret GithubClientSecret,
) []fw.Route {
	return routing.NewShort(
		logger,
		tracer,
		string(wwwRoot),
		urlRetriever,
		req,
		string(githubClientId),
		string(githubClientSecret),
	)
}

type serviceLauncher func(db *sql.DB)

func InitDB(
	host string,
	port int,
	user string,
	password string,
	dbName string,
	migrationRoot string,
	serviceLauncher serviceLauncher,
) {
	db, err := mddb.NewPostgresDb(host, port, user, password, dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = mddb.MigratePostgres(db, migrationRoot)
	if err != nil {
		panic(err)
	}

	serviceLauncher(db)
}
