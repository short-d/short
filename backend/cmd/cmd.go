package cmd

import (
	"fmt"
	"os"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app"
)

// NewRootCmd creates the base command.
func NewRootCmd(
	dbConfig fw.DBConfig,
	config app.ServiceConfig,
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
				config.MigrationRoot = migrationRoot
				app.Start(
					dbConfig,
					config,
					dbConnector,
					dbMigrationTool,
				)
			},
		},
	)
	startCmd.AddStringFlag(
		&migrationRoot,
		"migration",
		"app/adapter/db/migration",
		"migration migrations root directory",
	)

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

// Execute runs the root command.
func Execute(rootCmd fw.Command) {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
