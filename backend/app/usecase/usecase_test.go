package usecase

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/service"
)

var _ service.IdentityProvider = (*stubIDProvider)(nil)

type stubIDProvider struct {
	authorizationURL      string
	accessToken           string
	requestAccessTokenErr error
}

func (m stubIDProvider) GetAuthorizationURL() string {
	return m.authorizationURL
}

func (m stubIDProvider) RequestAccessToken(authorizationCode string) (accessToken string, err error) {
	return m.accessToken, m.requestAccessTokenErr
}

var _ service.SSOAccount = (*stubSSOAccount)(nil)

type stubSSOAccount struct {
	ssoUser       entity.SSOUser
	getSSOUserErr error
}

func (s stubSSOAccount) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	return s.ssoUser, s.getSSOUserErr
}

func TestShort_RequestGithubSignIn(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2020-01-26T08:32:40.759788656Z")
	mdtest.Equal(t, nil, err)

	testCases := []struct {
		name                             string
		now                              time.Time
		existingURLs                     map[string]entity.URL
		existingUsers                    []entity.User
		oauthURL                         string
		githubIDProvider                 stubIDProvider
		authToken                        string
		tokenValidDuration               time.Duration
		expectedShowHomeCallArgs         []showHomeCallArgs
		expectedShowUserHomeCallArgs     []showUserHomeCallArgs
		expectedShowExternalPageCallArgs []showExternalPageCallArgs
	}{
		{
			name: "user already signed in",
			now:  now,
			githubIDProvider: stubIDProvider{
				authorizationURL: "github_sign_in_link",
				accessToken:      "access_token",
			},
			authToken:                    "auth_token",
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "github_sign_in_link",
				},
			},
		},
		{
			name: "user has not signed in",
			now:  now,
			githubIDProvider: stubIDProvider{
				authorizationURL: "github_sign_in_link",
				accessToken:      "access_token",
			},
			authToken:                    "",
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "github_sign_in_link",
				},
			},
		},
		{
			name: "auth token has no email",
			now:  now,
			githubIDProvider: stubIDProvider{
				authorizationURL: "github_sign_in_link",
				accessToken:      "access_token",
			},
			authToken: `
{
  "issued_at": "2020-01-26T08:32:40.759788656Z"
}
`,
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "github_sign_in_link",
				},
			},
		},
		{
			name: "auth token has empty email",
			now:  now,
			githubIDProvider: stubIDProvider{
				authorizationURL: "github_sign_in_link",
				accessToken:      "access_token",
			},
			authToken: `
{
  "email": "",
  "issued_at": "2020-01-26T08:32:40.759788656Z"
}
`,
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "github_sign_in_link",
				},
			},
		},
		{
			name: "auth token expired",
			now:  now,
			githubIDProvider: stubIDProvider{
				authorizationURL: "github_sign_in_link",
				accessToken:      "access_token",
			},
			authToken: `
{
  "email": "byliuyang11@gmail.com",
  "issued_at": "2020-01-26T07:31:40.759788656Z"
}
`,
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "github_sign_in_link",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			useCase := newUseCase(
				testCase.now,
				testCase.tokenValidDuration,
				testCase.githubIDProvider,
				stubIDProvider{},
			)
			presenter := newMockPresenter()
			useCase.RequestGithubSignIn(testCase.authToken, &presenter)

			mdtest.Equal(t, testCase.expectedShowHomeCallArgs, presenter.showHomeCallArgs)
			mdtest.Equal(t, testCase.expectedShowUserHomeCallArgs, presenter.showUserHomeCallArgs)
			mdtest.Equal(t, testCase.expectedShowExternalPageCallArgs, presenter.showExternalPageCallArgs)
		})
	}
}

func TestShort_RequestFacebookSignIn(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2020-01-26T08:32:40.759788656Z")
	mdtest.Equal(t, nil, err)

	testCases := []struct {
		name                             string
		now                              time.Time
		existingURLs                     map[string]entity.URL
		existingUsers                    []entity.User
		oauthURL                         string
		facebookIDProvider               stubIDProvider
		authToken                        string
		tokenValidDuration               time.Duration
		expectedShowHomeCallArgs         []showHomeCallArgs
		expectedShowUserHomeCallArgs     []showUserHomeCallArgs
		expectedShowExternalPageCallArgs []showExternalPageCallArgs
	}{
		{
			name: "user already signed in",
			now:  now,
			facebookIDProvider: stubIDProvider{
				authorizationURL: "facebook_sign_in_link",
				accessToken:      "access_token",
			},
			authToken:                    "auth_token",
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "facebook_sign_in_link",
				},
			},
		},
		{
			name: "user has not signed in",
			now:  now,
			facebookIDProvider: stubIDProvider{
				authorizationURL: "facebook_sign_in_link",
				accessToken:      "access_token",
			},
			authToken:                    "",
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "facebook_sign_in_link",
				},
			},
		},
		{
			name: "auth token has no email",
			now:  now,
			facebookIDProvider: stubIDProvider{
				authorizationURL: "facebook_sign_in_link",
				accessToken:      "access_token",
			},
			authToken: `
{
  "issued_at": "2020-01-26T08:32:40.759788656Z"
}
`,
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "facebook_sign_in_link",
				},
			},
		},
		{
			name: "auth token has empty email",
			now:  now,
			facebookIDProvider: stubIDProvider{
				authorizationURL: "facebook_sign_in_link",
				accessToken:      "access_token",
			},
			authToken: `
{
  "email": "",
  "issued_at": "2020-01-26T08:32:40.759788656Z"
}
`,
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "facebook_sign_in_link",
				},
			},
		},
		{
			name: "auth token expired",
			now:  now,
			facebookIDProvider: stubIDProvider{
				authorizationURL: "facebook_sign_in_link",
				accessToken:      "access_token",
			},
			authToken: `
{
  "email": "byliuyang11@gmail.com",
  "issued_at": "2020-01-26T07:31:40.759788656Z"
}
`,
			tokenValidDuration:           time.Hour,
			expectedShowHomeCallArgs:     []showHomeCallArgs{},
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "facebook_sign_in_link",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			useCase := newUseCase(
				testCase.now,
				testCase.tokenValidDuration,
				stubIDProvider{},
				testCase.facebookIDProvider,
			)
			presenter := newMockPresenter()
			useCase.RequestFacebookSignIn(testCase.authToken, &presenter)

			mdtest.Equal(t, testCase.expectedShowHomeCallArgs, presenter.showHomeCallArgs)
			mdtest.Equal(t, testCase.expectedShowUserHomeCallArgs, presenter.showUserHomeCallArgs)
			mdtest.Equal(t, testCase.expectedShowExternalPageCallArgs, presenter.showExternalPageCallArgs)
		})
	}
}

func newUseCase(
	now time.Time,
	tokenValidDuration time.Duration,
	githubIDProvider GithubIDProvider,
	facebookIDProvider FacebookIDProvider,
) UseCase {
	logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
	timer := mdtest.NewTimerFake(now)
	tokenizer := mdtest.NewCryptoTokenizerFake()
	authenticator := auth.NewAuthenticator(tokenizer, timer, tokenValidDuration)

	useCase := NewUseCase(
		&logger,
		timer,
		authenticator,
		githubIDProvider,
		facebookIDProvider,
	)
	return useCase
}
