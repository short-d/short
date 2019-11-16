package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/usecase/auth"
	"short/app/usecase/url"
	"time"

	"github.com/byliuyang/app/fw"
)

// Query represents GraphQL query resolver
type Query struct {
	logger        fw.Logger
	tracer        fw.Tracer
	authenticator auth.Authenticator
	urlRetriever  url.Retriever
}

// NewQuery creates Query resolver
func NewQuery(
	logger fw.Logger,
	tracer fw.Tracer,
	authenticator auth.Authenticator,
	urlRetriever url.Retriever,
) Query {
	return Query{
		logger:        logger,
		tracer:        tracer,
		authenticator: authenticator,
		urlRetriever:  urlRetriever,
	}
}

// ViewerArgs represents possible arguments for viewer endpoint
type ViewerArgs struct {
	AuthToken *string
}

// Viewer extracts user information from authentication token
func (q Query) Viewer(args *ViewerArgs) (*User, error) {
	if args.AuthToken == nil {
		return nil, nil
	}

	authToken := *args.AuthToken
	user, err := q.authenticator.GetUser(authToken)
	if err != nil {
		return nil, err
	}
	return &User{
		user: user,
	}, nil
}

// URLArgs represents possible argument for URL endpoint
type URLArgs struct {
	Alias       string
	ExpireAfter *scalar.Time
}

// URL retrieves an URL persistent storage given alias and expiration time.
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
