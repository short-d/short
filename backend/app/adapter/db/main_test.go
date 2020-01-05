// +build integration all

package db_test

import (
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/short-d/short/dep"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/mdtest"
)

var dbConnector fw.DBConnector
var dbMigrationTool fw.DBMigrationTool

var dbConfig fw.DBConfig
var dbMigrationRoot string

func TestMain(m *testing.M) {
	env := dep.InjectEnvironment()
	env.AutoLoadDotEnvFile()

	host := env.GetEnv("DB_HOST", "")
	portStr := env.GetEnv("DB_PORT", "")
	port := mustInt(portStr)
	user := env.GetEnv("DB_USER", "")
	password := env.GetEnv("DB_PASSWORD", "")
	dbName := env.GetEnv("DB_NAME", "")

	dbConfig = fw.DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}

	dbMigrationRoot = path.Join(env.GetEnv("MIGRATION_ROOT", ""))

	dbConnector = dep.InjectDBConnector()
	dbMigrationTool = dep.InjectDBMigrationTool()

	m.Run()
}

func mustInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}

func mustParseTime(t *testing.T, timeString string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeString)
	mdtest.Equal(t, nil, err)
	return parsedTime.UTC()
}
