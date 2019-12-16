package google

// API represents Google API client.
type API struct {
	IdentityProvider IdentityProvider
	Account          Account
}

// NewAPI creates Google API client.
func NewAPI(identityProvider IdentityProvider, account Account) API {
	return API{
		IdentityProvider: identityProvider,
		Account:          account,
	}
}
