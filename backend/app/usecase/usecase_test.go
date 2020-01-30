package usecase

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/repository"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/url"
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
		name               string
		now                time.Time
		existingURLs       map[string]entity.URL
		existingUsers      []entity.User
		oauthURL           string
		githubIDProvider   stubIDProvider
		authToken          string
		tokenValidDuration time.Duration

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
			expectedShowUserHomeCallArgs: []showUserHomeCallArgs{},
			expectedShowExternalPageCallArgs: []showExternalPageCallArgs{
				{
					link: "github_sign_in_link",
				},
			},
		},
		{
			name: "auth token doesn't have email",
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
				testCase.existingURLs,
				testCase.tokenValidDuration,
				testCase.githubIDProvider,
			)
			presenter := newMockPresenter()
			useCase.RequestGithubSignIn(testCase.authToken, &presenter)

			mdtest.Equal(t, testCase.expectedShowUserHomeCallArgs, presenter.showUserHomeCallArgs)
			mdtest.Equal(t, testCase.expectedShowExternalPageCallArgs, presenter.showExternalPageCallArgs)
		})
	}
}

func newUseCase(
	now time.Time,
	existingURLs map[string]entity.URL,
	tokenValidDuration time.Duration,
	githubIDProvider GithubIDProvider,
) UseCase {
	urlRepo := repository.NewURLFake(existingURLs)

	logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
	timer := mdtest.NewTimerFake(now)
	urlRetriever := url.NewRetrieverPersist(&urlRepo)
	tokenizer := mdtest.NewCryptoTokenizerFake()
	authenticator := auth.NewAuthenticator(tokenizer, timer, tokenValidDuration)

	useCase := NewShort(
		&logger,
		timer,
		urlRetriever,
		authenticator,
		githubIDProvider,
	)
	return useCase
}
