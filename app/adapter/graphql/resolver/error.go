package resolver

type ErrCode string

const (
	ErrCodeUnknown            ErrCode = "unknown"
	ErrCodeAliasAlreadyExist          = "aliasAlreadyExist"
	ErrCodeRequesterNotHuman          = "requestNotHuman"
	ErrCodeInvalidLongLink            = "invalidLongLink"
	ErrCodeInvalidCustomAlias         = "invalidCustomAlias"
	ErrCodeInvalidAuthToken           = "invalidAuthToken"
)

type GraphQlError interface {
	Extensions() map[string]interface{}
	Error() string
}

type ErrUnknown struct{}

var _ GraphQlError = (*ErrUnknown)(nil)

func (e ErrUnknown) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeUnknown,
	}
}

func (e ErrUnknown) Error() string {
	return "unknown err"
}

type ErrURLAliasExist string

var _ GraphQlError = (*ErrURLAliasExist)(nil)

func (e ErrURLAliasExist) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":  ErrCodeAliasAlreadyExist,
		"alias": string(e),
	}
}

func (e ErrURLAliasExist) Error() string {
	return "url alias already exists"
}

type ErrNotHuman struct{}

var _ GraphQlError = (*ErrNotHuman)(nil)

func (e ErrNotHuman) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeRequesterNotHuman,
	}
}

func (e ErrNotHuman) Error() string {
	return "requester is not human"
}

type ErrInvalidLongLink string

var _ GraphQlError = (*ErrInvalidLongLink)(nil)

func (e ErrInvalidLongLink) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":     ErrCodeInvalidLongLink,
		"longLink": string(e),
	}
}

func (e ErrInvalidLongLink) Error() string {
	return "long link is invalid"
}

type ErrInvalidCustomAlias string

var _ GraphQlError = (*ErrInvalidCustomAlias)(nil)

func (e ErrInvalidCustomAlias) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":        ErrCodeInvalidCustomAlias,
		"customAlias": string(e),
	}
}

func (e ErrInvalidCustomAlias) Error() string {
	return "custom alias is invalid"
}

type ErrInvalidAuthToken string

var _ GraphQlError = (*ErrInvalidAuthToken)(nil)

func (e ErrInvalidAuthToken) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":      ErrCodeInvalidAuthToken,
		"authToken": string(e),
	}
}

func (e ErrInvalidAuthToken) Error() string {
	return "auth token is invalid"
}
