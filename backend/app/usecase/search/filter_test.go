package search

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/usecase/search/order"
)

func TestCreatedTime_ArrangeShortLinks(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		maxResults     int
		resources      []Resource
		orders         []order.By
		expectedHasErr bool
		expectedFilter Filter
	}{
		{
			name:           "more resources than orders",
			maxResults:     2,
			resources:      []Resource{ShortLink},
			orders:         nil,
			expectedHasErr: true,
			expectedFilter: Filter{},
		},
		{
			name:           "more orders than resources",
			maxResults:     2,
			resources:      nil,
			orders:         []order.By{order.ByCreatedTimeASC},
			expectedHasErr: true,
			expectedFilter: Filter{},
		},
		{
			name:           "valid filter",
			maxResults:     2,
			resources:      []Resource{ShortLink},
			orders:         []order.By{order.ByCreatedTimeASC},
			expectedHasErr: false,
			expectedFilter: Filter{
				maxResults: 2,
				resources:  []Resource{ShortLink},
				orders:     []order.By{order.ByCreatedTimeASC},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			filter, err := NewFilter(testCase.maxResults, testCase.resources, testCase.orders)
			if testCase.expectedHasErr {
				assert.NotEqual(t, nil, err)
				return
			}

			assert.Equal(t, testCase.expectedFilter, filter)
		})
	}
}
