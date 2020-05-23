package feature

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ DecisionMaker = (*DynamicDecisionMaker)(nil)

// DynamicDecisionMaker determines whether a feature should be turned on or off
// under certain conditions.
type DynamicDecisionMaker struct {
	instrumentation    instrumentation.Instrumentation
	featureToggleRepo  repository.FeatureToggle
	permissionCheckers map[PermissionToggle]PermissionChecker
}

// IsFeatureEnable determines whether a feature is enabled given featureID.
func (d DynamicDecisionMaker) IsFeatureEnable(featureID string, user *entity.User) bool {
	toggle, err := d.featureToggleRepo.FindToggleByID(featureID)
	if err != nil {
		d.instrumentation.FeatureToggleRetrievalFailed(err)
		d.instrumentation.MadeFeatureDecision(featureID, false)
		return false
	}
	d.instrumentation.FeatureToggleRetrievalSucceed()

	decision := d.makeDecision(toggle, user)
	d.instrumentation.MadeFeatureDecision(featureID, decision)
	return decision
}

func (d DynamicDecisionMaker) makeDecision(toggle entity.Toggle, user *entity.User) bool {
	if toggle.Type == entity.ManualToggle {
		return d.makeManualDecision(toggle)
	}

	if toggle.Type == entity.PermissionToggle {
		return d.makePermissionDecision(toggle, user)
	}

	// deny access by default if the toggle type's value is unexpected
	return false
}

func (d DynamicDecisionMaker) makeManualDecision(toggle entity.Toggle) bool {
	return toggle.IsEnabled
}

func (d DynamicDecisionMaker) makePermissionDecision(toggle entity.Toggle, user *entity.User) bool {
	// deny access by default unless the user exists and has necessary permissions
	if user == nil {
		return false
	}
	if !toggle.IsEnabled {
		return false
	}

	permissionToggle := PermissionToggle(toggle.ID)
	checker, ok := d.permissionCheckers[permissionToggle]
	if !ok {
		return false
	}

	decision, err := checker(*user)
	if err != nil {
		return false
	}
	return decision
}

var _ DecisionMakerFactory = (*DynamicDecisionMakerFactory)(nil)

// DynamicDecisionMakerFactory creates feature decision maker.
type DynamicDecisionMakerFactory struct {
	featureToggleRepo repository.FeatureToggle
	authorizer        authorizer.Authorizer
}

// NewDecision creates feature decision maker with instrumentation.
func (d DynamicDecisionMakerFactory) NewDecision(
	instrumentation instrumentation.Instrumentation,
) DecisionMaker {
	permissionCheckers := map[PermissionToggle]PermissionChecker{
		AdminPanel: d.authorizer.CanViewAdminPanel,
	}
	return DynamicDecisionMaker{
		instrumentation:    instrumentation,
		featureToggleRepo:  d.featureToggleRepo,
		permissionCheckers: permissionCheckers,
	}
}

// NewDynamicDecisionMakerFactory creates DynamicDecisionMakerFactory.
func NewDynamicDecisionMakerFactory(
	featureToggleRepo repository.FeatureToggle,
	authorizer authorizer.Authorizer,
) DynamicDecisionMakerFactory {
	return DynamicDecisionMakerFactory{
		featureToggleRepo: featureToggleRepo,
		authorizer:        authorizer,
	}
}
