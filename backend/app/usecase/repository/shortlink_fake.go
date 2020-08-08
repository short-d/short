package repository

import (
	"errors"
	"time"

	"github.com/short-d/short/backend/app/entity"
)

var _ ShortLink = (*ShortLinkFake)(nil)

// ShortLinkFake accesses ShortLink information in short_link table through SQL.
type ShortLinkFake struct {
	shortLinks map[string]entity.ShortLink
	// TODO(issue#958) use eventbus for propagating short link change to all related repos
	userShortLinkRepoFake *UserShortLinkFake
}

// IsAliasExist checks whether a given alias exist in short_link table.
func (s ShortLinkFake) IsAliasExist(alias string) (bool, error) {
	_, ok := s.shortLinks[alias]
	return ok, nil
}

// CreateShortLink inserts a new ShortLink into short_link table.
func (s *ShortLinkFake) CreateShortLink(shortLinkInput entity.ShortLinkInput) error {
	if shortLinkInput.CustomAlias == nil {
		return errors.New("alias empty")
	}
	customAlias := shortLinkInput.GetCustomAlias("")
	isExist, err := s.IsAliasExist(customAlias)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("alias exists")
	}
	s.shortLinks[customAlias] = entity.ShortLink{
		Alias:     customAlias,
		LongLink:  shortLinkInput.GetLongLink(""),
		ExpireAt:  shortLinkInput.ExpireAt,
		CreatedAt: shortLinkInput.CreatedAt,
	}
	return nil
}

// GetShortLinkByAlias finds an ShortLink in short_link table given alias.
func (s ShortLinkFake) GetShortLinkByAlias(alias string) (entity.ShortLink, error) {
	isExist, err := s.IsAliasExist(alias)
	if err != nil {
		return entity.ShortLink{}, err
	}
	if !isExist {
		return entity.ShortLink{}, errors.New("alias not found")
	}
	shortLink := s.shortLinks[alias]
	return shortLink, nil
}

// GetShortLinksByAliases finds all ShortLink for a list of aliases
func (s ShortLinkFake) GetShortLinksByAliases(aliases []string) ([]entity.ShortLink, error) {
	if len(aliases) == 0 {
		return []entity.ShortLink{}, nil
	}

	var shortLinks []entity.ShortLink
	for _, alias := range aliases {
		shortLink, err := s.GetShortLinkByAlias(alias)

		if err != nil {
			return shortLinks, err
		}
		shortLinks = append(shortLinks, shortLink)
	}
	return shortLinks, nil
}

// UpdateShortLink updates an existing ShortLink with new properties.
func (s ShortLinkFake) UpdateShortLink(oldAlias string, shortLinkInput entity.ShortLinkInput) (entity.ShortLink, error) {
	if shortLinkInput.CustomAlias == nil {
		return entity.ShortLink{}, errors.New("alias empty")
	}

	prevShortLink, ok := s.shortLinks[oldAlias]
	if !ok {
		return entity.ShortLink{}, errors.New("alias not found")
	}

	// TODO(issue#958) use eventbus for propagating short link change to all related repos
	err := s.userShortLinkRepoFake.UpdateAliasCascade(oldAlias, shortLinkInput)
	if err != nil {
		return entity.ShortLink{}, err
	}

	now := time.Now().UTC()
	createdBy := prevShortLink.CreatedBy
	createdAt := prevShortLink.CreatedAt
	return entity.ShortLink{
		Alias:     shortLinkInput.GetCustomAlias(""),
		LongLink:  shortLinkInput.GetLongLink(""),
		ExpireAt:  shortLinkInput.ExpireAt,
		CreatedBy: createdBy,
		CreatedAt: createdAt,
		UpdatedAt: &now,
	}, nil
}

// DeleteShortLink deletes an existing user short link.
func (s ShortLinkFake) DeleteShortLink(alias string) error {
	if alias == "" {
		return errors.New("empty alias")
	}
	_, ok := s.shortLinks[alias]
	if !ok {
		return errors.New("alias does not exist")
	}
	delete(s.shortLinks, alias)

	return nil
}

// NewShortLinkFake creates in memory ShortLink repository
func NewShortLinkFake(userShortLinkRepoFake *UserShortLinkFake, shortLinks map[string]entity.ShortLink) ShortLinkFake {
	return ShortLinkFake{
		shortLinks:            shortLinks,
		userShortLinkRepoFake: userShortLinkRepoFake,
	}
}
