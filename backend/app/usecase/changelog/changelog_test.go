// +build !integration all

package changelog

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestPersist_CreateChange(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	summaryMarkdown3 := "summary 3"
	testCases := []struct {
		name                  string
		changeLog             []entity.Change
		change                entity.Change
		expectedChange        entity.Change
		availableKeys         []keygen.Key
		roles                 map[string][]role.Role
		user                  entity.User
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
			availableKeys: []keygen.Key{"test"},
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
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
			expectedChange: entity.Change{},
			availableKeys:  []keygen.Key{},
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
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
			expectedChange: entity.Change{},
			availableKeys:  []keygen.Key{"12345"},
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
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
			availableKeys: []keygen.Key{"22222"},
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
			expectedChangeLogSize: 3,
			hasErr:                false,
		}, {
			name: "user is not allowed to create a change",
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
			availableKeys: []keygen.Key{"test"},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := keygen.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			tm := timer.NewStub(now)
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
				&userChangeLogRepo,
				au,
			)

			newChange, err := persist.CreateChange(testCase.change.Title, testCase.change.SummaryMarkdown, testCase.user)
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
		availableKeys []keygen.Key
		roles         map[string][]role.Role
		user          entity.User
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
			availableKeys: []keygen.Key{},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			user: entity.User{
				ID: "alpha",
			},
		}, {
			name:          "get empty changelog successfully",
			changeLog:     []entity.Change{},
			availableKeys: []keygen.Key{},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			user: entity.User{
				ID: "alpha",
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := keygen.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			tm := timer.NewStub(now)
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
				&userChangeLogRepo,
				au,
			)

			changeLog, err := persist.GetChangeLog()
			assert.Equal(t, nil, err)
			assert.SameElements(t, testCase.changeLog, changeLog)
		})
	}
}

func TestPersist_GetLastViewedAt(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	twoMonthsAgo := now.AddDate(0, -2, 0)
	testCases := []struct {
		name          string
		userChangeLog map[string]time.Time
		user          entity.User
		roles         map[string][]role.Role
		lastViewedAt  *time.Time
	}{
		{
			name:          "user never viewed the change log before",
			userChangeLog: map[string]time.Time{},
			user: entity.User{
				ID:    "alpha",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			lastViewedAt: nil,
		},
		{
			name:          "user viewed change log",
			userChangeLog: map[string]time.Time{"test@gmail.com": twoMonthsAgo},
			user: entity.User{
				ID:    "alpha",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			lastViewedAt: &twoMonthsAgo,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(now)
			userChangeLogRepo := repository.NewUserChangeLogFake(testCase.userChangeLog)

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
				&userChangeLogRepo,
				au,
			)

			lastViewedAt, err := persist.GetLastViewedAt(testCase.user)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.lastViewedAt, lastViewedAt)
		})
	}
}

func TestPersist_ViewChangeLog(t *testing.T) {
	t.Parallel()

	now := time.Now().UTC()
	twoMonthsAgo := now.AddDate(0, -2, 0)
	testCases := []struct {
		name          string
		userChangeLog map[string]time.Time
		user          entity.User
		roles         map[string][]role.Role
		lastViewedAt  time.Time
	}{
		{
			name:          "user viewed the change log the first time",
			userChangeLog: map[string]time.Time{},
			user: entity.User{
				ID:    "alpha",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			lastViewedAt: now,
		},
		{
			name:          "user has viewed the change log before",
			userChangeLog: map[string]time.Time{"test@gmail.com": twoMonthsAgo},
			user: entity.User{
				ID:    "alpha",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			lastViewedAt: now,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			tm := timer.NewStub(now)
			userChangeLogRepo := repository.NewUserChangeLogFake(testCase.userChangeLog)

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
				&userChangeLogRepo,
				au,
			)

			lastViewedAt, err := persist.ViewChangeLog(testCase.user)
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.lastViewedAt, lastViewedAt)
		})
	}
}

func TestPersist_GetAllChanges(t *testing.T) {
	t.Parallel()

	now := time.Now()
	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	testCases := []struct {
		name            string
		changes         []entity.Change
		expectedChanges []entity.Change
		availableKeys   []keygen.Key
		roles           map[string][]role.Role
		user            entity.User
		hasErr          bool
	}{
		{
			name: "get full changelog successfully",
			changes: []entity.Change{
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
			expectedChanges: []entity.Change{
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
			availableKeys: []keygen.Key{},
			roles: map[string][]role.Role{
				"alpha": {role.ChangeLogViewer},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: false,
		}, {
			name: "user is not allowed to get changes",
			changes: []entity.Change{
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
			expectedChanges: []entity.Change{},
			availableKeys:   []keygen.Key{},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changes)
			keyFetcher := keygen.NewKeyFetcherFake(testCase.availableKeys)
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			tm := timer.NewStub(now)
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
				&userChangeLogRepo,
				au,
			)

			changeLog, err := persist.GetAllChanges(testCase.user)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)
			assert.SameElements(t, testCase.expectedChanges, changeLog)
		})
	}
}

func TestPersist_DeleteChange(t *testing.T) {
	t.Parallel()

	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	testCases := []struct {
		name                  string
		changeLog             []entity.Change
		deleteChangeId        string
		expectedChangeLog     []entity.Change
		expectedChangeLogSize int
		roles                 map[string][]role.Role
		user                  entity.User
		hasErr                bool
	}{
		{
			name: "delete existing change successfully",
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
			deleteChangeId: "12345",
			expectedChangeLog: []entity.Change{
				{
					ID:              "54321",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
			},
			expectedChangeLogSize: 1,
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: false,
		},
		{
			name: "delete non existing change",
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
			deleteChangeId: "34567",
			expectedChangeLog: []entity.Change{
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
			expectedChangeLogSize: 2,
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: false,
		},
		{
			name: "user is not allowed to delete a change",
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
			deleteChangeId:        "12345",
			expectedChangeLogSize: 1,
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			tm := timer.NewStub(time.Now())
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
				&userChangeLogRepo,
				au,
			)

			err = persist.DeleteChange(testCase.deleteChangeId, testCase.user)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)

			changeLog, err := persist.GetChangeLog()

			assert.Equal(t, nil, err)
			assert.SameElements(t, testCase.expectedChangeLog, changeLog)
		})
	}
}

func TestPersist_UpdateChange(t *testing.T) {
	t.Parallel()

	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	summaryMarkdown3 := "summary 3"
	testCases := []struct {
		name              string
		changeLog         []entity.Change
		change            entity.Change
		roles             map[string][]role.Role
		user              entity.User
		expectedChangeLog []entity.Change
		hasErr            bool
	}{
		{
			name: "update existing change successfully",
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
				ID:              "54321",
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
			expectedChangeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
				},
				{
					ID:              "54321",
					Title:           "title 3",
					SummaryMarkdown: &summaryMarkdown3,
				},
			},
			hasErr: false,
		},
		{
			name: "update non existing change",
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
				ID:              "34567",
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			roles: map[string][]role.Role{
				"alpha": {role.Admin},
			},
			user: entity.User{
				ID: "alpha",
			},
			expectedChangeLog: []entity.Change{
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
			hasErr: false,
		},
		{
			name: "user is not allowed to update a change",
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
				ID:              "54321",
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			roles: map[string][]role.Role{
				"alpha": {role.Basic},
			},
			user: entity.User{
				ID: "alpha",
			},
			hasErr: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			changeLogRepo := repository.NewChangeLogFake(testCase.changeLog)
			keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
			keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
			assert.Equal(t, nil, err)

			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			rb := rbac.NewRBAC(fakeRolesRepo)
			au := authorizer.NewAuthorizer(rb)

			tm := timer.NewStub(time.Now())
			userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
			persist := NewPersist(
				keyGen,
				tm,
				&changeLogRepo,
				&userChangeLogRepo,
				au,
			)

			_, err = persist.UpdateChange(testCase.change.ID, testCase.change.Title, testCase.change.SummaryMarkdown, testCase.user)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)

			changeLog, err := persist.GetChangeLog()
			assert.Equal(t, nil, err)
			assert.SameElements(t, testCase.expectedChangeLog, changeLog)
		})
	}
}
