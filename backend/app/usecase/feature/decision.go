package feature

import (
	"github.com/short-d/short/app/usecase/instrumentation"
	"github.com/short-d/short/app/usecase/repository"
)

type Decision struct {
	instrumentation   instrumentation.Instrumentation
	featureToggleRepo repository.FeatureToggle
}

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

type DecisionFactory struct {
	featureToggleRepo repository.FeatureToggle
}

func (f DecisionFactory) NewDecision(instrumentation instrumentation.Instrumentation) Decision {
	return Decision{
		instrumentation:   instrumentation,
		featureToggleRepo: f.featureToggleRepo,
	}
}

func NewDecisionFactory(featureToggleRepo repository.FeatureToggle) DecisionFactory {
	return DecisionFactory{
		featureToggleRepo: featureToggleRepo,
	}
}
