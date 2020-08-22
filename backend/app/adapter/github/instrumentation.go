package github

import (
	"github.com/short-d/short/backend/app/usecase/instrumentation"
)

func NewInstrumentation() instrumentation.SSO {
	return instrumentation.NewSSO("github")
}
