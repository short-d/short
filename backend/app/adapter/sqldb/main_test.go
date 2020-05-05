// +build integration all

package sqldb_test

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/envconfig"
	"github.com/short-d/short/dep"
)

var dbConnector db.Connector
var dbMigrationTool db.MigrationTool

var dbConfig db.Config
var dbMigrationRoot string

func TestMain(m *testing.M) {
	env := dep.InjectEnvironment()
	env.AutoLoadDotEnvFile()
	envConfig := envconfig.NewEnvConfig(env)

	config := struct {
		DBHost        string `env:"DB_HOST" default:"localhost"`
		DBPort        int    `env:"DB_PORT" default:"5432"`
		DBUser        string `env:"DB_USER" default:"postgres"`
		DBPassword    string `env:"DB_PASSWORD" default:"password"`
		DBName        string `env:"DB_NAME" default:"short"`
		MigrationRoot string `env:"MIGRATION_ROOT" default:""`
	}{}

	err := envConfig.ParseConfigFromEnv(&config)
	if err != nil {
		panic(err)
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
