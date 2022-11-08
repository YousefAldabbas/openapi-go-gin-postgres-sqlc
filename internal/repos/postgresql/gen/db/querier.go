// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	// plpgsql-language-server:use-keyword-query-parameters
	GetUser(ctx context.Context, db DBTX, arg GetUserParams) (GetUserRow, error)
	//----------------------------
	// this below would be O2O
	GetUsersWithJoins(ctx context.Context, db DBTX, arg GetUsersWithJoinsParams) ([]GetUsersWithJoinsRow, error)
	// -- name: Test :exec
	// update
	//   users
	// set
	//   username = '@test'
	//   , email = COALESCE(LOWER(sqlc.narg('email')) , email)
	// where
	//   user_id = @user_id;
	ListAllUsers(ctx context.Context, db DBTX) ([]ListAllUsersRow, error)
	// if O2O. This is discovered from FK on user_api_keys.user_id being also a unique constraint
	UpdateUserById(ctx context.Context, db DBTX, arg UpdateUserByIdParams) error
}

var _ Querier = (*Queries)(nil)
