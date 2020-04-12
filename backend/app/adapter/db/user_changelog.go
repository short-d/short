package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ repository.UserChangeLog = (*UserChangeLogSQL)(nil)

// UserChangeLogSQL accesses UserChangeLog information in user_changelog table through SQL.
type UserChangeLogSQL struct {
	db *sql.DB
}

// GetLastViewedAt finds LastViewedAt time for given user.
func (u UserChangeLogSQL) GetLastViewedAt(user entity.User) (time.Time, error) {
	statement := fmt.Sprintf(`
SELECT "%s" 
FROM "%s"
WHERE "%s"=$1;`,
		table.UserChangeLog.ColumnLastViewedAt,
		table.UserChangeLog.TableName,
		table.UserChangeLog.ColumnEmail,
	)

	row := u.db.QueryRow(statement, user.Email)
	lastViewedAt := time.Time{}
	err := row.Scan(&lastViewedAt)
	return lastViewedAt, err
}

// UpdateLastViewedAt updates LastViewedAt time for given user to currentTime.
func (u UserChangeLogSQL) UpdateLastViewedAt(user entity.User, currentTime time.Time) (time.Time, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET %s=$1
WHERE %s=$2
`,
		table.UserChangeLog.TableName,
		table.UserChangeLog.ColumnLastViewedAt,
		table.UserChangeLog.ColumnEmail,
	)

	result, err := u.db.Exec(
		statement,
		currentTime,
		user.Email,
	)
	if err != nil {
		return time.Time{}, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return time.Time{}, err
	}

	if rowsUpdated == 0 {
		return time.Time{}, errors.New("sql: no rows updated")
	}

	return currentTime, err
}

// CreateRelation inserts a new entry into user_changelog table.
func (u UserChangeLogSQL) CreateRelation(user entity.User, currentTime time.Time) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s")
VALUES ($1, $2, $3);
`,
		table.UserChangeLog.TableName,
		table.UserChangeLog.ColumnUserID,
		table.UserChangeLog.ColumnEmail,
		table.UserChangeLog.ColumnLastViewedAt,
	)

	_, err := u.db.Exec(
		statement,
		user.ID,
		user.Email,
		currentTime,
	)

	return err
}

// NewUserChangeLogSQL creates UserChangeLogSQL
func NewUserChangeLogSQL(db *sql.DB) UserChangeLogSQL {
	return UserChangeLogSQL{
		db: db,
	}
}