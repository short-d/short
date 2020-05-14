package repository

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
)

var _ User = (*UserFake)(nil)

// UserFake represents in memory implementation of user repository.
type UserFake struct {
	users []entity.User
}

// IsIDExist checks whether a given user id exists in the repository.
func (u UserFake) IsIDExist(id string) (bool, error) {
	for _, user := range u.users {
		if user.ID == id {
			return true, nil
		}
	}
	return false, nil
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

// GetUserByID finds an user with a given user ID.
func (u UserFake) GetUserByID(id string) (entity.User, error) {
	for _, user := range u.users {
		if user.ID == id {
			return user, nil
		}
	}
	return entity.User{}, ErrEntryNotFound("ID not found")
}

// GetUserByEmail finds an user with a given email.
func (u UserFake) GetUserByEmail(email string) (entity.User, error) {
	for _, user := range u.users {
		if user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, ErrEntryNotFound("email not found")
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

// UpdateUserID updates the ID of a user in the repository given email.
func (u *UserFake) UpdateUserID(email string, userID string) error {
	for idx, user := range u.users {
		if user.Email == email {
			u.users[idx].ID = userID
			return nil
		}
	}
	return ErrEntryNotFound("email does not exist")
}

// NewUserFake create in memory user repository implementation.
func NewUserFake(users []entity.User) UserFake {
	return UserFake{
		users: users,
	}
}
