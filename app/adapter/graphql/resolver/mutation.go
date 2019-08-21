package resolver

import (
	"short/app/entity"
	"short/app/usecase/input"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"short/fw"
	"time"
)

type Mutation struct {
	logger            fw.Logger
	tracer            fw.Tracer
	urlCreator        url.Creator
	requesterVerifier requester.Verifier
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
	UserEmail       *string
}

func (m Mutation) CreateURL(args *CreateURLArgs) (*URL, error) {
	trace := m.tracer.BeginTrace("Mutation.CreateUrl")

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
	if !m.aliasValidator.IsValid(customAlias) {
		return nil, ErrInvalidCustomAlias{
			customAlias: customAlias,
		}
	}

	u := entity.URL{
		OriginalURL: args.URL.OriginalURL,
		ExpireAt:    args.URL.ExpireAt,
	}

	if args.URL.CustomAlias != nil {
		customAlias := *args.URL.CustomAlias

		trace1 := trace.Next("CreateUrlWithCustomAlias")
		newURL, err := m.urlCreator.CreateWithCustomAlias(u, customAlias)
		trace1.End()

		if err == nil {
			return &URL{url: newURL}, nil
		}
		m.logger.Error(err)

		switch err.(type) {
		case url.ErrAliasExist:
			return &URL{}, ErrURLAliasExist(customAlias)
		default:
			return nil, ErrUnknown{}
		}
	}

	trace1 := trace.Next("CreateUrl")
	newURL, err := m.urlCreator.Create(u)
	trace1.End()

	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	trace.End()
	return &URL{url: newURL}, nil
}

func NewMutation(
	logger fw.Logger,
	tracer fw.Tracer,
	urlCreator url.Creator,
	requesterVerifier requester.Verifier,
) Mutation {
	return Mutation{
		logger:            logger,
		tracer:            tracer,
		urlCreator:        urlCreator,
		requesterVerifier: requesterVerifier,
		longLinkValidator: input.NewLongLink(),
		aliasValidator:    input.NewCustomAlias(),
	}
}
