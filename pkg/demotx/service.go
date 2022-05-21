package demotx

import (
	"context"
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/person"
	"github.com/erodriguezg/chapter-golang/pkg/pet"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
)

type DemoTxService interface {
	ProcessWithTx(delete bool, fail bool) error
	ProcessWithoutTx(delete bool, fail bool) error
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

func (s *defaultService) ProcessWithTx(delete bool, fail bool) error {

	// With transaction!

	fmt.Println("\n\n==================================")
	fmt.Println("STARTING TX")
	fmt.Println("==================================")

	ctx, err := s.txManager.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("error starting the transaction: \n%v \n", err)
	}

	err = s.process(ctx, delete, fail)
	if err != nil {

		fmt.Println("\n\n==================================")
		fmt.Println("ROLLBACK TX")
		fmt.Println("==================================")

		if errRoll := s.txManager.Rollback(ctx); errRoll != nil {
			return errRoll
		}
		return err
	}

	fmt.Println("\n\n==================================")
	fmt.Println("COMMIT TX")
	fmt.Println("==================================")

	if errCommit := s.txManager.Commit(ctx); errCommit != nil {
		return errCommit
	}

	return nil
}

func (s *defaultService) ProcessWithoutTx(delete bool, fail bool) error {

	// No transaction!

	return s.process(context.Background(), delete, fail)
}

// private

func (s *defaultService) process(ctx context.Context, delete bool, fail bool) error {

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

	newPerson, err := s.personService.Save(ctx, &person)
	if err != nil {
		return fmt.Errorf("han error has occuried saving person: \n%v \n", err)
	}
	fmt.Printf("a person with id %d has been created \n", *newPerson.Id)

	fmt.Println("\n\n==================================")
	fmt.Println("UPDATE PERSON")
	fmt.Println("==================================")

	newPerson.Rut = 22222222
	newPerson.Active = false

	updatedPerson, err := s.personService.Save(ctx, newPerson)
	if err != nil {
		return fmt.Errorf("han error has occuried updating person: \n%v \n", err)
	}
	fmt.Printf("a person with id %d has been updated \n", *updatedPerson.Id)

	fmt.Println("\n\n==================================")
	fmt.Println("SEARCH ONE PERSON")
	fmt.Println("==================================")

	foundPerson, err := s.personService.FindByRut(ctx, updatedPerson.Rut)
	if err != nil {
		return fmt.Errorf("han error has occuried search one person: \n%v \n", err)
	}
	fmt.Printf("a person is found %v \n", *foundPerson)

	err = s.addPetsToPerson(ctx, foundPerson)
	if err != nil {
		return err
	}

	if fail == true {
		return fmt.Errorf("forcing the error!")
	}

	if delete == false {
		return nil
	}

	err = s.deletePetsFromPerson(ctx, foundPerson)
	if err != nil {
		return err
	}

	fmt.Println("\n\n==================================")
	fmt.Println("DELETE PERSON")
	fmt.Println("==================================")

	err = s.personService.Delete(ctx, updatedPerson)
	if err != nil {
		return fmt.Errorf("han error has occuried deleting person: \n%v \n", err)
	}
	fmt.Printf("a person with id %d has been deleted \n", *updatedPerson.Id)

	return nil

}

func (s *defaultService) addPetsToPerson(ctx context.Context, person *person.Person) error {
	fmt.Println("\n\n==================================")
	fmt.Println("ADD PETS TO PERSON")
	fmt.Println("==================================")

	pets := []pet.Pet{
		{
			OwnerRut: person.Rut,
			Name:     "Reptiliano",
			Race:     "perro",
		},
		{
			OwnerRut: person.Rut,
			Name:     "Bestia",
			Race:     "gato",
		},
	}

	for _, itPet := range pets {
		_, err := s.petService.Save(ctx, &itPet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *defaultService) deletePetsFromPerson(ctx context.Context, person *person.Person) error {

	fmt.Println("\n\n==================================")
	fmt.Println("FIND ALL PET FROM PERSON")
	fmt.Println("==================================")

	pets, err := s.petService.FindByOwnerRut(ctx, person.Rut)
	if err != nil {
		return nil
	}

	fmt.Println("\n\n==================================")
	fmt.Println("DELETE ALL PET FROM PERSON")
	fmt.Println("==================================")

	for _, itPet := range pets {
		if errDel := s.petService.Delete(ctx, &itPet); errDel != nil {
			return errDel
		}
	}
	return nil
}
