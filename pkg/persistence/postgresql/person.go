package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/person"
	"github.com/erodriguezg/chapter-golang/pkg/utils/sqlutils"
)

const personQueryProjection = "id, rut, first_name, last_name, birthday, active"

type personRepo struct {
	sqlTemplate sqlutils.SqlTemplate[person.Person]
}

func NewPersonRepository(sqlTemplate sqlutils.SqlTemplate[person.Person]) person.PersonRepository {
	return &personRepo{sqlTemplate}
}

func (r *personRepo) Insert(ctx context.Context, person *person.Person) (*person.Person, error) {
	query := `INSERT INTO persons (rut, first_name, last_name, birthday, active) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING %s`

	query = fmt.Sprintf(query, personQueryProjection)

	params := []interface{}{person.Rut, person.FirstName, person.LastName, person.BirthDay, person.Active}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.fullMapper)

}

func (r *personRepo) Update(ctx context.Context, person *person.Person) (*person.Person, error) {
	query :=
		`UPDATE persons SET
		rut=$1, 
		first_name=$2, 
		last_name=$3, 
		birthday=$4, 
		active=$5 
		WHERE id=$6 
		RETURNING %s`

	query = fmt.Sprintf(query, personQueryProjection)

	params := []interface{}{person.Rut, person.FirstName, person.LastName, person.BirthDay, person.Active, *person.Id}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.fullMapper)
}

func (r *personRepo) Delete(ctx context.Context, person *person.Person) error {
	query := fmt.Sprintf("DELETE FROM persons WHERE id=$1 RETURNING id")

	params := []interface{}{*person.Id}

	_, err := r.sqlTemplate.QueryForOne(ctx, query, params, r.idOnlyMapper)
	return err
}

func (r *personRepo) GetAll(ctx context.Context) ([]person.Person, error) {
	query := fmt.Sprintf("SELECT %s FROM persons", personQueryProjection)

	return r.sqlTemplate.QueryForArray(ctx, query, nil, r.fullMapper)
}

func (r *personRepo) FindByRut(ctx context.Context, rut int) (*person.Person, error) {
	query := fmt.Sprintf("SELECT %s FROM persons WHERE rut=$1", personQueryProjection)

	params := []interface{}{rut}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.fullMapper)
}

// privates

func (r *personRepo) fullMapper(rows *sql.Rows) (person.Person, error) {
	var p person.Person
	err := rows.Scan(&p.Id, &p.Rut, &p.FirstName, &p.LastName, &p.BirthDay, &p.Active)
	return p, err
}

func (r *personRepo) idOnlyMapper(rows *sql.Rows) (person.Person, error) {
	var p person.Person
	err := rows.Scan(&p.Id)
	return p, err
}
