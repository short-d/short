//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/account"
	"short/app/adapter/graphql"
	"short/app/adapter/oauth"
	"short/app/adapter/recaptcha"
	"short/app/adapter/repo"
	"short/app/adapter/routing"
	"short/app/usecase/auth"
	"short/app/usecase/keygen"
	"short/app/usecase/requester"
	"short/app/usecase/service"
	"short/app/usecase/url"
	"short/fw"
	"short/modern/mdcrypto"
	"short/modern/mddb"
	"short/modern/mdgraphql"
	"short/modern/mdhttp"
	"short/modern/mdlogger"
	"short/modern/mdrequest"
	"short/modern/mdrouting"
	"short/modern/mdservice"
	"short/modern/mdtimer"
	"short/modern/mdtracer"
	"time"

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
		mdrequest.NewHttp,

		repo.NewUrlSql,
		keygen.NewInMemory,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
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
	jwtSecret JwtSecret,
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
		NewJwtGo,

		repo.NewUrlSql,
		url.NewRetrieverPersist,
		NewGithubOAuth,
		account.NewGithub,
		NewAuthenticator,
		NewShortRoutes,
	)
	return mdservice.Service{}
}

type GraphQlPath string

func NewGraphGophers(graphqlPath GraphQlPath, logger fw.Logger, tracer fw.Tracer, g fw.GraphQlApi) fw.Server {
	return mdgraphql.NewGraphGophers(string(graphqlPath), logger, tracer, g)
}

type ReCaptchaSecret string

func NewReCaptchaService(req fw.HttpRequest, secret ReCaptchaSecret) service.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

type GithubClientId string
type GithubClientSecret string

func NewGithubOAuth(
	req fw.HttpRequest,
	clientId GithubClientId,
	clientSecret GithubClientSecret,
) oauth.Github {
	return oauth.NewGithub(req, string(clientId), string(clientSecret))
}

type JwtSecret string

func NewJwtGo(secret JwtSecret) fw.CryptoTokenizer {
	return mdcrypto.NewJwtGo(string(secret))
}

const oneDay = 24 * time.Hour
const oneWeek = 7 * oneDay

func NewAuthenticator(tokenizer fw.CryptoTokenizer, timer fw.Timer) auth.Authenticator {
	return auth.NewAuthenticator(tokenizer, timer, oneWeek)
}

type WwwRoot string

func NewShortRoutes(
	logger fw.Logger,
	tracer fw.Tracer,
	wwwRoot WwwRoot,
	timer fw.Timer,
	urlRetriever url.Retriever,
	githubOAuth oauth.Github,
	githubAccount account.Github,
	authenticator auth.Authenticator,
) []fw.Route {
	return routing.NewShort(
		logger,
		tracer,
		string(wwwRoot),
		timer,
		urlRetriever,
		githubOAuth,
		githubAccount,
		authenticator,
	)
}

type ServiceLauncher func(db *sql.DB)

func InitDB(
	host string,
	port int,
	user string,
	password string,
	dbName string,
	migrationRoot string,
	serviceLauncher ServiceLauncher,
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
