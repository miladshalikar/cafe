package aclservice

type Service struct {
	repo Repository
}

type Repository interface {
	HasPermission(userID uint, permissionTitle string) (int, error)
}

func New(r Repository) Service {
	return Service{repo: r}
}
