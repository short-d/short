// +build integration all

package sqldb_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db/dbtest"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
)

var insertUserRoleRowSQL = fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2)`,
	table.UserRole.TableName,
	table.UserRole.ColumnUserID,
	table.UserRole.ColumnRole,
)

type userRoleTableRow struct {
	userID string
	role   role.Role
}

func TestUserRoleSQL_GetRoles(t *testing.T) {
	testCases := []struct {
		name          string
		user          entity.User
		userRoleRows  []userRoleTableRow
		expectedRoles []role.Role
		hasErr        bool
	}{
		{
			name: "get roles for user has no role",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows:  []userRoleTableRow{},
			expectedRoles: []role.Role{},
			hasErr:        false,
		},
		{
			name: "get roles for user with existing roles",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows: []userRoleTableRow{
				{"1343", role.Basic},
				{"1343", role.ChangeLogViewer},
			},
			expectedRoles: []role.Role{role.Basic, role.ChangeLogViewer},
			hasErr:        false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertUserRoleRow(t, sqlDB, testCase.userRoleRows)

					userRoleRepo := sqldb.NewUserRoleSQL(sqlDB)
					roles, err := userRoleRepo.GetRoles(testCase.user)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedRoles, roles)
				})
		})
	}
}

func TestUserRoleSQL_AddRole(t *testing.T) {
	testCases := []struct {
		name          string
		user          entity.User
		userRoleRows  []userRoleTableRow
		newRoles      []role.Role
		expectedRoles []role.Role
		hasErr        bool
	}{
		{
			name: "add 1 role for nonexistent user",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows: []userRoleTableRow{
				{"4444", role.Basic},
			},
			newRoles:      []role.Role{role.ChangeLogViewer},
			expectedRoles: []role.Role{role.ChangeLogViewer},
			hasErr:        false,
		},
		{
			name: "add 1 role for user with existing roles",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows: []userRoleTableRow{
				{"1343", role.Basic},
				{"4444", role.Basic},
			},
			newRoles:      []role.Role{role.ChangeLogViewer},
			expectedRoles: []role.Role{role.Basic, role.ChangeLogViewer},
			hasErr:        false,
		},
		{
			name: "add multiple roles for a give user",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows: []userRoleTableRow{
				{"1343", role.Basic},
				{"4444", role.Basic},
			},
			newRoles:      []role.Role{role.ChangeLogViewer, role.Admin},
			expectedRoles: []role.Role{role.Admin, role.Basic, role.ChangeLogViewer},
			hasErr:        false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertUserRoleRow(t, sqlDB, testCase.userRoleRows)

					userRoleRepo := sqldb.NewUserRoleSQL(sqlDB)

					for _, newRole := range testCase.newRoles {
						err := userRoleRepo.AddRole(testCase.user, newRole)

						if testCase.hasErr {
							assert.NotEqual(t, nil, err)
							return
						}
					}

					roles, err := userRoleRepo.GetRoles(testCase.user)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedRoles, roles)
				})
		})
	}
}

func TestUserRoleSQL_DeleteRole(t *testing.T) {
	// TODO(issue#755) Add integration test for foreign key constraint
	testCases := []struct {
		name          string
		user          entity.User
		userRoleRows  []userRoleTableRow
		toDelete      role.Role
		expectedRoles []role.Role
		hasErr        bool
	}{
		{
			name: "should remove a role from the given",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows: []userRoleTableRow{
				{"1343", role.ChangeLogViewer},
			},
			toDelete:      role.ChangeLogViewer,
			expectedRoles: []role.Role{},
			hasErr:        false,
		},
		{
			name: "should do nothing if a user doesn't have the role",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows: []userRoleTableRow{
				{"1343", role.Admin},
				{"1343", role.Basic},
				{"1343", role.ChangeLogViewer},
			},
			toDelete:      role.ChangeLogEditor,
			expectedRoles: []role.Role{role.Admin, role.Basic, role.ChangeLogViewer},
			hasErr:        false,
		},
		{
			name: "should do nothing for nonexistent user",
			user: entity.User{
				ID: "1343",
			},
			userRoleRows: []userRoleTableRow{
				{"0000", role.Admin},
				{"0000", role.Basic},
				{"0000", role.ChangeLogViewer},
			},
			toDelete:      role.ChangeLogEditor,
			expectedRoles: []role.Role{},
			hasErr:        false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					userRoleRepo := sqldb.NewUserRoleSQL(sqlDB)
					insertUserRoleRow(t, sqlDB, testCase.userRoleRows)

					err := userRoleRepo.DeleteRole(testCase.user, testCase.toDelete)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					roles, err := userRoleRepo.GetRoles(testCase.user)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedRoles, roles)
				})
		})
	}
}

func insertUserRoleRow(
	t *testing.T,
	sqlDB *sql.DB,
	tableRows []userRoleTableRow,
) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertUserRoleRowSQL,
			tableRow.userID,
			tableRow.role,
		)
		assert.Equal(t, nil, err)
	}
}
