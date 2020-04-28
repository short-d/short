package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/feature"
	"github.com/short-d/short/app/usecase/repository"
	"github.com/short-d/short/env"
)

// NewFeatureDecisionMakerFactorySwitch creates FeatureDecisionFactory based on
// server environment.
func NewFeatureDecisionMakerFactorySwitch(
	serverEnv fw.ServerEnv,
	toggleRepo repository.FeatureToggle,
) feature.DecisionMakerFactory {
	if serverEnv == env.Development {
		return feature.NewStaticDecisionMakerFactory()
	}
	return feature.NewDynamicDecisionMakerFactory(toggleRepo)
}
