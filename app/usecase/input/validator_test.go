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

	validator := NewLongLinkValidator()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expIsValid, validator.IsValid(&testCase.longLink))
		})
	}
}

func TestCustomAlias_IsValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		alias      string
		expIsValid bool
	}{
		{
			name:       "empty string",
			alias:      "",
			expIsValid: true,
		},
		{
			name:       "alias too long",
			alias:      strings.Repeat("helloworld", 5),
			expIsValid: false,
		},
		{
			name:       "alias valid",
			alias:      "fb",
			expIsValid: true,
		},
	}

	validator := NewCustomAliasValidator()
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expIsValid, validator.IsValid(&testCase.alias))
		})
	}
}
