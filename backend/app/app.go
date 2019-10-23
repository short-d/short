package app

import (
	"short/dep"
	"short/dep/provider"

	"github.com/byliuyang/app/fw"
)

// Start launches the GraphQL & HTTP APIs
func Start(
	dbConfig fw.DBConfig,
	migrationRoot string,
	recaptchaSecret string,
	githubClientID string,
	githubClientSecret string,
	jwtSecret string,
	webFrontendURL string,
	graphQLAPIPort int,
	httpAPIPort int,
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

	graphqlAPI := dep.InjectGraphQlService(
		"GraphQL API",
		db,
		"/graphql",
		provider.ReCaptchaSecret(recaptchaSecret),
		provider.JwtSecret(jwtSecret),
	)
	graphqlAPI.Start(graphQLAPIPort)

	httpAPI := dep.InjectRoutingService(
		"Routing API",
		db,
		provider.GithubClientID(githubClientID),
		provider.GithubClientSecret(githubClientSecret),
		provider.JwtSecret(jwtSecret),
		provider.WebFrontendURL(webFrontendURL),
	)
	httpAPI.StartAndWait(httpAPIPort)
}
