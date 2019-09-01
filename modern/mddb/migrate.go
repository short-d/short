package mddb

import (
	"database/sql"

	migrate "github.com/rubenv/sql-migrate"
)

func MigratePostgres(db *sql.DB, migrationRoot string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationRoot,
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	return err
}
