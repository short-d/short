package resolver

import (
	"time"

	"github.com/short-d/short/app/adapter/graphql/scalar"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/url"
)

// AuthQuery represents GraphQL query resolver that acts differently based
// on the identify of the user
type AuthQuery struct {
	user         *entity.User
	urlRetriever url.Retriever
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

func newAuthQuery(user *entity.User, urlRetriever url.Retriever) AuthQuery {
	return AuthQuery{
		user:         user,
		urlRetriever: urlRetriever,
	}
}
