package repository

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/role"
)

type UserRole interface {
	GetUserRoles(user entity.User) ([]role.Role, error)
	AddRole(user entity.User, role role.Role) error
	DeleteRole(user entity.User, role role.Role) error
}
