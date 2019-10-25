package provider

import (
	"short/app/usecase/keygen"
	"short/app/usecase/service"
)

type KeyGenBufferSize int

func NewRemote(
	bufferSize KeyGenBufferSize,
	keyGenService service.KeyGen,
) (keygen.Remote, error) {
	return keygen.NewRemote(int(bufferSize), keyGenService)
}
