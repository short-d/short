package resolver

import (
	"time"

	"github.com/short-d/short/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/url"
)

// AuthQuery represents GraphQL query resolver that acts differently based
// on the identify of the user
type AuthQuery struct {
	authToken     *string
	authenticator authenticator.Authenticator
	changeLog     changelog.ChangeLog
	urlRetriever  url.Retriever
}

// URLArgs represents possible parameters for URL endpoint
type URLArgs struct {
	Alias       string
	ExpireAfter *scalar.Time
}

// URL retrieves an URL persistent storage given alias and expiration time.
func (v AuthQuery) URL(args *URLArgs) (*URL, error) {
	var expireAt *time.Time
	if args.ExpireAfter != nil {
		expireAt = &args.ExpireAfter.Time
	}

	u, err := v.urlRetriever.GetURL(args.Alias, expireAt)
	if err != nil {
		return nil, err
	}
	return &URL{url: u}, nil
}

// ChangeLog retrieves full ChangeLog from persistent storage
func (v AuthQuery) ChangeLog() (ChangeLog, error) {
	_, err := viewer(v.authToken, v.authenticator)
	if err != nil {
		return newChangeLog([]entity.Change{}, nil), ErrInvalidAuthToken{}
	}

	changeLog, err := v.changeLog.GetChangeLog()
	lastViewedAt := v.changeLog.GetLastViewedAt()
	return newChangeLog(changeLog, lastViewedAt), err
}

// URLs retrieves urls created by a given user from persistent storage
func (v AuthQuery) URLs() ([]URL, error) {
	user, err := viewer(v.authToken, v.authenticator)
	if err != nil {
		return []URL{}, ErrInvalidAuthToken{}
	}

	urls, err := v.urlRetriever.GetURLsByUser(user)
	if err != nil {
		return []URL{}, err
	}

	var gqlURLs []URL
	for _, v := range urls {
		gqlURLs = append(gqlURLs, newURL(v))
	}

	return gqlURLs, nil
}

func newAuthQuery(
	authToken *string,
	authenticator authenticator.Authenticator,
	changeLog changelog.ChangeLog,
	urlRetriever url.Retriever,
) AuthQuery {
	return AuthQuery{
		authToken:     authToken,
		authenticator: authenticator,
		changeLog:     changeLog,
		urlRetriever:  urlRetriever,
	}
}
