package shortlink

import (
	"github.com/short-d/short/backend/app/entity/metatag"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ MetaTag = (*MetaTagPersist)(nil)

// MetaTag represents a ShortLink MetaTags retriever and updater
type MetaTag interface {
	GetOpenGraphTags(alias string) (metatag.OpenGraph, error)
	GetTwitterTags(alias string) (metatag.Twitter, error)
}

// MetaTagPersist represents a ShortLink MetaTags retriever and updater which persist the generated
// changes in the repository
type MetaTagPersist struct {
	shortLinkRepo repository.ShortLink
}

const defaultTitle string = "Short: Free link shortening service"
const defaultDesc string = "Short enables people to type less for their favorite web sites"
const defaultImageURL string = "https://short-d.com/promo/small-tile.png"

// GetOpenGraphTags retrieves ShortLink Open Graph Tags from persistent storage given alias
func (m MetaTagPersist) GetOpenGraphTags(alias string) (metatag.OpenGraph, error) {
	shortLink, err := m.shortLinkRepo.GetShortLinkByAlias(alias)
	if err != nil {
		return metatag.OpenGraph{}, err
	}

	defaultTitle := defaultTitle
	defaultDesc := defaultDesc
	defaultImageURL := defaultImageURL

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

// GetTwitterTags retrieves ShortLink Twitter Tags from persistent storage given alias
func (m MetaTagPersist) GetTwitterTags(alias string) (metatag.Twitter, error) {
	shortLink, err := m.shortLinkRepo.GetShortLinkByAlias(alias)
	if err != nil {
		return metatag.Twitter{}, err
	}

	defaultTitle := defaultTitle
	defaultDesc := defaultDesc
	defaultImageURL := defaultImageURL

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

// NewMetaTagPersist creates persistent ShortLink meta tags retriever and updater
func NewMetaTagPersist(shortLinkRepo repository.ShortLink) MetaTagPersist {
	return MetaTagPersist{shortLinkRepo: shortLinkRepo}
}
