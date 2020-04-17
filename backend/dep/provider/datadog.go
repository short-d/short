package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdlogger"
	"github.com/short-d/app/modern/mdmetrics"
)

type DataDogAPIKey string

func NewDataDogEntryRepo(
	apiKey DataDogAPIKey,
	httpRequest fw.HTTPRequest,
	env fw.ServerEnv,
) mdlogger.DataDogEntryRepo {
	return mdlogger.NewDataDogEntryRepo(string(apiKey), httpRequest, env)
}

func NewDataDogMetrics(
	apiKey DataDogAPIKey,
	httpRequest fw.HTTPRequest,
	timer fw.Timer,
	env fw.ServerEnv,
) mdmetrics.DataDog {
	return mdmetrics.NewDataDog(string(apiKey), httpRequest, timer, env)
}
