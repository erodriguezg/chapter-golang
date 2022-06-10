package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/demotx"
	"github.com/erodriguezg/chapter-golang/pkg/persistence/mysql"
	"github.com/erodriguezg/chapter-golang/pkg/persistence/postgresql"
	"github.com/erodriguezg/chapter-golang/pkg/person"
	"github.com/erodriguezg/chapter-golang/pkg/pet"
	"github.com/erodriguezg/chapter-golang/pkg/utils/sqlutils"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

const (
	databaseTypePostgres = "postgres"
	databaseTypeMysql    = "mysql"

	errMsgInvalidDatabaseType = "invalid database type"
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

func ConfigDemoSqlAll(databaseType string) {

	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("No .env file found!"))
	}

	// Database
	if databaseType == "postgres" {
		sqlDatabase = configPostgresDatabase()

	} else if databaseType == "mysql" {
		sqlDatabase = configMysqlDatabase()
	}

	sqltxManager = configSqlTxManager()

	// Repositories
	personRepository = configPersonRepository(databaseType)
	petRepository = configPetRepository(databaseType)

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

func configMysqlDatabase() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DATABASE_URL"))
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(3)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func configSqlTxManager() transaction.TxManager {
	panicIfAnyNil(sqlDatabase)
	return transaction.NewSqlTxManager(sqlDatabase)
}

func configPersonRepository(databaseType string) person.PersonRepository {

	panicIfAnyNil(sqlDatabase, sqltxManager)

	if databaseType == databaseTypePostgres {
		sqlTemplate := sqlutils.NewDatabaseSqlTemplate[person.Person](sqlDatabase, sqltxManager)
		return postgresql.NewPersonRepository(sqlTemplate)
	} else if databaseType == databaseTypeMysql {
		sqlTemplate := sqlutils.NewDatabaseSqlTemplate[person.Person](sqlDatabase, sqltxManager)
		return mysql.NewPersonRepository(sqlTemplate)
	}

	panic(fmt.Errorf(errMsgInvalidDatabaseType))
}

func configPetRepository(databaseType string) pet.PetRepository {
	panicIfAnyNil(sqlDatabase, sqltxManager)

	if databaseType == databaseTypePostgres {
		sqlTemplate := sqlutils.NewDatabaseSqlTemplate[pet.Pet](sqlDatabase, sqltxManager)
		return postgresql.NewPetRepository(sqlTemplate)
	} else if databaseType == databaseTypeMysql {
		sqlTemplate := sqlutils.NewDatabaseSqlTemplate[pet.Pet](sqlDatabase, sqltxManager)
		return mysql.NewPetRepository(sqlTemplate)
	}

	panic(fmt.Errorf(errMsgInvalidDatabaseType))
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
