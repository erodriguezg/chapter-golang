package main

import (
	"fmt"
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/config"
	"github.com/erodriguezg/chapter-golang/pkg/demosql"
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

	personService := config.GetPersonService()

	person := demosql.Person{
		Rut:       11111111,
		FirstName: "Pedrito",
		LastName:  "Fuenzalida",
		BirthDay:  nil,
		Active:    true,
	}

	fmt.Println("\n\n==================================")
	fmt.Println("INSERT")
	fmt.Println("==================================")

	newPerson, err := personService.Save(nil, &person)
	if err != nil {
		fmt.Fprintf(os.Stderr, "han error has occuried saving person: \n%v \n", err)
		return
	}

	fmt.Printf("a person with id %d has been created \n", *newPerson.Id)

	fmt.Println("\n\n==================================")
	fmt.Println("UPDATE")
	fmt.Println("==================================")

	newPerson.Rut = 22222222
	newPerson.Active = false

	updatedPerson, err := personService.Save(nil, newPerson)
	if err != nil {
		fmt.Fprintf(os.Stderr, "han error has occuried updating person: \n%v \n", err)
		return
	}
	fmt.Printf("a person with id %d has been updated \n", *updatedPerson.Id)

	fmt.Println("\n\n==================================")
	fmt.Println("SEARCH ONE")
	fmt.Println("==================================")

	foundPerson, err := personService.FindByRut(nil, updatedPerson.Rut)
	if err != nil {
		fmt.Fprintf(os.Stderr, "han error has occuried search one person: \n%v \n", err)
		return
	}
	fmt.Printf("a person is found %v \n", *foundPerson)

	fmt.Println("\n\n==================================")
	fmt.Println("DELETE")
	fmt.Println("==================================")

	err = personService.Delete(nil, updatedPerson)
	if err != nil {
		fmt.Fprintf(os.Stderr, "han error has occuried deleting person: \n%v \n", err)
		return
	}
	fmt.Printf("a person with id %d has been deleted \n", *updatedPerson.Id)

}
