package resolver

import (
	"short/app/entity"
	"short/app/usecase/captcha"
	"short/app/usecase/url"
	"short/fw"
	"strings"
	"time"
)

type Mutation struct {
	logger          fw.Logger
	tracer          fw.Tracer
	urlCreator      url.Creator
	captchaVerifier captcha.Verifier
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

type ErrCode string

const (
	ErrCodeUnknown             ErrCode = "unknown"
	ErrCodeAliasAlreadyExist           = "aliasAlreadyExist"
	ErrCodeOriginalUrlTooShort         = "originalUrlTooShort"
	ErrCodeWrongUriFormat              = "wrongUriFormat"
	ErrCodeRequesterNotHuman           = "requestNotHuman"
)

const minUrlLength = 6

type ErrUnknown struct{}

func (e ErrUnknown) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeUnknown,
	}
}

func (e ErrUnknown) Error() string {
	return "unknown err"
}

type ErrUrlAliasExist string

func (e ErrUrlAliasExist) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":  ErrCodeAliasAlreadyExist,
		"alias": string(e),
	}
}

func (e ErrUrlAliasExist) Error() string {
	return "url alias already exists"
}

type ErrOriginalUrlTooShort string

func (e ErrOriginalUrlTooShort) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":        ErrCodeOriginalUrlTooShort,
		"originalUrl": string(e),
	}
}

func (e ErrOriginalUrlTooShort) Error() string {
	return "original url is too short"
}

type ErrWrongUriFormat string

func (e ErrWrongUriFormat) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":        ErrCodeWrongUriFormat,
		"originalUrl": string(e),
	}
}

func (e ErrWrongUriFormat) Error() string {
	return "url format is incorrect"
}

type ErrNotHuman struct{}

func (e ErrNotHuman) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeRequesterNotHuman,
	}
}

func (e ErrNotHuman) Error() string {
	return "requester is not human"
}

func isUri(text string) bool {
	return strings.Contains(text, "://")
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

	originalUrl := args.Url.OriginalUrl
	if len(originalUrl) < minUrlLength {
		return nil, ErrOriginalUrlTooShort(originalUrl)
	}

	if !isUri(originalUrl) {
		return nil, ErrWrongUriFormat(originalUrl)
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
		logger:          logger,
		tracer:          tracer,
		urlCreator:      urlCreator,
		captchaVerifier: captchaVerifier,
	}
}
