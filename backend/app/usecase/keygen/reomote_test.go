package keygen

import (
	"short/app/usecase/service"
	"testing"

	"github.com/byliuyang/app/mdtest"
	"github.com/byliuyang/kgs/app/entity"
)

func TestNewRemote(t *testing.T) {
	keyGenService := service.NewKeyGenFake([]entity.Key{})
	_, err := NewRemote(0, &keyGenService)
	mdtest.NotEqual(t, nil, err)
}

func TestRemote_NewKey(t *testing.T) {
	testCases := []struct {
		name              string
		availableKeys     []entity.Key
		bufferSize        int
		expectedGetKeyOps int
		expectedHasErrs   []bool
		expectedKeys      []entity.Key
	}{
		{
			name: "buffer size is 2",
			availableKeys: []entity.Key{
				entity.Key("0K"),
				entity.Key("0L"),
			},
			bufferSize:        2,
			expectedGetKeyOps: 2,
			expectedHasErrs: []bool{
				false,
				false,
			},
			expectedKeys: []entity.Key{
				entity.Key("0K"),
				entity.Key("0L"),
			},
		},
		{
			name:              "no key available at beginning",
			availableKeys:     []entity.Key{},
			bufferSize:        2,
			expectedGetKeyOps: 1,
			expectedHasErrs: []bool{
				true,
			},
			expectedKeys: []entity.Key{
				entity.Key(""),
			},
		},
		{
			name: "run out of key",
			availableKeys: []entity.Key{
				entity.Key("0K"),
				entity.Key("0L"),
			},
			bufferSize:        2,
			expectedGetKeyOps: 3,
			expectedHasErrs: []bool{
				false,
				false,
				true,
			},
			expectedKeys: []entity.Key{
				entity.Key("0K"),
				entity.Key("0L"),
				entity.Key(""),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			keyGenService := service.NewKeyGenFake(testCase.availableKeys)
			remote, err := NewRemote(testCase.bufferSize, &keyGenService)
			mdtest.Equal(t, nil, err)

			for idx := 0; idx < testCase.expectedGetKeyOps; idx++ {
				key, err := remote.NewKey()

				if testCase.expectedHasErrs[idx] {
					mdtest.NotEqual(t, nil, err)
					continue
				}
				mdtest.Equal(t, testCase.expectedKeys[idx], key)
			}
		})
	}
}
