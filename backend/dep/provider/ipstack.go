package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdgeo"
)

// IPStackAPIKey represents credential for IP Stack APIs.
type IPStackAPIKey string

// NewIPStack creates IPStack with IPStackAPIKey to uniquely identify apiKey
// during dependency injection.
func NewIPStack(
	apiKey IPStackAPIKey,
	httpRequest fw.HTTPRequest,
	logger fw.Logger,
) mdgeo.IPStack {
	return mdgeo.NewIPStack(string(apiKey), httpRequest, logger)
}
