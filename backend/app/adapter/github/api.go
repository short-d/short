package github

// API represents Github API client.
type API struct {
	IdentityProvider IdentityProvider
	Account          Account
}

// NewAPI creates Github API client.
func NewAPI(identityProvider IdentityProvider, account Account) API {
	return API{
		IdentityProvider: identityProvider,
		Account:          account,
	}
}
