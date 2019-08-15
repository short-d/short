package recaptcha

import "short/app/usecase/service"

type Fake struct {
}

func (v Fake) Verify(recaptchaResponse string) (service.VerifyResponse, error) {
	panic("implement me")
}

func NewFake() service.ReCaptcha {
	return Fake{}
}
