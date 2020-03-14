package repository

import "github.com/short-d/short/app/entity"

// User accesses users' information from storage, such as database.
type User interface {
	IsIDExist(id string) (bool, error)
	IsEmailExist(email string) (bool, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user entity.User) error
	UpdateUserID(email string, userID string) error
}
