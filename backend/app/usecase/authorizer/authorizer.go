package authorizer

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/permission"
	"github.com/short-d/short/app/usecase/repository"
)

type Authorizer struct {
	userRoleRepo repository.UserRole
}

// CanCreateChange tells if a user has a right to create a change
func (a Authorizer) CanCreateChange(user entity.User) (bool, error) {
	return a.hasPermission(user, permission.CreateChange)
}

// hasPermission tells if a user has a right to the given permission
func (a Authorizer) hasPermission(user entity.User, permission permission.Permission) (bool, error) {
	roles, err := a.userRoleRepo.GetUserRoles(user)

	if err != nil {
		return false, err
	}

	for _, role := range roles {
		if role.IsAllowed(permission) {
			return true, nil
		}
	}

	return false, nil
}

func NewAuthorizer(userRoleRepo repository.UserRole) Authorizer {
	return Authorizer{userRoleRepo: userRoleRepo}
}
