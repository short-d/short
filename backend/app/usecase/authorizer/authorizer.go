package authorizer

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/permission"
)

// Authorizer checks whether an user is granted required permissions in order
// to perform certain operations.
type Authorizer struct {
	rbac rbac.RBAC
}

// CanCreateChange decides whether an user to allowed to create a change.
func (a Authorizer) CanCreateChange(user entity.User) (bool, error) {
	return a.rbac.HasPermission(user, permission.CreateChange)
}

// NewAuthorizer creates a new Authorizer object
func NewAuthorizer(rbac rbac.RBAC) Authorizer {
	return Authorizer{rbac: rbac}
}
