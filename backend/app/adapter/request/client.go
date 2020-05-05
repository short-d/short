package request

import (
	"net/http"

	"github.com/short-d/app/fw/geo"
	"github.com/short-d/app/fw/network"
)

// Client retrieves user device info.
type Client struct {
	network network.Network
	geo     geo.Geo
}

// GetLocation extracts user's geo location from the HTTP request.
func (c Client) GetLocation(request *http.Request) (geo.Location, error) {
	connection := c.network.FromHTTP(request)
	clientIP := connection.ClientIP
	if clientIP == "" {
		return geo.Location{}, nil
	}
	return c.geo.GetLocation(clientIP)
}

// NewClient creates user device info retriever.
func NewClient(network network.Network, geo geo.Geo) Client {
	return Client{
		network: network,
		geo:     geo,
	}
}
