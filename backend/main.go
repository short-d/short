package main

import (
	"time"

	"github.com/short-d/short/app"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/cmd"
	"github.com/short-d/short/dep"
	"github.com/short-d/short/envconfig"
)

func main() {
	env := dep.InjectEnvironment()
	env.AutoLoadDotEnvFile()

	envConfig := envconfig.NewEnvConfig(env)

	config := struct {
		ServerEnv            string        `env:"ENV" default:"testing"`
		DBHost               string        `env:"DB_HOST" default:"localhost"`
		DBPort               int           `env:"DB_PORT" default:"5432"`
		DBUser               string        `env:"DB_USER" default:"postgres"`
		DBPassword           string        `env:"DB_PASSWORD" default:"password"`
		DBName               string        `env:"DB_NAME" default:"short"`
		ReCaptchaSecret      string        `env:"RECAPTCHA_SECRET" default:""`
		GithubClientID       string        `env:"GITHUB_CLIENT_ID" default:""`
		GithubClientSecret   string        `env:"GITHUB_CLIENT_SECRET" default:""`
		FacebookClientID     string        `env:"FACEBOOK_CLIENT_ID" default:""`
		FacebookClientSecret string        `env:"FACEBOOK_CLIENT_SECRET" default:""`
		FacebookRedirectURI  string        `env:"FACEBOOK_REDIRECT_URI" default:""`
		GoogleClientID       string        `env:"GOOGLE_CLIENT_ID" default:""`
		GoogleClientSecret   string        `env:"GOOGLE_CLIENT_SECRET" default:""`
		GoogleRedirectURI    string        `env:"GOOGLE_REDIRECT_URI" default:""`
		JWTSecret            string        `env:"JWT_SECRET" default:""`
		WebFrontendURL       string        `env:"WEB_FRONTEND_URL" default:""`
		KeyGenBufferSize     int           `env:"KEY_GEN_BUFFER_SIZE" default:"50"`
		KgsHostname          string        `env:"KEY_GEN_HOSTNAME" default:"localhost"`
		KgsPort              int           `env:"KEY_GEN_PORT" default:"8080"`
		GraphQLAPIPort       int           `env:"GRAPHQL_API_PORT" default:"8080"`
		HTTPAPIPort          int           `env:"HTTP_API_PORT" default:"80"`
		AuthTokenLifeTime    time.Duration `env:"AUTH_TOKEN_LIFETIME" default:"1w"`
		DataDogAPIKey        string        `env:"DATA_DOG_API_KEY" default:""`
		SegmentAPIKey        string        `env:"SEGMENT_API_KEY" default:""`
		IPStackAPIKey        string        `env:"IP_STACK_API_KEY" default:""`
	}{}

	err := envConfig.ParseConfigFromEnv(&config)
	if err != nil {
		panic(err)
	}

	cmdFactory := dep.InjectCommandFactory()
	dbConnector := dep.InjectDBConnector()
	dbMigrationTool := dep.InjectDBMigrationTool()

	dbConfig := fw.DBConfig{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DbName:   config.DBName,
	}

	serviceConfig := app.ServiceConfig{
		LogPrefix:            "Short",
		ServerEnv:            config.ServerEnv,
		LogLevel:             fw.LogTrace,
		MigrationRoot:        "",
		RecaptchaSecret:      config.ReCaptchaSecret,
		GithubClientID:       config.GithubClientID,
		GithubClientSecret:   config.GithubClientSecret,
		FacebookClientID:     config.FacebookClientID,
		FacebookClientSecret: config.FacebookClientSecret,
		FacebookRedirectURI:  config.FacebookRedirectURI,
		GoogleClientID:       config.GoogleClientID,
		GoogleClientSecret:   config.GoogleClientSecret,
		GoogleRedirectURI:    config.GoogleRedirectURI,
		JwtSecret:            config.JWTSecret,
		WebFrontendURL:       config.WebFrontendURL,
		GraphQLAPIPort:       config.GraphQLAPIPort,
		HTTPAPIPort:          config.HTTPAPIPort,
		KeyGenBufferSize:     config.KeyGenBufferSize,
		KgsHostname:          config.KgsHostname,
		KgsPort:              config.KgsPort,
		AuthTokenLifetime:    config.AuthTokenLifeTime,
		DataDogAPIKey:        config.DataDogAPIKey,
		SegmentAPIKey:        config.SegmentAPIKey,
		IPStackAPIKey:        config.IPStackAPIKey,
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
