package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/demosql"
	"github.com/erodriguezg/chapter-golang/pkg/persistence/postgresql"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var (

	// Database
	postgresqlDB *sql.DB
	txManager    transaction.TxManager[*sql.Tx]

	// Business
	personRepository demosql.PersonRepository
	personService    demosql.PersonService
)

func ConfigDemoSqlAll() {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("No .env file found!"))
	}

	// Database
	postgresqlDB = configPostgresDatabase()
	txManager = configTxManager()

	// Business
	personRepository = configPersonRepository()
	personService = configDemoService()

}

func GetPersonService() demosql.PersonService {
	return personService
}

func CloseDemoSqlAll() {
	err := postgresqlDB.Close()
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
	return transaction.NewTxManager[*sql.Tx](postgresqlDB)
}

func configPersonRepository() demosql.PersonRepository {
	return postgresql.NewPersonRepository(postgresqlDB, txManager)
}

func configDemoService() demosql.PersonService {
	return demosql.NewPersonService(personRepository)
}
