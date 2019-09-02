package inject

import (
	"short/app/adapter/oauth"

	"github.com/byliuyang/app/fw"
)

type GithubClientID string
type GithubClientSecret string

func GithubOAuth(
	req fw.HTTPRequest,
	clientID GithubClientID,
	clientSecret GithubClientSecret,
) oauth.Github {
	return oauth.NewGithub(req, string(clientID), string(clientSecret))
}
