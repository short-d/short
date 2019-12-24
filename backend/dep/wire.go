//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/db"
	"short/app/adapter/facebook"
	"short/app/adapter/github"
	"short/app/adapter/google"
	"short/app/adapter/graphql"
	"short/app/adapter/kgs"
	"short/app/usecase/account"
	"short/app/usecase/keygen"
	"short/app/usecase/repository"
	"short/app/usecase/requester"
	"short/app/usecase/service"
	"short/app/usecase/url"
	"short/app/usecase/validator"
	"short/dep/provider"
	"time"

	"github.com/byliuyang/app/fw"
	"github.com/byliuyang/app/modern/mdcli"
	"github.com/byliuyang/app/modern/mddb"
	"github.com/byliuyang/app/modern/mdenv"
	"github.com/byliuyang/app/modern/mdhttp"
	"github.com/byliuyang/app/modern/mdlogger"
	"github.com/byliuyang/app/modern/mdrequest"
	"github.com/byliuyang/app/modern/mdrouting"
	"github.com/byliuyang/app/modern/mdservice"
	"github.com/byliuyang/app/modern/mdtimer"
	"github.com/byliuyang/app/modern/mdtracer"
	"github.com/google/wire"
)

const oneDay = 24 * time.Hour

var authSet = wire.NewSet(
	provider.NewJwtGo,

	wire.Value(provider.TokenValidDuration(oneDay)),
	provider.NewAuthenticator,
)

var observabilitySet = wire.NewSet(
	mdlogger.NewLocal,
	mdtracer.NewLocal,
)

var githubAPISet = wire.NewSet(
	provider.NewGithubIdentityProvider,
	github.NewAccount,
	github.NewAPI,
)

var facebookAPISet = wire.NewSet(
	provider.NewFacebookIdentityProvider,
	facebook.NewAccount,
	facebook.NewAPI,
)

var googleAPISet = wire.NewSet(
	provider.NewGoogleIdentityProvider,
	google.NewAccount,
	google.NewAPI,
)

// InjectCommandFactory creates CommandFactory with configured dependencies.
func InjectCommandFactory() fw.CommandFactory {
	wire.Build(
		wire.Bind(new(fw.CommandFactory), new(mdcli.CobraFactory)),
		mdcli.NewCobraFactory,
	)
	return mdcli.CobraFactory{}
}

// InjectDBConnector creates DBConnector with configured dependencies.
func InjectDBConnector() fw.DBConnector {
	wire.Build(
		wire.Bind(new(fw.DBConnector), new(mddb.PostgresConnector)),
		mddb.NewPostgresConnector,
	)
	return mddb.PostgresConnector{}
}

// InjectDBMigrationTool creates DBMigrationTool with configured dependencies.
func InjectDBMigrationTool() fw.DBMigrationTool {
	wire.Build(
		wire.Bind(new(fw.DBMigrationTool), new(mddb.PostgresMigrationTool)),
		mddb.NewPostgresMigrationTool,
	)
	return mddb.PostgresMigrationTool{}
}

// InjectEnvironment creates Environment with configured dependencies.
func InjectEnvironment() fw.Environment {
	wire.Build(
		wire.Bind(new(fw.Environment), new(mdenv.GoDotEnv)),
		mdenv.NewGoDotEnv,
	)
	return mdenv.GoDotEnv{}
}

// InjectGraphQlService creates GraphQL service with configured dependencies.
func InjectGraphQlService(
	name string,
	sqlDB *sql.DB,
	graphqlPath provider.GraphQlPath,
	secret provider.ReCaptchaSecret,
	jwtSecret provider.JwtSecret,
	bufferSize provider.KeyGenBufferSize,
	kgsRPCConfig provider.KgsRPCConfig,
) (mdservice.Service, error) {
	wire.Build(
		wire.Bind(new(fw.GraphQlAPI), new(graphql.Short)),
		wire.Bind(new(url.Retriever), new(url.RetrieverPersist)),
		wire.Bind(new(url.Creator), new(url.CreatorPersist)),
		wire.Bind(new(repository.UserURLRelation), new(db.UserURLRelationSQL)),
		wire.Bind(new(repository.URL), new(*db.URLSql)),
		wire.Bind(new(keygen.KeyGenerator), new(keygen.Remote)),
		wire.Bind(new(service.KeyFetcher), new(kgs.RPC)),
		wire.Bind(new(fw.HTTPRequest), new(mdrequest.HTTP)),

		observabilitySet,
		authSet,

		mdservice.New,
		provider.NewGraphGophers,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdtimer.NewTimer,

		db.NewURLSql,
		db.NewUserURLRelationSQL,
		provider.NewRemote,
		validator.NewLongLink,
		validator.NewCustomAlias,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		provider.NewKgsRPC,
		provider.NewReCaptchaService,
		requester.NewVerifier,
		graphql.NewShort,
	)
	return mdservice.Service{}, nil
}

// InjectRoutingService creates routing service with configured dependencies.
func InjectRoutingService(
	name string,
	sqlDB *sql.DB,
	githubClientID provider.GithubClientID,
	githubClientSecret provider.GithubClientSecret,
	facebookClientID provider.FacebookClientID,
	facebookClientSecret provider.FacebookClientSecret,
	facebookRedirectURI provider.FacebookRedirectURI,
	googleClientID provider.GoogleClientID,
	googleClientSecret provider.GoogleClientSecret,
	googleRedirectURI provider.GoogleRedirectURI,
	jwtSecret provider.JwtSecret,
	webFrontendURL provider.WebFrontendURL,
) mdservice.Service {
	wire.Build(
		wire.Bind(new(url.Retriever), new(url.RetrieverPersist)),
		wire.Bind(new(repository.User), new(*(db.UserSQL))),
		wire.Bind(new(repository.URL), new(*db.URLSql)),
		wire.Bind(new(fw.HTTPRequest), new(mdrequest.HTTP)),

		observabilitySet,
		authSet,
		githubAPISet,
		facebookAPISet,
		googleAPISet,

		mdservice.New,
		mdrouting.NewBuiltIn,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdrequest.NewGraphQl,
		mdtimer.NewTimer,

		db.NewUserSQL,
		db.NewURLSql,
		url.NewRetrieverPersist,
		account.NewProvider,
		provider.NewShortRoutes,
	)
	return mdservice.Service{}
}
