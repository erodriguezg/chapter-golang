package pet

import "context"

type PetRepository interface {
	Insert(ctx context.Context, pet *Pet) (*Pet, error)
	Update(ctx context.Context, pet *Pet) (*Pet, error)
	Delete(ctx context.Context, pet *Pet) error
	FindByOwnerRut(ctx context.Context, ownerRut int) ([]Pet, error)
	FindById(ctx context.Context, id int64) (*Pet, error)
}
