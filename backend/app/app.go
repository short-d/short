package app

import (
	"short/dep"
	"short/dep/provider"

	"github.com/byliuyang/app/fw"
)

// ServiceConfig represents require parameters for the backend APIs
type ServiceConfig struct {
	MigrationRoot        string
	RecaptchaSecret      string
	GithubClientID       string
	GithubClientSecret   string
	FacebookClientID     string
	FacebookClientSecret string
	FacebookRedirectURI  string
	JwtSecret            string
	WebFrontendURL       string
	GraphQLAPIPort       int
	HTTPAPIPort          int
	KeyGenBufferSize     int
	KgsHostname          string
	KgsPort              int
}

// Start launches the GraphQL & HTTP APIs
func Start(
	dbConfig fw.DBConfig,
	config ServiceConfig,
	dbConnector fw.DBConnector,
	dbMigrationTool fw.DBMigrationTool,
) {
	db, err := dbConnector.Connect(dbConfig)
	if err != nil {
		panic(err)
	}

	err = dbMigrationTool.MigrateUp(db, config.MigrationRoot)
	if err != nil {
		panic(err)
	}

	graphqlAPI, err := dep.InjectGraphQlService(
		"GraphQL API",
		db,
		"/graphql",
		provider.ReCaptchaSecret(config.RecaptchaSecret),
		provider.JwtSecret(config.JwtSecret),
		provider.KeyGenBufferSize(config.KeyGenBufferSize),
		provider.KgsRPCConfig{
			Hostname: config.KgsHostname,
			Port:     config.KgsPort,
		},
	)
	if err != nil {
		panic(err)
	}
	graphqlAPI.Start(config.GraphQLAPIPort)

	httpAPI := dep.InjectRoutingService(
		"Routing API",
		db,
		provider.GithubClientID(config.GithubClientID),
		provider.GithubClientSecret(config.GithubClientSecret),
		provider.FacebookClientID(config.FacebookClientID),
		provider.FacebookClientSecret(config.FacebookClientSecret),
		provider.FacebookRedirectURI(config.FacebookRedirectURI),
		provider.JwtSecret(config.JwtSecret),
		provider.WebFrontendURL(config.WebFrontendURL),
	)
	httpAPI.StartAndWait(config.HTTPAPIPort)
}
