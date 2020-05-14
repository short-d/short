package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.User = (*UserSQL)(nil)

// UserSQL accesses User information in user table through SQL.
type UserSQL struct {
	db *sql.DB
}

// IsIDExist checks whether a given user ID exists in user table.
func (u UserSQL) IsIDExist(id string) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;
`,
		table.User.ColumnID,
		table.User.TableName,
		table.User.ColumnID,
	)

	err := u.db.QueryRow(query, id).Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsEmailExist checks whether a given email exists in user table.
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

// GetUserByID finds an User in user table given user ID.
func (u UserSQL) GetUserByID(id string) (entity.User, error) {
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
		table.User.ColumnID,
	)

	row := u.db.QueryRow(query, id)

	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.LastSignedInAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == nil {
		user.CreatedAt = utc(user.CreatedAt)
		user.UpdatedAt = utc(user.UpdatedAt)
		user.LastSignedInAt = utc(user.LastSignedInAt)
		return user, nil
	}

	if err == sql.ErrNoRows {
		return entity.User{}, repository.ErrEntryNotFound("user account not found")
	}

	return user, nil
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

	if err == nil {
		user.CreatedAt = utc(user.CreatedAt)
		user.UpdatedAt = utc(user.UpdatedAt)
		user.LastSignedInAt = utc(user.LastSignedInAt)
		return user, nil
	}

	if err == sql.ErrNoRows {
		return entity.User{}, repository.ErrEntryNotFound("user account not found")
	}

	return entity.User{}, err
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
		return repository.ErrEntryNotFound(fmt.Sprintf("email %s does not exist", email))
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
