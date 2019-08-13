package resolver

import (
	"short/app/entity"
	"short/app/usecase/captcha"
	"short/app/usecase/url"
	"short/app/usecase/validator"
	"short/fw"
	"time"
)

type Mutation struct {
	logger            fw.Logger
	tracer            fw.Tracer
	urlCreator        url.Creator
	captchaVerifier   captcha.Verifier
	longLinkValidator validator.Validator
	aliasValidator    validator.Validator
}

type UrlInput struct {
	OriginalUrl string
	CustomAlias *string
	ExpireAt    *time.Time
}

type CreateUrlArgs struct {
	CaptchaResponse string
	Url             UrlInput
	UserEmail       *string
}

func (m Mutation) CreateUrl(args *CreateUrlArgs) (*Url, error) {
	trace := m.tracer.BeginTrace("Mutation.CreateUrl")

	isHuman, err := m.captchaVerifier.IsHuman(args.CaptchaResponse)
	if err != nil {
		m.logger.Error(err)
		return nil, ErrUnknown{}
	}

	if !isHuman {
		return nil, ErrNotHuman{}
	}

	longLink := args.Url.OriginalUrl
	if !m.longLinkValidator.IsValid(&longLink) {
		return nil, ErrInvalidLongLink(longLink)
	}

	customAlias := args.Url.CustomAlias
	if !m.aliasValidator.IsValid(customAlias) {
		return nil, ErrInvalidCustomAlias{
			customAlias: customAlias,
		}
	}

	u := entity.Url{
		OriginalUrl: args.Url.OriginalUrl,
		ExpireAt:    args.Url.ExpireAt,
	}

	if args.Url.CustomAlias != nil {
		customAlias := *args.Url.CustomAlias

		trace1 := trace.Next("CreateUrlWithCustomAlias")
		newUrl, err := m.urlCreator.CreateWithCustomAlias(u, customAlias)
		trace1.End()

		if err == nil {
			return &Url{url: newUrl}, nil
		}
		m.logger.Error(err)

		switch err.(type) {
		case url.ErrAliasExist:
			return &Url{}, ErrUrlAliasExist(customAlias)
		default:
			return nil, ErrUnknown{}
		}
	}

	trace1 := trace.Next("CreateUrl")
	newUrl, err := m.urlCreator.Create(u)
	trace1.End()

	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	trace.End()
	return &Url{url: newUrl}, nil
}

func NewMutation(
	logger fw.Logger,
	tracer fw.Tracer,
	urlCreator url.Creator,
	captchaVerifier captcha.Verifier,
) Mutation {
	return Mutation{
		logger:            logger,
		tracer:            tracer,
		urlCreator:        urlCreator,
		captchaVerifier:   captchaVerifier,
		longLinkValidator: validator.NewLongLink(),
		aliasValidator:    validator.CustomAlias{},
	}
}
