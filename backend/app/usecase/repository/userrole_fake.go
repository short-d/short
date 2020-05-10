package repository

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
)

var _ UserRole = (*UserRoleFake)(nil)

// UserRoleFake represents in memory implementation of UserRole repository.
type UserRoleFake struct {
	userRoles map[string][]role.Role
}

// GetRoles fetches roles for the given user
func (u UserRoleFake) GetRoles(user entity.User) ([]role.Role, error) {
	roles, ok := u.userRoles[user.ID]
	if !ok {
		return nil, ErrEntryNotFound("user not found")
	}
	return roles, nil
}

// AddRole adds the given role to the user
func (u UserRoleFake) AddRole(user entity.User, r role.Role) error {
	if _, ok := u.userRoles[user.ID]; !ok {
		return errors.New("user not found")
	}
	u.userRoles[user.ID] = append(u.userRoles[user.ID], r)
	return nil
}

// DeleteRole removes the given role from the user
func (u UserRoleFake) DeleteRole(user entity.User, r role.Role) error {
	roles, ok := u.userRoles[user.ID]
	if !ok {
		return errors.New("user not found")
	}
	var newRoles []role.Role
	for _, v := range roles {
		if v != r {
			newRoles = append(newRoles, v)
		}
	}
	u.userRoles[user.ID] = newRoles
	return nil
}

// NewUserRoleFake creates a new instance of UserRoleFake
func NewUserRoleFake(userRoles map[string][]role.Role) UserRoleFake {
	return UserRoleFake{userRoles: userRoles}
}
