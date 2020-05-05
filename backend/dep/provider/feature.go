package provider

import (
	"github.com/short-d/app/fw/env"
	"github.com/short-d/short/app/usecase/feature"
	"github.com/short-d/short/app/usecase/repository"
)

// NewFeatureDecisionMakerFactorySwitch creates FeatureDecisionFactory based on
// server environment.
func NewFeatureDecisionMakerFactorySwitch(
	deployment env.Deployment,
	toggleRepo repository.FeatureToggle,
) feature.DecisionMakerFactory {
	if deployment.IsDevelopment() {
		return feature.NewStaticDecisionMakerFactory()
	}
	return feature.NewDynamicDecisionMakerFactory(toggleRepo)
}
