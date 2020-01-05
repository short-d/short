package facebook

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/service"

	fb "github.com/huandu/facebook"
)

var _ service.SSOAccount = (*Account)(nil)

// Account accesses user's account data through FB Graph API.
type Account struct {
}

// GetSingleSignOnUser retrieves user's email and name from Facebook API.
func (g Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
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

// NewAccount initializes Facebook account API client.
func NewAccount() Account {
	return Account{}
}
