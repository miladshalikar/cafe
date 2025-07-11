package aclservice

import "fmt"

func (s Service) HasPermission(userID uint, permissionTitle string) (bool, error) {

	permissionIDsList, err := s.repo.GetPermissionIDsByUserID(userID)
	if err != nil {
		return false, err
	}

	fmt.Println("1", permissionIDsList)

	roleIDsList, err := s.repo.GetRoleIDsByUserID(userID)
	if err != nil {
		return false, err
	}

	fmt.Println("2", roleIDsList)

	var allPermissionIDsList []uint

	for _, roleID := range roleIDsList {
		permissionIDs, pErr := s.repo.GetPermissionIDsByRoleID(roleID)
		if pErr != nil {
			return false, pErr
		}

		fmt.Println("33", permissionIDs)

		allPermissionIDsList = append(allPermissionIDsList, permissionIDs...)
	}

	allPermissionIDsList = append(allPermissionIDsList, permissionIDsList...)

	fmt.Println("4", allPermissionIDsList)

	PermissionID, iDrr := s.repo.GetPermissionIDByTitle(permissionTitle)
	if iDrr != nil {
		return false, iDrr
	}

	fmt.Println("5", PermissionID)

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
