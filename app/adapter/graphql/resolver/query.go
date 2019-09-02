package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
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

type URLArgs struct {
	Alias       string
	ExpireAfter *scalar.Time
}

func (q Query) URL(args *URLArgs) (*URL, error) {
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
		return &URL{url: u}, nil
	}

	trace1 := trace.Next("GetAfter")
	u, err := q.urlRetriever.GetAfter(trace1, args.Alias, args.ExpireAfter.Time)
	trace1.End()

	if err != nil {
		q.logger.Error(err)
		return nil, err
	}

	trace.End()
	return &URL{url: u}, nil
}
