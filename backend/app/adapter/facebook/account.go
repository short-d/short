package facebook

import (
	"net/http"
	"net/url"

	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/external"
)

const facebookAPI = "https://graph.facebook.com/me"

var _ external.SSOAccount = (*Account)(nil)

// Account accesses user's account data through FB Graph API.
type Account struct {
	httpRequest webreq.HTTP
}

// GetSingleSignOnUser retrieves user's email and name from Facebook API.
func (g Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	type response struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var fbResponse response

	u, err := url.Parse(facebookAPI)
	if err != nil {
		return entity.SSOUser{}, err
	}

	query := u.Query()
	query.Set("fields", "id,name,email")
	query.Set("access_token", accessToken)
	u.RawQuery = query.Encode()

	headers := map[string]string{}

	err = g.httpRequest.JSON(http.MethodGet, u.String(), headers, "", &fbResponse)

	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		ID:    fbResponse.ID,
		Email: fbResponse.Email,
		Name:  fbResponse.Name,
	}, nil
}

// NewAccount initializes Facebook account API client.
func NewAccount(httpRequest webreq.HTTP) Account {
	return Account{
		httpRequest: httpRequest,
	}
}
