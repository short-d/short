package resolver

import (
	"short/app/entity"
)

type User struct {
	user entity.User
}

func (u User) Email() *string {
	return &u.user.Email
}
