package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/persistence/postgresql"
	"github.com/erodriguezg/chapter-golang/pkg/person"
	"github.com/erodriguezg/chapter-golang/pkg/utils/sqlutils"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var (

	// Database
	sqlDatabase *sql.DB
	txManager   transaction.TxManager[*sql.Tx]

	// Business
	personRepository person.PersonRepository
	personService    person.PersonService
)

func ConfigDemoSqlAll() {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("No .env file found!"))
	}

	// Database
	sqlDatabase = configPostgresDatabase()
	txManager = configTxManager()

	// Business
	personRepository = configPersonRepository()
	personService = configDemoService()

}

func GetPersonService() person.PersonService {
	return personService
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

func configTxManager() transaction.TxManager[*sql.Tx] {
	panicIfAnyNil(sqlDatabase)
	return transaction.NewSqlTxManager(sqlDatabase)
}

func configPersonRepository() person.PersonRepository {
	panicIfAnyNil(sqlDatabase, txManager)
	sqlTemplate := sqlutils.NewDatabaseSqlTemplate[person.Person](sqlDatabase, txManager)
	return postgresql.NewPersonRepository(sqlTemplate)
}

func configDemoService() person.PersonService {
	panicIfAnyNil(personRepository)
	return person.NewPersonService(personRepository)
}
