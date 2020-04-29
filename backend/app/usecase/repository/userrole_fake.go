package repository

import (
	"errors"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/role"
)

var _ UserRole = (*UserRoleFake)(nil)

type UserRoleFake struct {
	userRoles map[string][]role.Role
}

func (u UserRoleFake) GetUserRoles(user entity.User) ([]role.Role, error) {
	roles, ok := u.userRoles[user.ID]
	if !ok {
		return nil, errors.New("user not found")
	}
	return roles, nil
}

func NewUserRoleFake(userRoles map[string][]role.Role) UserRoleFake {
	return UserRoleFake{userRoles:userRoles}
}
