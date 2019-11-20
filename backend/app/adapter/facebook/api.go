package facebook

// API represents Facebook API client.
type API struct {
	IdentityProvider IdentityProvider
	Account              Account
}

// NewAPI creates Facebook API client.
func NewAPI(identityProvider IdentityProvider, account Account) API {
	return API{
		IdentityProvider: identityProvider,
		Account:              account,
	}
}
