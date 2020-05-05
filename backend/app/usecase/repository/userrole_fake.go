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

func (u UserRoleFake) AddRole(user entity.User, r role.Role) error {
	if _, ok := u.userRoles[user.ID]; !ok {
		return errors.New("user not found")
	}

	u.userRoles[user.ID] = append(u.userRoles[user.ID], r)
	return nil
}

func (u UserRoleFake) DeleteRole(user entity.User, r role.Role) error {
	roles, ok := u.userRoles[user.ID]

	if !ok {
		return errors.New("user not found")
	}

	for i, v := range roles {
		if v == r {
			roles[i], roles[len(roles)-1] = roles[len(roles)-1], roles[i]
			roles = roles[:len(roles)-1]
			break
		}
	}

	u.userRoles[user.ID] = roles

	return nil
}

func NewUserRoleFake(userRoles map[string][]role.Role) UserRoleFake {
	return UserRoleFake{userRoles: userRoles}
}
