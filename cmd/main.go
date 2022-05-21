package main

import (
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/config"
	"github.com/erodriguezg/chapter-golang/pkg/problems"
)

func main() {

	switch os.Args[1] {

	case "problem-float32":
		problems.Float32ExampleProblem()

	case "problem-config":
		problems.Config()

	case "sqltemplate":
		mainSqlTemplate()
	}

}

func mainSqlTemplate() {

	defer config.CloseDemoSqlAll()

	config.ConfigDemoSqlAll()

}
