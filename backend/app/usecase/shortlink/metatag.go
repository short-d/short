package shortlink

import (
	"github.com/short-d/short/backend/app/entity/metatag"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ MetaTag = (*MetaTagPersist)(nil)

// MetaTag fetches and updates MetaTags for a short link.
type MetaTag interface {
	GetOpenGraphTags(alias string) (metatag.OpenGraph, error)
	GetTwitterTags(alias string) (metatag.Twitter, error)
}

// MetaTagPersist fetches and updates MetaTags for a short link from persistent storage.
type MetaTagPersist struct {
	shortLinkRepo repository.ShortLink
}

const (
	defaultTitle    = "Short: Free link shortening service"
	defaultDesc     = "Short enables people to type less for their favorite web sites"
	defaultImageURL = "https://short-d.com/promo/small-tile.png"
)

// GetOpenGraphTags retrieves Open Graph tags for a short link from persistent storage given alias.
func (m MetaTagPersist) GetOpenGraphTags(alias string) (metatag.OpenGraph, error) {
	shortLink, err := m.shortLinkRepo.GetShortLinkByAlias(alias)
	if err != nil {
		return metatag.OpenGraph{}, err
	}

	if shortLink.OpenGraphTags.Title == nil {
		title := defaultTitle
		shortLink.OpenGraphTags.Title = &title
	}

	if shortLink.OpenGraphTags.Description == nil {
		description := defaultDesc
		shortLink.OpenGraphTags.Description = &description
	}

	if shortLink.OpenGraphTags.ImageURL == nil {
		imageURL := defaultImageURL
		shortLink.OpenGraphTags.ImageURL = &imageURL
	}

	return shortLink.OpenGraphTags, nil
}

// GetTwitterTags retrieves Twitter tags for a short link from persistent storage given alias.
func (m MetaTagPersist) GetTwitterTags(alias string) (metatag.Twitter, error) {
	shortLink, err := m.shortLinkRepo.GetShortLinkByAlias(alias)
	if err != nil {
		return metatag.Twitter{}, err
	}

	if shortLink.TwitterTags.Title == nil {
		title := defaultTitle
		shortLink.TwitterTags.Title = &title
	}

	if shortLink.TwitterTags.Description == nil {
		description := defaultDesc
		shortLink.TwitterTags.Description = &description
	}

	if shortLink.TwitterTags.ImageURL == nil {
		imageURL := defaultImageURL
		shortLink.TwitterTags.ImageURL = &imageURL
	}

	return shortLink.TwitterTags, nil
}

// NewMetaTagPersist creates NewMetaTagPersist given repository.
func NewMetaTagPersist(shortLinkRepo repository.ShortLink) MetaTagPersist {
	return MetaTagPersist{shortLinkRepo: shortLinkRepo}
}
