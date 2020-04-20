package request

import (
	"net/http"

	"github.com/short-d/app/fw"
)

// Client retriever user device info.
type Client struct {
	network     fw.Network
	geoLocation fw.GeoLocation
}

// GetLocation extracts user's geo location from the HTTP request.
func (c Client) GetLocation(request *http.Request) (fw.Location, error) {
	connection := c.network.FromHTTP(request)
	clientIP := connection.ClientIP
	return c.geoLocation.GetLocation(clientIP)
}

// NewClient creates user device info retriever.
func NewClient(network fw.Network, geoLocation fw.GeoLocation) Client {
	return Client{
		network:     network,
		geoLocation: geoLocation,
	}
}
