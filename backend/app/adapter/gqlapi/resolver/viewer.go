package resolver

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
)

func viewer(authToken *string, auth authenticator.Authenticator) (entity.User, error) {
	if authToken == nil {
		return entity.User{}, errors.New("auth token can't be empty")
	}

	return auth.GetUser(*authToken)
}
