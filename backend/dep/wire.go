//+build wireinject

package dep

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/cli"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/env"
	"github.com/short-d/app/fw/geo"
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/io"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/network"
	"github.com/short-d/app/fw/rpc"
	"github.com/short-d/app/fw/runtime"
	"github.com/short-d/app/fw/security"
	"github.com/short-d/app/fw/service"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/adapter/facebook"
	"github.com/short-d/short/backend/app/adapter/github"
	"github.com/short-d/short/backend/app/adapter/google"
	"github.com/short-d/short/backend/app/adapter/gqlapi/resolver"
	"github.com/short-d/short/backend/app/adapter/grpcapi"
	"github.com/short-d/short/backend/app/adapter/kgs"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/fw/filesystem"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/shortlink"
	"github.com/short-d/short/backend/app/usecase/sso"
	"github.com/short-d/short/backend/app/usecase/validator"
	"github.com/short-d/short/backend/dep/provider"
	"github.com/short-d/short/backend/tool"
)

var authenticatorSet = wire.NewSet(
	provider.NewJwtGo,
	provider.NewAuthenticator,
)

var authorizerSet = wire.NewSet(
	wire.Bind(new(repository.UserRole), new(sqldb.UserRoleSQL)),
	sqldb.NewUserRoleSQL,
	rbac.NewRBAC,
	authorizer.NewAuthorizer,
)

var observabilitySet = wire.NewSet(
	wire.Bind(new(io.Output), new(io.StdOut)),
	wire.Bind(new(runtime.Runtime), new(runtime.Program)),
	wire.Bind(new(metrics.Metrics), new(metrics.DataDog)),
	wire.Bind(new(analytics.Analytics), new(analytics.Segment)),
	wire.Bind(new(network.Network), new(network.Proxy)),

	io.NewStdOut,
	provider.NewEntryRepositorySwitch,
	provider.NewLogger,
	runtime.NewProgram,
	provider.NewDataDogMetrics,
	provider.NewSegment,
	network.NewProxy,
	request.NewClient,
	request.NewInstrumentationFactory,
	request.NewSSOInstrumentationFactory,
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

var keyGenSet = wire.NewSet(
	wire.Bind(new(keygen.KeyFetcher), new(kgs.RPC)),
	provider.NewKgsRPC,
	provider.NewKeyGenerator,
)

var featureDecisionSet = wire.NewSet(
	wire.Bind(new(repository.FeatureToggle), new(sqldb.FeatureToggleSQL)),
	sqldb.NewFeatureToggleSQL,
	provider.NewFeatureDecisionMakerFactorySwitch,
)

// InjectCommandFactory creates CommandFactory with configured dependencies.
func InjectCommandFactory() cli.CommandFactory {
	wire.Build(
		wire.Bind(new(cli.CommandFactory), new(cli.CobraFactory)),
		cli.NewCobraFactory,
	)
	return cli.CobraFactory{}
}

// InjectDBConnector creates DBConnector with configured dependencies.
func InjectDBConnector() db.Connector {
	wire.Build(
		wire.Bind(new(db.Connector), new(db.PostgresConnector)),
		db.NewPostgresConnector,
	)
	return db.PostgresConnector{}
}

// InjectDBMigrationTool creates DBMigrationTool with configured dependencies.
func InjectDBMigrationTool() db.MigrationTool {
	wire.Build(
		wire.Bind(new(db.MigrationTool), new(db.PostgresMigrationTool)),
		db.NewPostgresMigrationTool,
	)
	return db.PostgresMigrationTool{}
}

// InjectEnv creates Environment with configured dependencies.
func InjectEnv() env.Env {
	wire.Build(
		wire.Bind(new(env.Env), new(env.GoDotEnv)),
		env.NewGoDotEnv,
	)
	return env.GoDotEnv{}
}

// InjectGRPCService creates gRPC service with configured dependencies.
func InjectGRPCService(
	runtime env.Runtime,
	prefix provider.LogPrefix,
	logLevel logger.LogLevel,
	sqlDB *sql.DB,
	securityPolicy security.Policy,
	dataDogAPIKey provider.DataDogAPIKey,
) (service.GRPC, error) {
	wire.Build(
		wire.Bind(new(timer.Timer), new(timer.System)),
		wire.Bind(new(repository.ShortLink), new(sqldb.ShortLinkSQL)),
		wire.Bind(new(rpc.API), new(grpcapi.Short)),
		wire.Bind(new(shortlink.MetaTag), new(shortlink.MetaTagPersist)),

		observabilitySet,

		timer.NewSystem,
		webreq.NewHTTPClient,
		webreq.NewHTTP,
		env.NewDeployment,
		service.NewGRPC,

		sqldb.NewShortLinkSQL,
		shortlink.NewMetaTagPersist,
		grpcapi.NewShort,

		grpcapi.NewMetaTagServer,
	)
	return service.GRPC{}, nil
}

// InjectGraphQLService creates GraphQL service with configured dependencies.
func InjectGraphQLService(
	runtime env.Runtime,
	prefix provider.LogPrefix,
	logLevel logger.LogLevel,
	sqlDB *sql.DB,
	graphqlSchemaPath provider.GraphQLSchemaPath,
	graphqlPath provider.GraphQLPath,
	graphiQLDefaultQuery provider.GraphiQLDefaultQuery,
	secret provider.ReCaptchaSecret,
	jwtSecret provider.JwtSecret,
	bufferSize provider.KeyGenBufferSize,
	kgsRPCConfig provider.KgsRPCConfig,
	tokenValidDuration provider.TokenValidDuration,
	dataDogAPIKey provider.DataDogAPIKey,
	segmentAPIKey provider.SegmentAPIKey,
	ipStackAPIKey provider.IPStackAPIKey,
	googleAPIKey provider.GoogleAPIKey,
) (service.GraphQL, error) {
	wire.Build(
		wire.Bind(new(timer.Timer), new(timer.System)),
		wire.Bind(new(graphql.Handler), new(graphql.GraphGopherHandler)),
		wire.Bind(new(graphql.WebUI), new(graphql.GraphiQL)),

		wire.Bind(new(filesystem.FileSystem), new(filesystem.Local)),
		wire.Bind(new(risk.BlackList), new(google.SafeBrowsing)),
		wire.Bind(new(repository.UserShortLink), new(sqldb.UserShortLinkSQL)),
		wire.Bind(new(repository.ChangeLog), new(sqldb.ChangeLogSQL)),
		wire.Bind(new(repository.UserChangeLog), new(sqldb.UserChangeLogSQL)),
		wire.Bind(new(repository.ShortLink), new(sqldb.ShortLinkSQL)),

		wire.Bind(new(changelog.ChangeLog), new(changelog.Persist)),
		wire.Bind(new(shortlink.Retriever), new(shortlink.RetrieverPersist)),
		wire.Bind(new(shortlink.Creator), new(shortlink.CreatorPersist)),
		wire.Bind(new(shortlink.Updater), new(shortlink.UpdaterPersist)),

		observabilitySet,
		authenticatorSet,
		authorizerSet,
		keyGenSet,

		env.NewDeployment,
		provider.NewGraphQLService,
		graphql.NewGraphGopherHandler,
		provider.NewGraphiQL,
		webreq.NewHTTPClient,
		webreq.NewHTTP,
		timer.NewSystem,

		filesystem.NewLocal,
		resolver.NewResolver,
		provider.NewShortGraphQLAPI,
		provider.NewSafeBrowsing,
		risk.NewDetector,
		provider.NewReCaptchaService,
		provider.NewVerifier,
		sqldb.NewChangeLogSQL,
		sqldb.NewUserChangeLogSQL,
		sqldb.NewShortLinkSQL,
		sqldb.NewUserShortLinkSQL,

		validator.NewLongLink,
		validator.NewCustomAlias,
		changelog.NewPersist,
		shortlink.NewRetrieverPersist,
		shortlink.NewCreatorPersist,
		shortlink.NewUpdaterPersist,
	)
	return service.GraphQL{}, nil
}

// InjectRoutingService creates routing service with configured dependencies.
func InjectRoutingService(
	runtime env.Runtime,
	prefix provider.LogPrefix,
	logLevel logger.LogLevel,
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
	bufferSize provider.KeyGenBufferSize,
	kgsRPCConfig provider.KgsRPCConfig,
	webFrontendURL provider.WebFrontendURL,
	tokenValidDuration provider.TokenValidDuration,
	searchTimeout provider.SearchTimeout,
	swaggerUIDir provider.SwaggerUIDir,
	openAPISpecPath provider.OpenAPISpecPath,
	dataDogAPIKey provider.DataDogAPIKey,
	segmentAPIKey provider.SegmentAPIKey,
	ipStackAPIKey provider.IPStackAPIKey,
) (service.Routing, error) {
	wire.Build(
		wire.Bind(new(timer.Timer), new(timer.System)),
		wire.Bind(new(geo.Geo), new(geo.IPStack)),

		wire.Bind(new(shortlink.Retriever), new(shortlink.RetrieverPersist)),
		wire.Bind(new(repository.UserShortLink), new(sqldb.UserShortLinkSQL)),
		wire.Bind(new(repository.User), new(sqldb.UserSQL)),
		wire.Bind(new(repository.ShortLink), new(sqldb.ShortLinkSQL)),

		observabilitySet,
		authenticatorSet,
		authorizerSet,
		githubAPISet,
		facebookAPISet,
		googleAPISet,
		keyGenSet,
		featureDecisionSet,

		service.NewRouting,
		webreq.NewHTTPClient,
		webreq.NewHTTP,
		graphql.NewClientFactory,
		timer.NewSystem,
		provider.NewIPStack,
		env.NewDeployment,

		provider.NewGithubAccountLinker,
		provider.NewGithubSSO,
		provider.NewFacebookAccountLinker,
		provider.NewFacebookSSO,
		provider.NewGoogleAccountLinker,
		provider.NewGoogleSSO,
		sqldb.NewGithubSSOSql,
		sqldb.NewFacebookSSOSql,
		sqldb.NewGoogleSSOSql,
		sqldb.NewUserSQL,
		sqldb.NewShortLinkSQL,
		sqldb.NewUserShortLinkSQL,

		sso.NewAccountLinkerFactory,
		sso.NewFactory,
		shortlink.NewRetrieverPersist,
		provider.NewSearch,
		provider.NewShortRoutes,
	)
	return service.Routing{}, nil
}

// InjectDataTool creates data tool with configured dependencies.
func InjectDataTool(
	prefix provider.LogPrefix,
	logLevel logger.LogLevel,
	dbConfig db.Config,
	dbConnector db.Connector,
	bufferSize provider.KeyGenBufferSize,
	kgsRPCConfig provider.KgsRPCConfig,
) (tool.Data, error) {
	wire.Build(
		wire.Bind(new(io.Output), new(io.StdOut)),
		wire.Bind(new(timer.Timer), new(timer.System)),
		wire.Bind(new(logger.EntryRepository), new(logger.Local)),

		keyGenSet,

		io.NewStdOut,
		runtime.NewProgram,
		provider.NewLocalEntryRepo,
		provider.NewLogger,
		timer.NewSystem,
		tool.NewData,
	)
	return tool.Data{}, nil
}
