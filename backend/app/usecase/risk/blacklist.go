package risk

// BlackList checks whether an item is acceptable
type BlackList interface {
	HasURL(shortLink string) (bool, error)
}
