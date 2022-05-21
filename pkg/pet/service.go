package pet

import (
	"context"
	"fmt"
)

type PetService interface {
	Save(ctx context.Context, pet *Pet) (*Pet, error)
	Delete(ctx context.Context, pet *Pet) error
	FindByOwnerRut(ctx context.Context, ownerRut int) ([]Pet, error)
	FindById(ctx context.Context, id int64) (*Pet, error)
}

type petServiceImpl struct {
	repository PetRepository
}

func NewService(repository PetRepository) PetService {
	return &petServiceImpl{repository}
}

func (s *petServiceImpl) Save(ctx context.Context, pet *Pet) (*Pet, error) {
	if pet == nil {
		return nil, fmt.Errorf("the pet is nil")
	}
	if pet.Id == nil {
		return s.repository.Insert(ctx, pet)
	} else {
		return s.repository.Update(ctx, pet)
	}
}

func (s *petServiceImpl) Delete(ctx context.Context, pet *Pet) error {
	return s.repository.Delete(ctx, pet)
}

func (s *petServiceImpl) FindByOwnerRut(ctx context.Context, ownerRut int) ([]Pet, error) {
	return s.repository.FindByOwnerRut(ctx, ownerRut)
}

func (s *petServiceImpl) FindById(ctx context.Context, id int64) (*Pet, error) {
	return s.repository.FindById(ctx, id)
}
