package resolver

import (
	"short/app/entity"
	"short/app/usecase/auth"
	"short/app/usecase/repo"
	"short/app/usecase/url"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
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

			fakeRepo := repo.NewURLFake(map[string]entity.URL{})
			authenticator := auth.NewAuthenticatorFake(time.Now(), time.Hour)
			retrieverFake := url.NewRetrieverPersist(&fakeRepo)
			logger := mdtest.NewLoggerFake()
			tracer := mdtest.NewTracerFake()
			query := NewQuery(&logger, &tracer, authenticator, retrieverFake)

			mdtest.Equal(t, nil, err)
			viewerArgs := ViewerQueryArgs{AuthToken: testCase.authToken}
			authQuery, err := query.AuthQuery(&viewerArgs)
			if testCase.expHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expUser, authQuery.user)
		})
	}
}
