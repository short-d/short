package repo

import (
	"database/sql"
	"fmt"
	"short/app/adapter/repo/table"
	"short/app/entity"
	"short/app/usecase/repo"
)

type UserSql struct {
	db *sql.DB
}

func (u UserSql) IsEmailExist(email string) (bool, error) {
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

func (u UserSql) GetByEmail(email string) (entity.User, error) {
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
	err := row.Scan(&user.Email, &user.Name, &user.LastSignedInAt, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserSql) Create(user entity.User) error {
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

func NewUserSql(db *sql.DB) repo.User {
	return &UserSql{
		db: db,
	}
}
