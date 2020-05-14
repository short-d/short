package feature

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/ctx"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestDynamicDecisionMaker_IsFeatureEnable(t *testing.T) {
	t.Parallel()
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
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			featureRepo := repository.NewFeatureToggleFake(testCase.toggles)

			entryRepo := logger.NewEntryRepoFake()
			lg, err := logger.NewFake(logger.LogOff, &entryRepo)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(time.Now())
			mt := metrics.NewFake()
			ana := analytics.NewFake()
			ctxCh := make(chan ctx.ExecutionContext)
			go func() {
				ctxCh <- ctx.ExecutionContext{}
			}()

			ins := instrumentation.NewInstrumentation(lg, tm, mt, ana, ctxCh)
			factory := NewDynamicDecisionMakerFactory(featureRepo)
			decision := factory.NewDecision(ins)
			gotIsEnabled := decision.IsFeatureEnable(testCase.featureID)
			assert.Equal(t, testCase.expectedIsEnabled, gotIsEnabled)
		})
	}
}
