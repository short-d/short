// +build !integration all

package resolver

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/external"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/repository"
	"github.com/short-d/short/app/usecase/url"
)

func TestQuery_AuthQuery(t *testing.T) {
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
		t.Run(testCase.name, func(t *testing.T) {
			fakeURLRepo := repository.NewURLFake(map[string]entity.URL{})
			fakeUserURLRelationRepo := repository.NewUserURLRepoFake(nil, nil)
			auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)
			retrieverFake := url.NewRetrieverPersist(&fakeURLRepo, &fakeUserURLRelationRepo)
			entryRepo := logger.NewEntryRepoFake()
			lg, err := logger.NewFake(logger.LogOff, &entryRepo)

			keyFetcher := external.NewKeyFetcherFake([]external.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(now)
			changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
			changeLog := changelog.NewPersist(keyGen, tm, &changeLogRepo)

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
