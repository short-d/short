package facebook

// Facebook groups Facebook oauth and public APIs.
type Facebook struct {
	IdentityProvider IdentityProvider
	API              API
}

// NewFacebook creates consumer of Facebook API.
func NewFacebook(identityProvider IdentityProvider, api API) Facebook {
	return Facebook{
		IdentityProvider: identityProvider,
		API:              api,
	}
}
