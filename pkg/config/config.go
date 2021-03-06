package config

import (
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/bar"
	"github.com/erodriguezg/chapter-golang/pkg/foo"
)

var (
	fooService1 foo.FooService
	fooService2 foo.FooService

	barService1 bar.BarService
	barService2 bar.BarService
)

func ConfigAll() {

	fooService1 = configFooService1()
	fooService2 = configFooService2()

	barService1 = configBarService1()
	barService2 = configBarService2()

}

// private config methods

func configFooService1() foo.FooService {
	return foo.NewService("Foo1!")
}

func configFooService2() foo.FooService {
	return foo.NewService("Foo2!")
}

func configBarService1() bar.BarService {
	panicIfAnyNil(fooService1)
	return bar.NewService(fooService1)
}

func configBarService2() bar.BarService {
	panicIfAnyNil(fooService2)
	return bar.NewService(fooService2)
}

// GET

func GetBarService1() bar.BarService {
	return barService1
}

func GetBarService2() bar.BarService {
	return barService2
}

// utils

func panicIfAnyNil(nilables ...interface{}) {
	if len(nilables) == 0 {
		return
	}
	for _, nilable := range nilables {
		if nilable == nil {
			panic(fmt.Errorf("an config element is nil!"))
		}
	}
}
