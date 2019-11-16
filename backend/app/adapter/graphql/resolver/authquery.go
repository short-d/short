package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/entity"
	"short/app/usecase/url"
	"time"
)

type AuthQuery struct {
	user         *entity.User
	urlRetriever url.Retriever
}

// URLArgs represents possible argument for URL endpoint
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
