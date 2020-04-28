// +build !integration all

package resolver

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/repository"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/url"
)

func TestQuery_AuthQuery(t *testing.T) {
	now := time.Now()
	auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)
	user := entity.User{
		Email: "alpha@example.com",
	}
	authToken, err := auth.GenerateToken(user)
	mdtest.Equal(t, nil, err)
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
			sqlDB, _, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer sqlDB.Close()

			fakeURLRepo := repository.NewURLFake(map[string]entity.URL{})
			fakeUserURLRelationRepo := repository.NewUserURLRepoFake(nil, nil)
			auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)
			retrieverFake := url.NewRetrieverPersist(&fakeURLRepo, &fakeUserURLRelationRepo)
			logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
			tracer := mdtest.NewTracerFake()

			keyFetcher := service.NewKeyFetcherFake([]service.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			mdtest.Equal(t, nil, err)

			timerFake := mdtest.NewTimerFake(now)
			changeLogRepo := db.NewChangeLogSQL(sqlDB)
			changeLog := changelog.NewPersist(keyGen, timerFake, changeLogRepo)

			query := newQuery(&logger, &tracer, auth, changeLog, retrieverFake)

			mdtest.Equal(t, nil, err)
			authQueryArgs := AuthQueryArgs{AuthToken: testCase.authToken}
			_, err = query.AuthQuery(&authQueryArgs)
			if testCase.expHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
		})
	}
}
