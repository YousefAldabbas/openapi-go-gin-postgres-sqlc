package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// User represents a row from 'public.users'.
type User struct {
	UserID     uuid.UUID  `json:"user_id" db:"user_id"`         // user_id
	Username   string     `json:"username" db:"username"`       // username
	Email      string     `json:"email" db:"email"`             // email
	FirstName  *string    `json:"first_name" db:"first_name"`   // first_name
	LastName   *string    `json:"last_name" db:"last_name"`     // last_name
	FullName   *string    `json:"full_name" db:"full_name"`     // full_name
	ExternalID string     `json:"external_id" db:"external_id"` // external_id
	APIKeyID   *int       `json:"api_key_id" db:"api_key_id"`   // api_key_id
	Scopes     []string   `json:"scopes" db:"scopes"`           // scopes
	RoleRank   int16      `json:"role_rank" db:"role_rank"`     // role_rank
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`   // created_at
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`   // updated_at
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at"`   // deleted_at

	TimeEntries *[]TimeEntry `json:"time_entries"` // O2M
	Teams       *[]Team      `json:"teams"`        // M2M
	WorkItems   *[]WorkItem  `json:"work_items"`   // M2M
	// xo fields
	_exists, _deleted bool
}

type UserSelectConfig struct {
	limit     string
	orderBy   string
	joins     UserJoins
	deletedAt string
}
type UserSelectConfigOption func(*UserSelectConfig)

// WithUserLimit limits row selection.
func WithUserLimit(limit int) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

// WithDeletedUserOnly limits result to records marked as deleted.
func WithDeletedUserOnly() UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.deletedAt = " not null "
	}
}

type UserOrderBy = string

const (
	UserCreatedAtDescNullsFirst UserOrderBy = " created_at DESC NULLS FIRST "
	UserCreatedAtDescNullsLast  UserOrderBy = " created_at DESC NULLS LAST "
	UserCreatedAtAscNullsFirst  UserOrderBy = " created_at ASC NULLS FIRST "
	UserCreatedAtAscNullsLast   UserOrderBy = " created_at ASC NULLS LAST "
	UserUpdatedAtDescNullsFirst UserOrderBy = " updated_at DESC NULLS FIRST "
	UserUpdatedAtDescNullsLast  UserOrderBy = " updated_at DESC NULLS LAST "
	UserUpdatedAtAscNullsFirst  UserOrderBy = " updated_at ASC NULLS FIRST "
	UserUpdatedAtAscNullsLast   UserOrderBy = " updated_at ASC NULLS LAST "
	UserDeletedAtDescNullsFirst UserOrderBy = " deleted_at DESC NULLS FIRST "
	UserDeletedAtDescNullsLast  UserOrderBy = " deleted_at DESC NULLS LAST "
	UserDeletedAtAscNullsFirst  UserOrderBy = " deleted_at ASC NULLS FIRST "
	UserDeletedAtAscNullsLast   UserOrderBy = " deleted_at ASC NULLS LAST "
)

// WithUserOrderBy orders results by the given columns.
func WithUserOrderBy(rows ...UserOrderBy) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type UserJoins struct {
	TimeEntries bool
	Teams       bool
	WorkItems   bool
}

// WithUserJoin orders results by the given columns.
func WithUserJoin(joins UserJoins) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the User has been marked for deletion from
// the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.users (` +
		`username, email, first_name, last_name, external_id, api_key_id, scopes, role_rank, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9` +
		`) RETURNING user_id, full_name, created_at, updated_at `
	// run
	logf(sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.DeletedAt)
	if err := db.QueryRow(ctx, sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.DeletedAt).Scan(&u.UserID, &u.FullName, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.users SET ` +
		`username = $1, email = $2, first_name = $3, last_name = $4, external_id = $5, api_key_id = $6, scopes = $7, role_rank = $8, deleted_at = $9 ` +
		`WHERE user_id = $10 ` +
		`RETURNING user_id, full_name, created_at, updated_at `
	// run
	logf(sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.UserID)
	if err := db.QueryRow(ctx, sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.DeletedAt, u.UserID).Scan(&u.UserID, &u.FullName, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) error {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.users (` +
		`user_id, username, email, first_name, last_name, full_name, external_id, api_key_id, scopes, role_rank, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)` +
		` ON CONFLICT (user_id) DO ` +
		`UPDATE SET ` +
		`username = EXCLUDED.username, email = EXCLUDED.email, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, external_id = EXCLUDED.external_id, api_key_id = EXCLUDED.api_key_id, scopes = EXCLUDED.scopes, role_rank = EXCLUDED.role_rank, deleted_at = EXCLUDED.deleted_at  `
	// run
	logf(sqlstr, u.UserID, u.Username, u.Email, u.FirstName, u.LastName, u.FullName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.DeletedAt)
	if _, err := db.Exec(ctx, sqlstr, u.UserID, u.Username, u.Email, u.FirstName, u.LastName, u.FullName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.DeletedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.users ` +
		`WHERE user_id = $1 `
	// run
	logf(sqlstr, u.UserID)
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UsersByCreatedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_created_at_idx'.
func UsersByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) ([]*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $3::boolean = true then joined_work_items.work_items end)::jsonb as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_id as teams_user_id
		, json_agg(teams.*) as teams
	from
		user_team
		join teams using (team_id)
	where
		user_id in (
			select
				user_id
			from
				user_team
			where
				team_id = any (
					select
						team_id
					from
						teams))
			group by
				user_id) joined_teams on joined_teams.teams_user_id = users.user_id
-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		member as work_items_user_id
		, json_agg(work_items.*) as work_items
	from
		work_item_member
		join work_items using (work_item_id)
	where
		member in (
			select
				member
			from
				work_item_member
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				member) joined_work_items on joined_work_items.work_items_user_id = users.user_id`+
		` WHERE users.created_at = $4  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.Teams, c.joins.WorkItems, createdAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*User
	for rows.Next() {
		u := User{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UsersByDeletedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_deleted_at_idx'.
func UsersByDeletedAt(ctx context.Context, db DB, deletedAt *time.Time, opts ...UserSelectConfigOption) ([]*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $3::boolean = true then joined_work_items.work_items end)::jsonb as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_id as teams_user_id
		, json_agg(teams.*) as teams
	from
		user_team
		join teams using (team_id)
	where
		user_id in (
			select
				user_id
			from
				user_team
			where
				team_id = any (
					select
						team_id
					from
						teams))
			group by
				user_id) joined_teams on joined_teams.teams_user_id = users.user_id
-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		member as work_items_user_id
		, json_agg(work_items.*) as work_items
	from
		work_item_member
		join work_items using (work_item_id)
	where
		member in (
			select
				member
			from
				work_item_member
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				member) joined_work_items on joined_work_items.work_items_user_id = users.user_id`+
		` WHERE users.deleted_at = $4  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, deletedAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.Teams, c.joins.WorkItems, deletedAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*User
	for rows.Next() {
		u := User{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UserByEmail retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_email_key'.
func UserByEmail(ctx context.Context, db DB, email string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $3::boolean = true then joined_work_items.work_items end)::jsonb as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_id as teams_user_id
		, json_agg(teams.*) as teams
	from
		user_team
		join teams using (team_id)
	where
		user_id in (
			select
				user_id
			from
				user_team
			where
				team_id = any (
					select
						team_id
					from
						teams))
			group by
				user_id) joined_teams on joined_teams.teams_user_id = users.user_id
-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		member as work_items_user_id
		, json_agg(work_items.*) as work_items
	from
		work_item_member
		join work_items using (work_item_id)
	where
		member in (
			select
				member
			from
				work_item_member
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				member) joined_work_items on joined_work_items.work_items_user_id = users.user_id`+
		` WHERE users.email = $4  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, email)
	u := User{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, c.joins.Teams, c.joins.WorkItems, email).Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.TimeEntries, &u.Teams, &u.WorkItems); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UserByExternalID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_external_id_key'.
func UserByExternalID(ctx context.Context, db DB, externalID string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $3::boolean = true then joined_work_items.work_items end)::jsonb as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_id as teams_user_id
		, json_agg(teams.*) as teams
	from
		user_team
		join teams using (team_id)
	where
		user_id in (
			select
				user_id
			from
				user_team
			where
				team_id = any (
					select
						team_id
					from
						teams))
			group by
				user_id) joined_teams on joined_teams.teams_user_id = users.user_id
-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		member as work_items_user_id
		, json_agg(work_items.*) as work_items
	from
		work_item_member
		join work_items using (work_item_id)
	where
		member in (
			select
				member
			from
				work_item_member
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				member) joined_work_items on joined_work_items.work_items_user_id = users.user_id`+
		` WHERE users.external_id = $4  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, externalID)
	u := User{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, c.joins.Teams, c.joins.WorkItems, externalID).Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.TimeEntries, &u.Teams, &u.WorkItems); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UserByUserID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $3::boolean = true then joined_work_items.work_items end)::jsonb as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_id as teams_user_id
		, json_agg(teams.*) as teams
	from
		user_team
		join teams using (team_id)
	where
		user_id in (
			select
				user_id
			from
				user_team
			where
				team_id = any (
					select
						team_id
					from
						teams))
			group by
				user_id) joined_teams on joined_teams.teams_user_id = users.user_id
-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		member as work_items_user_id
		, json_agg(work_items.*) as work_items
	from
		work_item_member
		join work_items using (work_item_id)
	where
		member in (
			select
				member
			from
				work_item_member
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				member) joined_work_items on joined_work_items.work_items_user_id = users.user_id`+
		` WHERE users.user_id = $4  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, userID)
	u := User{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, c.joins.Teams, c.joins.WorkItems, userID).Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.TimeEntries, &u.Teams, &u.WorkItems); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UsersByUpdatedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_updated_at_idx'.
func UsersByUpdatedAt(ctx context.Context, db DB, updatedAt time.Time, opts ...UserSelectConfigOption) ([]*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $3::boolean = true then joined_work_items.work_items end)::jsonb as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_id as teams_user_id
		, json_agg(teams.*) as teams
	from
		user_team
		join teams using (team_id)
	where
		user_id in (
			select
				user_id
			from
				user_team
			where
				team_id = any (
					select
						team_id
					from
						teams))
			group by
				user_id) joined_teams on joined_teams.teams_user_id = users.user_id
-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		member as work_items_user_id
		, json_agg(work_items.*) as work_items
	from
		work_item_member
		join work_items using (work_item_id)
	where
		member in (
			select
				member
			from
				work_item_member
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				member) joined_work_items on joined_work_items.work_items_user_id = users.user_id`+
		` WHERE users.updated_at = $4  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, updatedAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.Teams, c.joins.WorkItems, updatedAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*User
	for rows.Next() {
		u := User{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// UserByUsername retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_username_key'.
func UserByUsername(ctx context.Context, db DB, username string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries,
(case when $2::boolean = true then joined_teams.teams end)::jsonb as teams,
(case when $3::boolean = true then joined_work_items.work_items end)::jsonb as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , json_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_id as teams_user_id
		, json_agg(teams.*) as teams
	from
		user_team
		join teams using (team_id)
	where
		user_id in (
			select
				user_id
			from
				user_team
			where
				team_id = any (
					select
						team_id
					from
						teams))
			group by
				user_id) joined_teams on joined_teams.teams_user_id = users.user_id
-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		member as work_items_user_id
		, json_agg(work_items.*) as work_items
	from
		work_item_member
		join work_items using (work_item_id)
	where
		member in (
			select
				member
			from
				work_item_member
			where
				work_item_id = any (
					select
						work_item_id
					from
						work_items))
			group by
				member) joined_work_items on joined_work_items.work_items_user_id = users.user_id`+
		` WHERE users.username = $4  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, username)
	u := User{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.TimeEntries, c.joins.Teams, c.joins.WorkItems, username).Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.TimeEntries, &u.Teams, &u.WorkItems); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// FKUserAPIKey returns the UserAPIKey associated with the User's (APIKeyID).
//
// Generated from foreign key 'users_api_key_id_fkey'.
func (u *User) FKUserAPIKey(ctx context.Context, db DB) (*UserAPIKey, error) {
	return UserAPIKeyByUserAPIKeyID(ctx, db, *u.APIKeyID)
}
