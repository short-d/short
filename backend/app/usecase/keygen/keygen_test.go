// +build !integration all

package keygen

import (
	"testing"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/usecase/service"
)

func TestNewRemote(t *testing.T) {
	t.Parallel()

	keyFetcher := service.NewKeyFetcherFake([]service.Key{})
	_, err := NewKeyGenerator(0, &keyFetcher)
	mdtest.NotEqual(t, nil, err)
}

func TestRemote_NewKey(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		availableKeys     []service.Key
		bufferSize        int
		expectedGetKeyOps int
		expectedHasErrs   []bool
		expectedKeys      []service.Key
	}{
		{
			name: "buffer size is 2",
			availableKeys: []service.Key{
				service.Key("0K"),
				service.Key("0L"),
			},
			bufferSize:        2,
			expectedGetKeyOps: 2,
			expectedHasErrs: []bool{
				false,
				false,
			},
			expectedKeys: []service.Key{
				service.Key("0K"),
				service.Key("0L"),
			},
		},
		{
			name:              "no key available at beginning",
			availableKeys:     []service.Key{},
			bufferSize:        2,
			expectedGetKeyOps: 1,
			expectedHasErrs: []bool{
				true,
			},
			expectedKeys: []service.Key{
				service.Key(""),
			},
		},
		{
			name: "run out of key",
			availableKeys: []service.Key{
				service.Key("0K"),
				service.Key("0L"),
			},
			bufferSize:        2,
			expectedGetKeyOps: 3,
			expectedHasErrs: []bool{
				false,
				false,
				true,
			},
			expectedKeys: []service.Key{
				service.Key("0K"),
				service.Key("0L"),
				service.Key(""),
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			keyFetcher := service.NewKeyFetcherFake(testCase.availableKeys)
			remote, err := NewKeyGenerator(testCase.bufferSize, &keyFetcher)
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
