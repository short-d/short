package feature

import "github.com/short-d/short/app/usecase/instrumentation"

// Decision determines whether a feature should be turned on or off under
// certain conditions.
type DecisionMaker interface {
	IsFeatureEnable(featureID string) bool
}

// DecisionMakerFactory creates feature decision maker.
type DecisionMakerFactory interface {
	NewDecision(instrumentation instrumentation.Instrumentation) DecisionMaker
}
