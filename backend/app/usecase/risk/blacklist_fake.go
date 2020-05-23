package risk

var _ BlackList = (*BlackListFake)(nil)

// BlackListFake is a in memory implementation of a BlackList used for testing.
type BlackListFake struct {
	blacklist map[string]bool
}

// HasShortLink checks whether a given short link exists in the blacklist.
func (b BlackListFake) HasShortLink(shortLink string) (bool, error) {
	_, found := b.blacklist[shortLink]
	return found, nil
}

// NewBlackListFake initializes an in-memory blacklist.
func NewBlackListFake(blacklist map[string]bool) BlackListFake {
	return BlackListFake{
		blacklist,
	}
}
