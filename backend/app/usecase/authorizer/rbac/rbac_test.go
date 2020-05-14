package rbac

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/permission"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestRBAC_HasPermission(t *testing.T) {
	testCases := []struct {
		name                string
		user                entity.User
		userRoles           map[string][]role.Role
		permission          permission.Permission
		expectHasPermission bool
	}{
		{
			name:                "has no role",
			user:                entity.User{ID: "alpha"},
			userRoles:           map[string][]role.Role{},
			permission:          permission.CreateChange,
			expectHasPermission: false,
		},
		{
			name: "one of the roles has permission",
			user: entity.User{ID: "alpha"},
			userRoles: map[string][]role.Role{"alpha": {
				role.ChangeLogEditor,
				role.Admin,
			}},
			permission:          permission.CreateChange,
			expectHasPermission: true,
		},
		{
			name: "no role has permission",
			user: entity.User{ID: "alpha"},
			userRoles: map[string][]role.Role{"alpha": {
				role.Basic,
			}},
			permission:          permission.CreateChange,
			expectHasPermission: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fakeRolesRepo := repository.NewUserRoleFake(testCase.userRoles)
			ac := NewRBAC(fakeRolesRepo)

			gotHasPermission, err := ac.HasPermission(testCase.user, testCase.permission)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectHasPermission, gotHasPermission)
		})
	}
}
