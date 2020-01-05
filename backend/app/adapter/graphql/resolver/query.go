package resolver

import (
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/url"

	"github.com/short-d/app/fw"
)

// Query represents GraphQL query resolver
type Query struct {
	logger        fw.Logger
	tracer        fw.Tracer
	authenticator auth.Authenticator
	urlRetriever  url.Retriever
}

// AuthQueryArgs represents possible parameters for AuthQuery endpoint
type AuthQueryArgs struct {
	AuthToken *string
}

// AuthQuery extracts user information from authentication token
func (q Query) AuthQuery(args *AuthQueryArgs) (*AuthQuery, error) {
	user, err := viewer(args.AuthToken, q.authenticator)
	if err != nil {
		return nil, err
	}

	authQuery := newAuthQuery(user, q.urlRetriever)
	return &authQuery, nil
}

func newQuery(
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
