package cmd

import (
	"fmt"
	"os"

	"github.com/short-d/app/fw/cli"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/short/backend/app"
	"github.com/short-d/short/backend/dep"
	"github.com/short-d/short/backend/dep/provider"
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

	idCmd := cmdFactory.NewCommand(cli.CommandConfig{
		Usage:        "data-id",
		ShortHelpMsg: "Use user ID to uniquely identify a user",
		OnExecute: func(cmd *cli.Command, args []string) {
			kgsConfig := provider.KgsRPCConfig{
				Hostname: config.KgsHostname,
				Port:     config.KgsPort,
			}
			keyGenBufferSize := provider.KeyGenBufferSize(config.KeyGenBufferSize)
			dataTool, err := dep.InjectDataTool(dbConfig, dbConnector, keyGenBufferSize, kgsConfig)
			if err != nil {
				panic(err)
			}
			dataTool.EmailToID()
		},
	})

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
	err = rootCmd.AddSubCommand(idCmd)
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
