package routing

import (
	"net/url"
)

func getToken(params map[string]string) string {
	return params["token"]
}

func setToken(url url.URL, token string) url.URL {
	query := url.Query()
	query.Set("token", token)
	url.RawQuery = query.Encode()
	return url
}
