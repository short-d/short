// +build integration all

package db_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/role"
)

func TestUserRoleSql_AddRole(t *testing.T) {
	testCases := []struct {
		name          string
		user          entity.User
		toAdd         []role.Role
		expectedRoles []role.Role
		hasErr        bool
	}{
		{
			name: "simple add",
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
			expectedRoles: []role.Role{role.Basic, role.ChangeLogViewer, role.Admin},
			hasErr:        false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mdtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					userRoleRepo := db.NewUserRoleSQL(sqlDB)

					for _, toAdd := range testCase.toAdd {
						err := userRoleRepo.AddRole(testCase.user, toAdd)

						if testCase.hasErr {
							mdtest.NotEqual(t, nil, err)
							return
						}
					}

					roles, err := userRoleRepo.GetUserRoles(testCase.user)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}

					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedRoles, roles)

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
			name: "should delete the lowest role",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.Admin, role.ChangeLogViewer, role.Basic},
			toDelete:      role.Basic,
			expectedRoles: []role.Role{role.ChangeLogViewer, role.Admin},
			hasErr:        false,
		},
		{
			name: "should delete the middle role",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.Admin, role.ChangeLogViewer, role.Basic},
			toDelete:      role.ChangeLogViewer,
			expectedRoles: []role.Role{role.Basic, role.Admin},
			hasErr:        false,
		},
		{
			name: "should do nothing",
			user: entity.User{
				ID: "1343",
			},
			toAdd:         []role.Role{role.Admin, role.ChangeLogViewer, role.Basic},
			toDelete:      role.ChangeLogEditor,
			expectedRoles: []role.Role{role.Basic, role.ChangeLogViewer, role.Admin},
			hasErr:        false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mdtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					userRoleRepo := db.NewUserRoleSQL(sqlDB)

					for _, toAdd := range testCase.toAdd {
						err := userRoleRepo.AddRole(testCase.user, toAdd)

						if testCase.hasErr {
							mdtest.NotEqual(t, nil, err)
							return
						}
					}

					err := userRoleRepo.DeleteRole(testCase.user, testCase.toDelete)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}

					roles, err := userRoleRepo.GetUserRoles(testCase.user)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}

					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedRoles, roles)

					_, _ = sqlDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table.UserRole.TableName))
				})
		})
	}
}
