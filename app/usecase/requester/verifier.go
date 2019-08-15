package requester

import (
	"short/app/usecase/service"
)

type Verifier struct {
	service service.Recaptcha
}

func (r Verifier) IsHuman(recaptchaResponse string) (bool, error) {
	apiRes, err := r.service.Verify(recaptchaResponse)
	if err != nil {
		return false, err
	}
	return apiRes.Score > 0.7, nil
}

func NewVerifier(service service.Recaptcha) Verifier {
	return Verifier{
		service: service,
	}
}
