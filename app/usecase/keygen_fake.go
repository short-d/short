package usecase

type KeyGenFake struct {
}

func (KeyGenFake) NewKey() string {
	panic("implement me")
}

func NewKeyGenFake() KeyGenerator {
	return KeyGenFake{}
}
