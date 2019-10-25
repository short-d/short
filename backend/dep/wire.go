//+build wireinject

package dep

import (
	"database/sql"
	"short/app/adapter/db"
	"short/app/adapter/github"
	"short/app/adapter/graphql"
	"short/app/adapter/kgs"
	"short/app/usecase/account"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"short/app/usecase/requester"
	"short/app/usecase/service"
	"short/app/usecase/url"
	"short/dep/provider"
	"time"

	"github.com/byliuyang/app/fw"
	"github.com/byliuyang/app/modern/mdcli"
	"github.com/byliuyang/app/modern/mddb"
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

func InjectCommandFactory() fw.CommandFactory {
	wire.Build(
		wire.Bind(new(fw.CommandFactory), new(mdcli.CobraFactory)),
		mdcli.NewCobraFactory,
	)
	return mdcli.CobraFactory{}
}

func InjectDBConnector() fw.DBConnector {
	wire.Build(
		wire.Bind(new(fw.DBConnector), new(mddb.PostgresConnector)),
		mddb.NewPostgresConnector,
	)
	return mddb.PostgresConnector{}
}

func InjectDBMigrationTool() fw.DBMigrationTool {
	wire.Build(
		wire.Bind(new(fw.DBMigrationTool), new(mddb.PostgresMigrationTool)),
		mddb.NewPostgresMigrationTool,
	)
	return mddb.PostgresMigrationTool{}
}

func InjectGraphQlService(
	name string,
	sqlDB *sql.DB,
	graphqlPath provider.GraphQlPath,
	secret provider.ReCaptchaSecret,
	jwtSecret provider.JwtSecret,
	bufferSize provider.KeyGenBufferSize,
	kgsRpcConfig provider.KgsRpcConfig,
) (mdservice.Service, error) {
	wire.Build(
		wire.Bind(new(fw.GraphQlAPI), new(graphql.Short)),
		wire.Bind(new(url.Retriever), new(url.RetrieverPersist)),
		wire.Bind(new(url.Creator), new(url.CreatorPersist)),
		wire.Bind(new(repo.UserURLRelation), new(db.UserURLRelationSQL)),
		wire.Bind(new(repo.URL), new(*db.URLSql)),
		wire.Bind(new(keygen.KeyGenerator), new(keygen.Remote)),
		wire.Bind(new(service.KeyGen), new(kgs.Rpc)),

		observabilitySet,
		authSet,

		mdservice.New,
		provider.GraphGophers,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdtimer.NewTimer,

		db.NewURLSql,
		db.NewUserURLRelationSQL,
		provider.NewRemote,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		provider.NewKgsRpc,
		provider.NewReCaptchaService,
		requester.NewVerifier,
		graphql.NewShort,
	)
	return mdservice.Service{}, nil
}

func InjectRoutingService(
	name string,
	sqlDB *sql.DB,
	githubClientID provider.GithubClientID,
	githubClientSecret provider.GithubClientSecret,
	jwtSecret provider.JwtSecret,
	webFrontendURL provider.WebFrontendURL,
) mdservice.Service {
	wire.Build(
		wire.Bind(new(service.Account), new(account.RepoService)),
		wire.Bind(new(url.Retriever), new(url.RetrieverPersist)),
		wire.Bind(new(repo.User), new(*(db.UserSQL))),
		wire.Bind(new(repo.URL), new(*db.URLSql)),

		observabilitySet,
		authSet,

		mdservice.New,
		mdrouting.NewBuiltIn,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdrequest.NewGraphQl,
		mdtimer.NewTimer,

		db.NewUserSQL,
		db.NewURLSql,
		url.NewRetrieverPersist,
		account.NewRepoService,
		provider.NewGithubOAuth,
		github.NewAPI,
		provider.NewShortRoutes,
	)
	return mdservice.Service{}
}
