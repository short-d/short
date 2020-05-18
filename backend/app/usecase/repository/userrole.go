package repository

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
)

// UserRole accesses users' role information from storage, such as database.
type UserRole interface {
	GetRoles(user entity.User) ([]role.Role, error)
	AddRole(user entity.User, role role.Role) error
	DeleteRole(user entity.User, role role.Role) error
}
