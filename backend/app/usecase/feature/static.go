package feature

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
)

var _ DecisionMaker = (*StaticDecisionMaker)(nil)

// StaticDecisionMaker makes feature decisions based on hardcoded values.
type StaticDecisionMaker struct {
	instrumentation    instrumentation.Instrumentation
	decisions          map[string]bool
	permissionCheckers map[PermissionToggle]PermissionChecker
}

// IsFeatureEnable determines whether a feature is enabled given featureID.
func (s StaticDecisionMaker) IsFeatureEnable(featureID string, user *entity.User) bool {
	decision := s.makeDecision(featureID, user)
	s.instrumentation.MadeFeatureDecision(featureID, decision)
	return decision
}

func (s StaticDecisionMaker) makeDecision(featureID string, user *entity.User) bool {
	_, hasChecker := s.permissionCheckers[PermissionToggle(featureID)]
	if hasChecker {
		return s.makePermissionDecision(featureID, user)
	}

	isEnabled := s.decisions[featureID]
	return isEnabled
}

func (s StaticDecisionMaker) makePermissionDecision(featureID string, user *entity.User) bool {
	if user == nil {
		return false
	}

	isEnabled := s.decisions[featureID]
	if !isEnabled {
		return false
	}

	checker := s.permissionCheckers[PermissionToggle(featureID)]
	decision, err := checker(*user)
	if err != nil {
		return false
	}

	return decision
}

var _ DecisionMakerFactory = (*StaticDecisionMakerFactory)(nil)

// StaticDecisionMakerFactory creates static feature decision maker.
type StaticDecisionMakerFactory struct {
	authorizer authorizer.Authorizer
}

// NewDecision creates static feature decision maker with config map.
func (s StaticDecisionMakerFactory) NewDecision(
	instrumentation instrumentation.Instrumentation,
) DecisionMaker {
	permissionCheckers := map[PermissionToggle]PermissionChecker{
		AdminPanel: s.authorizer.CanViewAdminPanel,
	}
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
			"admin-panel":              true,
		},
		permissionCheckers: permissionCheckers,
	}
}

// NewStaticDecisionMakerFactory creates StaticDecisionMakerFactory.
func NewStaticDecisionMakerFactory(authorizer authorizer.Authorizer) StaticDecisionMakerFactory {
	return StaticDecisionMakerFactory{
		authorizer: authorizer,
	}
}
