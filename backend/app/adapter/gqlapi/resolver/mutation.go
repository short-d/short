package resolver

import (
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/emotic"
	"github.com/short-d/short/backend/app/usecase/requester"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

// Mutation represents GraphQL mutation resolver
type Mutation struct {
	logger            logger.Logger
	shortLinkCreator  shortlink.Creator
	shortLinkUpdater  shortlink.Updater
	requesterVerifier requester.Verifier
	authenticator     authenticator.Authenticator
	cloudApiAuth      authenticator.CloudAPI
	changeLog         changelog.ChangeLog
	feedback          emotic.Feedback
}

// AuthMutationArgs represents possible parameters for AuthMutation endpoint
type AuthMutationArgs struct {
	AuthToken       *string
	CaptchaResponse *string
	ApiKey          *string
}

// AuthMutation extracts user information from authentication token
func (m Mutation) AuthMutation(args *AuthMutationArgs) (*AuthMutation, error) {
	if args.ApiKey == nil {
		if args.CaptchaResponse == nil {
			return nil, ErrNotHuman{}
		}

		isHuman, err := m.requesterVerifier.IsHuman(*args.CaptchaResponse)
		if err != nil {
			return nil, ErrUnknown{}
		}

		if !isHuman {
			return nil, ErrNotHuman{}
		}
	}

	authMutation := newAuthMutation(
		args.AuthToken,
		args.ApiKey,
		m.authenticator,
		m.cloudApiAuth,
		m.changeLog,
		m.shortLinkCreator,
		m.shortLinkUpdater,
		m.feedback,
	)
	return &authMutation, nil
}

func newMutation(
	logger logger.Logger,
	changeLog changelog.ChangeLog,
	shortLinkCreator shortlink.Creator,
	shortLinkUpdater shortlink.Updater,
	feedback emotic.Feedback,
	requesterVerifier requester.Verifier,
	authenticator authenticator.Authenticator,
	cloudApiAuth authenticator.CloudAPI,
) Mutation {
	return Mutation{
		logger:            logger,
		changeLog:         changeLog,
		shortLinkCreator:  shortLinkCreator,
		shortLinkUpdater:  shortLinkUpdater,
		feedback: feedback,
		requesterVerifier: requesterVerifier,
		authenticator:     authenticator,
		cloudApiAuth:      cloudApiAuth,
	}
}
