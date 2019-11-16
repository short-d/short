package resolver

import (
	"short/app/entity"
)

// User stores user information used by GraphQL API
type User struct {
	user entity.User
}

// Email retrieves user's email address
func (u User) Email() *string {
	return &u.user.Email
}
