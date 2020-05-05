package provider

import (
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/adapter/github"
)

// GithubClientID represents client ID used for Github OAuth.
type GithubClientID string

// GithubClientSecret represents client secret used for Github OAuth.
type GithubClientSecret string

// NewGithubIdentityProvider creates a new Github OAuth client with
// GithubClientID and GithubClientSecret to uniquely identify clientID and
// clientSecret during dependency injection.
func NewGithubIdentityProvider(
	req webreq.HTTP,
	clientID GithubClientID,
	clientSecret GithubClientSecret,
) github.IdentityProvider {
	return github.NewIdentityProvider(req, string(clientID), string(clientSecret))
}
