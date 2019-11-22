package repo

import (
	"errors"
	"short/app/entity"
)

var _ User = (*UserFake)(nil)

type UserFake struct {
	users []entity.User
}

func (u UserFake) IsEmailExist(email string) (bool, error) {
	for _, user := range u.users {
		if user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

func (u UserFake) GetUserByEmail(email string) (entity.User, error) {
	for _, user := range u.users {
		if user.Email == email {
			return user, nil
		}
	}
	return entity.User{}, errors.New("email not found")
}

func (u *UserFake) CreateUser(user entity.User) error {
	for _, user := range u.users {
		if user.Email == user.Email {
			return errors.New("user exists")
		}
	}
	u.users = append(u.users, user)
	return nil
}

func NewUserFake(users []entity.User) UserFake {
	return UserFake{
		users: users,
	}
}
