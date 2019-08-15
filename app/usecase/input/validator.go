package input

type Validator interface {
	IsValid(text *string) bool
}
