package main

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/cmd"
	"github.com/short-d/short/dep"
)

func main() {
	env := dep.InjectEnvironment()
	env.AutoLoadDotEnvFile()

	host := env.GetEnv("DB_HOST", "localhost")
	portStr := env.GetEnv("DB_PORT", "5432")
	port := dep.MustInt(portStr)
	user := env.GetEnv("DB_USER", "postgres")
	password := env.GetEnv("DB_PASSWORD", "password")
	dbName := env.GetEnv("DB_NAME", "short")

	recaptchaSecret := env.GetEnv("RECAPTCHA_SECRET", "")
	githubClientID := env.GetEnv("GITHUB_CLIENT_ID", "")
	githubClientSecret := env.GetEnv("GITHUB_CLIENT_SECRET", "")
	jwtSecret := env.GetEnv("JWT_SECRET", "")
	webFrontendURL := env.GetEnv("WEB_FRONTEND_URL", "")
	graphQLAPIPort := dep.MustInt(env.GetEnv("GRAPHQL_API_PORT", "8080"))
	httpAPIPort := dep.MustInt(env.GetEnv("HTTP_API_PORT", "80"))

	keyGenBufferSize := dep.MustInt(env.GetEnv("KEY_GEN_BUFFER_SIZE", "50"))
	kgsHostname := env.GetEnv("KEY_GEN_HOSTNAME", "localhost")
	kgsPort := dep.MustInt(env.GetEnv("KEY_GEN_PORT", "8080"))

	facebookClientID := env.GetEnv("FACEBOOK_CLIENT_ID", "")
	facebookClientSecret := env.GetEnv("FACEBOOK_CLIENT_SECRET", "")
	facebookRedirectURI := env.GetEnv("FACEBOOK_REDIRECT_URI", "")

	googleClientID := env.GetEnv("GOOGLE_CLIENT_ID", "")
	googleClientSecret := env.GetEnv("GOOGLE_CLIENT_SECRET", "")
	googleRedirectURI := env.GetEnv("GOOGLE_REDIRECT_URI", "")

	cmdFactory := dep.InjectCommandFactory()
	dbConnector := dep.InjectDBConnector()
	dbMigrationTool := dep.InjectDBMigrationTool()

	dbConfig := fw.DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}

	serviceConfig := cmd.ServiceConfig{
		LogPrefix:            "Short",
		LogLevel:             fw.LogTrace,
		RecaptchaSecret:      recaptchaSecret,
		GithubClientID:       githubClientID,
		GithubClientSecret:   githubClientSecret,
		FacebookClientID:     facebookClientID,
		FacebookClientSecret: facebookClientSecret,
		FacebookRedirectURI:  facebookRedirectURI,
		GoogleClientID:       googleClientID,
		GoogleClientSecret:   googleClientSecret,
		GoogleRedirectURI:    googleRedirectURI,
		JwtSecret:            jwtSecret,
		WebFrontendURL:       webFrontendURL,
		GraphQLAPIPort:       graphQLAPIPort,
		HTTPAPIPort:          httpAPIPort,
		KeyGenBufferSize:     keyGenBufferSize,
		KgsHostname:          kgsHostname,
		KgsPort:              kgsPort,
	}

	rootCmd := cmd.NewRootCmd(
		dbConfig,
		serviceConfig,
		cmdFactory,
		dbConnector,
		dbMigrationTool,
	)
	cmd.Execute(rootCmd)
}
