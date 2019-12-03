package cmd

import (
	"fmt"
	"github.com/byliuyang/app/fw"
	"os"
	"short/app"
)

// ServiceConfig represents necessary parameters needed to initialize the
// backend APIs.
type ServiceConfig struct {
	RecaptchaSecret      string
	GithubClientID       string
	GithubClientSecret   string
	FacebookClientID     string
	FacebookClientSecret string
	FacebookRedirectURI  string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string
	JwtSecret            string
	WebFrontendURL       string
	GraphQLAPIPort       int
	HTTPAPIPort          int
	KeyGenBufferSize     int
	KgsHostname          string
	KgsPort              int
}

func NewRootCmd(
	dbConfig fw.DBConfig,
	config ServiceConfig,
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

				serviceConfig := app.ServiceConfig{
					MigrationRoot:        migrationRoot,
					RecaptchaSecret:      config.RecaptchaSecret,
					GithubClientID:       config.GithubClientID,
					GithubClientSecret:   config.GithubClientSecret,
					FacebookClientID:     config.FacebookClientID,
					FacebookClientSecret: config.FacebookClientSecret,
					FacebookRedirectURI:  config.FacebookRedirectURI,
					GoogleClientID:		  config.GoogleClientID,
					GoogleClientSecret:	  config.GoogleClientSecret,
					GoogleRedirectURI:	  config.GoogleRedirectURI,
					JwtSecret:            config.JwtSecret,
					WebFrontendURL:       config.WebFrontendURL,
					GraphQLAPIPort:       config.GraphQLAPIPort,
					HTTPAPIPort:          config.HTTPAPIPort,
					KeyGenBufferSize:     config.KeyGenBufferSize,
					KgsHostname:          config.KgsHostname,
					KgsPort:              config.KgsPort,
				}

				app.Start(
					dbConfig,
					serviceConfig,
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
