package risk

// BlackList checks whether an item is acceptable
type BlackList interface {
	HasURL(url string) (bool, error)
}
