package provider

import (
	"time"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/search"
)

// SearchTimeout represents timeout duration of a search request.
type SearchTimeout time.Duration

// NewSearch creates Search with its dependencies.
func NewSearch(
	logger logger.Logger,
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
	timeout SearchTimeout,
) search.Search {
	return search.NewSearch(logger, shortLinkRepo, userShortLinkRepo, time.Duration(timeout))
}
