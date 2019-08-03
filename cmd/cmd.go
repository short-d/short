package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"tinyURL/app"
	"tinyURL/dep"
	"tinyURL/modern"
)

func Execute() {
	var host string
	var port int
	var user string
	var password string
	var dbName string

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start TinyUrl service",
		Run: func(cmd *cobra.Command, args []string) {
			start(host, port, user, password, dbName)
		},
	}

	startCmd.Flags().StringVar(&host, "host", "localhost", "hostname of db server")
	startCmd.Flags().IntVar(&port, "port", 5432, "port of db server")
	startCmd.Flags().StringVar(&user, "user", "postgres", "username of database")
	startCmd.Flags().StringVar(&password, "password", "password", "password of database")
	startCmd.Flags().StringVar(&dbName, "db", "tinyurl", "name of database")

	rootCmd := &cobra.Command{Use: "tinyurl"}
	rootCmd.AddCommand(startCmd)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func start(host string, port int, user string, password string, dbName string) {
	db, err := modern.NewPostgresDb(host, port, user, password, dbName)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = modern.MigratePostgres(db, "app/db")
	if err != nil {
		panic(err)
	}

	service := dep.InitGraphQlService("TinyUrl GraphQL API", db, modern.GraphQlPath("/graphql"))
	service.Start(8080)

	service = dep.InitRoutingService("TinyUrl Routing API", db, app.WwwRoot("app/web/build"))
	service.StartAndWait(80)
}
