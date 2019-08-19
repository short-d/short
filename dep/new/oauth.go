package new

import (
	"short/app/adapter/oauth"
	"short/fw"
)

type GithubClientId string
type GithubClientSecret string

func GithubOAuth(
	req fw.HttpRequest,
	clientId GithubClientId,
	clientSecret GithubClientSecret,
) oauth.Github {
	return oauth.NewGithub(req, string(clientId), string(clientSecret))
}
