package requester

var _ Verifier = (*ReCaptchaVerifier)(nil)

// ReCaptchaVerifier verifies incoming network using ReCaptcha to prevent spamming attacks.
type ReCaptchaVerifier struct {
	service ReCaptcha
}

// IsHuman checks whether the request is sent by a human user.
func (r ReCaptchaVerifier) IsHuman(recaptchaResponse string) (bool, error) {
	apiRes, err := r.service.Verify(recaptchaResponse)
	if err != nil {
		return false, err
	}
	return apiRes.Score > 0.7, nil
}

// NewReCaptchaVerifier creates new ReCaptcha-backed request verifier.
func NewReCaptchaVerifier(service ReCaptcha) ReCaptchaVerifier {
	return ReCaptchaVerifier{
		service: service,
	}
}
