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
	"github.com/short-d/short/backend/app/usecase/authorizer/role"
)

func TestUserRoleSQL_AddRole(t *testing.T) {
	testCases := []struct {
		name          string
		user          entity.User
		toAdd         []role.Role
		expectedRoles []role.Role
		hasErr        bool
	}{
		{
			name: "add 1 role",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.ChangeLogViewer},
			expectedRoles: []role.Role{role.ChangeLogViewer},
			hasErr:        false,
		},
		{
			name: "add multiple",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.ChangeLogViewer, role.Basic, role.Admin},
			expectedRoles: []role.Role{role.Admin, role.Basic, role.ChangeLogViewer},
			hasErr:        false,
		},
		{
			name: "nonexistent user",
			user: entity.User{
				ID: "0000",
			},
			toAdd:         []role.Role{},
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

					for _, toAdd := range testCase.toAdd {
						err := userRoleRepo.AddRole(testCase.user, toAdd)

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

					_, _ = sqlDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table.UserRole.TableName))
				})
		})
	}
}

func TestUserRoleSql_DeleteRole(t *testing.T) {
	testCases := []struct {
		name          string
		user          entity.User
		toAdd         []role.Role
		toDelete      role.Role
		expectedRoles []role.Role
		hasErr        bool
	}{
		{
			name: "should delete the record",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.ChangeLogViewer},
			toDelete:      role.ChangeLogViewer,
			expectedRoles: []role.Role{},
			hasErr:        false,
		},
		{
			name: "should delete the higher role",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.Admin, role.ChangeLogViewer, role.Basic},
			toDelete:      role.Admin,
			expectedRoles: []role.Role{role.Basic, role.ChangeLogViewer},
			hasErr:        false,
		},
		{
			name: "should do nothing",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.Admin, role.Basic, role.ChangeLogViewer},
			toDelete:      role.ChangeLogEditor,
			expectedRoles: []role.Role{role.Admin, role.Basic, role.ChangeLogViewer},
			hasErr:        false,
		},
		{
			name: "should do nothing for nonexistent user",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{},
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

					for _, toAdd := range testCase.toAdd {
						err := userRoleRepo.AddRole(testCase.user, toAdd)

						if testCase.hasErr {
							assert.NotEqual(t, nil, err)
							return
						}
					}

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

					_, _ = sqlDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table.UserRole.TableName))
				})
		})
	}
}
