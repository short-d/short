package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdlogger"
	"github.com/short-d/short/env"
)

func NewEntryRepositorySwitch(
	serverEnv fw.ServerEnv,
	stdOut fw.StdOut,
	dataDogAPIKey DataDogAPIKey,
	httpRequest fw.HTTPRequest,
) mdlogger.EntryRepository {
	if serverEnv == env.Development {
		return mdlogger.NewLocal(stdOut)
	}
	return NewDataDogEntryRepo(dataDogAPIKey, httpRequest, serverEnv)
}
