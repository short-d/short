package observability

type Observability interface {
	LongLinkRetrievalFailed(err error)
	LongLinkRetrievalSucceed()
}
