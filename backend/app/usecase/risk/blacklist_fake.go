package risk

var _ BlackList = (*BlackListFake)(nil)

// BlackListFake is a immemory implementation of a BlackList used for testing.
type BlackListFake struct {
	blacklist map[string]bool
}

// HasURL checks whether a given url exists in the blacklist.
func (b BlackListFake) HasURL(url string) (bool, error) {
	_, found := b.blacklist[url]
	return found, nil
}

// NewBlackListFake initializes a BlackListFake instance.
func NewBlackListFake(blacklist map[string]bool) BlackListFake {
	return BlackListFake{
		blacklist,
	}
}
