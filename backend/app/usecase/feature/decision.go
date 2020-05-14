package feature

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
)

type PermissionChecker func(user entity.User) (bool, error)

// DecisionMaker determines whether a feature should be turned on or off under
// certain conditions.
type DecisionMaker interface {
	IsFeatureEnable(featureID string, user *entity.User) bool
}

// DecisionMakerFactory creates feature decision maker.
type DecisionMakerFactory interface {
	NewDecision(instrumentation instrumentation.Instrumentation) DecisionMaker
}
