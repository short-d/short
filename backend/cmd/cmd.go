package cmd

import (
	"fmt"
	"os"

	"github.com/short-d/app/fw/cli"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/short/app"
)

// NewRootCmd creates the base command.
func NewRootCmd(
	dbConfig db.Config,
	config app.ServiceConfig,
	cmdFactory cli.CommandFactory,
	dbConnector db.Connector,
	dbMigrationTool db.MigrationTool,
) cli.Command {
	var migrationRoot string

	startCmd := cmdFactory.NewCommand(
		cli.CommandConfig{
			Usage:        "start",
			ShortHelpMsg: "Start service",
			OnExecute: func(cmd *cli.Command, args []string) {
				config.MigrationRoot = migrationRoot
				app.Start(
					dbConfig,
					dbConnector,
					dbMigrationTool,
					config,
				)
			},
		},
	)
	startCmd.AddStringFlag(
		&migrationRoot,
		"migration",
		"app/adapter/sqldb/migration",
		"migration migrations root directory",
	)

	rootCmd := cmdFactory.NewCommand(
		cli.CommandConfig{
			Usage:     "short",
			OnExecute: func(cmd *cli.Command, args []string) {},
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
func Execute(rootCmd cli.Command) {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
