package resolver

import (
	"short/app/entity"
	"short/app/usecase"
	"short/fw"
	"strings"
	"time"
)

type Mutation struct {
	logger     fw.Logger
	tracer     fw.Tracer
	urlCreator usecase.UrlCreator
}

type UrlInput struct {
	OriginalUrl string
	CustomAlias *string
	ExpireAt    *time.Time
}

type CreateUrlArgs struct {
	Url       *UrlInput
	UserEmail *string
}

type ErrCode string

const (
	ErrCodeUnknown             ErrCode = "unknown"
	ErrCodeAliasAlreadyExist           = "aliasAlreadyExist"
	ErrCodeOriginalUrlTooShort         = "originalUrlTooShort"
	ErrCodeWrongUriFormat              = "wrongUriFormat"
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

func isUri(text string) bool {
	return strings.Contains(text, "://")
}

func (m Mutation) CreateUrl(args *CreateUrlArgs) (*Url, error) {
	trace := m.tracer.BeginTrace("Mutation.CreateUrl")

	originalUrl := args.Url.OriginalUrl
	if len(originalUrl) < minUrlLength {
		return nil, ErrOriginalUrlTooShort(originalUrl)
	}

	if !isUri(originalUrl) {
		return nil, ErrWrongUriFormat(originalUrl)
	}

	url := entity.Url{
		OriginalUrl: args.Url.OriginalUrl,
		ExpireAt:    args.Url.ExpireAt,
	}

	if args.Url.CustomAlias != nil {
		customAlias := *args.Url.CustomAlias

		trace1 := trace.Next("CreateUrlWithCustomAlias")
		newUrl, err := m.urlCreator.CreateUrlWithCustomAlias(url, customAlias)
		trace1.End()

		if err == nil {
			return &Url{url: newUrl}, nil
		}
		m.logger.Error(err)

		switch err.(type) {
		case usecase.ErrAliasExist:
			return &Url{}, ErrUrlAliasExist(customAlias)
		default:
			return nil, ErrUnknown{}
		}
	}

	trace1 := trace.Next("CreateUrl")
	newUrl, err := m.urlCreator.CreateUrl(url)
	trace1.End()

	if err != nil {
		m.logger.Error(err)
		return nil, err
	}

	trace.End()
	return &Url{url: newUrl}, nil
}

func NewMutation(logger fw.Logger, tracer fw.Tracer, urlCreator usecase.UrlCreator) Mutation {
	return Mutation{
		logger:     logger,
		tracer:     tracer,
		urlCreator: urlCreator,
	}
}
