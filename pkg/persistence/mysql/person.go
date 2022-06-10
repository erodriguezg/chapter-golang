package mysql

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
			  VALUES (?, ?, ?, ?, ?)`

	params := []interface{}{person.Rut, person.FirstName, person.LastName, person.BirthDay, person.Active}

	id, err := r.sqlTemplate.Exec(ctx, query, params)
	if err != nil {
		return nil, err
	}

	person.Id = &id
	return person, nil
}

func (r *personRepo) Update(ctx context.Context, person *person.Person) (*person.Person, error) {
	query :=
		`UPDATE persons SET
		rut=?, 
		first_name=?, 
		last_name=?, 
		birthday=?, 
		active=? 
		WHERE id=?`

	params := []interface{}{person.Rut, person.FirstName, person.LastName, person.BirthDay, person.Active, *person.Id}

	_, err := r.sqlTemplate.Exec(ctx, query, params)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (r *personRepo) Delete(ctx context.Context, person *person.Person) error {
	query := fmt.Sprintf("DELETE FROM persons WHERE id=?")

	params := []interface{}{*person.Id}

	_, err := r.sqlTemplate.Exec(ctx, query, params)
	return err
}

func (r *personRepo) GetAll(ctx context.Context) ([]person.Person, error) {
	query := fmt.Sprintf("SELECT %s FROM persons", personQueryProjection)

	return r.sqlTemplate.QueryForArray(ctx, query, nil, r.mapper)
}

func (r *personRepo) FindByRut(ctx context.Context, rut int) (*person.Person, error) {
	query := fmt.Sprintf("SELECT %s FROM persons WHERE rut=?", personQueryProjection)

	params := []interface{}{rut}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.mapper)
}

// privates

func (r *personRepo) mapper(rows *sql.Rows) (person.Person, error) {
	var p person.Person
	err := rows.Scan(&p.Id, &p.Rut, &p.FirstName, &p.LastName, &p.BirthDay, &p.Active)
	return p, err
}
