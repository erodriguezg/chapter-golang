package demosql

import (
	"context"
	"fmt"
)

type PersonService interface {
	Save(ctx context.Context, person *Person) (*Person, error)
	GetAll(ctx context.Context) ([]Person, error)
	Delete(ctx context.Context, person *Person) error
}

type defaulService struct {
	repository PersonRepository
}

func NewPersonService(repository PersonRepository) PersonService {
	return &defaulService{repository}
}

func (s *defaulService) Save(ctx context.Context, person *Person) (*Person, error) {
	if person == nil {
		return nil, fmt.Errorf("the input person is nil")
	}
	if person.Id == nil {
		return s.repository.Insert(ctx, person)
	} else {
		return s.repository.Update(ctx, person)
	}
}

func (s *defaulService) GetAll(ctx context.Context) ([]Person, error) {
	return s.repository.GetAll(ctx)
}

func (s *defaulService) Delete(ctx context.Context, person *Person) error {
	return s.repository.Delete(ctx, person)
}
