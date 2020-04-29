package risk

var _ BlackList = (*BlackListFake)(nil)

type BlackListFake struct {
	blacklist map[string]bool
}

func (b BlackListFake) HasURL(url string) (bool, error) {
	_, found := b.blacklist[url]
	return found, nil
}

func NewBlackListFake(blacklist map[string]bool) BlackListFake {
	return BlackListFake{
		blacklist,
	}
}
