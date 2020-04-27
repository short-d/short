package resolver

import (
	"github.com/short-d/app/fw"

	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/requester"
	"github.com/short-d/short/app/usecase/url"
)

// Mutation represents GraphQL mutation resolver
type Mutation struct {
	logger            fw.Logger
	tracer            fw.Tracer
	urlCreator        url.Creator
	requesterVerifier requester.Verifier
	authenticator     authenticator.Authenticator
	changeLog         changelog.ChangeLog
}

// AuthMutationArgs represents possible parameters for AuthMutation endpoint
type AuthMutationArgs struct {
	AuthToken       *string
	CaptchaResponse string
}

// AuthMutation extracts user information from authentication token
func (m Mutation) AuthMutation(args *AuthMutationArgs) (*AuthMutation, error) {
	isHuman, err := m.requesterVerifier.IsHuman(args.CaptchaResponse)

	if err != nil {
		return nil, ErrUnknown{}
	}

	if !isHuman {
		return nil, ErrNotHuman{}
	}

	authMutation := newAuthMutation(args.AuthToken, m.authenticator, m.changeLog, m.urlCreator)
	return &authMutation, nil
}

func newMutation(
	logger fw.Logger,
	tracer fw.Tracer,
	changeLog changelog.ChangeLog,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator authenticator.Authenticator,
) Mutation {
	return Mutation{
		logger:            logger,
		tracer:            tracer,
		changeLog:         changeLog,
		urlCreator:        urlCreator,
		requesterVerifier: requesterVerifier,
		authenticator:     authenticator,
	}
}
