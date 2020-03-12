package search

import (
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

// Search fetches URLs and users from persistent
// storage, such as database
type Search struct {
	userURLRepo repository.UserURLRelation
	urlRepo     repository.URL
}

// SearchForURLs fetches all URLs for a given user
func (s Search) SearchForURLs(user entity.User) ([]entity.URL, error) {
	aliases, err := s.userURLRepo.FindAliasesByUser(user)
	if err != nil {
		return nil, err
	}

	urls, err := s.urlRepo.GetByAliases(aliases)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

// NewSearch creates Search API
func NewSearch(urlRepo repository.URL, userUrlRepo repository.UserURLRelation) Search {
	return Search{
		userURLRepo: userUrlRepo,
		urlRepo:     urlRepo,
	}
}
