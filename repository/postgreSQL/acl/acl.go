package aclpostgresql

func (d *DB) GetPermissionIDsByUserID(userID uint) ([]uint, error) {

	query := `
		SELECT permission_id
		FROM permission_user
		WHERE user_id = $1 AND deleted_at IS NULL
		`

	rows, rErr := d.conn.Query(query, userID)
	if rErr != nil {
		return nil, rErr
	}
	defer rows.Close()

	var permissionIDs []uint
	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		permissionIDs = append(permissionIDs, id)
	}
	return permissionIDs, nil
}

func (d *DB) GetRoleIDsByUserID(userID uint) ([]uint, error) {
	query := `
		SELECT role_id
		FROM role_user
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	rows, rErr := d.conn.Query(query, userID)
	if rErr != nil {
		return nil, rErr
	}
	defer rows.Close()

	var roleIDs []uint
	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		roleIDs = append(roleIDs, id)
	}
	return roleIDs, nil
}

func (d *DB) GetPermissionIDsByRoleID(roleID uint) ([]uint, error) {
	query := `
		SELECT permission_id
		FROM permission_role
		WHERE role_id = $1 AND deleted_at IS NULL
	`

	rows, rErr := d.conn.Query(query, roleID)
	if rErr != nil {
		return nil, rErr
	}
	defer rows.Close()

	var permissionIDs []uint
	for rows.Next() {
		var id uint
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		permissionIDs = append(permissionIDs, id)
	}
	return permissionIDs, nil
}

func (d *DB) GetPermissionIDByTitle(title string) (uint, error) {
	query := `
		SELECT id
		FROM permissions
		WHERE title = $1 AND deleted_at IS NULL
	`

	var id uint
	err := d.conn.QueryRow(query, title).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
