package aclservice

type Service struct {
	repo Repository
}

type Repository interface {
	GetPermissionIDsByUserID(userID uint) ([]uint, error)
	GetRoleIDsByUserID(userID uint) ([]uint, error)
	GetPermissionIDsByRoleID(roleID uint) ([]uint, error)
	GetPermissionIDByTitle(title string) (uint, error)
}

func New(r Repository) Service {
	return Service{repo: r}
}
