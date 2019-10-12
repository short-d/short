package resolver

import (
	"github.com/byliuyang/app/mdtest"
	"testing"
)

func TestErrUnknown_Extensions(t *testing.T) {
	err := ErrUnknown{}

	expected := map[string]interface{}{
		"code": ErrCodeUnknown,
	}
	result := err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)
}

func TestErrUnknown_Error(t *testing.T) {
	err := ErrUnknown{}

	expected := "unknown err"
	result := err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)
}

func TestErrURLAliasExist_Extensions(t *testing.T) {
	var err ErrURLAliasExist

	expected := map[string]interface{}{
		"code":  ErrCodeAliasAlreadyExist,
		"alias": "",
	}
	result := err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)

	err = "testAlias"
	expected["alias"] = "testAlias"
	result = err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)
}

func TestErrURLAliasExist_Error(t *testing.T) {
	var err ErrURLAliasExist
	expected := "url alias already exists"

	result := err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)

	err = "testErr"
	result = err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)
}

func TestErrNotHuman_Extensions(t *testing.T) {
	err := ErrNotHuman{}

	expected := map[string]interface{}{
		"code": ErrCodeRequesterNotHuman,
	}
	result := err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)
}

func TestErrNotHuman_Error(t *testing.T) {
	err := ErrNotHuman{}

	expected := "requester is not human"
	result := err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)
}

func TestErrInvalidLongLink_Extensions(t *testing.T) {
	var err ErrInvalidLongLink

	expected := map[string]interface{}{
		"code":     ErrCodeInvalidLongLink,
		"longLink": "",
	}
	result := err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)

	err = "testLongLink"
	expected["longLink"] = "testLongLink"
	result = err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)
}

func TestErrInvalidLongLink_Error(t *testing.T) {
	var err ErrInvalidLongLink
	expected := "long link is invalid"

	result := err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)

	err = "testErr"
	result = err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)
}

func TestErrInvalidCustomAlias_Extensions(t *testing.T) {
	var err ErrInvalidCustomAlias

	expected := map[string]interface{}{
		"code":        ErrCodeInvalidCustomAlias,
		"customAlias": "",
	}
	result := err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)

	err = "testCustomAlias"
	expected["customAlias"] = "testCustomAlias"
	result = err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)
}

func TestErrInvalidCustomAlias_Error(t *testing.T) {
	var err ErrInvalidCustomAlias
	expected := "custom alias is invalid"

	result := err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)

	err = "testErr"
	result = err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)
}

func TestErrInvalidAuthToken_Extensions(t *testing.T) {
	var err ErrInvalidAuthToken

	expected := map[string]interface{}{
		"code":      ErrCodeInvalidAuthToken,
		"authToken": "",
	}
	result := err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)

	err = "testAuthToken"
	expected["authToken"] = "testAuthToken"
	result = err.Extensions()
	mdtest.Equal(t, result, expected, "err.Extensions() is '%v', but should be '%v'", result, expected)
}

func TestErrInvalidAuthToken_Error(t *testing.T) {
	var err ErrInvalidAuthToken
	expected := "auth token is invalid"

	result := err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)

	err = "testErr"
	result = err.Error()
	mdtest.Equal(t, expected, result, "err.Error() is '%v', but should be '%v'", result, expected)
}
