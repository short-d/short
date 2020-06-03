package search

import (
	"time"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// Search finds different types of resources matching certain criteria and sort them based on predefined orders.
type Search struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	timeout           time.Duration
}

// Result represents the result of a search query.
type Result struct {
	shortLinks []entity.ShortLink
	users      []entity.User
}

// Search finds resources based on specified criteria.
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
