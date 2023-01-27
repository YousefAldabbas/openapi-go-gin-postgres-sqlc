// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const GetUser = `-- name: GetUser :one
select
  username
  , email
  , role_rank
  , created_at
  , updated_at
  , user_id
  -- case when @get_db_data::boolean then
  --   (user_id)
  -- end as user_id, -- TODO sqlc.yaml overrides sql.NullInt64
from
  users
where (email = LOWER($1)::text
  or $1::text is null)
and (username = $2::text
  or $2::text is null)
and (user_id = $3::uuid
  or $3::uuid is null)
limit 1
`

type GetUserParams struct {
	Email    *string    `db:"email" json:"email"`
	Username *string    `db:"username" json:"username"`
	UserID   *uuid.UUID `db:"user_id" json:"user_id"`
}

type GetUserRow struct {
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	RoleRank  int16     `db:"role_rank" json:"role_rank"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
}

func (q *Queries) GetUser(ctx context.Context, db DBTX, arg GetUserParams) (GetUserRow, error) {
	row := db.QueryRow(ctx, GetUser, arg.Email, arg.Username, arg.UserID)
	var i GetUserRow
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.RoleRank,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
	)
	return i, err
}

const RegisterNewUser = `-- name: RegisterNewUser :one
insert into users (
  username
  , email
  , role_rank)
values (
  $1
  , $2
  , $3)
returning
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
`

type RegisterNewUserParams struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	RoleRank int16  `db:"role_rank" json:"role_rank"`
}

type RegisterNewUserRow struct {
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	RoleRank  int16     `db:"role_rank" json:"role_rank"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// plpgsql-language-server:disable
func (q *Queries) RegisterNewUser(ctx context.Context, db DBTX, arg RegisterNewUserParams) (RegisterNewUserRow, error) {
	row := db.QueryRow(ctx, RegisterNewUser, arg.Username, arg.Email, arg.RoleRank)
	var i RegisterNewUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.Email,
		&i.RoleRank,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const Test = `-- name: Test :exec
select
  user_id
  , username
  , email
  , role_rank
  , created_at
  , updated_at
from
  users
`

type TestRow struct {
	UserID    uuid.UUID `db:"user_id" json:"user_id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	RoleRank  int16     `db:"role_rank" json:"role_rank"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// update
//
//	users
//
// set
//
//	username = null
//	, email = COALESCE(LOWER(sqlc.narg('email')) , email)
//
// where
//
//	user_id = @user_id;
func (q *Queries) Test(ctx context.Context, db DBTX) error {
	_, err := db.Exec(ctx, Test)
	return err
}
