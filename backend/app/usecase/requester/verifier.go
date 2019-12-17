package requester

import (
	"short/app/usecase/service"
)

// Verifier verifies in coming network to prevents cyber attacks.
type Verifier struct {
	service service.ReCaptcha
}

// IsHuman checks whether the request is sent by a human user.
func (r Verifier) IsHuman(recaptchaResponse string) (bool, error) {
	apiRes, err := r.service.Verify(recaptchaResponse)
	if err != nil {
		return false, err
	}
	return apiRes.Score > 0.7, nil
}

// NewVerifier creates new request verifier.
func NewVerifier(service service.ReCaptcha) Verifier {
	return Verifier{
		service: service,
	}
}
