package resolver

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth"
)

func viewer(authToken *string, authenticator auth.Authenticator) (*entity.User, error) {
	if authToken == nil {
		return nil, nil
	}

	user, err := authenticator.GetUser(*authToken)
	return &user, err
}
