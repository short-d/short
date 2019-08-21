package inject

import (
	"short/app/usecase/repo"
	"short/app/usecase/url"
)

func URLRetrieverPersist(urlRepo repo.URL) url.Retriever {
	return url.NewRetrieverPersist(urlRepo)
}
