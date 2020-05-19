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
	permissionCheckers map[string]PermissionChecker
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

	if toggle.IsEnabled && toggle.Type == entity.PermissionToggle {
		decision := d.makePermissionDecision(toggle, user)

		d.instrumentation.MadeFeatureDecision(featureID, decision)
		return decision
	}

	d.instrumentation.MadeFeatureDecision(featureID, toggle.IsEnabled)
	return toggle.IsEnabled
}

func (d DynamicDecisionMaker) makePermissionDecision(toggle entity.Toggle, user *entity.User) bool {
	checker, ok := d.permissionCheckers[toggle.ID]
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
	permissionCheckers := map[string]PermissionChecker{
		IncludeAdminPanel: d.authorizer.CanViewAdminPanel,
	}
	return &DynamicDecisionMaker{
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
