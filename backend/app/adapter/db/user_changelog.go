package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ repository.UserChangeLog = (*UserChangeLogSQL)(nil)

type UserChangeLogSQL struct {
	db    *sql.DB
	timer fw.Timer
}

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
	if err == sql.ErrNoRows {
		return lastViewedAt, nil
	} else if err != nil {
		return lastViewedAt, err
	}

	return lastViewedAt, nil
}

func (u UserChangeLogSQL) UpdateLastViewedAt(user entity.User) (time.Time, error) {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s")
VALUES ($1, $2, $3)
ON CONFLICT (%s)
DO UPDATE SET %s=$3;
`,
		table.UserChangeLog.TableName,
		table.UserChangeLog.ColumnUserID,
		table.UserChangeLog.ColumnEmail,
		table.UserChangeLog.ColumnLastViewedAt,
		table.UserChangeLog.ColumnEmail,
		table.UserChangeLog.ColumnLastViewedAt,
	)

	now := u.timer.Now()
	_, err := u.db.Exec(
		statement,
		user.ID,
		user.Email,
		now,
	)

	if err != nil {
		return time.Time{}, err
	}

	return now, nil
}

func NewUserChangeLogSql(db *sql.DB, timer fw.Timer) *UserChangeLogSQL {
	return &UserChangeLogSQL{
		db:    db,
		timer: timer,
	}
}
