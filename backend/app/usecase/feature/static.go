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
	isEnabled := s.decisions[featureID]

	_, hasPermissionCheck := s.permissionCheckers[featureID]
	if isEnabled && hasPermissionCheck {
		decision := s.makePermissionDecision(featureID, user)

		s.instrumentation.MadeFeatureDecision(featureID, decision)
		return decision
	}

	s.instrumentation.MadeFeatureDecision(featureID, isEnabled)
	return isEnabled
}


func (s StaticDecisionMaker) makePermissionDecision(featureID string, user *entity.User) bool {
	checker, ok := s.permissionCheckers[featureID]
	if !ok {
		return false
	}
	if user == nil {
		return false
	}

	isEnabled, err := checker(*user)
	if err != nil {
		return false
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
		IncludeAdminPanel: s.authorizer.CanViewAdminPanel,
	}
	return &StaticDecisionMaker{
		instrumentation: instrumentation,
		decisions: map[string]bool{
			ChangeLog:             true,
			FacebookSignIn:        true,
			GithubSignIn:          true,
			GoogleSignIn:          true,
			SearchBar:             true,
			UserShortLinksSection: true,
			PreferenceToggles:     true,
			IncludeAdminPanel:     true,
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
