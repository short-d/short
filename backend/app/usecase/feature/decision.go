package feature

import (
	"github.com/short-d/short/app/usecase/instrumentation"
	"github.com/short-d/short/app/usecase/repository"
)

// Decision determines whether a feature should be turned on or off under
// certain situations.
type Decision struct {
	instrumentation   instrumentation.Instrumentation
	featureToggleRepo repository.FeatureToggle
}

// IsFeatureEnable determines whether a feature is enabled given featureID.
func (f Decision) IsFeatureEnable(featureID string) bool {
	toggle, err := f.featureToggleRepo.FindToggleByID(featureID)
	if err != nil {
		f.instrumentation.FeatureToggleRetrievalFailed(err)
		f.instrumentation.MadeFeatureDecision(featureID, false)
		return false
	}
	f.instrumentation.FeatureToggleRetrievalSucceed()
	f.instrumentation.MadeFeatureDecision(featureID, toggle.IsEnabled)
	return toggle.IsEnabled
}

// DecisionFactory creates feature decision maker.
type DecisionFactory struct {
	featureToggleRepo repository.FeatureToggle
}

// NewDecision creates feature decision maker with instrumentation.
func (f DecisionFactory) NewDecision(instrumentation instrumentation.Instrumentation) Decision {
	return Decision{
		instrumentation:   instrumentation,
		featureToggleRepo: f.featureToggleRepo,
	}
}

// NewDecisionFactory creates DecisionFactory.
func NewDecisionFactory(featureToggleRepo repository.FeatureToggle) DecisionFactory {
	return DecisionFactory{
		featureToggleRepo: featureToggleRepo,
	}
}
