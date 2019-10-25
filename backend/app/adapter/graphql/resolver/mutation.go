package resolver

import (
	"short/app/entity"
	"short/app/usecase/auth"
	"short/app/usecase/input"
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
	longLinkValidator input.Validator
	aliasValidator    input.Validator
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
	trace := m.tracer.BeginTrace("Mutation.CreateUrl")
	defer trace.End()

	isHuman, err := m.requesterVerifier.IsHuman(args.CaptchaResponse)
	if err != nil {
		m.logger.Error(err)
		return nil, ErrUnknown{}
	}

	if !isHuman {
		return nil, ErrNotHuman{}
	}

	longLink := args.URL.OriginalURL
	if !m.longLinkValidator.IsValid(&longLink) {
		return nil, ErrInvalidLongLink(longLink)
	}

	customAlias := args.URL.CustomAlias
	if customAlias != nil && !m.aliasValidator.IsValid(customAlias) {
		return nil, ErrInvalidCustomAlias(*customAlias)
	}

	u := entity.URL{
		OriginalURL: args.URL.OriginalURL,
		ExpireAt:    args.URL.ExpireAt,
	}

	authToken := args.AuthToken
	userEmail, err := m.authenticator.GetUserEmail(authToken)
	if err != nil {
		m.logger.Error(err)
		return nil, ErrInvalidAuthToken(authToken)
	}

	if customAlias == nil {
		trace1 := trace.Next("CreateURL")
		defer trace1.End()

		newURL, err := m.urlCreator.CreateURL(u, userEmail)
		if err != nil {
			m.logger.Error(err)
			return nil, ErrUnknown{}
		}

		return &URL{url: newURL}, nil
	}

	trace1 := trace.Next("CreateURLWithCustomAlias")
	defer trace1.End()

	newURL, err := m.urlCreator.CreateURLWithCustomAlias(u, *customAlias, userEmail)
	if err == nil {
		return &URL{url: newURL}, nil
	}

	m.logger.Error(err)
	switch err.(type) {
	case url.ErrAliasExist:
		return &URL{}, ErrURLAliasExist(*customAlias)
	default:
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
		longLinkValidator: input.NewLongLink(),
		aliasValidator:    input.NewCustomAlias(),
	}
}
