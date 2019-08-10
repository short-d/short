package usecase

type KeyGenFake struct {
	keys       []string
	currKeyIdx int
}

func (k *KeyGenFake) NewKey() string {
	if k.currKeyIdx >= len(k.keys) {
		return ""
	}

	key := k.keys[k.currKeyIdx]
	k.currKeyIdx++
	return key
}

func NewKeyGenFake(keys []string) KeyGenerator {
	return &KeyGenFake{
		keys: keys,
	}
}
