package cmd

import (
	"short/dep"
	"short/dep/provider"

	"github.com/byliuyang/app/fw"
)

type GithubConfig struct {
	ClientID     string
	ClientSecret string
}

type FacebookConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

type KgsRPCConfig struct {
	Hostname string
	Port     int
}

func start(
	dbConfig fw.DBConfig,
	migrationRoot string,
	recaptchaSecret string,
	githubConfig GithubConfig,
	facebookConfig FacebookConfig,
	jwtSecret string,
	KeyGenBufferSize int,
	KgsRPCConfig KgsRPCConfig,
	webFrontendURL string,
	dbConnector fw.DBConnector,
	dbMigrationTool fw.DBMigrationTool,
) {
	db, err := dbConnector.Connect(dbConfig)
	if err != nil {
		panic(err)
	}

	err = dbMigrationTool.Migrate(db, migrationRoot)
	if err != nil {
		panic(err)
	}

	graphqlAPI, err := dep.InjectGraphQlService(
		"GraphQL API",
		db,
		"/graphql",
		provider.ReCaptchaSecret(recaptchaSecret),
		provider.JwtSecret(jwtSecret),
		provider.KeyGenBufferSize(KeyGenBufferSize),
		provider.KgsRPCConfig(KgsRPCConfig),
	)
	if err != nil {
		panic(err)
	}
	graphqlAPI.Start(8080)

	httpAPI := dep.InjectRoutingService(
		"Routing API",
		db,
		provider.GithubClientID(githubConfig.ClientID),
		provider.GithubClientSecret(githubConfig.ClientSecret),
		provider.FacebookClientID(facebookConfig.ClientID),
		provider.FacebookClientSecret(facebookConfig.ClientSecret),
		provider.FacebookRedirectURI(facebookConfig.RedirectURI),
		provider.JwtSecret(jwtSecret),
		provider.WebFrontendURL(webFrontendURL),
	)
	httpAPI.StartAndWait(80)
}
