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

// CanCreateChange decides whether a user is allowed to create a change.
func (a Authorizer) CanCreateChange(user entity.User) (bool, error) {
	return a.rbac.HasPermission(user, permission.CreateChange)
}

// CanGetChanges decides whether a user is allowed to get changes.
func (a Authorizer) CanGetChanges(user entity.User) (bool, error) {
	return a.rbac.HasPermission(user, permission.ViewChange)
}

// CanDeleteChange decides whether a user is allowed to delete a change.
func (a Authorizer) CanDeleteChange(user entity.User) (bool, error) {
	return a.rbac.HasPermission(user, permission.DeleteChange)
}

// CanUpdateChange decides whether a user is allowed to update a change.
func (a Authorizer) CanUpdateChange(user entity.User) (bool, error) {
	return a.rbac.HasPermission(user, permission.EditChange)
}

// CanViewAdminPanel decides whether a user is allowed to view admin panel.
func (a Authorizer) CanViewAdminPanel(user entity.User) (bool, error) {
	return a.rbac.HasPermission(user, permission.ViewAdminPanel)
}

// CanGenerateAPIKey decides whether a user is allowed to generate a new api key.
func (a Authorizer) CanGenerateAPIKey(user entity.User) (bool, error) {
	return a.rbac.HasPermission(user, permission.CreateAPIKey)
}

// NewAuthorizer creates a new Authorizer object
func NewAuthorizer(rbac rbac.RBAC) Authorizer {
	return Authorizer{rbac: rbac}
}
