// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: animals.sql

package db

import (
	"context"
)

const GetPetById = `-- name: GetPetById :one
select
  pet_id,
  color,
  metadata
from
  pets
left join animals using (animal_id)
where pet_id = $1
limit 1
`

func (q *Queries) GetPetById(ctx context.Context, petID int32) (Pets, error) {
	row := q.db.QueryRow(ctx, GetPetById, petID)
	var i Pets
	err := row.Scan(&i.PetID, &i.Color, &i.Metadata)
	return i, err
}
