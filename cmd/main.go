package main

import (
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/problems"
)

func main() {

	switch os.Args[1] {
	case "problem-float32":
		problems.Float32ExampleProblem()

	case "problem-config-1":
		problems.ConfigWithoutContainer()
	}

}
