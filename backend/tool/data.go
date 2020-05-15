package tool

import (
	"database/sql"

	"github.com/short-d/app/fw/db"
	"github.com/short-d/short/backend/app/usecase/keygen"
)

// Data transform existing data to the new format and move them to the correct
// location.
type Data struct {
	keyGen keygen.KeyGenerator
	db     *sql.DB
}

// EmailToID generates IDs for users and changes DB tables reference those IDs.
func (d Data) EmailToID() {
	rows, err := d.db.Query(`
SELECT email FROM "user";
`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			panic(err)
		}
		key, err := d.keyGen.NewKey()
		if err != nil {
			panic(err)
		}
		_, err = d.db.Exec(`
UPDATE "user" SET id=$1 WHERE email=$2;
`, key, email)
		_, err = d.db.Exec(`
UPDATE "user_url_relation" SET user_id=$1 WHERE user_email=$2;
`, key, email)
	}
}

// NewData creates data management tool.
func NewData(
	dbConfig db.Config,
	dbConnector db.Connector,
	keyGen keygen.KeyGenerator,
) (Data, error) {
	sqlDB, err := dbConnector.Connect(dbConfig)
	if err != nil {
		return Data{}, err
	}
	return Data{
		keyGen: keyGen,
		db:     sqlDB,
	}, nil
}
