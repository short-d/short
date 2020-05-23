package resolver

import (
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

// Query represents GraphQL query resolver
type Query struct {
	logger        logger.Logger
	authenticator authenticator.Authenticator
	changeLog     changelog.ChangeLog
	urlRetriever  shortlink.Retriever
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
	logger logger.Logger,
	authenticator authenticator.Authenticator,
	changeLog changelog.ChangeLog,
	urlRetriever shortlink.Retriever,
) Query {
	return Query{
		logger:        logger,
		authenticator: authenticator,
		changeLog:     changeLog,
		urlRetriever:  urlRetriever,
	}
}
