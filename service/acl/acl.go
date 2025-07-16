package aclservice

func (s Service) HasPermission(userID uint, permissionTitle string) (bool, error) {

	permissionIDsList, err := s.repo.GetPermissionIDsByUserID(userID)
	if err != nil {
		return false, err
	}

	roleIDsList, err := s.repo.GetRoleIDsByUserID(userID)
	if err != nil {
		return false, err
	}

	var allPermissionIDsList []uint

	for _, roleID := range roleIDsList {
		permissionIDs, pErr := s.repo.GetPermissionIDsByRoleID(roleID)
		if pErr != nil {
			return false, pErr
		}

		allPermissionIDsList = append(allPermissionIDsList, permissionIDs...)
	}

	allPermissionIDsList = append(allPermissionIDsList, permissionIDsList...)

	PermissionID, iDrr := s.repo.GetPermissionIDByTitle(permissionTitle)
	if iDrr != nil {
		return false, iDrr
	}

	if hasID(allPermissionIDsList, PermissionID) {
		return true, nil
	} else {
		return false, nil
	}

}

func hasID(list []uint, target uint) bool {

	for _, id := range list {
		if id == target {
			return true
		}
	}
	return false

}
