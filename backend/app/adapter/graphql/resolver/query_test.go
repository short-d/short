// +build !integration all

package resolver

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/repository"
	"github.com/short-d/short/app/usecase/url"
)

func TestQuery_AuthQuery(t *testing.T) {
	authenticator := auth.NewAuthenticatorFake(time.Now(), time.Hour)
	user := entity.User{
		Email: "alpha@example.com",
	}
	authToken, err := authenticator.GenerateToken(user)
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
			expHasErr: true,
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

			fakeRepo := repository.NewURLFake(map[string]entity.URL{})
			authenticator := auth.NewAuthenticatorFake(time.Now(), time.Hour)
			retrieverFake := url.NewRetrieverPersist(&fakeRepo)
			logger := mdtest.NewLoggerFake()
			tracer := mdtest.NewTracerFake()
			query := newQuery(&logger, &tracer, authenticator, retrieverFake)

			mdtest.Equal(t, nil, err)
			authQueryArgs := AuthQueryArgs{AuthToken: testCase.authToken}
			authQuery, err := query.AuthQuery(&authQueryArgs)
			if testCase.expHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expUser, authQuery.user)
		})
	}
}
