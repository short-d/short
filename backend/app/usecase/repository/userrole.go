package repository

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/role"
)

type UserRole interface {
	GetUserRoles(user entity.User) ([]role.Role, error)
}
