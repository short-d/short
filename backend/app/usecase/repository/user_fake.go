package repository

import (
	"errors"

	"github.com/short-d/short/app/entity"
)

var _ User = (*UserFake)(nil)

// UserFake represents in memory implementation of user repository.
type UserFake struct {
	users []entity.User
}

// IsEmailExist checks whether an user with given email exists in the repository.
func (u UserFake) IsEmailExist(email string) (bool, error) {
	for _, user := range u.users {
		if user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

// IsUserIDExist checks whether an user with given ID exists in the repository.
func (u UserFake) IsUserIDExist(userID string) bool {
	for _, user := range u.users {
		if user.ID == userID {
			return true
		}
	}
	return false
}

// GetUserByEmail finds an user with a given email.
func (u UserFake) GetUserByEmail(email string) (entity.User, error) {
	for _, user := range u.users {
		if user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, errors.New("email not found")
}

// CreateUser creates and persists user in the repository for future access.
func (u *UserFake) CreateUser(user entity.User) error {
	for _, user := range u.users {
		if user.Email == user.Email {
			return errors.New("user exists")
		}
	}
	u.users = append(u.users, user)
	return nil
}

// UpdateUserID updates the ID of an user in the repository.
func (u *UserFake) UpdateUserID(email string, userID string) error {
	for idx, user := range u.users {
		if user.Email == email {
			u.users[idx].ID = userID
			return nil
		}
	}
	return errors.New("email does not exist")
}

// NewUserFake create in memory user repository implementation.
func NewUserFake(users []entity.User) UserFake {
	return UserFake{
		users: users,
	}
}
