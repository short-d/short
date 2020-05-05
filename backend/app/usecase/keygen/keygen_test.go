// +build !integration all

package keygen

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/usecase/external"
)

func TestNewRemote(t *testing.T) {
	t.Parallel()

	keyFetcher := external.NewKeyFetcherFake([]external.Key{})
	_, err := NewKeyGenerator(0, &keyFetcher)
	assert.NotEqual(t, nil, err)
}

func TestRemote_NewKey(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		availableKeys     []external.Key
		bufferSize        int
		expectedGetKeyOps int
		expectedHasErrs   []bool
		expectedKeys      []external.Key
	}{
		{
			name: "buffer size is 2",
			availableKeys: []external.Key{
				external.Key("0K"),
				external.Key("0L"),
			},
			bufferSize:        2,
			expectedGetKeyOps: 2,
			expectedHasErrs: []bool{
				false,
				false,
			},
			expectedKeys: []external.Key{
				external.Key("0K"),
				external.Key("0L"),
			},
		},
		{
			name:              "no key available at beginning",
			availableKeys:     []external.Key{},
			bufferSize:        2,
			expectedGetKeyOps: 1,
			expectedHasErrs: []bool{
				true,
			},
			expectedKeys: []external.Key{
				external.Key(""),
			},
		},
		{
			name: "run out of key",
			availableKeys: []external.Key{
				external.Key("0K"),
				external.Key("0L"),
			},
			bufferSize:        2,
			expectedGetKeyOps: 3,
			expectedHasErrs: []bool{
				false,
				false,
				true,
			},
			expectedKeys: []external.Key{
				external.Key("0K"),
				external.Key("0L"),
				external.Key(""),
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			keyFetcher := external.NewKeyFetcherFake(testCase.availableKeys)
			remote, err := NewKeyGenerator(testCase.bufferSize, &keyFetcher)
			assert.Equal(t, nil, err)

			for idx := 0; idx < testCase.expectedGetKeyOps; idx++ {
				key, err := remote.NewKey()

				if testCase.expectedHasErrs[idx] {
					assert.NotEqual(t, nil, err)
					continue
				}
				assert.Equal(t, testCase.expectedKeys[idx], key)
			}
		})
	}
}
