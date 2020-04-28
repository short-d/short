package risk

// URLBlackList is the interface that wraps the IsBlacklisted method.
type URLBlackList interface {
	IsBlacklisted(url string) (bool, error)
}
