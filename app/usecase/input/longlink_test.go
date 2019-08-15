package input

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongLink_IsValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		longLink   string
		expIsValid bool
	}{
		{
			name:       "empty string",
			longLink:   "",
			expIsValid: false,
		},
		{
			name:       "no ://",
			longLink:   "randomLink",
			expIsValid: false,
		},
		{
			name:       "only ://",
			longLink:   "://",
			expIsValid: false,
		},
		{
			name:       "no hostname",
			longLink:   "http://",
			expIsValid: false,
		},
		{
			name:       "link too long",
			longLink:   strings.Repeat("helloworld", 20),
			expIsValid: false,
		},
		{
			name:       "link valid",
			longLink:   "https://google.com",
			expIsValid: true,
		},
	}

	validator := NewLongLink()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expIsValid, validator.IsValid(&testCase.longLink))
		})
	}
}
