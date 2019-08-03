package modern

import (
	"database/sql"
	"github.com/rubenv/sql-migrate"
)

func MigratePostgres(db *sql.DB, migrationFilepath string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationFilepath,
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	return err
}
