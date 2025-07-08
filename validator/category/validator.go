package categoryvalidator

type Validator struct {
	repo Repository
}

type Repository interface {
	//CheckCategoryIsExistByID(ctx context.Context, id uint) (bool, error)
}

func New(r Repository) Validator {
	return Validator{repo: r}
}
