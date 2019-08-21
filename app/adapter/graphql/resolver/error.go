package resolver

type ErrCode string

const (
	ErrCodeUnknown            ErrCode = "unknown"
	ErrCodeAliasAlreadyExist          = "aliasAlreadyExist"
	ErrCodeRequesterNotHuman          = "requestNotHuman"
	ErrCodeInvalidLongLink            = "invalidLongLink"
	ErrCodeInvalidCustomAlias         = "invalidCustomAlias"
)

type ErrUnknown struct{}

func (e ErrUnknown) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeUnknown,
	}
}

func (e ErrUnknown) Error() string {
	return "unknown err"
}

type ErrURLAliasExist string

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

func (e ErrNotHuman) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code": ErrCodeRequesterNotHuman,
	}
}

func (e ErrNotHuman) Error() string {
	return "requester is not human"
}

type ErrInvalidLongLink string

func (e ErrInvalidLongLink) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"code":     ErrCodeInvalidLongLink,
		"longLink": string(e),
	}
}

func (e ErrInvalidLongLink) Error() string {
	return "long link is invalid"
}

type ErrInvalidCustomAlias struct {
	customAlias *string
}

func (e ErrInvalidCustomAlias) Extensions() map[string]interface{} {
	var alias string
	if e.customAlias != nil {
		alias = *e.customAlias
	}
	return map[string]interface{}{
		"code":        ErrCodeInvalidCustomAlias,
		"customAlias": alias,
	}
}

func (e ErrInvalidCustomAlias) Error() string {
	return "custom alias is invalid"
}
