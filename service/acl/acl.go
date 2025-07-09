package aclservice

func (s Service) HasPermission(userID uint, permissionTitle string) (bool, error) {
	count, err := s.repo.HasPermission(userID, permissionTitle)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
