package resolver

import (
	"short/app/usecase/auth"
	"short/app/usecase/requester"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

type Mutation struct {
	logger            fw.Logger
	tracer            fw.Tracer
	urlCreator        url.Creator
	requesterVerifier requester.Verifier
	authenticator     auth.Authenticator
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

	authMutation := newAuthMutation(user, m.urlCreator)
	return &authMutation, nil
}

func newMutation(
	logger fw.Logger,
	tracer fw.Tracer,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
	authenticator auth.Authenticator,
) Mutation {
	return Mutation{
		logger:            logger,
		tracer:            tracer,
		urlCreator:        urlCreator,
		requesterVerifier: requesterVerifier,
		authenticator:     authenticator,
	}
}
