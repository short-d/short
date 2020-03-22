package db

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/app/adapter/db/table"

	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ repository.Changelog = (*ChangeLogSQL)(nil)

type ChangeLogSQL struct {
	db *sql.DB
}

func (c ChangeLogSQL) GetChangeByID(id string) (entity.Change, error) {
	statement := fmt.Sprintf(`SELECT "%s","%s","%s","%s" FROM "%s" WHERE "%s"=$1;`,
		table.ChangeLog.ColumnID,
		table.ChangeLog.ColumnTitle,
		table.ChangeLog.ColumnSummaryMarkdown,
		table.ChangeLog.ColumnReleasedAt,
		table.ChangeLog.TableName,
		table.ChangeLog.ColumnID,
	)

	row := c.db.QueryRow(statement, id)
	var change = entity.Change{}

	err := row.Scan(&change.ID, &change.Title, &change.SummaryMarkdown, &change.ReleasedAt)
	return change, err
}

func (c ChangeLogSQL) GetChangeLog() ([]entity.Change, error) {
	statement := fmt.Sprintf(`SELECT "%s","%s","%s","%s" FROM "%s";`,
		table.ChangeLog.ColumnID,
		table.ChangeLog.ColumnTitle,
		table.ChangeLog.ColumnSummaryMarkdown,
		table.ChangeLog.ColumnReleasedAt,
		table.ChangeLog.TableName,
	)

	var changelog []entity.Change
	rows, err := c.db.Query(statement)

	if err != nil {
		return changelog, err
	}

	for rows.Next() {
		change := entity.Change{}
		err = rows.Scan(&change.ID, &change.Title, &change.SummaryMarkdown, &change.ReleasedAt)
		if err != nil {
			return changelog, err
		}
		change.ReleasedAt = utc(change.ReleasedAt)
		changelog = append(changelog, change)
	}

	return changelog, nil
}

func (c ChangeLogSQL) CreateChange(newChange entity.Change) (entity.Change, error) {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s","%s","%s")
VALUES ($1, $2, $3, $4);
`,
		table.ChangeLog.TableName,
		table.ChangeLog.ColumnID,
		table.ChangeLog.ColumnTitle,
		table.ChangeLog.ColumnSummaryMarkdown,
		table.ChangeLog.ColumnReleasedAt,
	)

	_, err := c.db.Exec(
		statement,
		newChange.ID,
		newChange.Title,
		newChange.SummaryMarkdown,
		newChange.ReleasedAt,
	)

	if err != nil {
		return entity.Change{}, err
	}

	return c.GetChangeByID(newChange.ID)
}

func NewChangeLogSQL(db *sql.DB) *ChangeLogSQL {
	return &ChangeLogSQL{
		db: db,
	}
}
