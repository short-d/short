package resolver

import (
	"errors"

	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authenticator"
)

func viewer(authToken *string, authenticator authenticator.Authenticator) (entity.User, error) {
	if authToken == nil {
		return entity.User{}, errors.New("auth token can't be empty")
	}

	return authenticator.GetUser(*authToken)
}
