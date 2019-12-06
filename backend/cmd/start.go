package cmd

import (
	"short/dep"
	"short/dep/provider"

	"github.com/byliuyang/app/fw"
)

// GithubConfig contains Github OAuth config
type GithubConfig struct {
	ClientID     string
	ClientSecret string
}

// FacebookConfig contains Facebook OAuth config
type FacebookConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

// GoogleConfig contains Facebook OAuth config
type GoogleConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func start(
	dbConfig fw.DBConfig,
	migrationRoot string,
	recaptchaSecret string,
	githubConfig GithubConfig,
	facebookConfig FacebookConfig,
	googleConfig GoogleConfig,
	jwtSecret string,
	KeyGenBufferSize int,
	KgsRPCConfig provider.KgsRPCConfig,
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
		provider.GoogleClientID(googleConfig.ClientID),
		provider.GoogleClientSecret(googleConfig.ClientSecret),
		provider.GoogleRedirectURI(googleConfig.RedirectURI),
		provider.JwtSecret(jwtSecret),
		provider.WebFrontendURL(webFrontendURL),
	)
	httpAPI.StartAndWait(80)
}
