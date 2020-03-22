package resolver

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/auth"
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
	authenticator     auth.Authenticator
	changeLogControl  changelog.ChangeLog
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

	user, err := viewer(args.AuthToken, m.authenticator)
	if err != nil {
		return nil, err
	}

	authMutation := newAuthMutation(user, m.changeLogControl, m.urlCreator)
	return &authMutation, nil
}

func newMutation(
	logger fw.Logger,
	tracer fw.Tracer,
	changeLogControl changelog.ChangeLog,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) Mutation {
	return Mutation{
		logger:            logger,
		tracer:            tracer,
		changeLogControl:  changeLogControl,
		urlCreator:        urlCreator,
		requesterVerifier: requesterVerifier,
		authenticator:     authenticator,
	}
}
