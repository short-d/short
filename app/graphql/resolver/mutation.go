package resolver

import (
	"short/app/entity"
	"short/app/usecase"
	"short/fw"
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

func (m Mutation) CreateUrl(args *CreateUrlArgs) (*Url, error) {
	trace := m.tracer.BeginTrace("CreateUrl")

	url := entity.Url{
		OriginalUrl: args.Url.OriginalUrl,
		ExpireAt:    args.Url.ExpireAt,
	}

	if args.Url.CustomAlias != nil {
		customAlias := *args.Url.CustomAlias

		trace1 := trace.Next("CreateUrlWithCustomAlias")
		newUrl, err := m.urlCreator.CreateUrlWithCustomAlias(url, customAlias)
		trace1.End()

		if err != nil {
			m.logger.Error(err)
			return nil, err
		}

		return &Url{url: newUrl}, nil
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
