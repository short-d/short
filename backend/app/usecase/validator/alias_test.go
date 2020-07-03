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
		violation  Violation
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
			violation:  AliasTooLong,
		},
		{
			name:       "alias valid",
			alias:      "fb",
			expIsValid: true,
		},
		{
			name:       "alias contains hash tag",
			alias:      "#fb",
			expIsValid: false,
			violation:  HasFragmentCharacter,
		},
	}

	validator := NewCustomAlias()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			valid, violation := validator.IsValid(&testCase.alias)
			if !testCase.expIsValid {
				assert.Equal(t, testCase.violation, violation)
			}
			assert.Equal(t, testCase.expIsValid, valid)
		})
	}
}

func TestCustomAlias_hasFragmentCharacter(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name       string
		alias      string
		expIsValid bool
	}{
		{
			name:       "empty string",
			alias:      "",
			expIsValid: false,
		},
		{
			name:       "alias with fragment",
			alias:      "fb#home",
			expIsValid: true,
		},
		{
			name:       "alias without fragment",
			alias:      "fb",
			expIsValid: false,
		},
	}

	validator := NewCustomAlias()
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, testCase.expIsValid, validator.hasFragmentCharacter(testCase.alias))
		})
	}
}
