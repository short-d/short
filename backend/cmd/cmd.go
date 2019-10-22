package cmd

import (
	"fmt"
	"os"
	"short/app"

	"github.com/byliuyang/app/fw"
)

func NewRootCmd(
	dbConfig fw.DBConfig,
	recaptchaSecret string,
	githubClientID string,
	githubClientSecret string,
	jwtSecret string,
	webFrontendURL string,
	graphQLAPIPort int,
	httpAPIPort int,
	cmdFactory fw.CommandFactory,
	dbConnector fw.DBConnector,
	dbMigrationTool fw.DBMigrationTool,
) fw.Command {
	var migrationRoot string

	startCmd := cmdFactory.NewCommand(
		fw.CommandConfig{
			Usage:        "start",
			ShortHelpMsg: "Start service",
			OnExecute: func(cmd *fw.Command, args []string) {
				app.Start(
					dbConfig,
					migrationRoot,
					recaptchaSecret,
					githubClientID,
					githubClientSecret,
					jwtSecret,
					webFrontendURL,
					graphQLAPIPort,
					httpAPIPort,
					dbConnector,
					dbMigrationTool,
				)
			},
		},
	)
	startCmd.AddStringFlag(&migrationRoot, "migration", "app/adapter/migration", "migration migrations root directory")

	rootCmd := cmdFactory.NewCommand(
		fw.CommandConfig{
			Usage:     "short",
			OnExecute: func(cmd *fw.Command, args []string) {},
		},
	)
	err := rootCmd.AddSubCommand(startCmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return rootCmd
}

func Execute(rootCmd fw.Command) {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
