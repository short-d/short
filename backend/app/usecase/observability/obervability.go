package observability

// Observability measures the internal operation of the system.
type Observability interface {
	LongLinkRetrievalFailed(err error)
	LongLinkRetrievalSucceed()
}
