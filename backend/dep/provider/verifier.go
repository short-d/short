package provider

import (
	"github.com/short-d/app/fw/env"
	"github.com/short-d/short/backend/app/usecase/requester"
)

// NewVerifier creates Verifier based on
// server environment.
func NewVerifier(
	deployment env.Deployment,
	service requester.ReCaptcha,
) requester.Verifier {
	if deployment.IsDevelopment() {
		return requester.NewVerifierFake()
	}
	return requester.NewVerifierReCaptcha(service)
}
