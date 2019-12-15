// +build integration

package db_test

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"path"
	"short/app/adapter/db"
	"short/app/adapter/db/sqltest"
	"short/app/adapter/db/table"
	"short/app/entity"
	"short/dep"
	"strconv"
	"testing"
	"time"

	"github.com/byliuyang/app/fw"
	"github.com/byliuyang/app/mdtest"
)

var dbConnector fw.DBConnector
var dbMigrationTool fw.DBMigrationTool

var dbConfig fw.DBConfig
var dbMigrationRoot string

var insertRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s, %s, %s, %s)
VALUES ($1, $2, $3, $4, $5)`,
	table.URL.TableName,
	table.URL.ColumnAlias,
	table.URL.ColumnOriginalURL,
	table.URL.ColumnCreatedAt,
	table.URL.ColumnExpireAt,
	table.URL.ColumnUpdatedAt,
)

type tableRow struct {
	alias     string
	longLink  string
	createdAt time.Time
	expireAt  time.Time
	updatedAt time.Time
}

func TestURLSql_IsAliasExist(t *testing.T) {
	testCases := []struct {
		name       string
		tableRows  []tableRow
		alias      string
		expIsExist bool
	}{
		{
			name:       "alias doesn't exist",
			alias:      "gg",
			tableRows:  []tableRow{},
			expIsExist: false,
		},
		{
			name:  "alias found",
			alias: "gg",
			tableRows: []tableRow{
				{alias: "gg"},
			},
			expIsExist: true,
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
					for _, tableRow := range testCase.tableRows {
						_, err := sqlDB.Exec(
							insertRowSQL,
							tableRow.alias,
							tableRow.longLink,
							tableRow.createdAt,
							tableRow.expireAt,
							tableRow.updatedAt,
						)
						mdtest.Equal(t, nil, err)
					}

					urlRepo := db.NewURLSql(sqlDB)
					gotIsExist, err := urlRepo.IsAliasExist(testCase.alias)
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expIsExist, gotIsExist)
				})
		})
	}
}

func TestURLSql_GetByAlias(t *testing.T) {
	testCases := []struct {
		name        string
		tableRows   *mdtest.TableRows
		alias       string
		hasErr      bool
		expectedURL entity.URL
	}{
		{
			name: "alias not found",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
				table.URL.ColumnOriginalURL,
				table.URL.ColumnExpireAt,
				table.URL.ColumnCreatedAt,
				table.URL.ColumnUpdatedAt,
			}),
			alias:  "220uFicCJj",
			hasErr: true,
		},
		{
			name: "found url",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
				table.URL.ColumnOriginalURL,
				table.URL.ColumnExpireAt,
				table.URL.ColumnCreatedAt,
				table.URL.ColumnUpdatedAt,
			}).AddRow(
				"220uFicCJj",
				"http://www.google.com",
				sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
				sqltest.MustParseSQLTime("2017-05-01 08:02:16"),
				nil,
			).AddRow(
				"yDOBcj5HIPbUAsw",
				"http://www.facebook.com",
				sqltest.MustParseSQLTime("2018-04-02 08:02:16"),
				sqltest.MustParseSQLTime("2017-05-01 08:02:16"),
				nil,
			),
			alias:  "220uFicCJj",
			hasErr: false,
			expectedURL: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
				CreatedAt:   sqltest.MustParseSQLTime("2017-05-01 08:02:16"),
				UpdatedAt:   nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sqlDB, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer sqlDB.Close()

			statement := fmt.Sprintf(`^SELECT .+ FROM "%s" WHERE "%s"=.+$`, table.URL.TableName, table.URL.ColumnAlias)
			stub.ExpectQuery(statement).WillReturnRows(testCase.tableRows)

			urlRepo := db.NewURLSql(sqlDB)
			url, err := urlRepo.GetByAlias("220uFicCJj")

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, url)
		})
	}
}

func TestURLSql_Create(t *testing.T) {
	testCases := []struct {
		name     string
		url      entity.URL
		rowExist bool
		hasErr   bool
	}{
		{
			name: "url exists",
			url: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
			},
			rowExist: true,
			hasErr:   true,
		},
		{
			name: "successfully create url",
			url: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
			},
			rowExist: false,
			hasErr:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			sqlDB, stub, err := mdtest.NewSQLStub()

			mdtest.Equal(t, nil, err)
			defer sqlDB.Close()

			statement := fmt.Sprintf(`INSERT INTO "%s" .+ VALUES .+`, table.URL.TableName)

			if testCase.rowExist {
				stub.ExpectExec(statement).WillReturnError(errors.New("row exists"))
			} else {
				stub.ExpectExec(statement).WillReturnResult(driver.ResultNoRows)
			}

			urlRepo := db.NewURLSql(sqlDB)
			err = urlRepo.Create(testCase.url)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
		})
	}
}

func TestMain(m *testing.M) {
	env := dep.InjectEnvironment()
	env.AutoLoadDotEnvFile()

	host := env.GetEnv("DB_HOST", "")
	portStr := env.GetEnv("DB_PORT", "")
	port := mustInt(portStr)
	user := env.GetEnv("DB_USER", "")
	password := env.GetEnv("DB_PASSWORD", "")
	dbName := env.GetEnv("DB_NAME", "")

	dbConfig = fw.DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}

	dbMigrationRoot = path.Join(env.GetEnv("MIGRATION_ROOT", ""))

	dbConnector = dep.InjectDBConnector()
	dbMigrationTool = dep.InjectDBMigrationTool()

	m.Run()
}

func mustInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return num
}
