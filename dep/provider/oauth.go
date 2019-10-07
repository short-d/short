package provider

import (
	"short/app/adapter/oauth"

	"github.com/byliuyang/app/fw"
)

// GithubClientID Github client ID credential.
type GithubClientID string

// GithubClientSecret Github client secret credential.
type GithubClientSecret string

// GithubOAuth create a new Github struct to get Authorization infos.
func GithubOAuth(
	req fw.HTTPRequest,
	clientID GithubClientID,
	clientSecret GithubClientSecret,
) oauth.Github {
	return oauth.NewGithub(req, string(clientID), string(clientSecret))
}
