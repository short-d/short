package reposql

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"short/app/adapter/reposql/table"
	"short/app/entity"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUserSql_IsEmailExist(t *testing.T) {
	testCases := []struct {
		name       string
		tableRows  *sqlmock.Rows
		email      string
		expIsExist bool
	}{
		{
			name:  "email doesn't exist",
			email: "user@example.com",
			tableRows: sqlmock.NewRows([]string{
				table.User.ColumnEmail,
			}),
			expIsExist: false,
		},
		{
			name:  "email found",
			email: "user@example.com",
			tableRows: sqlmock.
				NewRows([]string{
					table.User.ColumnEmail,
				}).
				AddRow("user@example.com"),
			expIsExist: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nil(t, err)
			defer db.Close()

			expQuery := fmt.Sprintf(`^SELECT ".+" FROM "%s" WHERE "%s"=.+$`, table.User.TableName, table.User.ColumnEmail)
			mock.ExpectQuery(expQuery).WillReturnRows(testCase.tableRows)

			userRepo := NewUser(db)
			gotIsExist, err := userRepo.IsEmailExist(testCase.email)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expIsExist, gotIsExist)
		})
	}
}

func TestUserSql_GetByEmail(t *testing.T) {
	testCases := []struct {
		name      string
		tableRows *sqlmock.Rows
		email     string
		hasErr    bool
		expUser   entity.User
	}{
		{
			name:  "email doesn't exist",
			email: "user@example.com",
			tableRows: sqlmock.NewRows([]string{
				table.User.ColumnEmail,
				table.User.ColumnName,
				table.User.ColumnLastSignedInAt,
				table.User.ColumnUpdatedAt,
				table.User.ColumnCreatedAt,
			}),
			hasErr: true,
		},
		{
			name:  "email found",
			email: "user@example.com",
			tableRows: sqlmock.
				NewRows([]string{
					table.User.ColumnEmail,
					table.User.ColumnName,
					table.User.ColumnLastSignedInAt,
					table.User.ColumnUpdatedAt,
					table.User.ColumnCreatedAt,
				}).
				AddRow("alpha@example.com", "Alpha", nil, nil, nil),
			hasErr: false,
			expUser: entity.User{
				Name:  "Alpha",
				Email: "alpha@example.com",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nil(t, err)
			defer db.Close()

			expQuery := fmt.Sprintf(`^SELECT ".+",".+",".+",".+",".+" FROM "%s" WHERE "%s"=.+$`, table.User.TableName, table.User.ColumnEmail)
			mock.ExpectQuery(expQuery).WillReturnRows(testCase.tableRows)

			userRepo := NewUser(db)

			gotUser, err := userRepo.GetByEmail(testCase.email)
			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, testCase.expUser, gotUser)
		})
	}
}

func TestUserSql_Create(t *testing.T) {
	testCases := []struct {
		name     string
		user     entity.User
		rowExist bool
		hasErr   bool
	}{
		{
			name: "email exists",
			user: entity.User{
				Email: "alpha@example.com",
			},
			rowExist: true,
			hasErr:   true,
		},
		{
			name: "email available",
			user: entity.User{
				Email: "user@example.com",
				Name:  "Test User",
			},
			rowExist: false,
			hasErr:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nil(t, err)
			defer db.Close()

			expStatement := fmt.Sprintf(`INSERT\s*INTO\s*"%s"`, table.User.TableName)
			if testCase.rowExist {
				mock.ExpectExec(expStatement).WillReturnError(errors.New("row exists"))
			} else {
				mock.ExpectExec(expStatement).WillReturnResult(driver.ResultNoRows)
			}

			userRepo := NewUser(db)

			err = userRepo.Create(testCase.user)
			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
