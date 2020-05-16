package tool

import (
	"database/sql"
	"fmt"

	"github.com/short-d/app/fw/logger"

	"github.com/short-d/app/fw/db"
	"github.com/short-d/short/backend/app/usecase/keygen"
)

// Data transform existing data to the new format and move them to the correct
// location.
type Data struct {
	keyGen keygen.KeyGenerator
	db     *sql.DB
	logger logger.Logger
}

// EmailToID generates IDs for users and changes DB tables reference those IDs.
func (d Data) EmailToID(batchSize int) {
	d.logger.Info(fmt.Sprintf("Migrating upto %d accounts in the database", batchSize))
	rows, err := d.db.Query(`
SELECT email FROM "user" WHERE id IS NULL LIMIT $1;
`, batchSize)
	if err != nil {
		panic(err)
	}
	count := 0
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
		count++
	}
	d.logger.Info(fmt.Sprintf("Migrated %d accounts.", count))
}

// NewData creates data manage
func NewData(
	dbConfig db.Config,
	dbConnector db.Connector,
	keyGen keygen.KeyGenerator,
	logger logger.Logger,
) (Data, error) {
	sqlDB, err := dbConnector.Connect(dbConfig)
	if err != nil {
		return Data{}, err
	}
	return Data{
		keyGen: keyGen,
		db:     sqlDB,
		logger: logger,
	}, nil
}
