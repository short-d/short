package requester

// Verifier verifies incoming network to prevent cyber attacks.
type Verifier interface {
	IsHuman(recaptchaResponse string) (bool, error)
}

// VerifierReCaptcha verifies incoming network using ReCaptcha to prevent cyber attacks.
type VerifierReCaptcha struct {
	service ReCaptcha
}

// IsHuman checks whether the request is sent by a human user.
func (r VerifierReCaptcha) IsHuman(recaptchaResponse string) (bool, error) {
	apiRes, err := r.service.Verify(recaptchaResponse)
	if err != nil {
		return false, err
	}
	return apiRes.Score > 0.7, nil
}

// NewVerifierReCaptcha creates new ReCaptcha-backed request verifier.
func NewVerifierReCaptcha(service ReCaptcha) VerifierReCaptcha {
	return VerifierReCaptcha{
		service: service,
	}
}
