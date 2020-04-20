package request

import (
	"net/http"

	"github.com/short-d/app/fw"
)

type Client struct {
	network     fw.Network
	geoLocation fw.GeoLocation
}

func (c Client) GetLocation(request *http.Request) (fw.Location, error) {
	connection := c.network.FromHTTP(request)
	clientIP := connection.ClientIP
	return c.geoLocation.GetLocation(clientIP)
}

func NewClient(network fw.Network, geoLocation fw.GeoLocation) Client {
	return Client{
		network:     network,
		geoLocation: geoLocation,
	}
}
