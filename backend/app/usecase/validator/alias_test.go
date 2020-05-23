// +build !integration all

package validator

import (
	"strings"
	"testing"

	"github.com/short-d/app/fw/assert"
)

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

	validator := NewCustomAlias()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expIsValid, validator.IsValid(testCase.alias))
		})
	}
}
