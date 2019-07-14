package resolver

import (
	"tinyURL/app/entity"
	"tinyURL/app/graphql/scalar"
	"tinyURL/app/usecase"
	"tinyURL/fw"
)

type Query struct {
	logger       fw.Logger
	tracer       fw.Tracer
	urlRetriever usecase.UrlRetriever
}

func NewQuery(logger fw.Logger, tracer fw.Tracer, urlRetriever usecase.UrlRetriever) Query {
	return Query{
		logger:       logger,
		tracer:       tracer,
		urlRetriever: urlRetriever,
	}
}

type UrlArgs struct {
	Alias       string
	ExpireAfter *scalar.Time
}

func (q Query) Url(args *UrlArgs) (*Url, error) {
	var url entity.Url
	var err error

	if args.ExpireAfter == nil {
		finish := q.tracer.Begin()
		url, err = q.urlRetriever.GetUrl(args.Alias)
		finish("usecase.UrlRetriever.GetUrl")
	} else {
		finish := q.tracer.Begin()
		url, err = q.urlRetriever.GetUrlAfter(args.Alias, args.ExpireAfter.Time)
		finish("usecase.UrlRetriever.GetUrlAfter")
	}

	if err != nil {
		q.logger.Error(err)
		return nil, err
	}

	return &Url{url:url}, nil
}
