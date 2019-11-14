package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/usecase/url"
	"time"

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
	trace := q.tracer.BeginTrace("Query.URL")
	defer trace.End()

	var expireAt *time.Time
	if args.ExpireAfter != nil {
		expireAt = &args.ExpireAfter.Time
	}

	u, err := q.urlRetriever.GetURL(args.Alias, expireAt)
	if err != nil {
		q.logger.Error(err)
		return nil, err
	}
	return &URL{url: u}, nil
}
