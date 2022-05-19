package demosql

import (
	"context"
	"fmt"
)

type PersonService interface {
	Save(ctx context.Context, person *Person) (*Person, error)
	Delete(ctx context.Context, person *Person) error
	GetAll(ctx context.Context) ([]Person, error)
	FindByRut(ctx context.Context, rut int) (*Person, error)
}

type personServiceImpl struct {
	repository PersonRepository
}

func NewPersonService(repository PersonRepository) PersonService {
	return &personServiceImpl{repository}
}

func (s *personServiceImpl) Save(ctx context.Context, person *Person) (*Person, error) {
	if person == nil {
		return nil, fmt.Errorf("the input person is nil")
	}
	if person.Id == nil {
		return s.repository.Insert(ctx, person)
	} else {
		return s.repository.Update(ctx, person)
	}
}

func (s *personServiceImpl) GetAll(ctx context.Context) ([]Person, error) {
	return s.repository.GetAll(ctx)
}

func (s *personServiceImpl) FindByRut(ctx context.Context, rut int) (*Person, error) {
	return s.repository.FindByRut(ctx, rut)
}

func (s *personServiceImpl) Delete(ctx context.Context, person *Person) error {
	return s.repository.Delete(ctx, person)
}
