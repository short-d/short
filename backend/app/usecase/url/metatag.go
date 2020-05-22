package url

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ MetaTag = (*MetaTagPersist)(nil)

type MetaTag interface {
	GetOGTags(alias string) (entity.OpenGraphTags, error)
	GetTwitterTags(alias string) (entity.TwitterTags, error)
	UpdateOGTags(alias string, title string, description string, imageURL string) (entity.OpenGraphTags, error)
	UpdateTwitterTags(alias string, title string, description string, imageURL string) (entity.TwitterTags, error)
}

type MetaTagPersist struct {
	urlRepo repository.URL
}

func (m MetaTagPersist) UpdateTwitterTags(alias string, title string, description string, imageURL string) (entity.TwitterTags, error) {
	url, err := m.urlRepo.UpdateTwitterMetaTags(alias, entity.TwitterTags{
		TwitterTitle:       &title,
		TwitterDescription: &description,
		TwitterImageURL:    &imageURL,
	})

	if err != nil {
		return entity.TwitterTags{}, err
	}

	return entity.TwitterTags{
		TwitterTitle:       url.TwitterTitle,
		TwitterDescription: url.TwitterDescription,
		TwitterImageURL:    url.TwitterImageURL,
	}, nil
}

func (m MetaTagPersist) UpdateOGTags(alias string, title string, description string, imageURL string) (entity.OpenGraphTags, error) {
	url, err := m.urlRepo.UpdateOGMetaTags(alias, entity.OpenGraphTags{
		OpenGraphTitle:       &title,
		OpenGraphDescription: &description,
		OpenGraphImageURL:    &imageURL,
	})

	if err != nil {
		return entity.OpenGraphTags{}, err
	}

	return entity.OpenGraphTags{
		OpenGraphTitle:       url.OpenGraphTitle,
		OpenGraphDescription: url.OpenGraphDescription,
		OpenGraphImageURL:    url.OpenGraphImageURL,
	}, nil
}

func (m MetaTagPersist) GetOGTags(alias string) (entity.OpenGraphTags, error) {
	url, err := m.urlRepo.GetByAlias(alias)
	if err != nil {
		return entity.OpenGraphTags{}, err
	}

	defaultTitle := "Short: Free link shortening service"
	defaultDesc := "Short enables people to type less for their favorite web sites"
	defaultImageURL := "https://short-d.com/promo/small-tile.png"

	if url.OpenGraphTitle == nil {
		url.OpenGraphTitle = &defaultTitle
	}

	if url.OpenGraphDescription == nil {
		url.OpenGraphDescription = &defaultDesc
	}

	if url.OpenGraphImageURL == nil {
		url.OpenGraphImageURL = &defaultImageURL
	}

	return entity.OpenGraphTags{
		OpenGraphTitle:       url.OpenGraphTitle,
		OpenGraphDescription: url.OpenGraphDescription,
		OpenGraphImageURL:    url.OpenGraphImageURL,
	}, nil
}

func (m MetaTagPersist) GetTwitterTags(alias string) (entity.TwitterTags, error) {
	url, err := m.urlRepo.GetByAlias(alias)
	if err != nil {
		return entity.TwitterTags{}, err
	}

	defaultTitle := "Short: Free link shortening service"
	defaultDesc := "Short enables people to type less for their favorite web sites"
	defaultImageURL := "https://short-d.com/promo/small-tile.png"

	if url.TwitterTitle == nil {
		url.TwitterTitle = &defaultTitle
	}

	if url.TwitterDescription == nil {
		url.TwitterDescription = &defaultDesc
	}

	if url.TwitterImageURL == nil {
		url.TwitterImageURL = &defaultImageURL
	}

	return entity.TwitterTags{
		TwitterTitle:       url.TwitterTitle,
		TwitterDescription: url.TwitterDescription,
		TwitterImageURL:    url.TwitterImageURL,
	}, nil
}

func NewMetaTagPersist(urlRepo repository.URL) MetaTagPersist {
	return MetaTagPersist{urlRepo: urlRepo}
}
