package recaptcha

import "short/app/usecase/service"

var _ service.ReCaptcha = (*Fake)(nil)

type Fake struct {
}

func (v Fake) Verify(recaptchaResponse string) (service.VerifyResponse, error) {
	panic("implement me")
}

func NewFake() Fake {
	return Fake{}
}
