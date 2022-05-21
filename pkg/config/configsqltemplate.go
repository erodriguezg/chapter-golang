package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/demotx"
	"github.com/erodriguezg/chapter-golang/pkg/persistence/postgresql"
	"github.com/erodriguezg/chapter-golang/pkg/person"
	"github.com/erodriguezg/chapter-golang/pkg/pet"
	"github.com/erodriguezg/chapter-golang/pkg/utils/sqlutils"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var (

	// Database
	sqlDatabase  *sql.DB
	sqltxManager transaction.TxManager

	// Repositories
	personRepository person.PersonRepository
	petRepository    pet.PetRepository

	// Services
	personService person.PersonService
	petService    pet.PetService
	demoTxService demotx.DemoTxService
)

func ConfigDemoSqlAll() {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("No .env file found!"))
	}

	// Database
	sqlDatabase = configPostgresDatabase()
	sqltxManager = configSqlTxManager()

	// Repositories
	personRepository = configPersonRepository()
	petRepository = configPetRepository()

	// Services
	personService = configDemoService()
	petService = configPetService()
	demoTxService = configDemoTxService()

}

func GetDemoTxService() demotx.DemoTxService {
	return demoTxService
}

func CloseDemoSqlAll() {
	err := sqlDatabase.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error closing postgresql DB: %v\n", err)
	}
}

// privates

func configPostgresDatabase() *sql.DB {
	db, err := sql.Open("pgx", os.Getenv("POSTGRESQL_DATABASE_URL"))
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(3)
	if err != nil {
		panic(err)
	}
	return db
}

func configSqlTxManager() transaction.TxManager {
	panicIfAnyNil(sqlDatabase)
	return transaction.NewSqlTxManager(sqlDatabase)
}

func configPersonRepository() person.PersonRepository {
	panicIfAnyNil(sqlDatabase, sqltxManager)
	sqlTemplate := sqlutils.NewDatabaseSqlTemplate[person.Person](sqlDatabase, sqltxManager)
	return postgresql.NewPersonRepository(sqlTemplate)
}

func configPetRepository() pet.PetRepository {
	panicIfAnyNil(sqlDatabase, sqltxManager)
	sqlTemplate := sqlutils.NewDatabaseSqlTemplate[pet.Pet](sqlDatabase, sqltxManager)
	return postgresql.NewPetRepository(sqlTemplate)
}

func configDemoService() person.PersonService {
	panicIfAnyNil(personRepository)
	return person.NewPersonService(personRepository)
}

func configPetService() pet.PetService {
	panicIfAnyNil(petRepository)
	return pet.NewService(petRepository)
}

func configDemoTxService() demotx.DemoTxService {
	panicIfAnyNil(personService, petService, sqltxManager)
	return demotx.NewService(personService, petService, sqltxManager)
}
