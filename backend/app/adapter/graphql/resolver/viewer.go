package resolver

import (
	"errors"

	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth"
)

func viewer(authToken *string, authenticator auth.Authenticator) (entity.User, error) {
	if authToken == nil {
		return entity.User{}, errors.New("auth token can't be empty")
	}

	return authenticator.GetUser(*authToken)
}
