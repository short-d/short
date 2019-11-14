package resolver

import (
	"short/app/entity"
	"short/app/usecase/auth"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"time"

	"github.com/byliuyang/app/fw"
)

type Mutation struct {
	logger            fw.Logger
	tracer            fw.Tracer
	urlCreator        url.Creator
	requesterVerifier requester.Verifier
	authenticator     auth.Authenticator
}

type URLInput struct {
	OriginalURL string
	CustomAlias *string
	ExpireAt    *time.Time
}

type CreateURLArgs struct {
	CaptchaResponse string
	URL             URLInput
	AuthToken       string
}

func (m Mutation) CreateURL(args *CreateURLArgs) (*URL, error) {
	trace := m.tracer.BeginTrace("Mutation.CreateURL")
	defer trace.End()

	isHuman, err := m.requesterVerifier.IsHuman(args.CaptchaResponse)
	if err != nil {
		m.logger.Error(err)
		return nil, ErrUnknown{}
	}

	if !isHuman {
		return nil, ErrNotHuman{}
	}

	u := entity.URL{
		OriginalURL: args.URL.OriginalURL,
		ExpireAt:    args.URL.ExpireAt,
	}

	authToken := args.AuthToken
	user, err := m.authenticator.GetUser(authToken)
	if err != nil {
		m.logger.Error(err)
		return nil, ErrInvalidAuthToken(authToken)
	}

	customAlias := args.URL.CustomAlias

	trace1 := trace.Next("CreateURL")
	defer trace1.End()

	newURL, err := m.urlCreator.CreateURL(u, customAlias, user)
	if err == nil {
		return &URL{url: newURL}, nil
	}

	switch err.(type) {
	case url.ErrAliasExist:
		m.logger.Error(err)
		return nil, ErrURLAliasExist(*customAlias)
	case url.ErrInvalidLongLink:
		m.logger.Error(err)
		return nil, ErrInvalidLongLink(u.OriginalURL)
	case url.ErrInvalidCustomAlias:
		m.logger.Error(err)
		return nil, ErrInvalidCustomAlias(*customAlias)
	default:
		m.logger.Error(err)
		return nil, ErrUnknown{}
	}
}

func NewMutation(
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
