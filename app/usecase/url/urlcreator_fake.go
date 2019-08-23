package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
)

var _ Creator = (*FakeCreator)(nil)

type FakeCreator struct {
	urls   map[string]entity.URL
	keyGen keygen.Fake
}

func (f FakeCreator) Create(url entity.URL, userEmail string) (entity.URL, error) {
	randomAlias := f.keyGen.NewKey()
	return f.CreateWithCustomAlias(url, randomAlias, userEmail)
}

func (f FakeCreator) CreateWithCustomAlias(url entity.URL, alias string, userEmail string) (entity.URL, error) {
	url.Alias = alias

	_, ok := f.urls[alias]
	if ok {
		return entity.URL{}, ErrAliasExist("usecase: url alias already exist")
	}

	f.urls[alias] = url
	return url, nil
}

func NewCreatorFake(urls map[string]entity.URL, availableAlias []string) FakeCreator {
	return FakeCreator{
		urls:   urls,
		keyGen: keygen.NewFake(availableAlias),
	}
}
