package risk

var _ BlackList = (*BlackListFake)(nil)

var blacklist = map[string]struct{}{
	"http://malware.wicar.org/data/ms14_064_ole_not_xp.html": struct {
	}{},
}

type BlackListFake struct {
	blacklist map[string]struct{}
}

func (b BlackListFake) HasURL(url string) (bool, error) {
	_, found := b.blacklist[url]
	if found {
		return true, nil
	}
	return false, nil
}

func NewBlackListFake() BlackListFake {
	return BlackListFake{
		blacklist,
	}
}
