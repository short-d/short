package db

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db/table"
	"short/app/entity"
	"short/app/usecase/repo"
)

var _ repo.User = (*UserSQL)(nil)

// UserSQL accesses User information in user table through SQL.
type UserSQL struct {
	db *sql.DB
}

// IsEmailExist checks whether a given email exist in user table.
func (u UserSQL) IsEmailExist(email string) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;
`,
		table.User.ColumnEmail,
		table.User.TableName,
		table.User.ColumnEmail,
	)

	err := u.db.QueryRow(query, email).Scan(&email)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetUserByEmail finds an User in user table given email.
func (u UserSQL) GetUserByEmail(email string) (entity.User, error) {
	query := fmt.Sprintf(`
SELECT "%s","%s","%s","%s","%s"
FROM "%s" 
WHERE "%s"=$1;
`,
		table.User.ColumnEmail,
		table.User.ColumnName,
		table.User.ColumnLastSignedInAt,
		table.User.ColumnCreatedAt,
		table.User.ColumnUpdatedAt,
		table.User.TableName,
		table.User.ColumnEmail,
	)

	row := u.db.QueryRow(query, email)

	user := entity.User{}
	err := row.Scan(
		&user.Email,
		&user.Name,
		&user.LastSignedInAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser inserts a new User into user table.
func (u *UserSQL) CreateUser(user entity.User) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s","%s","%s")
VALUES ($1, $2, $3, $4, $5)
`,
		table.User.TableName,
		table.User.ColumnEmail,
		table.User.ColumnName,
		table.User.ColumnLastSignedInAt,
		table.User.ColumnCreatedAt,
		table.User.ColumnUpdatedAt,
	)

	_, err := u.db.Exec(statement, user.Email, user.Name, user.LastSignedInAt, user.CreatedAt, user.UpdatedAt)
	return err
}

// NewUserSQL creates UserSQL
func NewUserSQL(db *sql.DB) *UserSQL {
	return &UserSQL{
		db: db,
	}
}
