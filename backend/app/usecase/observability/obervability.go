package observability

import "github.com/short-d/short/app/entity"

// Observability measures the internal operation of the system.
type Observability interface {
	RedirectingAliasToLongLink(user *entity.User)
	RedirectedAliasToLongLink(user *entity.User)
	LongLinkRetrievalSucceed()
	LongLinkRetrievalFailed(err error)
}
