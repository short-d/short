package routing

import (
	"net/url"

	"github.com/short-d/app/fw"
)

func getToken(params fw.Params) string {
	return params["token"]
}

func setToken(url url.URL, token string) url.URL {
	query := url.Query()
	query.Set("token", token)
	url.RawQuery = query.Encode()
	return url
}
