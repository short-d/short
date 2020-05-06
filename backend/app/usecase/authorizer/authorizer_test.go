package authorizer

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/role"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestAuthorizer_hasPermission(t *testing.T) {
	testCases := []struct {
		name     string
		roles    map[string][]role.Role
		user     entity.User
		hasRight bool
	}{
		{
			name: "no access",
			roles: map[string][]role.Role{
				"id": {role.Basic},
			},
			user: entity.User{
				ID: "id",
			},
			hasRight: false,
		},
		{
			name: "has access",
			roles: map[string][]role.Role{
				"id": {role.Admin},
			},
			user: entity.User{
				ID: "id",
			},
			hasRight: true,
		},
		{
			name: "has access multiple roles",
			roles: map[string][]role.Role{
				"id": {role.Basic, role.Admin},
			},
			user: entity.User{
				ID: "id",
			},
			hasRight: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			authorizer := NewAuthorizer(fakeRolesRepo)

			canChange, err := authorizer.CanCreateChange(testCase.user)

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.hasRight, canChange)
		})
	}
}
