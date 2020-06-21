package shortlink

import (
	"github.com/short-d/short/backend/app/entity/metatag"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ MetaTag = (*MetaTagPersist)(nil)

type MetaTag interface {
	GetOGTags(alias string) (metatag.OpenGraph, error)
	GetTwitterTags(alias string) (metatag.Twitter, error)
}

type MetaTagPersist struct {
	shortLinkRepo repository.ShortLink
}

func (m MetaTagPersist) GetOGTags(alias string) (metatag.OpenGraph, error) {
	shortLink, err := m.shortLinkRepo.GetShortLinkByAlias(alias)
	if err != nil {
		return metatag.OpenGraph{}, err
	}

	defaultTitle := "Short: Free link shortening service"
	defaultDesc := "Short enables people to type less for their favorite web sites"
	defaultImageURL := "https://short-d.com/promo/small-tile.png"

	if shortLink.OpenGraphTags.Title == nil {
		shortLink.OpenGraphTags.Title = &defaultTitle
	}

	if shortLink.OpenGraphTags.Description == nil {
		shortLink.OpenGraphTags.Description = &defaultDesc
	}

	if shortLink.OpenGraphTags.ImageURL == nil {
		shortLink.OpenGraphTags.ImageURL = &defaultImageURL
	}

	return shortLink.OpenGraphTags, nil
}

func (m MetaTagPersist) GetTwitterTags(alias string) (metatag.Twitter, error) {
	shortLink, err := m.shortLinkRepo.GetShortLinkByAlias(alias)
	if err != nil {
		return metatag.Twitter{}, err
	}

	defaultTitle := "Short: Free link shortening service"
	defaultDesc := "Short enables people to type less for their favorite web sites"
	defaultImageURL := "https://short-d.com/promo/small-tile.png"

	if shortLink.TwitterTags.Title == nil {
		shortLink.TwitterTags.Title = &defaultTitle
	}

	if shortLink.TwitterTags.Description == nil {
		shortLink.TwitterTags.Description = &defaultDesc
	}

	if shortLink.TwitterTags.ImageURL == nil {
		shortLink.TwitterTags.ImageURL = &defaultImageURL
	}

	return shortLink.TwitterTags, nil
}

func NewMetaTagPersist(shortLinkRepo repository.ShortLink) MetaTagPersist {
	return MetaTagPersist{shortLinkRepo: shortLinkRepo}
}
