package authenticator

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/fw/must"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestThirdPartyApp_GetApp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		existingAPIKeys []entity.APIKey
		existingApps    []entity.App
		apiKey          string
		expectHasErr    bool
		expectApp       entity.App
	}{
		{
			name:         "api key doesn't contain app ID",
			apiKey:       `{"key": "12345"}`,
			expectHasErr: true,
		},
		{
			name: "api key doesn't contain key",
			existingApps: []entity.App{
				{ID: "alpha"},
			},
			apiKey:       `{"app_id": "alpha"}`,
			expectHasErr: true,
		},
		{
			name: "app not found",
			existingApps: []entity.App{
				{ID: "alpha"},
			},
			existingAPIKeys: []entity.APIKey{
				{
					AppID: "alpha",
					Key:   "secret",
				},
			},
			apiKey:       `{"app_id": "beta","key":"secret"}`,
			expectHasErr: true,
		},
		{
			name: "key not found",
			existingApps: []entity.App{
				{ID: "alpha"},
			},
			existingAPIKeys: []entity.APIKey{
				{
					AppID: "alpha",
					Key:   "different",
				},
			},
			apiKey:       `{"app_id": "alpha","key":"secret"}`,
			expectHasErr: true,
		},
		{
			name: "disable API key",
			existingApps: []entity.App{
				{ID: "alpha"},
			},
			existingAPIKeys: []entity.APIKey{
				{
					AppID:      "alpha",
					Key:        "secret",
					IsDisabled: true,
				},
			},
			apiKey:       `{"app_id": "alpha","key":"secret"}`,
			expectHasErr: true,
		},
		{
			name: "valid API key",
			existingApps: []entity.App{
				{
					ID:        "alpha",
					Name:      "Alpha",
					CreatedAt: must.Time(t, "2020-07-17T15:04:05+07:00"),
				},
			},
			existingAPIKeys: []entity.APIKey{
				{
					AppID: "alpha",
					Key:   "secret",
				},
			},
			apiKey:       `{"app_id": "alpha","key":"secret"}`,
			expectHasErr: false,
			expectApp: entity.App{
				ID:        "alpha",
				Name:      "Alpha",
				CreatedAt: must.Time(t, "2020-07-17T15:04:05+07:00"),
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			userRoleRepo := repository.NewUserRoleFake(map[string][]role.Role{})
			r := rbac.NewRBAC(userRoleRepo)
			auth := authorizer.NewAuthorizer(r)

			tokenizer := crypto.NewTokenizerFake()
			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})

			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(time.Now())
			apiKeyRepo := repository.NewAPIKeyFake(testCase.existingAPIKeys)
			appRepo := repository.NewAppFake(testCase.existingApps)
			thirdPartyApp := NewThirdPartyApp(auth, tokenizer, keyGen, tm, &apiKeyRepo, appRepo)

			cred := Credential{APIKey: &testCase.apiKey}
			gotApp, err := thirdPartyApp.GetApp(cred)
			if testCase.expectHasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, testCase.expectApp, gotApp)
		})
	}
}

func TestThirdPartyApp_GenerateAPIKey(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		existingApps    []entity.App
		existingAPIKeys []entity.APIKey
		userRoles       map[string][]role.Role
		availableKeys   []keygen.Key
		now             time.Time
		user            entity.User
		app             entity.App
		expectHasErr    bool
		expectKeyExist  bool
		expectAPIKey    entity.APIKey
		expectKey       string
	}{
		{
			name: "App doesn't exist",
			existingApps: []entity.App{
				{ID: "app"},
			},
			existingAPIKeys: []entity.APIKey{},
			userRoles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			availableKeys: []keygen.Key{"secret"},
			user:          entity.User{ID: "alpha"},
			app:           entity.App{ID: "coder"},
			expectHasErr:  true,
		},
		{
			name: "Duplicate key for the same app",
			existingApps: []entity.App{
				{ID: "app"},
			},
			existingAPIKeys: []entity.APIKey{
				{AppID: "app", Key: "secret"},
			},
			userRoles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			availableKeys:  []keygen.Key{"secret"},
			user:           entity.User{ID: "alpha"},
			app:            entity.App{ID: "app"},
			expectHasErr:   true,
			expectKeyExist: true,
		},
		{
			name: "Out of available keys",
			existingApps: []entity.App{
				{ID: "app"},
			},
			existingAPIKeys: []entity.APIKey{},
			userRoles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			availableKeys: []keygen.Key{},
			user:          entity.User{ID: "alpha"},
			app:           entity.App{ID: "app"},
			expectHasErr:  true,
		},
		{
			name: "No permission to generate API key",
			existingApps: []entity.App{
				{ID: "app"},
			},
			existingAPIKeys: []entity.APIKey{},
			userRoles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			availableKeys: []keygen.Key{"alpha"},
			user:          entity.User{ID: "alpha"},
			app:           entity.App{ID: "app"},
			expectHasErr:  true,
		},
		{
			name: "Successfully generated API key",
			existingApps: []entity.App{
				{ID: "app"},
			},
			existingAPIKeys: []entity.APIKey{},
			userRoles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			availableKeys: []keygen.Key{"secret"},
			now:           must.Time(t, "2020-07-17T15:04:05+07:00"),
			user:          entity.User{ID: "alpha"},
			app:           entity.App{ID: "app"},
			expectHasErr:  false,
			expectAPIKey: entity.APIKey{
				AppID:      "app",
				Key:        "secret",
				IsDisabled: false,
				CreatedAt:  must.Time(t, "2020-07-17T15:04:05+07:00"),
			},
			expectKey: `{"app_id":"app","key":"secret"}`,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			userRoleRepo := repository.NewUserRoleFake(testCase.userRoles)
			r := rbac.NewRBAC(userRoleRepo)
			auth := authorizer.NewAuthorizer(r)

			tokenizer := crypto.NewTokenizerFake()
			keyFetcher := keygen.NewKeyFetcherFake(testCase.availableKeys)

			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(testCase.now)
			apiKeyRepo := repository.NewAPIKeyFake(testCase.existingAPIKeys)
			appRepo := repository.NewAppFake(testCase.existingApps)
			thirdPartyApp := NewThirdPartyApp(auth, tokenizer, keyGen, tm, &apiKeyRepo, appRepo)

			gotKey, err := thirdPartyApp.GenerateAPIKey(testCase.user, testCase.app)
			if testCase.expectHasErr {
				assert.NotEqual(t, nil, err)
				return
			}

			assert.Equal(t, testCase.expectKey, gotKey)

			apiKey, err := apiKeyRepo.GetAPIKey(testCase.app.ID, string(testCase.availableKeys[0]))
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectAPIKey, apiKey)
		})
	}
}
