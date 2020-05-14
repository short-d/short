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
	permissionCheckers map[string]PermissionChecker
}

func (s StaticDecisionMaker) IsFeatureEnable(featureID string, user *entity.User) bool {
	checker, ok := s.permissionCheckers[featureID]
	if ok {
		if user == nil {
			return s.decisions[featureID]
		}
		isEnabled, err := checker(*user)
		if err != nil {
			return s.decisions[featureID]
		}
		return isEnabled
	}
	isEnabled := s.decisions[featureID]
	s.instrumentation.MadeFeatureDecision(featureID, isEnabled)
	return isEnabled
}

func (s StaticDecisionMaker) makePermissionDecision(toggle entity.Toggle, user *entity.User) bool {
	checker, ok := s.permissionCheckers[toggle.ID]
	if !ok {
		return toggle.IsEnabled
	}
	if user == nil {
		return toggle.IsEnabled
	}
	isEnabled, err := checker(*user)
	if err != nil {
		return toggle.IsEnabled
	}
	return isEnabled
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
	permissionCheckers := map[string]PermissionChecker{
		"include-admin-panel": s.authorizer.CanViewAdminPanel,
	}
	return &StaticDecisionMaker{
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
		permissionCheckers: permissionCheckers,
	}
}

// NewStaticDecisionMakerFactory creates StaticDecisionMakerFactory.
func NewStaticDecisionMakerFactory(authorizer authorizer.Authorizer) StaticDecisionMakerFactory {
	return StaticDecisionMakerFactory{
		authorizer: authorizer,
	}
}
