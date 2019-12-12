package google

import (
	"net/http"
	"short/app/entity"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const googleAPI = "https://www.googleapis.com/oauth2/v3/userinfo"

var _ service.SSOAccount = (*Account)(nil)

// Account accesses user's account data through Google API.
type Account struct {
	http fw.HTTPRequest
}

// GetSingleSignOnUser retrieves user's email and name from Google API.
func (a Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	type response struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var result response

	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}

	err := a.http.JSON(http.MethodGet, googleAPI, headers, "", &result)
	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		Email: result.Email,
		Name:  result.Name,
	}, nil
}

// NewAccount initializes Google account API client.
func NewAccount(http fw.HTTPRequest) Account {
	return Account{
		http: http,
	}
}
