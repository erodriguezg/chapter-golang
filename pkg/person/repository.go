package person

import "context"

type PersonRepository interface {
	Insert(ctx context.Context, person *Person) (*Person, error)
	Update(ctx context.Context, person *Person) (*Person, error)
	Delete(ctx context.Context, person *Person) error
	GetAll(ctx context.Context) ([]Person, error)
	FindByRut(ctx context.Context, rut int) (*Person, error)
}
