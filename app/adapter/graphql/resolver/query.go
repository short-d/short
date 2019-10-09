package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/usecase/auth"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

type Query struct {
	logger        fw.Logger
	tracer        fw.Tracer
	urlRetriever  url.Retriever
	authenticator auth.Authenticator
}

func NewQuery(logger fw.Logger, tracer fw.Tracer, urlRetriever url.Retriever, authenticator auth.Authenticator) Query {
	return Query{
		logger:        logger,
		tracer:        tracer,
		urlRetriever:  urlRetriever,
		authenticator: authenticator,
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

type ListURLsArgs struct {
	AuthToken string
}

func (q Query) ListURLs(args *ListURLsArgs) ([]*URL, error) {
	trace := q.tracer.BeginTrace("Query.listURLs")

	authToken := args.AuthToken
	userEmail, err := q.authenticator.GetUserEmail(authToken)
	if err != nil {
		q.logger.Error(err)
		return nil, ErrInvalidAuthToken(authToken)
	}

	trace1 := trace.Next("Get")
	list, err := q.urlRetriever.GetList(trace1, userEmail)
	trace1.End()

	if err != nil {
		q.logger.Error(err)
		return nil, err
	}

	urls := make([]*URL, 0)
	for _, u := range list {
		urls = append(urls, &URL{url: u})
	}

	trace.End()
	return urls, nil
}
