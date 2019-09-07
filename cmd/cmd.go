package cmd

import (
	"fmt"
	"os"
	"short/dep"
	"strconv"

	"github.com/byliuyang/app/fw"
)

func NewRootCmd(
	host string,
	portStr string,
	user string,
	password string,
	dbName string,
	recaptchaSecret string,
	githubClientID string,
	githubClientSecret string,
	jwtSecret string,
) fw.Command {
	var migrationRoot string
	var wwwRoot string

	cmdFactory := dep.InjectCommandFactory()

	startCmd := cmdFactory.NewCommand(
		fw.CommandConfig{
			Usage:        "start",
			ShortHelpMsg: "Start service",
			OnExecute: func(cmd *fw.Command, args []string) {
				port := MustInt(portStr)

				start(
					host,
					port,
					user,
					password,
					dbName,
					migrationRoot,
					wwwRoot,
					recaptchaSecret,
					githubClientID,
					githubClientSecret,
					jwtSecret,
				)
			},
		},
	)
	startCmd.AddStringFlag(&migrationRoot, "migration", "app/adapter/migration", "migration migrations root directory")
	startCmd.AddStringFlag(&wwwRoot, "www", "public", "www root directory")

	rootCmd := cmdFactory.NewCommand(
		fw.CommandConfig{
			Usage: "short",
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

func MustInt(numStr string) int {
	num, err := strconv.Atoi(numStr)

	if err != nil {
		panic(err)
	}

	return num
}
