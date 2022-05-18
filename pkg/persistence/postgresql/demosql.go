package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/demosql"
	"github.com/erodriguezg/chapter-golang/pkg/utils/sqlutils"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
)

const personQueryProjection = "ID, FIRST_NAME, LAST_NAME, BIRTHDAY, ACTIVE"

type personRepo struct {
	sqlTemplate sqlutils.SqlTemplate[demosql.Person]
}

func NewPersonRepository(db *sql.DB, txManager transaction.TxManager[*sql.Tx]) demosql.PersonRepository {
	sqlTemplate := sqlutils.NewSqlTemplate[demosql.Person](db, txManager)
	return &personRepo{sqlTemplate}
}

func (r *personRepo) Insert(ctx context.Context, person *demosql.Person) (*demosql.Person, error) {

	query := `INSERT INTO PERSON (FIRST_NAME, LAST_NAME, BIRTHDAY, ACTIVE) 
			  VALUES ($1, $2, $3, $4)`

	params := []interface{}{person.FirstName, person.LastName, person.BirthDay, person.Active}

	id, err := r.sqlTemplate.Exec(ctx, query, params)
	if err != nil {
		return nil, err
	}

	person.Id = &id
	return person, nil
}

func (r *personRepo) Update(ctx context.Context, person *demosql.Person) (*demosql.Person, error) {
	return nil, nil
}

func (r *personRepo) GetAll(ctx context.Context) ([]demosql.Person, error) {
	query := fmt.Sprintf("SELECT %s FROM PERSON", personQueryProjection)
	args := []interface{}{}
	return r.sqlTemplate.QueryForArray(nil, query, args, r.mapper)
}

func (r *personRepo) Delete(ctx context.Context, person *demosql.Person) error {
	return nil
}

// privates

func (r *personRepo) mapper(rows *sql.Rows) (demosql.Person, error) {
	var p demosql.Person
	err := rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.BirthDay, &p.Active)
	return p, err
}
