package rbac

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/permission"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type RBAC struct {
	userRoleRepo repository.UserRole
}

func (a RBAC) HasPermission(user entity.User, permission permission.Permission) (bool, error) {
	roles, err := a.userRoleRepo.GetRoles(user)
	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role.HasPermission(permission) {
			return true, nil
		}
	}
	return false, nil
}

func NewRBAC(userRoleRepo repository.UserRole) RBAC {
	return RBAC{userRoleRepo: userRoleRepo}
}
