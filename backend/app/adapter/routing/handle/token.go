package handle

import (
	"net/http"
	"net/url"
	"strings"
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

// getBearerToken parses Authorization token with format "Bearer <token>"
func getBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < 1 {
		return ""
	}
	words := strings.Split(authHeader, " ")
	if len(words) != 2 {
		return ""
	}
	if words[0] != "Bearer" {
		return ""
	}
	return words[1]
}
