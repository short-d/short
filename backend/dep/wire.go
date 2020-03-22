//+build wireinject

package dep

import (
	"database/sql"
	"time"

	"github.com/short-d/short/app/usecase/changelog"

	"github.com/google/wire"
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdcli"
	"github.com/short-d/app/modern/mddb"
	"github.com/short-d/app/modern/mdenv"
	"github.com/short-d/app/modern/mdhttp"
	"github.com/short-d/app/modern/mdio"
	"github.com/short-d/app/modern/mdlogger"
	"github.com/short-d/app/modern/mdrequest"
	"github.com/short-d/app/modern/mdrouting"
	"github.com/short-d/app/modern/mdruntime"
	"github.com/short-d/app/modern/mdservice"
	"github.com/short-d/app/modern/mdtimer"
	"github.com/short-d/app/modern/mdtracer"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/adapter/facebook"
	"github.com/short-d/short/app/adapter/github"
	"github.com/short-d/short/app/adapter/google"
	"github.com/short-d/short/app/adapter/graphql"
	"github.com/short-d/short/app/adapter/kgs"
	"github.com/short-d/short/app/usecase/account"
	"github.com/short-d/short/app/usecase/repository"
	"github.com/short-d/short/app/usecase/requester"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/url"
	"github.com/short-d/short/app/usecase/validator"
	"github.com/short-d/short/dep/provider"
)

const oneDay = 24 * time.Hour

var authSet = wire.NewSet(
	provider.NewJwtGo,

	wire.Value(provider.TokenValidDuration(oneDay)),
	provider.NewAuthenticator,
)

var observabilitySet = wire.NewSet(
	wire.Bind(new(fw.Logger), new(mdlogger.Local)),
	provider.NewLocalLogger,
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

// InjectGraphQLService creates GraphQL service with configured dependencies.
func InjectGraphQLService(
	name string,
	prefix provider.LogPrefix,
	logLevel fw.LogLevel,
	sqlDB *sql.DB,
	graphqlPath provider.GraphQlPath,
	secret provider.ReCaptchaSecret,
	jwtSecret provider.JwtSecret,
	bufferSize provider.KeyGenBufferSize,
	kgsRPCConfig provider.KgsRPCConfig,
) (mdservice.Service, error) {
	wire.Build(
		wire.Bind(new(fw.StdOut), new(mdio.StdOut)),
		wire.Bind(new(fw.ProgramRuntime), new(mdruntime.BuildIn)),
		wire.Bind(new(fw.GraphQLAPI), new(graphql.Short)),
		wire.Bind(new(changelog.Retriever), new(changelog.RetrieverPersist)),
		wire.Bind(new(changelog.Creator), new(changelog.CreatorPersist)),
		wire.Bind(new(url.Retriever), new(url.RetrieverPersist)),
		wire.Bind(new(url.Creator), new(url.CreatorPersist)),
		wire.Bind(new(repository.UserURLRelation), new(db.UserURLRelationSQL)),
		wire.Bind(new(repository.Changelog), new(*db.ChangeLogSql)),
		wire.Bind(new(repository.URL), new(*db.URLSql)),
		wire.Bind(new(service.KeyFetcher), new(kgs.RPC)),
		wire.Bind(new(fw.HTTPRequest), new(mdrequest.HTTP)),

		observabilitySet,
		authSet,

		mdio.NewBuildInStdOut,
		mdruntime.NewBuildIn,
		mdservice.New,
		provider.NewGraphGophers,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdtimer.NewTimer,

		db.NewChangeLogSql,
		db.NewURLSql,
		db.NewUserURLRelationSQL,
		provider.NewKeyGenerator,
		validator.NewLongLink,
		validator.NewCustomAlias,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		changelog.NewRetrieverPersist,
		changelog.NewCreatorPersist,
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
	prefix provider.LogPrefix,
	logLevel fw.LogLevel,
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
		wire.Bind(new(fw.StdOut), new(mdio.StdOut)),
		wire.Bind(new(fw.ProgramRuntime), new(mdruntime.BuildIn)),
		wire.Bind(new(url.Retriever), new(url.RetrieverPersist)),
		wire.Bind(new(repository.User), new(*(db.UserSQL))),
		wire.Bind(new(repository.URL), new(*db.URLSql)),
		wire.Bind(new(fw.HTTPRequest), new(mdrequest.HTTP)),
		wire.Bind(new(fw.GraphQlRequest), new(mdrequest.GraphQL)),

		observabilitySet,
		authSet,
		githubAPISet,
		facebookAPISet,
		googleAPISet,

		mdio.NewBuildInStdOut,
		mdruntime.NewBuildIn,
		mdservice.New,
		mdrouting.NewBuiltIn,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdrequest.NewGraphQL,
		mdtimer.NewTimer,

		db.NewUserSQL,
		db.NewURLSql,
		url.NewRetrieverPersist,
		account.NewProvider,
		provider.NewShortRoutes,
	)
	return mdservice.Service{}
}
