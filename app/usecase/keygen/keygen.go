package keygen

type KeyGenerator interface {
	NewKey() string
}
