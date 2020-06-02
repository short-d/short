package order

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
)

func TestCreatedTime_ArrangeShortLinks(t *testing.T) {
	t.Parallel()

	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name               string
		shortLinks         []entity.ShortLink
		expectedShortLinks []entity.ShortLink
	}{
		{
			name:               "empty short links",
			shortLinks:         []entity.ShortLink{},
			expectedShortLinks: []entity.ShortLink{},
		},
		{
			name: "ordered short links",
			shortLinks: []entity.ShortLink{
				{Alias: "google", CreatedAt: &before},
				{Alias: "facebook", CreatedAt: &now},
				{Alias: "short", CreatedAt: &after},
			},
			expectedShortLinks: []entity.ShortLink{
				{Alias: "google", CreatedAt: &before},
				{Alias: "facebook", CreatedAt: &now},
				{Alias: "short", CreatedAt: &after},
			},
		},
		{
			name: "unordered short links",
			shortLinks: []entity.ShortLink{
				{Alias: "short", CreatedAt: &after},
				{Alias: "google", CreatedAt: &before},
				{Alias: "facebook", CreatedAt: &now},
			},
			expectedShortLinks: []entity.ShortLink{
				{Alias: "google", CreatedAt: &before},
				{Alias: "facebook", CreatedAt: &now},
				{Alias: "short", CreatedAt: &after},
			},
		},
		{
			name: "tied short links",
			shortLinks: []entity.ShortLink{
				{Alias: "short", CreatedAt: &after},
				{Alias: "google", CreatedAt: &before},
				{Alias: "facebook", CreatedAt: &now},
				{Alias: "youtube", CreatedAt: &after},
			},
			expectedShortLinks: []entity.ShortLink{
				{Alias: "google", CreatedAt: &before},
				{Alias: "facebook", CreatedAt: &now},
				{Alias: "short", CreatedAt: &after},
				{Alias: "youtube", CreatedAt: &after},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			orderByTime := NewOrder(ByCreatedTimeASC)
			result := orderByTime.ArrangeShortLinks(testCase.shortLinks)

			assert.Equal(t, testCase.expectedShortLinks, result)
		})
	}
}

func TestCreatedTime_ArrangeShortUsers(t *testing.T) {
	t.Parallel()

	now := time.Now()
	before := now.Add(-5 * time.Second)
	after := now.Add(5 * time.Second)

	testCases := []struct {
		name          string
		users         []entity.User
		expectedUsers []entity.User
	}{
		{
			name:          "empty users",
			users:         []entity.User{},
			expectedUsers: []entity.User{},
		},
		{
			name: "ordered users",
			users: []entity.User{
				{ID: "alpha", CreatedAt: &before},
				{ID: "beta", CreatedAt: &now},
				{ID: "gamma", CreatedAt: &after},
			},
			expectedUsers: []entity.User{
				{ID: "alpha", CreatedAt: &before},
				{ID: "beta", CreatedAt: &now},
				{ID: "gamma", CreatedAt: &after},
			},
		},
		{
			name: "unordered users",
			users: []entity.User{
				{ID: "gamma", CreatedAt: &after},
				{ID: "alpha", CreatedAt: &before},
				{ID: "beta", CreatedAt: &now},
			},
			expectedUsers: []entity.User{
				{ID: "alpha", CreatedAt: &before},
				{ID: "beta", CreatedAt: &now},
				{ID: "gamma", CreatedAt: &after},
			},
		},
		{
			name: "tied short links",
			users: []entity.User{
				{ID: "gamma", CreatedAt: &after},
				{ID: "alpha", CreatedAt: &before},
				{ID: "beta", CreatedAt: &now},
				{ID: "delta", CreatedAt: &after},
			},
			expectedUsers: []entity.User{
				{ID: "alpha", CreatedAt: &before},
				{ID: "beta", CreatedAt: &now},
				{ID: "gamma", CreatedAt: &after},
				{ID: "delta", CreatedAt: &after},
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			orderByTime := NewOrder(ByCreatedTimeASC)
			result := orderByTime.ArrangeUsers(testCase.users)

			assert.Equal(t, testCase.expectedUsers, result)
		})
	}
}
