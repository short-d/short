package search

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type Search struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
}

type Result struct {
	shortLinks []entity.ShortLink
	users      []entity.User
}

func (s Search) Search(query Query, filter Filter) Result {
	return Result{}
}

func NewSearch(
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
) Search {
	return Search{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
	}
}
