package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.ChangeLog = (*ChangeLogSQL)(nil)

// ChangeLogSQL accesses ChangeLog information in change_log table through SQL.
type ChangeLogSQL struct {
	db *sql.DB
}

// GetChangeLog retrieves full changelog from change_log table.
func (c ChangeLogSQL) GetChangeLog() ([]entity.Change, error) {
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s" 
FROM "%s";`,
		table.ChangeLog.ColumnID,
		table.ChangeLog.ColumnTitle,
		table.ChangeLog.ColumnSummaryMarkdown,
		table.ChangeLog.ColumnReleasedAt,
		table.ChangeLog.TableName,
	)

	rows, err := c.db.Query(statement)
	if err != nil {
		return []entity.Change{}, err
	}

	// the consumer of GetChangeLog expects empty slice instead of `nil` if there are no records
	changeLog := []entity.Change{}
	for rows.Next() {
		change := entity.Change{}
		err = rows.Scan(&change.ID, &change.Title, &change.SummaryMarkdown, &change.ReleasedAt)
		if err != nil {
			return changeLog, err
		}
		change.ReleasedAt = change.ReleasedAt.UTC()
		changeLog = append(changeLog, change)
	}

	return changeLog, nil
}

// CreateChange adds a new Change into change_log table.
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

	return newChange, nil
}

// DeleteChange removes a change with the given id
func (c ChangeLogSQL) DeleteChange(id string) error {
	statement := fmt.Sprintf(`
DELETE FROM "%s"
WHERE "%s" = $1;
`,
		table.ChangeLog.TableName,
		table.ChangeLog.ColumnID,
	)

	_, err := c.db.Exec(statement, id)
	return err
}

// NewChangeLogSQL creates ChangeLogSQL
func NewChangeLogSQL(db *sql.DB) ChangeLogSQL {
	return ChangeLogSQL{
		db: db,
	}
}
