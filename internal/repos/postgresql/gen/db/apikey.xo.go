package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/google/uuid"
)

// APIKey represents a row from 'public.api_keys'.
type APIKey struct {
	APIKeyID int       `json:"api_key_id"` // api_key_id
	APIKey   string    `json:"api_key"`    // api_key
	UserID   uuid.UUID `json:"user_id"`    // user_id
	// xo fields
	_exists, _deleted bool
}

// GetMostRecentAPIKey returns n most recent rows from 'api_keys',
// ordered by "created_at" in descending order.
func GetMostRecentAPIKey(ctx context.Context, db DB, n int) ([]*APIKey, error) {
	// list
	const sqlstr = `SELECT ` +
		`api_key_id, api_key, user_id ` +
		`FROM public.api_keys ` +
		`ORDER BY created_at DESC LIMIT $1`
	// run
	logf(sqlstr, n)

	rows, err := db.Query(ctx, sqlstr, n)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	// load results
	var res []*APIKey
	for rows.Next() {
		ak := APIKey{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ak.APIKeyID, &ak.APIKey, &ak.UserID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ak)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Exists returns true when the APIKey exists in the database.
func (ak *APIKey) Exists() bool {
	return ak._exists
}

// Deleted returns true when the APIKey has been marked for deletion from
// the database.
func (ak *APIKey) Deleted() bool {
	return ak._deleted
}

// Insert inserts the APIKey to the database.
func (ak *APIKey) Insert(ctx context.Context, db DB) error {
	switch {
	case ak._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ak._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.api_keys (` +
		`api_key, user_id` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING api_key_id`
	// run
	logf(sqlstr, ak.APIKey, ak.UserID)
	if err := db.QueryRow(ctx, sqlstr, ak.APIKey, ak.UserID).Scan(&ak.APIKeyID); err != nil {
		return logerror(err)
	}
	// set exists
	ak._exists = true
	return nil
}

// Update updates a APIKey in the database.
func (ak *APIKey) Update(ctx context.Context, db DB) error {
	switch {
	case !ak._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ak._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.api_keys SET ` +
		`api_key = $1, user_id = $2 ` +
		`WHERE api_key_id = $3`
	// run
	logf(sqlstr, ak.APIKey, ak.UserID, ak.APIKeyID)
	if _, err := db.Exec(ctx, sqlstr, ak.APIKey, ak.UserID, ak.APIKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the APIKey to the database.
func (ak *APIKey) Save(ctx context.Context, db DB) error {
	if ak.Exists() {
		return ak.Update(ctx, db)
	}
	return ak.Insert(ctx, db)
}

// Upsert performs an upsert for APIKey.
func (ak *APIKey) Upsert(ctx context.Context, db DB) error {
	switch {
	case ak._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.api_keys (` +
		`api_key_id, api_key, user_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (api_key_id) DO ` +
		`UPDATE SET ` +
		`api_key = EXCLUDED.api_key, user_id = EXCLUDED.user_id `
	// run
	logf(sqlstr, ak.APIKeyID, ak.APIKey, ak.UserID)
	if _, err := db.Exec(ctx, sqlstr, ak.APIKeyID, ak.APIKey, ak.UserID); err != nil {
		return logerror(err)
	}
	// set exists
	ak._exists = true
	return nil
}

// Delete deletes the APIKey from the database.
func (ak *APIKey) Delete(ctx context.Context, db DB) error {
	switch {
	case !ak._exists: // doesn't exist
		return nil
	case ak._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.api_keys ` +
		`WHERE api_key_id = $1`
	// run
	logf(sqlstr, ak.APIKeyID)
	if _, err := db.Exec(ctx, sqlstr, ak.APIKeyID); err != nil {
		return logerror(err)
	}
	// set deleted
	ak._deleted = true
	return nil
}

// APIKeyByAPIKey retrieves a row from 'public.api_keys' as a APIKey.
//
// Generated from index 'api_keys_api_key_key'.
func APIKeyByAPIKey(ctx context.Context, db DB, apiKey string) (*APIKey, error) {
	// query
	const sqlstr = `SELECT ` +
		`api_key_id, api_key, user_id ` +
		`FROM public.api_keys ` +
		`WHERE api_key = $1`
	// run
	logf(sqlstr, apiKey)
	ak := APIKey{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, apiKey).Scan(&ak.APIKeyID, &ak.APIKey, &ak.UserID); err != nil {
		return nil, logerror(err)
	}
	return &ak, nil
}

// APIKeyByAPIKeyID retrieves a row from 'public.api_keys' as a APIKey.
//
// Generated from index 'api_keys_pkey'.
func APIKeyByAPIKeyID(ctx context.Context, db DB, apiKeyID int) (*APIKey, error) {
	// query
	const sqlstr = `SELECT ` +
		`api_key_id, api_key, user_id ` +
		`FROM public.api_keys ` +
		`WHERE api_key_id = $1`
	// run
	logf(sqlstr, apiKeyID)
	ak := APIKey{
		_exists: true,
	}
	if err := db.QueryRow(ctx, sqlstr, apiKeyID).Scan(&ak.APIKeyID, &ak.APIKey, &ak.UserID); err != nil {
		return nil, logerror(err)
	}
	return &ak, nil
}

// User returns the User associated with the APIKey's (UserID).
//
// Generated from foreign key 'api_keys_user_id_fkey'.
func (ak *APIKey) User(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, ak.UserID)
}
