package resolver

import (
	"short/app/entity"
	"short/app/graphql/scalar"
	"short/app/usecase"
	"short/fw"
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
	trace := q.tracer.BeginTrace("Url")

	var url entity.Url
	var err error

	if args.ExpireAfter == nil {
		trace1 := trace.Next("GetUrl")
		url, err = q.urlRetriever.GetUrl(trace1, args.Alias)
		trace1.End()
	} else {
		trace1 := trace.Next("GetUrlAfter")
		url, err = q.urlRetriever.GetUrlAfter(trace1, args.Alias, args.ExpireAfter.Time)
		trace1.End()
	}

	if err != nil {
		q.logger.Error(err)
		return nil, err
	}

	trace.End()
	return &Url{url: url}, nil
}
