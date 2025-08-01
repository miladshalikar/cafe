package aclservice

import "github.com/miladshalikar/cafe/pkg/richerror"

func (s Service) HasPermission(userID uint, permissionTitle string) (bool, error) {
	const op = "aclservice.HasPermission"

	permissionIDsList, err := s.repo.GetPermissionIDsByUserID(userID)
	if err != nil {
		return false, richerror.New(op).WithWarpError(err)
	}

	roleIDsList, rErr := s.repo.GetRoleIDsByUserID(userID)
	if rErr != nil {
		return false, richerror.New(op).WithWarpError(rErr)
	}

	var allPermissionIDsList []uint

	for _, roleID := range roleIDsList {
		permissionIDs, pErr := s.repo.GetPermissionIDsByRoleID(roleID)
		if pErr != nil {
			return false, richerror.New(op).WithWarpError(pErr)
		}

		allPermissionIDsList = append(allPermissionIDsList, permissionIDs...)
	}

	allPermissionIDsList = append(allPermissionIDsList, permissionIDsList...)

	PermissionID, iErr := s.repo.GetPermissionIDByTitle(permissionTitle)
	if iErr != nil {
		return false, richerror.New(op).WithWarpError(iErr)
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
