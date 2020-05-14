package feature

import (
	"github.com/short-d/short/backend/app/usecase/instrumentation"
)

var _ DecisionMaker = (*StaticDecisionMaker)(nil)

// StaticDecisionMaker makes feature decisions based on hardcoded values.
type StaticDecisionMaker struct {
	instrumentation instrumentation.Instrumentation
	decisions       map[string]bool
}

// IsFeatureEnable determines whether a feature is enabled given featureID.
func (s StaticDecisionMaker) IsFeatureEnable(featureID string) bool {
	isEnabled := s.decisions[featureID]
	s.instrumentation.MadeFeatureDecision(featureID, isEnabled)
	return isEnabled
}

var _ DecisionMakerFactory = (*StaticDecisionMakerFactory)(nil)

// StaticDecisionMakerFactory creates static feature decision maker.
type StaticDecisionMakerFactory struct {
}

// NewDecision creates static feature decision maker with config map.
func (s StaticDecisionMakerFactory) NewDecision(
	instrumentation instrumentation.Instrumentation,
) DecisionMaker {
	return StaticDecisionMaker{
		instrumentation: instrumentation,
		decisions: map[string]bool{
			"change-log":               true,
			"facebook-sign-in":         true,
			"github-sign-in":           true,
			"google-sign-in":           true,
			"search-bar":               true,
			"user-short-links-section": true,
			"preference-toggles":       true,
		},
	}
}

// NewStaticDecisionMakerFactory creates StaticDecisionMakerFactory.
func NewStaticDecisionMakerFactory() StaticDecisionMakerFactory {
	return StaticDecisionMakerFactory{}
}
