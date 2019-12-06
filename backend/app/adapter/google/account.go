package google

import (
	"net/http"
	"short/app/entity"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const (
	googleApiURL = "https://www.googleapis.com/drive/v2/files"
)

var _ service.SSOAccount = (*Account)(nil)

// Account accesses user's account data through Google API.
type Account struct {
	http fw.HTTPRequest
}

// GetSingleSignOnUser retrieves user's email and name from Google API.
func (g Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	type response struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var result response

	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
	}
	err := g.http.JSON(http.MethodGet, googleApiURL, headers, "", &result)
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
