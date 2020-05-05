// +build integration all

package sqldb_test

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/envconfig"
	"github.com/short-d/short/dep"
)

var dbConnector db.Connector
var dbMigrationTool db.MigrationTool

var dbConfig db.Config
var dbMigrationRoot = "./migration"

func TestMain(m *testing.M) {
	env := dep.InjectEnv()
	env.AutoLoadDotEnvFile()

	envConfig := envconfig.NewEnvConfig(env)

	config := struct {
		DBHost     string `env:"DB_HOST" default:"localhost"`
		DBPort     int    `env:"DB_PORT" default:"5432"`
		DBUser     string `env:"DB_USER" default:"postgres"`
		DBPassword string `env:"DB_PASSWORD" default:"password"`
		DBName     string `env:"DB_NAME" default:"short"`
	}{}

	err := envConfig.ParseConfigFromEnv(&config)
	if err != nil {
		panic(err)
	}

	dbConfig = db.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Password: config.DBPassword,
		DbName:   config.DBName,
	}

	dbConnector = dep.InjectDBConnector()
	dbMigrationTool = dep.InjectDBMigrationTool()

	m.Run()
}

func mustParseTime(t *testing.T, timeString string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeString)
	assert.Equal(t, nil, err)
	return parsedTime.UTC()
}
