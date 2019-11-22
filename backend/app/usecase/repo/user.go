package repo

import "short/app/entity"

// User access users' information from storage, such as database.
type User interface {
	IsEmailExist(email string) (bool, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user entity.User) error
}
