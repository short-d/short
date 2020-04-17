package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdgeo"
)

type IPStackAPIKey string

func NewIPStack(
	apiKey IPStackAPIKey,
	httpRequest fw.HTTPRequest,
	logger fw.Logger,
) mdgeo.IPStack {
	return mdgeo.NewIPStack(string(apiKey), httpRequest, logger)
}
