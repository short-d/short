package cmd

import (
	"fmt"
	"os"

	"github.com/byliuyang/app/fw"
)

func NewRootCmd(
	dbConfig fw.DBConfig,
	recaptchaSecret string,
	githubConfig GithubConfig,
	jwtSecret string,
	cmdFactory fw.CommandFactory,
	dbConnector fw.DBConnector,
	dbMigrationTool fw.DBMigrationTool,
) fw.Command {
	var migrationRoot string
	var wwwRoot string

	startCmd := cmdFactory.NewCommand(
		fw.CommandConfig{
			Usage:        "start",
			ShortHelpMsg: "Start service",
			OnExecute: func(cmd *fw.Command, args []string) {
				start(
					dbConfig,
					migrationRoot,
					wwwRoot,
					recaptchaSecret,
					githubConfig,
					jwtSecret,
					dbConnector,
					dbMigrationTool,
				)
			},
		},
	)
	startCmd.AddStringFlag(&migrationRoot, "migration", "app/adapter/migration", "migration migrations root directory")
	startCmd.AddStringFlag(&wwwRoot, "www", "public", "www root directory")

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
