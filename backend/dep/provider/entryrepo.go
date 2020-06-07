package provider

import (
	"github.com/short-d/app/fw/env"
	"github.com/short-d/app/fw/io"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/webreq"
)

// NewEntryRepositorySwitch swaps between different entry repository
// implementations based on server environment.
func NewEntryRepositorySwitch(
	runtime env.Runtime,
	deployment env.Deployment,
	stdOut io.StdOut,
	dataDogAPIKey DataDogAPIKey,
	httpRequest webreq.HTTP,
) logger.EntryRepository {
	if deployment.IsDevelopment() {
		return NewLocalEntryRepo(stdOut)
	}
	return NewDataDogEntryRepo(dataDogAPIKey, httpRequest, runtime)
}
