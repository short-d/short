package feature

import "github.com/short-d/short/backend/app/entity"

// PermissionChecker checks whether the user is allowed to access a feature
type PermissionChecker func(user entity.User) (bool, error)

type PermissionToggle string

const (
	AdminPanel PermissionToggle = "admin-panel"
)
