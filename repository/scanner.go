package postgresql

type Scanner interface {
	Scan(dest ...any) error
}
