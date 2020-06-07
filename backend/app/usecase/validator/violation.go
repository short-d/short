package validator

// Violation represents a type of invalid error encountered by the validator.
type Violation string

const (
	Valid                Violation = "Valid"
	EmptyLongLink                  = "EmptyLongLink"
	LongLinkNotURL                 = "LongLinkNotURL"
	AliasTooLong                   = "AliasTooLong"
	LongLinkTooLong                = "LongLinkTooLong"
	HasFragmentCharacter           = "HasFragmentCharacter"
)
