package facebook

import (
	"short/app/entity"
	"short/app/usecase/service"

	fb "github.com/huandu/facebook"
)

var _ service.SSOAccount = (*API)(nil)

// API accesses user's account data through FB Graph API.
type API struct {
}

// GetSingleSignOnUser retrieves user's email and name from Facebook.
func (g API) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	type response struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var fbResponse response

	res, err := fb.Get("/me", fb.Params{
		"fields":       "name,email",
		"access_token": accessToken,
	})

	if err != nil {
		return entity.SSOUser{}, err
	}

	err = res.Decode(&fbResponse)
	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		Email: fbResponse.Email,
		Name:  fbResponse.Name,
	}, nil
}

// NewAPI initializes Facebook API access service.
func NewAPI() API {
	return API{}
}
