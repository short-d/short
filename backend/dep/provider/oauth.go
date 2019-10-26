package provider

import (
	"short/app/adapter/oauth"

	"github.com/byliuyang/app/fw"
)

// GithubClientID represents client ID used for Github OAuth.
type GithubClientID string

// GithubClientSecret represents client secret used for Github OAuth.
type GithubClientSecret string

// NewGithubOAuth creates a new Github OAuth client with GithubClientID and GithubClientSecret to uniquely identify clientID and clientSecret during dependency injection.
func NewGithubOAuth(
	req fw.HTTPRequest,
	clientID GithubClientID,
	clientSecret GithubClientSecret,
) oauth.Github {
	return oauth.NewGithub(req, string(clientID), string(clientSecret))
}

// FacebookClientID represents client ID used for Facebook OAuth.
type FacebookClientID string

// FacebookClientSecret represents client secret used for Facebook OAuth.
type FacebookClientSecret string

// FacebookRedirectURI represents redirect_uri for facebook
type FacebookRedirectURI string

// NewFacebookOAuth creates a new Facebook OAuth client with FacebookClientID and FacebookClientSecret to uniquely identify clientID and clientSecret during dependency injection.
func NewFacebookOAuth(
	req fw.HTTPRequest,
	clientID FacebookClientID,
	clientSecret FacebookClientSecret,
	redirectURI FacebookRedirectURI,
) oauth.Facebook {
	return oauth.NewFacebook(req, string(clientID), string(clientSecret), string(redirectURI))
}
