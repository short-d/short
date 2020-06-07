package provider

import (
	"github.com/short-d/app/fw/env"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/app/fw/webreq"
)

// DataDogAPIKey represents credential for DataDog APIs.
type DataDogAPIKey string

// NewDataDogEntryRepo creates new DataDogEntryRepo with DataDogAPIKey to uniquely
// identify apiKey during dependency injection.
func NewDataDogEntryRepo(
	apiKey DataDogAPIKey,
	httpRequest webreq.HTTP,
	runtime env.Runtime,
) logger.DataDogEntryRepo {
	return logger.NewDataDogEntryRepo(string(apiKey), httpRequest, runtime)
}

// NewDataDogMetrics creates new DataDog Metrics with DataDogAPIKey to uniquely
// identify apiKey during dependency injection.
func NewDataDogMetrics(
	apiKey DataDogAPIKey,
	httpRequest webreq.HTTP,
	timer timer.Timer,
	runtime env.Runtime,
) metrics.DataDog {
	return metrics.NewDataDog(string(apiKey), httpRequest, timer, runtime)
}
