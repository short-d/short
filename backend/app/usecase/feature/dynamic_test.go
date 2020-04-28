package feature

import (
	"testing"
	"time"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/instrumentation"
	"github.com/short-d/short/app/usecase/repository"
)

func TestDynamicDecisionMaker_IsFeatureEnable(t *testing.T) {
	testCases := []struct {
		name              string
		toggles           map[string]entity.Toggle
		featureID         string
		expectedIsEnabled bool
	}{
		{
			name:              "toggle not found",
			toggles:           map[string]entity.Toggle{},
			featureID:         "example-feature",
			expectedIsEnabled: false,
		},
		{
			name: "toggle disabled",
			toggles: map[string]entity.Toggle{
				"example-feature": {
					ID:        "example-feature",
					IsEnabled: false,
				},
			},
			featureID:         "example-feature",
			expectedIsEnabled: false,
		},
		{
			name: "toggle enabled",
			toggles: map[string]entity.Toggle{
				"example-feature": {
					ID:        "example-feature",
					IsEnabled: true,
				},
			},
			featureID:         "example-feature",
			expectedIsEnabled: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			featureRepo := repository.NewFeatureToggleFake(testCase.toggles)

			logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
			tracer := mdtest.NewTracerFake()
			timer := mdtest.NewTimerFake(time.Now())
			metrics := mdtest.NewMetricsFake()
			analytics := mdtest.NewAnalyticsFake()
			ctxCh := make(chan fw.ExecutionContext)
			go func() {
				ctxCh <- fw.ExecutionContext{}
			}()

			ins := instrumentation.NewInstrumentation(&logger, &tracer, timer, metrics, analytics, ctxCh)
			factory := NewDynamicDecisionMakerFactory(featureRepo)
			decision := factory.NewDecision(ins)
			gotIsEnabled := decision.IsFeatureEnable(testCase.featureID)
			mdtest.Equal(t, testCase.expectedIsEnabled, gotIsEnabled)
		})
	}
}
