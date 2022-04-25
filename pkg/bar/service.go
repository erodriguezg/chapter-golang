package bar

import (
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/foo"
)

type BarService interface {
	DoBar()
}

type defaultService struct {
	fooService foo.FooService
}

func NewService(fooService foo.FooService) BarService {
	return &defaultService{fooService}
}

func (s *defaultService) DoBar() {
	fmt.Printf("bar %s! \n", s.fooService.DoFoo())
}
