package unit

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
)

func TestParseDuration(t *testing.T) {
	t.Parallel()

	t.Run("second", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "regular seconds",
				time:             "3s",
				expectedHasError: false,
				expectedDuration: 3 * time.Second,
			},
			{
				name:             "irregular seconds",
				time:             "3ss",
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

	t.Run("minute", func(t *testing.T) {
		testCases := []struct {
			name             string
			time             string
			expectedHasError bool
			expectedDuration time.Duration
		}{
			{
				name:             "regular minutes",
				time:             "5m",
				expectedHasError: false,
				expectedDuration: 5 * time.Minute,
			},
			{
				name:             "irregular minutes",
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
				name:             "regular hours",
				time:             "6h",
				expectedHasError: false,
				expectedDuration: 6 * time.Hour,
			},
			{
				name:             "irregular hours",
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
				name:             "regular days",
				time:             "2d",
				expectedHasError: false,
				expectedDuration: 2 * oneDay,
			},
			{
				name:             "irregular days",
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
				name:             "regular weeks",
				time:             "1w",
				expectedHasError: false,
				expectedDuration: oneWeak,
			},
			{
				name:             "irregular weeks",
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
