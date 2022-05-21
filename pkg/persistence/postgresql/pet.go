package postgresql

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
	VALUES ($1, $2, $3) RETURNING %s`
	query = fmt.Sprintf(query, petQueryProjection)

	params := []interface{}{pet.OwnerRut, pet.Name, pet.Race}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.mapper)
}

func (r *petRepo) Update(ctx context.Context, pet *pet.Pet) (*pet.Pet, error) {
	query :=
		`UPDATE pets SET
	owner_rut=$1, 
	name=$2, 
	race=$3
	WHERE id=$4 
	RETURNING %s`
	query = fmt.Sprintf(query, petQueryProjection)

	params := []interface{}{pet.OwnerRut, pet.Name, pet.Race, *pet.Id}

	return r.sqlTemplate.QueryForOne(ctx, query, params, r.mapper)
}

func (r *petRepo) Delete(ctx context.Context, pet *pet.Pet) error {
	query := fmt.Sprintf("DELETE FROM pets WHERE id=$1 RETURNING id")
	params := []interface{}{*pet.Id}
	_, err := r.sqlTemplate.QueryForOne(ctx, query, params, r.mapper)
	return err
}

func (r *petRepo) FindByOwnerRut(ctx context.Context, ownerRut int) ([]pet.Pet, error) {
	query :=
		`SELECT %s FROM pets pet 
	JOIN persons per ON (pet.owner_rut = per.rut) 
	WHERE per.rut = $1 
	ORDER BY pet.name ASC`
	query = fmt.Sprintf(query, petQueryProjectionWithAlias)

	params := []interface{}{ownerRut}

	return r.sqlTemplate.QueryForArray(ctx, query, params, r.mapper)
}

func (r *petRepo) FindById(ctx context.Context, id int64) (*pet.Pet, error) {
	query := `SELECT %s FROM pets WHERE id = $1`
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
