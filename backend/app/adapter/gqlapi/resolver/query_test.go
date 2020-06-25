// +build !integration all

package resolver

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

func TestQuery_AuthQuery(t *testing.T) {
	t.Parallel()
	now := time.Now()
	auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)
	user := entity.User{
		Email: "alpha@example.com",
	}
	authToken, err := auth.GenerateToken(user)
	assert.Equal(t, nil, err)
	randomToken := "random_token"

	testCases := []struct {
		name      string
		authToken *string
		expHasErr bool
		expUser   *entity.User
	}{
		{
			name:      "with valid auth token",
			authToken: &authToken,
			expHasErr: false,
			expUser: &entity.User{
				Email: "alpha@example.com",
			},
		},
		{
			name:      "with invalid auth token",
			authToken: &randomToken,
			expHasErr: false,
		},
		{
			name:      "without auth token",
			authToken: nil,
			expHasErr: false,
			expUser:   nil,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			fakeShortLinkRepo := repository.NewShortLinkFake(map[string]entity.ShortLink{})
			fakeUserShortLinkRepo := repository.NewUserShortLinkRepoFake(nil, nil)
			auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)
			retrieverFake := shortlink.NewRetrieverPersist(&fakeShortLinkRepo, &fakeUserShortLinkRepo)
			entryRepo := logger.NewEntryRepoFake()
			lg, err := logger.NewFake(logger.LogOff, &entryRepo)
			assert.Equal(t, nil, err)

			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(now)
			changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})

			fakeRolesRepo := repository.NewUserRoleFake(map[string][]role.Role{})
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			changeLog := changelog.NewPersist(keyGen, tm, &changeLogRepo, &userChangeLogRepo, au)

			query := newQuery(lg, auth, changeLog, retrieverFake)

			assert.Equal(t, nil, err)
			authQueryArgs := AuthQueryArgs{AuthToken: testCase.authToken}
			_, err = query.AuthQuery(&authQueryArgs)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)
		})
	}
}
