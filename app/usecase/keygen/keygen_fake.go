package keygen

var _ KeyGenerator = (*Fake)(nil)

type Fake struct {
	keys       []string
	currKeyIdx int
}

func (k *Fake) NewKey() string {
	if k.currKeyIdx >= len(k.keys) {
		return ""
	}

	key := k.keys[k.currKeyIdx]
	k.currKeyIdx++
	return key
}

func NewFake(keys []string) Fake {
	return Fake{
		keys: keys,
	}
}
