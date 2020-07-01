package sqldb

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.App = (*AppSQL)(nil)

// AppSQL access third party app info through SQL.
type AppSQL struct {
	db *sql.DB
}

// FindAppByID  fetches an app with given ID from SQL database.
func (a AppSQL) FindAppByID(id string) (entity.App, error) {
	query := fmt.Sprintf(`
SELECT "%s", "%s" 
FROM "%s" WHERE "%s"=$1;
`,
		table.App.ColumnName,
		table.App.ColumnCreatedAt,
		table.App.TableName,
		table.App.ColumnID,
	)
	app := entity.App{}
	err := a.db.QueryRow(query, id).Scan(&app.Name, &app.CreatedAt)
	if err == nil {
		app.ID = id
		return app, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return entity.App{},
			repository.ErrEntryNotFound(fmt.Sprintf("ID(%s)", id))
	}
	return entity.App{}, err
}

// NewAppSQL creates AppSQL.
func NewAppSQL(db *sql.DB) AppSQL {
	return AppSQL{
		db: db,
	}
}
