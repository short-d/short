package resolver

import (
	"time"

	"github.com/short-d/short/app/adapter/graphql/scalar"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/url"
)

// AuthQuery represents GraphQL query resolver that acts differently based
// on the identify of the user
type AuthQuery struct {
	authToken     *string
	authenticator auth.Authenticator
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
	changeLog, err := v.changeLog.GetChangeLog()
	lastViewedAt := v.changeLog.GetLastViewedAt()
	return newChangeLog(changeLog, lastViewedAt), err
}

// URLs retrieves urls created by a given user from persistent storage
func (v AuthQuery) URLs() ([]URL, error) {
	user, err := viewer(v.authToken, v.authenticator)
	if err != nil {
		return []URL{}, err
	}

	urls, err := v.urlRetriever.GetURLsByUser(user)
	if err != nil {
		return []URL{}, err
	}

	listOfURLs := []URL{}
	for _, v := range urls {
		listOfURLs = append(listOfURLs, newURL(v))
	}

	return listOfURLs, nil
}

func newAuthQuery(
	authToken *string,
	authenticator auth.Authenticator,
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
