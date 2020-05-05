package requester

import "github.com/short-d/short/app/usecase/external"

// Verifier verifies in coming network to prevents cyber attacks.
type Verifier struct {
	service external.ReCaptcha
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
func NewVerifier(service external.ReCaptcha) Verifier {
	return Verifier{
		service: service,
	}
}
