package aclpostgresql

func (d *DB) HasPermission(userID uint, permissionTitle string) (int, error) {

	query := `SELECT COUNT(1)
		FROM permissions p
		WHERE p.title = $1 AND p.id IN (
			SELECT pu.permission_id FROM permission_user pu WHERE pu.user_id = $2
			UNION
			SELECT pr.permission_id
			FROM permission_role pr
			JOIN role_user ru ON ru.role_id = pr.role_id
			WHERE ru.user_id = $2
		)`

	var count int
	err := d.conn.QueryRow(query, permissionTitle, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
