package resolver

import (
	"github.com/short-d/app/fw"

	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/url"
)

// Query represents GraphQL query resolver
type Query struct {
	logger        fw.Logger
	tracer        fw.Tracer
	authenticator authenticator.Authenticator
	changeLog     changelog.ChangeLog
	urlRetriever  url.Retriever
}

// AuthQueryArgs represents possible parameters for AuthQuery endpoint
type AuthQueryArgs struct {
	AuthToken *string
}

// AuthQuery extracts user information from authentication token
func (q Query) AuthQuery(args *AuthQueryArgs) (*AuthQuery, error) {
	authQuery := newAuthQuery(args.AuthToken, q.authenticator, q.changeLog, q.urlRetriever)
	return &authQuery, nil
}

func newQuery(
	logger fw.Logger,
	tracer fw.Tracer,
	authenticator authenticator.Authenticator,
	changeLog changelog.ChangeLog,
	urlRetriever url.Retriever,
) Query {
	return Query{
		logger:        logger,
		tracer:        tracer,
		authenticator: authenticator,
		changeLog:     changeLog,
		urlRetriever:  urlRetriever,
	}
}
