package authorizer

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/permission"
	"github.com/short-d/short/app/usecase/repository"
)

type Authorizer struct {
	userRoleRepo repository.UserRole
}

// HasAccess tells if a user has a right to the given permission
func (a Authorizer) HasAccess(user entity.User, permission permission.Permission) (bool, error) {
	roles, err := a.userRoleRepo.GetUserRoles(user)

	if err != nil {
		return false, err
	}

	hasAccess := false

	for _, role := range roles {
		if role.IsAllowed(permission) {
			hasAccess = true
			break
		}
	}

	return hasAccess, nil
}

func NewAuthorizer(userRoleRepo repository.UserRole) Authorizer {
	return Authorizer{userRoleRepo: userRoleRepo}
}
