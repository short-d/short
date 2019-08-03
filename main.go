package main

import (
	"os"
	"strconv"
	"tinyURL/app"
	"tinyURL/dep"
	"tinyURL/modern"
)

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}
	return val
}

func main() {
	host := getEnv("TINYURL_DB_HOST", "localhost")

	port, err := strconv.Atoi(getEnv("TINYURL_DB_PORT", "5432"))
	if err != nil {
		panic(err)
	}

	user := getEnv("TINYURL_DB_USER", "postgres")
	password := getEnv("TINYURL_DB_PASSWORD", "password")
	dbName := getEnv("TINYURL_DB_NAME", "tinyurl")

	db, err := modern.NewPostgresDb(host, port, user, password, dbName)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	service := dep.InitGraphQlService("TinyUrl GraphQL API", db, modern.GraphQlPath("/graphql"))
	service.Start(8080)

	service = dep.InitRoutingService("TinyUrl Routing API", db, app.WwwRoot("app/web/build"))
	service.StartAndWait(80)
}
