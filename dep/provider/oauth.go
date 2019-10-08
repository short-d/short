package provider

import (
	"short/app/adapter/oauth"

	"github.com/byliuyang/app/fw"
)

// GithubClientID represents client ID used for Github OAuth.
type GithubClientID string

// GithubClientSecret represents client secret used for Github OAuth.
type GithubClientSecret string

// GithubOAuth creates a new Github OAuth client with GithubClientID and GithubClientSecret to uniquely identify clientID and clientSecret during dependency injection.
func GithubOAuth(
	req fw.HTTPRequest,
	clientID GithubClientID,
	clientSecret GithubClientSecret,
) oauth.Github {
	return oauth.NewGithub(req, string(clientID), string(clientSecret))
}
