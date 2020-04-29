package authorizer

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/permission"
	"github.com/short-d/short/app/usecase/authorizer/role"
	"github.com/short-d/short/app/usecase/repository"
)

type Authorizer struct {
	userRoleRepo repository.UserRole
}

func (a Authorizer) CanCreateChange(user entity.User) bool {
	permissions := a.getUserPermissions(user)
	_, ok := permissions[permission.CreateChange]
	return ok
}

func (a Authorizer) getUserPermissions(user entity.User) map[permission.Permission]entity.Empty {
	permissions := make(map[permission.Permission]entity.Empty)
	roles, err := a.userRoleRepo.GetUserRoles(user)
	if err != nil {
		return permissions
	}

	for _, userRole := range roles {
		rolePermissions := role.Permissions[userRole]
		for _, perm := range rolePermissions {
			permissions[perm] = entity.Empty{}
		}
	}
	return permissions
}

func NewAuthorizer(userRoleRepo repository.UserRole) Authorizer {
	return Authorizer{userRoleRepo: userRoleRepo}
}
