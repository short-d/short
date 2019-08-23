package inject

import (
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
	"short/app/usecase/url"
)

func URLCreatorPersist(
	urlRepo repo.URL,
	userUrlRepo repo.UserURL,
	keyGen keygen.KeyGenerator,
) url.Creator {
	return url.NewCreatorPersist(urlRepo, userUrlRepo, keyGen)
}
