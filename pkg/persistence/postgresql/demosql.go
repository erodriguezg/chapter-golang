package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/demosql"
	"github.com/erodriguezg/chapter-golang/pkg/utils/sqlutils"
	"github.com/erodriguezg/chapter-golang/pkg/utils/transaction"
)

const personQueryProjection = "id, rut, first_name, last_name, birthday, active"

type personRepo struct {
	sqlTemplate sqlutils.SqlTemplate[demosql.Person]
}

func NewPersonRepository(db *sql.DB, txManager transaction.TxManager[*sql.Tx]) demosql.PersonRepository {
	sqlTemplate := sqlutils.NewSqlTemplate[demosql.Person](db, txManager)
	return &personRepo{sqlTemplate}
}

func (r *personRepo) Insert(ctx context.Context, person *demosql.Person) (*demosql.Person, error) {
	query := `INSERT INTO persons (rut, first_name, last_name, birthday, active) 
			  VALUES ($1, $2, $3, $4, $5)`

	params := []interface{}{person.Rut, person.FirstName, person.LastName, person.BirthDay, person.Active}

	id, err := r.sqlTemplate.Exec(ctx, query, params)
	if err != nil {
		return nil, err
	}

	person.Id = &id
	return person, nil
}

func (r *personRepo) Update(ctx context.Context, person *demosql.Person) (*demosql.Person, error) {
	query :=
		`UPDATE persons SET
	rut=$1, 
	first_name=$2, 
	last_name=$3, 
	birthday=$4, 
	active=$5 
	WHERE id=$6`

	params := []interface{}{person.Rut, person.FirstName, person.LastName, person.BirthDay, person.Active}

	_, err := r.sqlTemplate.Exec(ctx, query, params)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (r *personRepo) Delete(ctx context.Context, person *demosql.Person) error {
	query := fmt.Sprintf("DELETE FROM persons WHERE id=$1")

	params := []interface{}{*person.Id}

	_, err := r.sqlTemplate.Exec(ctx, query, params)
	return err
}

func (r *personRepo) GetAll(ctx context.Context) ([]demosql.Person, error) {
	query := fmt.Sprintf("SELECT %s FROM persons", personQueryProjection)

	return r.sqlTemplate.QueryForArray(ctx, query, nil, r.mapper)
}

func (r *personRepo) FindByRut(ctx context.Context, rut int) (*demosql.Person, error) {
	query := fmt.Sprintf("SELECT %s FROM persons WHERE rut=$1", personQueryProjection)

	params := []interface{}{rut}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.mapper)
}

// privates

func (r *personRepo) mapper(rows *sql.Rows) (demosql.Person, error) {
	var p demosql.Person
	err := rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.BirthDay, &p.Active)
	return p, err
}