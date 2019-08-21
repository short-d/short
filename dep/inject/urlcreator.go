package inject

import (
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"short/app/usecase/url"
)

func UrlCreatorPersist(urlRepo repo.Url, keyGen keygen.KeyGenerator) url.Creator {
	return url.NewCreatorPersist(urlRepo, keyGen)
}
