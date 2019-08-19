package new

import (
	"database/sql"
	"short/modern/mddb"
)

type ServiceLauncher func(db *sql.DB)

func DB(
	host string,
	port int,
	user string,
	password string,
	dbName string,
	migrationRoot string,
	serviceLauncher ServiceLauncher,
) {
	db, err := mddb.NewPostgresDb(host, port, user, password, dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = mddb.MigratePostgres(db, migrationRoot)
	if err != nil {
		panic(err)
	}
	serviceLauncher(db)
}
