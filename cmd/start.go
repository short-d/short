package cmd

import (
	"database/sql"
	"short/dep"
	"short/dep/provider"
)

func start(
	host string,
	port int,
	user string,
	password string,
	dbName string,
	migrationRoot string,
	wwwRoot string,
	recaptchaSecret string,
	githubClientID string,
	githubClientSecret string,
	jwtSecret string,
) {
	provider.DB(host, port, user, password, dbName, migrationRoot, func(db *sql.DB) {
		service := dep.InjectGraphQlService(
			"GraphQL API",
			db,
			"/graphql",
			provider.ReCaptchaSecret(recaptchaSecret),
			provider.JwtSecret(jwtSecret),
		)
		service.Start(8080)

		service = dep.InjectRoutingService(
			"Routing API",
			db,
			provider.WwwRoot(wwwRoot),
			provider.GithubClientID(githubClientID),
			provider.GithubClientSecret(githubClientSecret),
			provider.JwtSecret(jwtSecret),
		)
		service.StartAndWait(80)
	})
}
