package cmd

import (
	"fmt"
	"os"
	"tinyURL/app"
	"tinyURL/dep"
	"tinyURL/modern"

	"github.com/spf13/cobra"
)

func Execute() {
	var host string
	var port int
	var user string
	var password string
	var dbName string

	var migrationRoot string
	var wwwRoot string

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start TinyUrl service",
		Run: func(cmd *cobra.Command, args []string) {
			start(host, port, user, password, dbName, migrationRoot, wwwRoot)
		},
	}

	startCmd.Flags().StringVar(&host, "host", "localhost", "hostname of db server")
	startCmd.Flags().IntVar(&port, "port", 5432, "port of db server")
	startCmd.Flags().StringVar(&user, "user", "postgres", "username of database")
	startCmd.Flags().StringVar(&password, "password", "password", "password of database")
	startCmd.Flags().StringVar(&dbName, "db", "tinyurl", "name of database")

	startCmd.Flags().StringVar(&migrationRoot, "migration", "app/db", "db migrations root directory")
	startCmd.Flags().StringVar(&wwwRoot, "www", "app/web/build", "www root directory")

	rootCmd := &cobra.Command{Use: "tinyurl"}
	rootCmd.AddCommand(startCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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

	service := dep.InitGraphQlService("TinyUrl GraphQL API", db, modern.GraphQlPath("/graphql"))
	service.Start(8080)

	service = dep.InitRoutingService("TinyUrl Routing API", db, app.WwwRoot(wwwRoot))
	service.StartAndWait(80)
}
