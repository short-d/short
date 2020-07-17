package validator

// Validator is a common interface for validators.
type Validator interface {
	IsValid(entry string) (bool, Violation)
}
