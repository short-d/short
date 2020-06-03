package search

import (
	"time"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// Search represents the search handler of short links and users
// from a persistent storage
type Search struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	timeout           time.Duration
}

// Result represents the result of a search
type Result struct {
	shortLinks []entity.ShortLink
	users      []entity.User
}

// Search searches the short links and users for given query and filter
func (s Search) Search(query Query, filter Filter) (Result, error) {
	return Result{}, nil
}

// NewSearch creates Search
func NewSearch(
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
	timeout time.Duration,
) Search {
	return Search{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
		timeout:           timeout,
	}
}
