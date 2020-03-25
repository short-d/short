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

	testCases := []struct {
		name          string
		changeLog     []entity.Change
		change        entity.Change
		availableKeys []service.Key
		hasErr        bool
	}{
		{
			name: "Create change successfully",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: "Summary 1",
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: "Summary 2",
				},
			},
			change: entity.Change{
				Title:           "Title 3",
				SummaryMarkdown: "Summary 3",
			},
			availableKeys: []service.Key{"test"},
			hasErr:        false,
		}, {
			name: "no available key",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: "Summary 1",
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: "Summary 2",
				},
			},
			change: entity.Change{
				Title:           "Title 3",
				SummaryMarkdown: "Summary 3",
			},
			availableKeys: []service.Key{},
			hasErr:        true,
		}, {
			name: "ID already exists",
			changeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "Title 1",
					SummaryMarkdown: "Summary 1",
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: "Summary 2",
				},
			},
			change: entity.Change{
				Title:           "Title 3",
				SummaryMarkdown: "Summary 3",
			},
			availableKeys: []service.Key{"12345"},
			hasErr:        true,
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

			fakeTimer := mdtest.NewTimerFake(time.Now())
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
			mdtest.Equal(t, testCase.change.Title, newChange.Title)
			mdtest.Equal(t, testCase.change.SummaryMarkdown, newChange.SummaryMarkdown)
			mdtest.Equal(t, fakeTimer.Now(), *newChange.ReleasedAt)

			changeLog, err := persist.GetChangeLog()

			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, len(testCase.changeLog)+1, len(changeLog))
		})
	}
}

func TestPersist_GetChangeLog(t *testing.T) {
	t.Parallel()

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
					SummaryMarkdown: "Summary 1",
				},
				{
					ID:              "54321",
					Title:           "Title 2",
					SummaryMarkdown: "Summary 2",
				},
			},
			availableKeys: []service.Key{
				"11111",
				"22222",
				"33333",
			},
		}, {
			name:      "get empty changelog successfully",
			changeLog: []entity.Change{},
			availableKeys: []service.Key{
				"11111",
				"22222",
				"33333",
			},
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

			fakeTimer := mdtest.NewTimerFake(time.Now())
			persist := NewPersist(
				keyGen,
				fakeTimer,
				&changeLogRepo,
			)

			changeLog, err := persist.GetChangeLog()

			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, len(testCase.changeLog), len(changeLog))
		})
	}
}
