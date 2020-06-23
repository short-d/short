package handle

import (
	"encoding/json"
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/feature"
)

// Feature retrieves the status of feature toggle.
func Feature(
	instrumentationFactory request.InstrumentationFactory,
	featureDecisionMakerFactory feature.DecisionMakerFactory,
	authenticator authenticator.Authenticator,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		i := instrumentationFactory.NewRequest()
		featureID := params["featureID"]
		user := getUser(r, authenticator)

		decision := featureDecisionMakerFactory.NewDecision(i)
		isEnable := decision.IsFeatureEnable(featureID, user)

		body, err := json.Marshal(isEnable)
		if err != nil {
			return
		}

		w.Write(body)
	}
}
