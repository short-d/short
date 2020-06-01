package feature

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/ctx"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestDynamicDecisionMaker_IsFeatureEnable(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name              string
		toggles           map[string]entity.Toggle
		featureID         string
		expectedIsEnabled bool
		roles             map[string][]role.Role
		user              entity.User
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
					Type:      entity.ManualToggle,
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
					Type:      entity.ManualToggle,
				},
			},
			featureID:         "example-feature",
			expectedIsEnabled: true,
		},
		{
			name: "permission toggle enabled, permission checker not defined",
			toggles: map[string]entity.Toggle{
				"example-feature": {
					ID:        "example-feature",
					IsEnabled: true,
					Type:      entity.PermissionToggle,
				},
			},
			featureID: "example-feature",
			roles: map[string][]role.Role{
				"id": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
			expectedIsEnabled: false,
		},
		{
			name: "permission toggle enabled, user has no permission",
			toggles: map[string]entity.Toggle{
				"admin-panel": {
					ID:        "admin-panel",
					IsEnabled: true,
					Type:      entity.PermissionToggle,
				},
			},
			featureID: "admin-panel",
			roles: map[string][]role.Role{
				"id": {role.Basic},
			},
			user: entity.User{
				ID: "alpha",
			},
			expectedIsEnabled: false,
		},
		{
			name: "permission toggle enabled, user has permission",
			toggles: map[string]entity.Toggle{
				"admin-panel": {
					ID:        "admin-panel",
					IsEnabled: true,
					Type:      entity.PermissionToggle,
				},
			},
			featureID: "admin-panel",
			roles: map[string][]role.Role{
				"id": {role.Admin},
			},
			user: entity.User{
				ID: "id",
			},
			expectedIsEnabled: true,
		},
		{
			name: "permission toggle disabled, user has permission",
			toggles: map[string]entity.Toggle{
				"admin-panel": {
					ID:        "admin-panel",
					IsEnabled: false,
					Type:      entity.PermissionToggle,
				},
			},
			featureID: "admin-panel",
			roles: map[string][]role.Role{
				"id": {role.Admin},
			},
			user: entity.User{
				ID: "id",
			},
			expectedIsEnabled: false,
		},
		{
			name: "permission toggle, feature enabled, user is nil",
			toggles: map[string]entity.Toggle{
				"admin-panel": {
					ID:        "admin-panel",
					IsEnabled: true,
					Type:      entity.PermissionToggle,
				},
			},
			featureID: "admin-panel",
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

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			ins := instrumentation.NewInstrumentation(lg, tm, mt, ana, ctxCh)
			factory := NewDynamicDecisionMakerFactory(featureRepo, au)
			decision := factory.NewDecision(ins)
			gotIsEnabled := decision.IsFeatureEnable(testCase.featureID, &testCase.user)
			assert.Equal(t, testCase.expectedIsEnabled, gotIsEnabled)
		})
	}
}
