package feature

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/instrumentation"
	"github.com/short-d/short/app/usecase/repository"
)

var _ DecisionMaker = (*DynamicDecisionMaker)(nil)

// DynamicDecisionMaker determines whether a feature should be turned on or off
// under certain conditions.
type DynamicDecisionMaker struct {
	instrumentation   instrumentation.Instrumentation
	featureToggleRepo repository.FeatureToggle
}

// IsFeatureEnable determines whether a feature is enabled given featureID.
func (f DynamicDecisionMaker) IsFeatureEnable(featureID string) bool {
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

var _ DecisionMakerFactory = (*DynamicDecisionMakerFactory)(nil)

// DynamicDecisionMakerFactory creates feature decision maker.
type DynamicDecisionMakerFactory struct {
	serverEnv         fw.ServerEnv
	featureToggleRepo repository.FeatureToggle
}

// NewDecision creates feature decision maker with instrumentation.
func (f DynamicDecisionMakerFactory) NewDecision(
	instrumentation instrumentation.Instrumentation,
) DecisionMaker {
	return DynamicDecisionMaker{
		instrumentation:   instrumentation,
		featureToggleRepo: f.featureToggleRepo,
	}
}

// NewDynamicDecisionMakerFactory creates DynamicDecisionMakerFactory.
func NewDynamicDecisionMakerFactory(
	featureToggleRepo repository.FeatureToggle,
) DynamicDecisionMakerFactory {
	return DynamicDecisionMakerFactory{
		featureToggleRepo: featureToggleRepo,
	}
}
