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

func (f FakeCreator) CreateURL(url entity.URL, userEmail string) (entity.URL, error) {
	key, err := f.keyGen.NewKey()
	if err != nil {
		return entity.URL{}, err
	}
	randomAlias := string(key)
	return f.CreateURLWithCustomAlias(url, randomAlias, userEmail)
}

func (f FakeCreator) CreateURLWithCustomAlias(
	url entity.URL,
	alias string,
	userEmail string,
) (entity.URL, error) {
	url.Alias = alias

	_, ok := f.urls[alias]
	if ok {
		return entity.URL{}, ErrAliasExist("usecase: url alias already exist")
	}

	f.urls[alias] = url
	return url, nil
}

func NewCreatorFake(
	urls map[string]entity.URL,
	availableAlias []string,
) FakeCreator {
	return FakeCreator{
		urls:   urls,
		keyGen: keygen.NewFake(availableAlias),
	}
}
