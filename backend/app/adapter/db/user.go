package db

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db/table"
	"short/app/entity"
	"short/app/usecase/repository"
)

var _ repository.User = (*UserSQL)(nil)

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
SELECT "%s","%s","%s","%s","%s", "%s"
FROM "%s" 
WHERE "%s"=$1;
`,
		table.User.ColumnID,
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
		&user.ID,
		&user.Email,
		&user.Name,
		&user.LastSignedInAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	user.CreatedAt = utc(user.CreatedAt)
	user.UpdatedAt = utc(user.UpdatedAt)
	user.LastSignedInAt = utc(user.LastSignedInAt)

	return user, nil
}

// CreateUser inserts a new User into user table.
func (u *UserSQL) CreateUser(user entity.User) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s","%s","%s","%s","%s")
VALUES ($1, $2, $3, $4, $5, $6)
`,
		table.User.TableName,
		table.User.ColumnID,
		table.User.ColumnEmail,
		table.User.ColumnName,
		table.User.ColumnLastSignedInAt,
		table.User.ColumnCreatedAt,
		table.User.ColumnUpdatedAt,
	)

	_, err := u.db.Exec(
		statement,
		user.ID,
		user.Email,
		user.Name,
		user.LastSignedInAt,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

// UpdateUserID updates the ID of an user in user table with given email address.
func (u UserSQL) UpdateUserID(email string, userID string) error {
	isExist, err := u.IsEmailExist(email)
	if err != nil {
		return err
	}

	if !isExist {
		return fmt.Errorf("email %s does not exist", email)
	}
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1
WHERE "%s"=$2
`,
		table.User.TableName,
		table.User.ColumnID,
		table.User.ColumnEmail)
	_, err = u.db.Exec(statement, userID, email)
	return err
}

// NewUserSQL creates UserSQL
func NewUserSQL(db *sql.DB) *UserSQL {
	return &UserSQL{
		db: db,
	}
}
