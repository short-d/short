package inject

import (
	"short/app/usecase/repo"
	"short/app/usecase/url"
)

func UrlRetrieverPersist(urlRepo repo.Url) url.Retriever {
	return url.NewRetrieverPersist(urlRepo)
}
