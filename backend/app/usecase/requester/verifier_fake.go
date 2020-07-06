package requester

// VerifierFake turns off human verification during local development.
type VerifierFake struct{}

// IsHuman checks whether the request is sent by a human user.
func (r VerifierFake) IsHuman(recaptchaResponse string) (bool, error) {
	return true, nil
}

// NewVerifierFake creates new fake request verifier.
func NewVerifierFake() VerifierFake {
	return VerifierFake{}
}
