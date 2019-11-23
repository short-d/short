package repository

import "short/app/entity"

// User accesses users' information from storage, such as database.
type User interface {
	IsEmailExist(email string) (bool, error)
	GetUserByEmail(email string) (entity.User, error)
	CreateUser(user entity.User) error
	UpdateUserID(email string, userID string) error
}
