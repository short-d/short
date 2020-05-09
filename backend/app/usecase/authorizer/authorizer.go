package authorizer

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/permission"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// Authorizer provides the API for checking user's permissions
type Authorizer struct {
	userRoleRepo repository.UserRole
}

// CanCreateChange tells if a user has a right to create a change
func (a Authorizer) CanCreateChange(user entity.User) (bool, error) {
	return a.hasPermission(user, permission.CreateChange)
}

// hasPermission tells if a user has a right to the given permission
func (a Authorizer) hasPermission(user entity.User, permission permission.Permission) (bool, error) {
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

// NewAuthorizer creates a new Authorizer object
func NewAuthorizer(userRoleRepo repository.UserRole) Authorizer {
	return Authorizer{userRoleRepo: userRoleRepo}
}
