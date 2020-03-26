// +build !integration all

package changelog

import (
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/repository"
	"github.com/short-d/short/app/usecase/service"
)

func TestPersist_CreateChange(t *testing.T) {
	t.Parallel()

	now := time.Now()
	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	summaryMarkdown3 := "summary 3"
	testCases := []struct {
		name                  string
		changeLog             []entity.Change
		change                entity.Change
		expectedChange        entity.Change
		availableKeys         []service.Key
		expectedChangeLogSize int
		hasErr                bool
	}{
		{
			name: "create change successfully",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "Title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			expectedChange: entity.Change{
				ID:              "test",
				Title:           "Title 3",
				SummaryMarkdown: &summaryMarkdown3,
				ReleasedAt:      &now,
			},
			availableKeys:         []service.Key{"test"},
			expectedChangeLogSize: 3,
			hasErr:                false,
		}, {
			name: "no available key",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "Title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			expectedChange:        entity.Change{},
			availableKeys:         []service.Key{},
			expectedChangeLogSize: 2,
			hasErr:                true,
		}, {
			name: "ID already exists",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "Title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			expectedChange:        entity.Change{},
			availableKeys:         []service.Key{"12345"},
			expectedChangeLogSize: 2,
			hasErr:                true,
		}, {
			name: "Allow summary to be nil",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "Title 3",
				SummaryMarkdown: nil,
			},
			expectedChange: entity.Change{
				ID:              "22222",
				Title:           "Title 3",
				SummaryMarkdown: nil,
				ReleasedAt:      &now,
			},
			availableKeys:         []service.Key{"22222"},
			expectedChangeLogSize: 3,
			hasErr:                false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := service.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)

			mdtest.Equal(t, nil, err)

			fakeTimer := mdtest.NewTimerFake(now)
			persist := NewPersist(
				keyGen,
				fakeTimer,
				&changeLogRepo,
			)

			newChange, err := persist.CreateChange(testCase.change.Title, testCase.change.SummaryMarkdown)
			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)

			mdtest.Equal(t, testCase.expectedChange, newChange)

			changeLog, err := persist.GetChangeLog()
			mdtest.Equal(t, nil, err)

			mdtest.Equal(t, testCase.expectedChangeLogSize, len(changeLog))
		})
	}
}

func TestPersist_GetChangeLog(t *testing.T) {
	t.Parallel()

	now := time.Now()
	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	testCases := []struct {
		name          string
		changeLog     []entity.Change
		availableKeys []service.Key
	}{
		{
			name: "get full changelog successfully",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			availableKeys: []service.Key{},
		}, {
			name:          "get empty changelog successfully",
			changeLog:     []entity.Change{},
			availableKeys: []service.Key{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := service.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			mdtest.Equal(t, nil, err)

			fakeTimer := mdtest.NewTimerFake(now)
			persist := NewPersist(
				keyGen,
				fakeTimer,
				&changeLogRepo,
			)

			changeLog, err := persist.GetChangeLog()

			mdtest.Equal(t, nil, err)
			mdtest.SameElements(t, testCase.changeLog, changeLog)
		})
	}
}
