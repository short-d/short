package facebook

import (
	"net/http"
	"net/url"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/service"
)

const facebookAPI = "https://graph.facebook.com/me"

var _ service.SSOAccount = (*Account)(nil)

// Account accesses user's account data through FB Graph API.
type Account struct {
	httpRequest fw.HTTPRequest
}

// GetSingleSignOnUser retrieves user's email and name from Facebook API.
func (g Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	type response struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var fbResponse response

	u, err := url.Parse(facebookAPI)
	if err != nil {
		return entity.SSOUser{}, err
	}

	query := u.Query()
	query.Set("fields", "name,email")
	query.Set("access_token", accessToken)
	u.RawQuery = query.Encode()

	headers := map[string]string{}

	err = g.httpRequest.JSON(http.MethodGet, u.String(), headers, "", &fbResponse)

	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		Email: fbResponse.Email,
		Name:  fbResponse.Name,
	}, nil
}

// NewAccount initializes Facebook account API client.
func NewAccount(http fw.HTTPRequest) Account {
	return Account{
		httpRequest: http,
	}
}
