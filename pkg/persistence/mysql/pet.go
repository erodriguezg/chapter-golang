package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/erodriguezg/chapter-golang/pkg/pet"
	"github.com/erodriguezg/chapter-golang/pkg/utils/sqlutils"
)

const petQueryProjection = "id, owner_rut, name, race"
const petQueryProjectionWithAlias = "pet.id, pet.owner_rut, pet.name, pet.race"

type petRepo struct {
	sqlTemplate sqlutils.SqlTemplate[pet.Pet]
}

func NewPetRepository(sqlTemplate sqlutils.SqlTemplate[pet.Pet]) pet.PetRepository {
	return &petRepo{sqlTemplate}
}

func (r *petRepo) Insert(ctx context.Context, pet *pet.Pet) (*pet.Pet, error) {
	query :=
		`INSERT INTO pets (owner_rut, name, race) 
	VALUES (?, ?, ?)`

	params := []interface{}{pet.OwnerRut, pet.Name, pet.Race}

	id, err := r.sqlTemplate.Exec(ctx, query, params)
	if err != nil {
		return nil, err
	}

	pet.Id = &id
	return pet, nil
}

func (r *petRepo) Update(ctx context.Context, pet *pet.Pet) (*pet.Pet, error) {
	query :=
		`UPDATE pets SET
		owner_rut=?, 
		name=?, 
		race=? 
		WHERE id=?`

	params := []interface{}{pet.OwnerRut, pet.Name, pet.Race, *pet.Id}

	_, err := r.sqlTemplate.Exec(ctx, query, params)
	if err != nil {
		return nil, err
	}
	return pet, nil
}

func (r *petRepo) Delete(ctx context.Context, pet *pet.Pet) error {
	query := fmt.Sprintf("DELETE FROM pets WHERE id=?")
	params := []interface{}{*pet.Id}
	_, err := r.sqlTemplate.Exec(ctx, query, params)
	return err
}

func (r *petRepo) FindByOwnerRut(ctx context.Context, ownerRut int) ([]pet.Pet, error) {
	query :=
		`SELECT %s FROM pets pet 
	JOIN persons per ON (pet.owner_rut = per.rut) 
	WHERE per.rut = ? 
	ORDER BY pet.name ASC`
	query = fmt.Sprintf(query, petQueryProjectionWithAlias)

	params := []interface{}{ownerRut}

	return r.sqlTemplate.QueryForArray(ctx, query, params, r.mapper)
}

func (r *petRepo) FindById(ctx context.Context, id int64) (*pet.Pet, error) {
	query := `SELECT %s FROM pets WHERE id = ?`
	query = fmt.Sprintf(query, petQueryProjection)

	params := []interface{}{id}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.mapper)
}

// privates

func (r *petRepo) mapper(rows *sql.Rows) (pet.Pet, error) {
	var p pet.Pet
	err := rows.Scan(&p.Id, &p.OwnerRut, &p.Name, &p.Race)
	return p, err
}
