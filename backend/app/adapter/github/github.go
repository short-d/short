package github

// Github groups Github oauth and public APIs together.
type Github struct {
	IdentityProvider IdentityProvider
	API              API
}

// NewFacebook creates consumer of Github API.
func NewGithub(identityProvider IdentityProvider, api API) Github {
	return Github{
		IdentityProvider: identityProvider,
		API:              api,
	}
}
