package url

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ MetaTag = (*MetaTagPersist)(nil)

type MetaTag interface {
	FetchOGTags(alias string) (entity.MetaOGTags, error)
	FetchTwitterTags(alias string) (entity.MetaTwitterTags, error)
	UpdateOGTags(alias string, ogTitle string, ogDescription string, ogImageURL string) (entity.MetaOGTags, error)
	UpdateTwitterTags(alias string, twitterTitle string, twitterDescription string, twitterImageURL string) (entity.MetaTwitterTags, error)
}

type MetaTagPersist struct {
	urlRepo repository.URL
}

func (m MetaTagPersist) UpdateTwitterTags(alias string, twitterTitle string, twitterDescription string, twitterImageURL string) (entity.MetaTwitterTags, error) {
	url, err := m.urlRepo.UpdateTwitterMetaTags(alias, entity.MetaTwitterTags{
		TwitterTitle:       twitterTitle,
		TwitterDescription: twitterDescription,
		TwitterImageURL:    twitterImageURL,
	})

	if err != nil {
		return entity.MetaTwitterTags{}, err
	}

	return entity.MetaTwitterTags{
		TwitterTitle:       url.TwitterTitle,
		TwitterDescription: url.TwitterDescription,
		TwitterImageURL:    url.TwitterImageURL,
	}, nil
}

func (m MetaTagPersist) UpdateOGTags(alias string, ogTitle string, ogDescription string, ogImageURL string) (entity.MetaOGTags, error) {
	url, err := m.urlRepo.UpdateOGMetaTags(alias, entity.MetaOGTags{
		OGTitle:       ogTitle,
		OGDescription: ogDescription,
		OGImageURL:    ogImageURL,
	})

	if err != nil {
		return entity.MetaOGTags{}, err
	}

	return entity.MetaOGTags{
		OGTitle:       url.OGTitle,
		OGDescription: url.OGDescription,
		OGImageURL:    url.OGImageURL,
	}, nil
}

func (m MetaTagPersist) FetchOGTags(alias string) (entity.MetaOGTags, error) {
	url, err := m.urlRepo.GetByAlias(alias)
	if err != nil {
		return entity.MetaOGTags{}, err
	}

	return entity.MetaOGTags{
		OGTitle:       url.OGTitle,
		OGDescription: url.OGDescription,
		OGImageURL:    url.OGImageURL,
	}, nil
}

func (m MetaTagPersist) FetchTwitterTags(alias string) (entity.MetaTwitterTags, error) {
	url, err := m.urlRepo.GetByAlias(alias)
	if err != nil {
		return entity.MetaTwitterTags{}, err
	}

	return entity.MetaTwitterTags{
		TwitterTitle:       url.TwitterTitle,
		TwitterDescription: url.TwitterDescription,
		TwitterImageURL:    url.TwitterImageURL,
	}, nil
}

func NewMetaTagPersist(urlRepo repository.URL) MetaTagPersist {
	return MetaTagPersist{urlRepo: urlRepo}
}
