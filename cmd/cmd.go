package cmd

import (
	"fmt"
	"os"
	"short/app"
	"short/dep"
	"short/modern"
	"strconv"

	"github.com/spf13/cobra"
)

func Execute() {
	var migrationRoot string
	var wwwRoot string

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start service",
		Run: func(cmd *cobra.Command, args []string) {
			host := getEnv("DB_HOST", "localhost")

			portStr := getEnv("DB_PORT", "5432")
			port := MustInt(portStr)

			user := getEnv("DB_USER", "postgres")
			password := getEnv("DB_PASSWORD", "password")
			dbName := getEnv("DB_NAME", "short")

			start(host, port, user, password, dbName, migrationRoot, wwwRoot)
		},
	}

	startCmd.Flags().StringVar(&migrationRoot, "migration", "app/db", "db migrations root directory")
	startCmd.Flags().StringVar(&wwwRoot, "www", "build/web", "www root directory")

	rootCmd := &cobra.Command{Use: "short"}
	rootCmd.AddCommand(startCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getEnv(varName string, defaultVal string) string {
	val := os.Getenv(varName)

	if val == "" {
		return defaultVal
	}

	return val
}

func MustInt(numStr string) int {
	num, err := strconv.Atoi(numStr)

	if err != nil {
		panic(err)
	}

	return num
}

func start(host string, port int, user string, password string, dbName string, migrationRoot string, wwwRoot string) {
	db, err := modern.NewPostgresDb(host, port, user, password, dbName)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = modern.MigratePostgres(db, migrationRoot)
	if err != nil {
		panic(err)
	}

	service := dep.InitGraphQlService("GraphQL API", db, modern.GraphQlPath("/graphql"))
	service.Start(8080)

	service = dep.InitRoutingService("Routing API", db, app.WwwRoot(wwwRoot))
	service.StartAndWait(80)
}
