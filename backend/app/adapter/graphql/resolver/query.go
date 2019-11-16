package resolver

import (
	"short/app/usecase/auth"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

// Query represents GraphQL query resolver
type Query struct {
	logger        fw.Logger
	tracer        fw.Tracer
	authenticator auth.Authenticator
	urlRetriever  url.Retriever
}

// ViewerQueryArgs represents possible arguments for viewer endpoint
type ViewerQueryArgs struct {
	AuthToken *string
}

// AuthQuery extracts user information from authentication token
func (q Query) AuthQuery(args *ViewerQueryArgs) (*AuthQuery, error) {
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
