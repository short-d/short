package risk

var _ BlackList = (*BlackListFake)(nil)

// BlackListFake implements a BlackList that can be used for testing.
type BlackListFake struct {
	blacklist map[string]bool
}

// HasURL checks if a given url is found within the blacklist.
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
