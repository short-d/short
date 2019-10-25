package db

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"short/app/adapter/db/table"
	"short/app/entity"
	"testing"

	"github.com/byliuyang/app/mdtest"
)

func TestUserSql_IsEmailExist(t *testing.T) {
	testCases := []struct {
		name       string
		tableRows  *mdtest.TableRows
		email      string
		expIsExist bool
	}{
		{
			name:  "email doesn't exist",
			email: "user@example.com",
			tableRows: mdtest.NewTableRows([]string{
				table.User.ColumnEmail,
			}),
			expIsExist: false,
		},
		{
			name:  "email found",
			email: "user@example.com",
			tableRows: mdtest.NewTableRows([]string{
				table.User.ColumnEmail,
			}).
				AddRow("user@example.com"),
			expIsExist: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			expQuery := fmt.Sprintf(`^SELECT ".+" FROM "%s" WHERE "%s"=.+$`, table.User.TableName, table.User.ColumnEmail)
			stub.ExpectQuery(expQuery).WillReturnRows(testCase.tableRows)

			userRepo := NewUserSQL(db)
			gotIsExist, err := userRepo.IsEmailExist(testCase.email)
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expIsExist, gotIsExist)
		})
	}
}

func TestUserSql_GetUserByEmail(t *testing.T) {
	testCases := []struct {
		name      string
		tableRows *mdtest.TableRows
		email     string
		hasErr    bool
		expUser   entity.User
	}{
		{
			name:  "email doesn't exist",
			email: "user@example.com",
			tableRows: mdtest.NewTableRows([]string{
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
			tableRows: mdtest.NewTableRows([]string{
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
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			expQuery := fmt.Sprintf(`^SELECT ".+",".+",".+",".+",".+" FROM "%s" WHERE "%s"=.+$`,
				table.User.TableName, table.User.ColumnEmail)
			stub.ExpectQuery(expQuery).WillReturnRows(testCase.tableRows)

			userRepo := NewUserSQL(db)

			gotUser, err := userRepo.GetUserByEmail(testCase.email)
			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expUser, gotUser)
		})
	}
}

func TestUserSql_CreateUser(t *testing.T) {
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
			name: "no given email",
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
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			expStatement := fmt.Sprintf(`INSERT\s+INTO\s+"%s"`, table.User.TableName)
			if testCase.rowExist {
				stub.ExpectExec(expStatement).WillReturnError(errors.New("row exists"))
			} else {
				stub.ExpectExec(expStatement).WillReturnResult(driver.ResultNoRows)
			}

			userRepo := NewUserSQL(db)

			err = userRepo.CreateUser(testCase.user)
			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
		})
	}
}

func TestUserSQL_UpdateUserID(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		rowExist bool
		hasErr   bool
	}{
		{
			name:     "user not found",
			email:    "alpha@example.com",
			rowExist: false,
			hasErr:   true,
		},
		{
			name:     "update user ID successfully",
			email:    "alpha@example.com",
			rowExist: true,
			hasErr:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			expStatement := fmt.Sprintf(
				`UPDATE\s+"%s"\s+SET\s+"%s"=.+ WHERE\s+"%s"=.+`,
				table.User.TableName,
				table.User.ColumnID,
				table.User.ColumnEmail)
			if testCase.rowExist {
				stub.ExpectExec(expStatement).WillReturnResult(driver.ResultNoRows)
			} else {
				stub.ExpectExec(expStatement).WillReturnError(errors.New("row not found"))
			}

			userRepo := NewUserSQL(db)

			err = userRepo.UpdateUserID(testCase.email, "userID")
			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
		})
	}
}
