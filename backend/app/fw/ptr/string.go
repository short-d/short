package ptr

// String returns the address of a string literal.
// This is needed because Golang does not allow taking the address of a string literal directly.
// Please see https://golang.org/ref/spec#Address_operators for more information.
func String(str string) *string {
	return &str
}
