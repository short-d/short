package rbac

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/permission"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// RBAC represents Role-based access control authorization policy.
type RBAC struct {
	userRoleRepo repository.UserRole
}

// HasPermission checks whether an user has a the given permission.
func (a RBAC) HasPermission(user entity.User, permission permission.Permission) (bool, error) {
	roles, err := a.userRoleRepo.GetRoles(user)
	var entryErr repository.ErrEntryNotFound
	if errors.As(err, &entryErr) {
		return false, nil
	}
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

// NewRBAC create RBAC.
func NewRBAC(userRoleRepo repository.UserRole) RBAC {
	return RBAC{userRoleRepo: userRoleRepo}
}
