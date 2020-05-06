package repository

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/role"
)

type UserRole interface {
	GetRoles(user entity.User) ([]role.Role, error)
	AddRole(user entity.User, role role.Role) error
	DeleteRole(user entity.User, role role.Role) error
}
