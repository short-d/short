package requester

// Verifier verifies incoming network to prevent cyber attacks.
type Verifier interface {
	IsHuman(recaptchaResponse string) (bool, error)
}
