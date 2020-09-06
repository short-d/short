package facebook

import (
	"github.com/short-d/short/backend/app/usecase/instrumentation"
)

func NewInstrumentationFactory() instrumentation.SSOInstrumentation {
	return instrumentation.NewSSO("facebook")
}
