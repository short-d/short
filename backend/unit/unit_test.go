package unit

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
)

type fakeDuration int64

func TestParseDuration(t *testing.T) {
	t.Parallel()

	t.Run("second", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
			hasFakeType      bool
			fakeDuration     fakeDuration
		}{
			{
				name:             "correct format",
				time:             "3s",
				expectedHasError: false,
				expectedDuration: 3 * time.Second,
			},
			{
				name:             "multi digits",
				time:             "120s",
				expectedHasError: false,
				expectedDuration: 120 * time.Second,
			},
			{
				name:             "empty string",
				time:             "",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:             "incorrect format",
				time:             "3ss",
				expectedHasError: true,
				expectedDuration: 0,
			},
			{
				name:         "custom Duration type",
				time:         "6s",
				hasFakeType:  true,
				fakeDuration: fakeDuration(6 * time.Second),
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					mdtest.NotEqual(t, nil, err)
					return
				} else if testCase.hasFakeType {
					mdtest.NotEqual(t, testCase.fakeDuration, duration)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("minute", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "5m",
				expectedHasError: false,
				expectedDuration: 5 * time.Minute,
			},
			{
				name:             "incorrect format",
				time:             "m5",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("hour", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "6h",
				expectedHasError: false,
				expectedDuration: 6 * time.Hour,
			},
			{
				name:             "incorrect format",
				time:             "6sh",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("day", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "2d",
				expectedHasError: false,
				expectedDuration: 2 * oneDay,
			},
			{
				name:             "incorrect format",
				time:             "d",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})

	t.Run("week", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "correct format",
				time:             "1w",
				expectedHasError: false,
				expectedDuration: oneWeak,
			},
			{
				name:             "incorrect format",
				time:             "w2w",
				expectedHasError: true,
				expectedDuration: 0,
			},
		}

		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				duration, err := ParseDuration(testCase.time)
				if testCase.expectedHasError {
					mdtest.NotEqual(t, nil, err)
					return
				}
				mdtest.Equal(t, nil, err)
				mdtest.Equal(t, testCase.expectedDuration, duration)
			})
		}
	})
}
