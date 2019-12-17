package db_test

import (
	"github.com/byliuyang/app/fw"
	"path"
	"short/dep"
	"strconv"
	"testing"
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
