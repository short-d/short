package google

import (
	"fmt"
	"net/http"

	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/sso"
)

const userInfoAPI = "https://openidconnect.googleapis.com/v1/userinfo"

var _ sso.Account = (*Account)(nil)

// Account accesses user's account data through Google API.
type Account struct {
	http webreq.HTTP
}

// GetSingleSignOnUser retrieves user's email and name from Google API.
func (a Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	// https://developers.google.com/identity/protocols/OpenIDConnect#obtainuserinfo
	type response struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		ID    string `json:"sub"`
	}

	var res response
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
	}

	err := a.http.JSON(http.MethodGet, userInfoAPI, headers, "", &res)
	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		Email: res.Email,
		Name:  res.Name,
		ID:    res.ID,
	}, nil
}

// NewAccount initializes Google account API client.
func NewAccount(http webreq.HTTP) Account {
	return Account{
		http: http,
	}
}
