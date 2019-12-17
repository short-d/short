package service

var _ ReCaptcha = (*Fake)(nil)

type Fake struct {
}

func (v Fake) Verify(recaptchaResponse string) (VerifyResponse, error) {
	panic("implement me")
}

func NewFake() Fake {
	return Fake{}
}
