// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

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
)

// Injectors from wire.go:

func InitGraphQlService(name string, db *sql.DB, graphqlPath GraphQlPath, secret ReCaptchaSecret) mdservice.Service {
	logger := mdlogger.NewLocal()
	tracer := mdtracer.NewLocal()
	repoUrl := repo.NewUrlSql(db)
	retriever := url.NewRetrieverPersist(repoUrl)
	keyGenerator := keygen.NewInMemory()
	creator := url.NewCreatorPersist(repoUrl, keyGenerator)
	client := mdhttp.NewClient()
	http := request.NewHttp(client)
	reCaptcha := NewReCaptchaService(http, secret)
	verifier := requester.NewVerifier(reCaptcha)
	graphQlApi := graphql.NewShort(logger, tracer, retriever, creator, verifier)
	server := NewGraphGophers(graphqlPath, logger, tracer, graphQlApi)
	service := mdservice.New(name, server, logger)
	return service
}

func InitRoutingService(name string, db *sql.DB, wwwRoot WwwRoot, githubClientId GithubClientId, githubClientSecret GithubClientSecret) mdservice.Service {
	logger := mdlogger.NewLocal()
	tracer := mdtracer.NewLocal()
	repoUrl := repo.NewUrlSql(db)
	retriever := url.NewRetrieverPersist(repoUrl)
	client := mdhttp.NewClient()
	http := request.NewHttp(client)
	v := NewShortRoutes(logger, tracer, wwwRoot, retriever, http, githubClientId, githubClientSecret)
	server := mdrouting.NewBuiltIn(logger, tracer, v)
	service := mdservice.New(name, server, logger)
	return service
}

// wire.go:

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
	migrationRoot string, serviceLauncher2 serviceLauncher,
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
	serviceLauncher2(db)
}
