package db

import (
	"database/sql"
	"fmt"
	"sort"

	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authorizer/role"
	"github.com/short-d/short/app/usecase/repository"
)

var _ repository.UserRole = (*UserRoleSQL)(nil)

type UserRoleSQL struct {
	db *sql.DB
}

func (u UserRoleSQL) GetUserRoles(user entity.User) ([]role.Role, error) {
	query := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;
`,
		table.UserRole.ColumnRoles,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
	)

	var encoded string

	err := u.db.QueryRow(query, user.ID).Scan(&encoded)

	if err == sql.ErrNoRows {
		return []role.Role{}, nil
	}

	if err != nil {
		return nil, err
	}

	return decodeRoles(encoded), nil
}

func (u UserRoleSQL) AddRole(user entity.User, r role.Role) error {
	return u.saveRoles(user, []role.Role{r}, nil)
}

func (u UserRoleSQL) DeleteRole(user entity.User, r role.Role) error {
	return u.saveRoles(user, nil, []role.Role{r})
}

func NewUserRoleSQL(db *sql.DB) UserRoleSQL {
	return UserRoleSQL{db: db}
}

func (u UserRoleSQL) saveRoles(user entity.User, toAdd, toDelete []role.Role) error {
	tx, err := u.db.Begin()

	if err != nil {
		return nil
	}

	defer tx.Rollback()

	query := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1 FOR UPDATE;
`,
		table.UserRole.ColumnRoles,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
	)

	encoded := ""
	err = tx.QueryRow(query, user.ID).Scan(&encoded)
	isNew := err == sql.ErrNoRows

	if err != nil && !isNew {
		return err
	}

	roles := modifyRoles(decodeRoles(encoded), toAdd, toDelete)

	if len(roles) > 0 {
		if isNew {
			if err := insertExec(tx, user.ID, roles); err != nil {
				return err
			}
		} else {
			if err := updateExec(tx, user.ID, roles); err != nil {
				return err
			}
		}
	} else {
		if err := deleteExec(tx, user.ID); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertExec(tx *sql.Tx, userID string, roles []role.Role) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2);
`,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
		table.UserRole.ColumnRoles,
	)

	if _, err := tx.Exec(statement, userID, encodeRoles(roles)); err != nil {
		return err
	}

	return nil
}

func updateExec(tx *sql.Tx, userID string, roles []role.Role) error {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1
WHERE "%s"=$2;
`,
		table.UserRole.TableName,
		table.UserRole.ColumnRoles,
		table.UserRole.ColumnUserID,
	)

	if _, err := tx.Exec(statement, encodeRoles(roles), userID); err != nil {
		return err
	}

	return nil
}

func deleteExec(tx *sql.Tx, userID string) error {
	statement := fmt.Sprintf(`
DELETE FROM "%s"
WHERE "%s"=$1;
`,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
	)

	if _, err := tx.Exec(statement, userID); err != nil {
		return err
	}

	return nil
}

func modifyRoles(roles []role.Role, toAdd, toDelete []role.Role) []role.Role {
	for _, add := range toAdd {
		roles = append(roles, add)
	}

	for _, delete := range toDelete {
		for i, r := range roles {
			if r == delete {
				roles[i], roles[len(roles)-1] = roles[len(roles)-1], roles[i]
				roles = roles[:len(roles)-1]
				break
			}
		}
	}

	return roles
}

func encodeRoles(roles []role.Role) string {
	if len(roles) == 0 {
		return ""
	}

	sort.Slice(roles, func(i, j int) bool {
		return roles[i] > roles[j]
	})

	var encoded string

	for r := roles[0]; r >= 0; r-- {
		if len(roles) > 0 && r == roles[0] {
			encoded += "1"
			roles = roles[1:]
		} else {
			encoded += "0"
		}
	}

	return encoded
}

func decodeRoles(encoded string) []role.Role {
	roles := []role.Role{}
	n := len(encoded)

	for i := n - 1; i >= 0; i-- {
		if encoded[i] == '1' {
			roles = append(roles, role.Role(n-1-i))
		}
	}

	return roles
}
