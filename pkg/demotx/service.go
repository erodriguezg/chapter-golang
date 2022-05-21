package demotx

import (
	"context"
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/person"
	"github.com/erodriguezg/chapter-golang/pkg/pet"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
)

type DemoTxService interface {
	ProcessWithTx(fail bool) error
	ProcessWithoutTx(fail bool) error
}

type defaultService struct {
	personService person.PersonService
	petService    pet.PetService
	txManager     transaction.TxManager
}

func NewService(
	personService person.PersonService,
	petService pet.PetService,
	txManager transaction.TxManager) DemoTxService {
	return &defaultService{personService, petService, txManager}
}

func (s *defaultService) ProcessWithTx(fail bool) error {
	ctx, err := s.txManager.Begin(context.TODO())
	if err != nil {
		return fmt.Errorf("error starting the transaction: \n%v \n", err)
	}
	return s.process(ctx, fail)
}

func (s *defaultService) ProcessWithoutTx(fail bool) error {

	// No transaction!

	return s.process(context.TODO(), fail)
}

// private

func (s *defaultService) process(ctx context.Context, fail bool) error {

	person := person.Person{
		Rut:       11111111,
		FirstName: "Pedrito",
		LastName:  "Fuenzalida",
		BirthDay:  nil,
		Active:    true,
	}

	fmt.Println("\n\n==================================")
	fmt.Println("INSERT PERSON")
	fmt.Println("==================================")

	newPerson, err := s.personService.Save(nil, &person)
	if err != nil {
		return fmt.Errorf("han error has occuried saving person: \n%v \n", err)
	}
	fmt.Printf("a person with id %d has been created \n", *newPerson.Id)

	err = s.addPetsToPerson(ctx, newPerson)
	if err != nil {
		return err
	}

	fmt.Println("\n\n==================================")
	fmt.Println("UPDATE PERSON")
	fmt.Println("==================================")

	newPerson.Rut = 22222222
	newPerson.Active = false

	updatedPerson, err := s.personService.Save(nil, newPerson)
	if err != nil {
		return fmt.Errorf("han error has occuried updating person: \n%v \n", err)
	}
	fmt.Printf("a person with id %d has been updated \n", *updatedPerson.Id)

	fmt.Println("\n\n==================================")
	fmt.Println("SEARCH ONE PERSON")
	fmt.Println("==================================")

	foundPerson, err := s.personService.FindByRut(nil, updatedPerson.Rut)
	if err != nil {
		return fmt.Errorf("han error has occuried search one person: \n%v \n", err)
	}
	fmt.Printf("a person is found %v \n", *foundPerson)

	err = s.deletePetsFromPerson(foundPerson)
	if err != nil {
		return err
	}

	fmt.Println("\n\n==================================")
	fmt.Println("DELETE PERSON")
	fmt.Println("==================================")

	err = s.personService.Delete(nil, updatedPerson)
	if err != nil {
		fmt.Errorf("han error has occuried deleting person: \n%v \n", err)
	}
	fmt.Printf("a person with id %d has been deleted \n", *updatedPerson.Id)

	return nil

}

func (s *defaultService) addPetsToPerson(ctx context.Context, person *person.Person) error {
	fmt.Println("\n\n==================================")
	fmt.Println("ADD PETS TO PERSON")
	fmt.Println("==================================")

	newPet := pet.Pet{
		OwnerRut: person.Rut,
		Name:     "Reptiliano",
		Race:     "perro",
	}

	otherPet := pet.Pet{
		OwnerRut: person.Rut,
		Name:     "Bestia",
		Race:     "gato",
	}

	pets := []pet.Pet{newPet, otherPet}

	for _, itPet := range pets {
		_, err := s.petService.Save(ctx, &itPet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *defaultService) deletePetsFromPerson(person *person.Person) error {

	fmt.Println("\n\n==================================")
	fmt.Println("FIND ALL PET FROM PERSON")
	fmt.Println("==================================")

	fmt.Println("\n\n==================================")
	fmt.Println("DELETE ALL PET FROM PERSON")
	fmt.Println("==================================")

	return nil
}
