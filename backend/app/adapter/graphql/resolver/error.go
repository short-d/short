package resolver

// ErrCode represents an unique string identifying a GraphQL api error.
type ErrCode string

const (
	ErrCodeUnknown            ErrCode = "unknown"
	ErrCodeAliasAlreadyExist          = "aliasAlreadyExist"
	ErrCodeRequesterNotHuman          = "requesterNotHuman"
	ErrCodeInvalidLongLink            = "invalidLongLink"
	ErrCodeInvalidCustomAlias         = "invalidCustomAlias"
	ErrCodeInvalidAuthToken           = "invalidAuthToken"
)

// GraphQlError represents a GraphAPI error.
type GraphQlError interface {
	Extensions() map[string]interface{}
	Error() string
}

// ErrUnknown represents an unclassified error. ErrUnknown maybe returned in
// order to prevent hackers from guessing security vulnerabilities.
type ErrUnknown struct{}

var _ GraphQlError = (*ErrUnknown)(nil)

// Extensions keeps structured error metadata so that the clients can reliably
// handle the error.
func (e ErrUnknown) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeUnknown,
	}
}

// Error retrieves the human readable error message.
func (e ErrUnknown) Error() string {
	return "unknown err"
}

type ErrURLAliasExist string

var _ GraphQlError = (*ErrURLAliasExist)(nil)

// Extensions keeps structured error metadata so that the clients can reliably
// handle the error.
func (e ErrURLAliasExist) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":  ErrCodeAliasAlreadyExist,
		"alias": string(e),
	}
}

// Error retrieves the human readable error message.
func (e ErrURLAliasExist) Error() string {
	return "url alias already exists"
}

// ErrNotHuman signifies that the API consumer is not human.
type ErrNotHuman struct{}

var _ GraphQlError = (*ErrNotHuman)(nil)

// Extensions keeps structured error metadata so that the clients can reliably
// handle the error.
func (e ErrNotHuman) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeRequesterNotHuman,
	}
}

// Error retrieves the human readable error message.
func (e ErrNotHuman) Error() string {
	return "requester is not human"
}

// ErrInvalidLongLink signifies that the provided long link has incorrect format.
type ErrInvalidLongLink string

var _ GraphQlError = (*ErrInvalidLongLink)(nil)

// Extensions keeps structured error metadata so that the clients can gracefully
// handle the error.
func (e ErrInvalidLongLink) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":     ErrCodeInvalidLongLink,
		"longLink": string(e),
	}
}

// Error retrieves the human readable error message.
func (e ErrInvalidLongLink) Error() string {
	return "long link is invalid"
}

// ErrInvalidCustomAlias signifies that the provided custom alias has incorrect
// format.
type ErrInvalidCustomAlias string

var _ GraphQlError = (*ErrInvalidCustomAlias)(nil)

// Extensions keeps structured error metadata so that the clients can gracefully
// handle the error.
func (e ErrInvalidCustomAlias) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":        ErrCodeInvalidCustomAlias,
		"customAlias": string(e),
	}
}

// Error retrieves the human readable error message.
func (e ErrInvalidCustomAlias) Error() string {
	return "custom alias is invalid"
}

// ErrInvalidAuthToken signifies that the provided authentication is invalid.
type ErrInvalidAuthToken string

var _ GraphQlError = (*ErrInvalidAuthToken)(nil)

// Extensions keeps structured error metadata so that the clients can gracefully
// handle the error.
func (e ErrInvalidAuthToken) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":      ErrCodeInvalidAuthToken,
		"authToken": string(e),
	}
}

// Error retrieves the human readable error message.
func (e ErrInvalidAuthToken) Error() string {
	return "auth token is invalid"
}
