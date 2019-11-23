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

func TestNewUserURLRelationSQL(t *testing.T) {
	testCases := []struct {
		name      string
		user      entity.User
		alias     string
		expHasErr bool
	}{
		{
			name:      "",
			expHasErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			statement := fmt.Sprintf(
				`INSERT INTO "%s" (.+) VALUES (.+)`,
				table.UserURLRelation.TableName,
			)

			if testCase.expHasErr {
				stub.ExpectExec(statement).WillReturnError(errors.New("row exists"))
			} else {
				stub.ExpectExec(statement).WillReturnResult(driver.ResultNoRows)
			}

			userURLRelationRepo := NewUserURLRelationSQL(db)

			url := entity.URL{
				Alias: testCase.alias,
			}
			err = userURLRelationRepo.CreateRelation(testCase.user, url)

			if testCase.expHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
		})
	}
}
