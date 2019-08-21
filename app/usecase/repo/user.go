package repo

import "short/app/entity"

type User interface {
	IsEmailExist(email string) (bool, error)
	GetByEmail(email string) (entity.User, error)
	Create(user entity.User) error
}
