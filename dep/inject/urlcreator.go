package inject

import (
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"short/app/usecase/url"
)

func URLCreatorPersist(urlRepo repo.URL, keyGen keygen.KeyGenerator) url.Creator {
	return url.NewCreatorPersist(urlRepo, keyGen)
}
