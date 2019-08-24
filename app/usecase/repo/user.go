package repo

import "short/app/entity"

// User access users' information from storage, such as database.
type User interface {
	IsEmailExist(email string) (bool, error)
	GetByEmail(email string) (entity.User, error)
	Create(user entity.User) error
}
