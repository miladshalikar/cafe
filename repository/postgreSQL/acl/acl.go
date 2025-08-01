package aclpostgresql

import (
	errmsg "github.com/miladshalikar/cafe/pkg/err_msg"
	"github.com/miladshalikar/cafe/pkg/richerror"
)

func (d *DB) GetPermissionIDsByUserID(userID uint) ([]uint, error) {
	const op = "aclpostgresql.GetPermissionIDsByUserID"

	query := `
		SELECT permission_id
		FROM permission_user
		WHERE user_id = $1 AND deleted_at IS NULL
		`

	rows, rErr := d.conn.Query(query, userID)
	if rErr != nil {
		return nil, richerror.New(op).
			WithWarpError(rErr).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var permissionIDs []uint
	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return nil, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).
				WithKind(richerror.KindUnexpected)
		}
		permissionIDs = append(permissionIDs, id)
	}
	if wErr := rows.Err(); wErr != nil {
		return nil, richerror.New(op).
			WithWarpError(wErr).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return permissionIDs, nil
}

func (d *DB) GetRoleIDsByUserID(userID uint) ([]uint, error) {
	const op = "aclpostgresql.GetRoleIDsByUserID"

	query := `
		SELECT role_id
		FROM role_user
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	rows, rErr := d.conn.Query(query, userID)
	if rErr != nil {
		return nil, richerror.New(op).
			WithWarpError(rErr).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var roleIDs []uint
	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return nil, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).
				WithKind(richerror.KindUnexpected)
		}
		roleIDs = append(roleIDs, id)
	}
	if wErr := rows.Err(); wErr != nil {
		return nil, richerror.New(op).
			WithWarpError(wErr).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return roleIDs, nil
}

func (d *DB) GetPermissionIDsByRoleID(roleID uint) ([]uint, error) {
	const op = "aclpostgresql.GetPermissionIDsByRoleID"

	query := `
		SELECT permission_id
		FROM permission_role
		WHERE role_id = $1 AND deleted_at IS NULL
	`

	rows, rErr := d.conn.Query(query, roleID)
	if rErr != nil {
		return nil, richerror.New(op).
			WithWarpError(rErr).
			WithMessage(errmsg.ErrorMsgSomethingWentWrong).
			WithKind(richerror.KindUnexpected)
	}
	defer rows.Close()

	var permissionIDs []uint
	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return nil, richerror.New(op).
				WithWarpError(err).
				WithMessage(errmsg.ErrorMsgCantScanQueryResult).
				WithKind(richerror.KindUnexpected)
		}
		permissionIDs = append(permissionIDs, id)
	}
	if wErr := rows.Err(); wErr != nil {
		return nil, richerror.New(op).
			WithWarpError(wErr).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return permissionIDs, nil
}

func (d *DB) GetPermissionIDByTitle(title string) (uint, error) {
	const op = "aclpostgresql.GetPermissionIDByTitle"

	query := `
		SELECT id
		FROM permissions
		WHERE title = $1 AND deleted_at IS NULL
	`

	var id uint
	err := d.conn.QueryRow(query, title).Scan(&id)
	if err != nil {
		return 0, richerror.New(op).
			WithWarpError(err).
			WithMessage(errmsg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}
	return id, nil
}
