package validator

// Violation represents a type of invalid error encountered by the validator.
type Violation int

const (
	Valid Violation = iota
	InvalidAlias
	InvalidLongLink
	AliasTooLong
	LongLinkTooLong
	HasFragmentCharacter
)

// Returns a string representation of the violation.
func (v Violation) String() string {
	return [...]string{
		"Valid",
		"InvalidAlias",
		"InvalidLongLink",
		"AliasTooLong",
		"LongLinkTooLong",
		"HasFragmentCharacter",
	}[v]
}
