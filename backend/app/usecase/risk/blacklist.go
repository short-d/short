package risk

// BlackList checks whether an item is acceptable
type BlackList interface {
	HasShortLink(shortLink string) (bool, error)
}
