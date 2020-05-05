// +build !integration all

package changelog

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/external"
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/repository"
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
		availableKeys         []external.Key
		expectedChangeLogSize int
		hasErr                bool
	}{
		{
			name: "create change successfully",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			expectedChange: entity.Change{
				ID:              "test",
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
				ReleasedAt:      now,
			},
			availableKeys:         []external.Key{"test"},
			expectedChangeLogSize: 3,
			hasErr:                false,
		}, {
			name: "no available key",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			expectedChange:        entity.Change{},
			availableKeys:         []external.Key{},
			expectedChangeLogSize: 2,
			hasErr:                true,
		}, {
			name: "ID already exists",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			expectedChange:        entity.Change{},
			availableKeys:         []external.Key{"12345"},
			expectedChangeLogSize: 2,
			hasErr:                true,
		}, {
			name: "allow summary to be nil",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			change: entity.Change{
				Title:           "title 3",
				SummaryMarkdown: nil,
			},
			expectedChange: entity.Change{
				ID:              "22222",
				Title:           "title 3",
				SummaryMarkdown: nil,
				ReleasedAt:      now,
			},
			availableKeys:         []external.Key{"22222"},
			expectedChangeLogSize: 3,
			hasErr:                false,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := external.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			fakeTimer := timer.NewStub(now)
			persist := NewPersist(
				keyGen,
				fakeTimer,
				&changeLogRepo,
			)

			newChange, err := persist.CreateChange(testCase.change.Title, testCase.change.SummaryMarkdown)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)

			assert.Equal(t, testCase.expectedChange, newChange)

			changeLog, err := persist.GetChangeLog()
			assert.Equal(t, nil, err)

			assert.Equal(t, testCase.expectedChangeLogSize, len(changeLog))
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
		availableKeys []external.Key
	}{
		{
			name: "get full changelog successfully",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			availableKeys: []external.Key{},
		}, {
			name:          "get empty changelog successfully",
			changeLog:     []entity.Change{},
			availableKeys: []external.Key{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := external.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(now)
			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
			)

			changeLog, err := persist.GetChangeLog()
			assert.Equal(t, nil, err)
			assert.SameElements(t, testCase.changeLog, changeLog)
		})
	}
}
