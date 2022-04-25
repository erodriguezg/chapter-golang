package foo

import "fmt"

type FooService interface {
	DoFoo() string
}

type defaultService struct {
	name string
}

func NewService(name string) FooService {
	return &defaultService{name}
}

func (s *defaultService) DoFoo() string {
	return fmt.Sprintf("foo %s!", s.name)
}
