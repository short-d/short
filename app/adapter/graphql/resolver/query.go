package resolver

import (
	scalar2 "short/app/adapter/graphql/scalar"
	"short/app/usecase/url"
	"short/fw"
)

type Query struct {
	logger       fw.Logger
	tracer       fw.Tracer
	urlRetriever url.Retriever
}

func NewQuery(logger fw.Logger, tracer fw.Tracer, urlRetriever url.Retriever) Query {
	return Query{
		logger:       logger,
		tracer:       tracer,
		urlRetriever: urlRetriever,
	}
}

type UrlArgs struct {
	Alias       string
	ExpireAfter *scalar2.Time
}

func (q Query) Url(args *UrlArgs) (*Url, error) {
	trace := q.tracer.BeginTrace("Query.Url")

	if args.ExpireAfter == nil {
		trace1 := trace.Next("Get")
		u, err := q.urlRetriever.Get(trace1, args.Alias)
		trace1.End()

		if err != nil {
			q.logger.Error(err)
			return nil, err
		}

		trace.End()
		return &Url{url: u}, nil
	}

	trace1 := trace.Next("GetAfter")
	u, err := q.urlRetriever.GetAfter(trace1, args.Alias, args.ExpireAfter.Time)
	trace1.End()

	if err != nil {
		q.logger.Error(err)
		return nil, err
	}

	trace.End()
	return &Url{url: u}, nil
}
