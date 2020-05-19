package feature

import (
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"testing"
	"time"

	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/ctx"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestDynamicDecisionMaker_IsFeatureEnable(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name              string
		toggles           map[string]entity.Toggle
		user              entity.User
		userRoles         map[string][]role.Role
		featureID         string
		expectedIsEnabled bool
	}{
		{
			name:              "toggle not found",
			toggles:           map[string]entity.Toggle{},
			featureID:         "example-feature",
			expectedIsEnabled: false,
		},
		{
			name: "toggle disabled",
			toggles: map[string]entity.Toggle{
				"example-feature": {
					ID:        "example-feature",
					IsEnabled: false,
				},
			},
			featureID:         "example-feature",
			expectedIsEnabled: false,
		},
		{
			name: "toggle enabled",
			toggles: map[string]entity.Toggle{
				"example-feature": {
					ID:        "example-feature",
					IsEnabled: true,
				},
			},
			featureID:         "example-feature",
			expectedIsEnabled: true,
		},
		{
			name: "permission enabled",
			toggles: map[string]entity.Toggle{
				"include-admin-panel": {
					ID:        "include-admin-panel",
					IsEnabled: true,
					Type:      entity.PermissionToggle,
				},
			},
			user: entity.User{
				ID: "user-id",
			},
			userRoles: map[string][]role.Role{
				"user-id": {role.Admin},
			},
			featureID:         "include-admin-panel",
			expectedIsEnabled: true,
		},
		{
			name: "permission disabled",
			toggles: map[string]entity.Toggle{
				"include-admin-panel": {
					ID:        "include-admin-panel",
					IsEnabled: true,
					Type:      entity.PermissionToggle,
				},
			},
			user: entity.User{
				ID: "user-id",
			},
			userRoles: map[string][]role.Role{
				"user-id": {role.Basic},
			},
			featureID:         "include-admin-panel",
			expectedIsEnabled: false,
		},
		{
			name: "permission enabled, toggle disabled",
			toggles: map[string]entity.Toggle{
				"include-admin-panel": {
					ID:        "include-admin-panel",
					IsEnabled: false,
					Type:      entity.PermissionToggle,
				},
			},
			user: entity.User{
				ID: "user-id",
			},
			userRoles: map[string][]role.Role{
				"user-id": {role.Admin},
			},
			featureID:         "include-admin-panel",
			expectedIsEnabled: false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			featureRepo := repository.NewFeatureToggleFake(testCase.toggles)

			entryRepo := logger.NewEntryRepoFake()
			lg, err := logger.NewFake(logger.LogOff, &entryRepo)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(time.Now())
			mt := metrics.NewFake()
			ana := analytics.NewFake()
			ctxCh := make(chan ctx.ExecutionContext)
			go func() {
				ctxCh <- ctx.ExecutionContext{}
			}()

			currentUser := testCase.user

			fakeRolesRepo := repository.NewUserRoleFake(testCase.userRoles)
			ac := rbac.NewRBAC(fakeRolesRepo)

			ins := instrumentation.NewInstrumentation(lg, tm, mt, ana, ctxCh)
			factory := NewDynamicDecisionMakerFactory(featureRepo, authorizer.NewAuthorizer(ac))
			decision := factory.NewDecision(ins)

			gotIsEnabled := decision.IsFeatureEnable(testCase.featureID, &currentUser)
			assert.Equal(t, testCase.expectedIsEnabled, gotIsEnabled)
		})
	}
}
