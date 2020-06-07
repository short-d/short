package provider

import (
	"github.com/short-d/app/fw/geo"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/webreq"
)

// IPStackAPIKey represents credential for IP Stack APIs.
type IPStackAPIKey string

// NewIPStack creates IPStack with IPStackAPIKey to uniquely identify apiKey
// during dependency injection.
func NewIPStack(
	apiKey IPStackAPIKey,
	httpRequest webreq.HTTP,
	logger logger.Logger,
) geo.IPStack {
	return geo.NewIPStack(string(apiKey), httpRequest, logger)
}
