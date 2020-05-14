package app

import (
	"time"

	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/env"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/dep"
	"github.com/short-d/short/backend/dep/provider"
)

// ServiceConfig represents require parameters for the backend APIs
type ServiceConfig struct {
	Runtime              string
	LogPrefix            string
	LogLevel             logger.LogLevel
	MigrationRoot        string
	RecaptchaSecret      string
	GithubClientID       string
	GithubClientSecret   string
	FacebookClientID     string
	FacebookClientSecret string
	FacebookRedirectURI  string
	GoogleClientID       string
	GoogleClientSecret   string
	GoogleRedirectURI    string
	JwtSecret            string
	WebFrontendURL       string
	GraphQLAPIPort       int
	HTTPAPIPort          int
	KeyGenBufferSize     int
	KgsHostname          string
	KgsPort              int
	AuthTokenLifetime    time.Duration
	DataDogAPIKey        string
	SegmentAPIKey        string
	IPStackAPIKey        string
	GoogleAPIKey         string
}

// Start launches the GraphQL & HTTP APIs
func Start(
	dbConfig db.Config,
	dbConnector db.Connector,
	dbMigrationTool db.MigrationTool,
	config ServiceConfig,
) {
	sqlDB, err := dbConnector.Connect(dbConfig)
	if err != nil {
		panic(err)
	}

	err = dbMigrationTool.MigrateUp(sqlDB, config.MigrationRoot)
	if err != nil {
		panic(err)
	}

	kgsBufferSize := provider.KeyGenBufferSize(config.KeyGenBufferSize)
	kgsRPCConfig := provider.KgsRPCConfig{
		Hostname: config.KgsHostname,
		Port:     config.KgsPort,
	}

	dataDogAPIKey := provider.DataDogAPIKey(config.DataDogAPIKey)
	segmentAPIKey := provider.SegmentAPIKey(config.SegmentAPIKey)
	ipStackAPIKey := provider.IPStackAPIKey(config.IPStackAPIKey)
	googleAPIKey := provider.GoogleAPIKey(config.GoogleAPIKey)

	graphqlAPI, err := dep.InjectGraphQLService(
		env.Runtime(config.Runtime),
		provider.LogPrefix(config.LogPrefix),
		config.LogLevel,
		sqlDB,
		"/graphql",
		provider.ReCaptchaSecret(config.RecaptchaSecret),
		provider.JwtSecret(config.JwtSecret),
		kgsBufferSize,
		kgsRPCConfig,
		provider.TokenValidDuration(config.AuthTokenLifetime),
		dataDogAPIKey,
		segmentAPIKey,
		ipStackAPIKey,
		googleAPIKey,
	)
	if err != nil {
		panic(err)
	}

	graphqlAPI.StartAsync(config.GraphQLAPIPort)

	httpAPI, err := dep.InjectRoutingService(
		env.Runtime(config.Runtime),
		provider.LogPrefix(config.LogPrefix),
		config.LogLevel,
		sqlDB,
		provider.GithubClientID(config.GithubClientID),
		provider.GithubClientSecret(config.GithubClientSecret),
		provider.FacebookClientID(config.FacebookClientID),
		provider.FacebookClientSecret(config.FacebookClientSecret),
		provider.FacebookRedirectURI(config.FacebookRedirectURI),
		provider.GoogleClientID(config.GoogleClientID),
		provider.GoogleClientSecret(config.GoogleClientSecret),
		provider.GoogleRedirectURI(config.GoogleRedirectURI),
		provider.JwtSecret(config.JwtSecret),
		kgsBufferSize,
		kgsRPCConfig,
		provider.WebFrontendURL(config.WebFrontendURL),
		provider.TokenValidDuration(config.AuthTokenLifetime),
		dataDogAPIKey,
		segmentAPIKey,
		ipStackAPIKey,
	)
	if err != nil {
		panic(err)
	}

	httpAPI.StartAndWait(config.HTTPAPIPort)
}
