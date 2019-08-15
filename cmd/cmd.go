package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"short/dep"
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

			recaptchaSecret := getEnv("RECAPTCHA_SECRET", "")
			githubClientId := getEnv("GITHUB_CLIENT_ID", "")
			githubClientSecret := getEnv("GITHUB_CLIENT_SECRET", "")

			start(
				host,
				port,
				user,
				password,
				dbName,
				migrationRoot,
				wwwRoot,
				recaptchaSecret,
				githubClientId,
				githubClientSecret,
			)
		},
	}

	startCmd.Flags().StringVar(&migrationRoot, "migration", "app/adapter/migration", "migration migrations root directory")
	startCmd.Flags().StringVar(&wwwRoot, "www", "public", "www root directory")

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

func start(
	host string,
	port int,
	user string,
	password string,
	dbName string,
	migrationRoot string,
	wwwRoot string,
	recaptchaSecret string,
	githubClientId string,
	githubClientSecret string,
) {
	dep.InitDB(host, port, user, password, dbName, migrationRoot, func(db *sql.DB) {
		service := dep.InitGraphQlService(
			"GraphQL API",
			db,
			"/graphql",
			dep.ReCaptchaSecret(recaptchaSecret),
		)
		service.Start(8080)

		service = dep.InitRoutingService(
			"Routing API",
			db,
			dep.WwwRoot(wwwRoot),
			dep.GithubClientId(githubClientId),
			dep.GithubClientSecret(githubClientSecret),
		)
		service.StartAndWait(80)
	})
}
