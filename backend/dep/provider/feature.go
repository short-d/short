package provider

import (
	"github.com/short-d/app/fw/env"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/feature"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// NewFeatureDecisionMakerFactorySwitch creates FeatureDecisionFactory based on
// server environment.
func NewFeatureDecisionMakerFactorySwitch(
	deployment env.Deployment,
	toggleRepo repository.FeatureToggle,
	authorizer authorizer.Authorizer,
) feature.DecisionMakerFactory {
	if deployment.IsDevelopment() {
		return feature.NewStaticDecisionMakerFactory(authorizer)
	}
	return feature.NewDynamicDecisionMakerFactory(toggleRepo, authorizer)
}
