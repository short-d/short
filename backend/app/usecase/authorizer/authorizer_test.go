package authorizer

import (
	"testing"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/role"
	"github.com/short-d/short/app/usecase/repository"
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

			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.hasRight, canChange)
		})
	}
}
