package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/role"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.UserRole = (*UserRoleSQL)(nil)

type UserRoleSQL struct {
	db *sql.DB
}

func (u UserRoleSQL) GetRoles(user entity.User) ([]role.Role, error) {
	statement := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;
`,
		table.UserRole.ColumnRole,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
	)

	roles := []role.Role{}
	rows, err := u.db.Query(statement, user.ID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var r role.Role

		if err = rows.Scan(&r); err != nil {
			return nil, err
		}

		roles = append(roles, r)
	}

	return roles, nil
}

func (u UserRoleSQL) AddRole(user entity.User, r role.Role) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2);
`,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
		table.UserRole.ColumnRole,
	)

	_, err := u.db.Exec(statement, user.ID, r)
	return err
}

func (u UserRoleSQL) DeleteRole(user entity.User, r role.Role) error {
	statement := fmt.Sprintf(`
DELETE FROM "%s"
WHERE "%s"=$1 AND "%s"=$2;
`,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
		table.UserRole.ColumnRole,
	)

	_, err := u.db.Exec(statement, user.ID, r)
	return err
}

func NewUserRoleSQL(db *sql.DB) UserRoleSQL {
	return UserRoleSQL{db: db}
}
