package google

import (
	"short/app/entity"
	"short/app/usecase/service"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/byliuyang/app/fw"
)

var _ service.SSOAccount = (*Account)(nil)

// Account accesses user's account data through Google API.
type Account struct {
	http fw.HTTPRequest
}

// GetSingleSignOnUser retrieves user's email and name from ID token.
func (a Account) GetSingleSignOnUser(IDToken string) (entity.SSOUser, error) {
	claims := jwt.MapClaims{}
	_, _, err := new(jwt.Parser).ParseUnverified(IDToken, claims)

	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
		ID:    claims["sub"].(string),
	}, nil
}

// NewAccount initializes Google account API client.
func NewAccount(http fw.HTTPRequest) Account {
	return Account{
		http: http,
	}
}
