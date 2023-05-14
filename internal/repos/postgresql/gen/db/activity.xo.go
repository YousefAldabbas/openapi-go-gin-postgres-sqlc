package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// Activity represents a row from 'public.activities'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type Activity struct {
	ActivityID   int    `json:"activityID" db:"activity_id" required:"true"`     // activity_id
	ProjectID    int    `json:"projectID" db:"project_id" required:"true"`       // project_id
	Name         string `json:"name" db:"name" required:"true"`                  // name
	Description  string `json:"description" db:"description" required:"true"`    // description
	IsProductive bool   `json:"isProductive" db:"is_productive" required:"true"` // is_productive

	ProjectJoin             *Project     `json:"-" db:"project_project_id" openapi-go:"ignore"` // O2O projects (generated from M2O)
	ActivityTimeEntriesJoin *[]TimeEntry `json:"-" db:"time_entries" openapi-go:"ignore"`       // M2O activities

}

// ActivityCreateParams represents insert params for 'public.activities'.
type ActivityCreateParams struct {
	ProjectID    int    `json:"projectID" required:"true"`    // project_id
	Name         string `json:"name" required:"true"`         // name
	Description  string `json:"description" required:"true"`  // description
	IsProductive bool   `json:"isProductive" required:"true"` // is_productive
}

// CreateActivity creates a new Activity in the database with the given params.
func CreateActivity(ctx context.Context, db DB, params *ActivityCreateParams) (*Activity, error) {
	a := &Activity{
		ProjectID:    params.ProjectID,
		Name:         params.Name,
		Description:  params.Description,
		IsProductive: params.IsProductive,
	}

	return a.Insert(ctx, db)
}

// ActivityUpdateParams represents update params for 'public.activities'
type ActivityUpdateParams struct {
	ProjectID    *int    `json:"projectID" required:"true"`    // project_id
	Name         *string `json:"name" required:"true"`         // name
	Description  *string `json:"description" required:"true"`  // description
	IsProductive *bool   `json:"isProductive" required:"true"` // is_productive
}

// SetUpdateParams updates public.activities struct fields with the specified params.
func (a *Activity) SetUpdateParams(params *ActivityUpdateParams) {
	if params.ProjectID != nil {
		a.ProjectID = *params.ProjectID
	}
	if params.Name != nil {
		a.Name = *params.Name
	}
	if params.Description != nil {
		a.Description = *params.Description
	}
	if params.IsProductive != nil {
		a.IsProductive = *params.IsProductive
	}
}

type ActivitySelectConfig struct {
	limit   string
	orderBy string
	joins   ActivityJoins
}
type ActivitySelectConfigOption func(*ActivitySelectConfig)

// WithActivityLimit limits row selection.
func WithActivityLimit(limit int) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ActivityOrderBy = string

const ()

type ActivityJoins struct {
	Project     bool // O2O projects
	TimeEntries bool // M2O time_entries
}

// WithActivityJoin joins with the given tables.
func WithActivityJoin(joins ActivityJoins) ActivitySelectConfigOption {
	return func(s *ActivitySelectConfig) {
		s.joins = ActivityJoins{
			Project:     s.joins.Project || joins.Project,
			TimeEntries: s.joins.TimeEntries || joins.TimeEntries,
		}
	}
}

// Insert inserts the Activity to the database.
func (a *Activity) Insert(ctx context.Context, db DB) (*Activity, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.activities (` +
		`project_id, name, description, is_productive` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING * `
	// run
	logf(sqlstr, a.ProjectID, a.Name, a.Description, a.IsProductive)

	rows, err := db.Query(ctx, sqlstr, a.ProjectID, a.Name, a.Description, a.IsProductive)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Insert/db.Query: %w", err))
	}
	newa, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Insert/pgx.CollectOneRow: %w", err))
	}

	*a = newa

	return a, nil
}

// Update updates a Activity in the database.
func (a *Activity) Update(ctx context.Context, db DB) (*Activity, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.activities SET ` +
		`project_id = $1, name = $2, description = $3, is_productive = $4 ` +
		`WHERE activity_id = $5 ` +
		`RETURNING * `
	// run
	logf(sqlstr, a.ProjectID, a.Name, a.Description, a.IsProductive, a.ActivityID)

	rows, err := db.Query(ctx, sqlstr, a.ProjectID, a.Name, a.Description, a.IsProductive, a.ActivityID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Update/db.Query: %w", err))
	}
	newa, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Update/pgx.CollectOneRow: %w", err))
	}
	*a = newa

	return a, nil
}

// Upsert upserts a Activity in the database.
// Requires appropiate PK(s) to be set beforehand.
func (a *Activity) Upsert(ctx context.Context, db DB, params *ActivityCreateParams) (*Activity, error) {
	var err error

	a.ProjectID = params.ProjectID
	a.Name = params.Name
	a.Description = params.Description
	a.IsProductive = params.IsProductive

	a, err = a.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			a, err = a.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return a, err
}

// Delete deletes the Activity from the database.
func (a *Activity) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.activities ` +
		`WHERE activity_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, a.ActivityID); err != nil {
		return logerror(err)
	}
	return nil
}

// ActivityPaginatedByActivityIDAsc returns a cursor-paginated list of Activity in Asc order.
func ActivityPaginatedByActivityIDAsc(ctx context.Context, db DB, activityID int, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.activity_id > $3 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id ORDER BY 
		activity_id Asc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, activityID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// ActivityPaginatedByProjectIDAsc returns a cursor-paginated list of Activity in Asc order.
func ActivityPaginatedByProjectIDAsc(ctx context.Context, db DB, projectID int, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.project_id > $3 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id ORDER BY 
		project_id Asc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// ActivityPaginatedByActivityIDDesc returns a cursor-paginated list of Activity in Desc order.
func ActivityPaginatedByActivityIDDesc(ctx context.Context, db DB, activityID int, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.activity_id < $3 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id ORDER BY 
		activity_id Desc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, activityID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// ActivityPaginatedByProjectIDDesc returns a cursor-paginated list of Activity in Desc order.
func ActivityPaginatedByProjectIDDesc(ctx context.Context, db DB, projectID int, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.project_id < $3 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id ORDER BY 
		project_id Desc `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// ActivityByNameProjectID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_project_id_key'.
func ActivityByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...ActivitySelectConfigOption) (*Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.name = $3 AND activities.project_id = $4 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, name, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByNameProjectID/db.Query: %w", err))
	}
	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByNameProjectID/pgx.CollectOneRow: %w", err))
	}

	return &a, nil
}

// ActivitiesByName retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_project_id_key'.
func ActivitiesByName(ctx context.Context, db DB, name string, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.name = $3 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// ActivitiesByProjectID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_name_project_id_key'.
func ActivitiesByProjectID(ctx context.Context, db DB, projectID int, opts ...ActivitySelectConfigOption) ([]Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.project_id = $3 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("Activity/ActivityByNameProjectID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// ActivityByActivityID retrieves a row from 'public.activities' as a Activity.
//
// Generated from index 'activities_pkey'.
func ActivityByActivityID(ctx context.Context, db DB, activityID int, opts ...ActivitySelectConfigOption) (*Activity, error) {
	c := &ActivitySelectConfig{joins: ActivityJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`activities.activity_id,
activities.project_id,
activities.name,
activities.description,
activities.is_productive,
(case when $1::boolean = true and _activities_project_id.project_id is not null then row(_activities_project_id.*) end) as project_project_id,
(case when $2::boolean = true then COALESCE(joined_time_entries.time_entries, '{}') end) as time_entries ` +
		`FROM public.activities ` +
		`-- O2O join generated from "activities_project_id_fkey (Generated from M2O)"
left join projects as _activities_project_id on _activities_project_id.project_id = activities.project_id
-- M2O join generated from "time_entries_activity_id_fkey"
left join (
  select
  activity_id as time_entries_activity_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
  group by
        activity_id) joined_time_entries on joined_time_entries.time_entries_activity_id = activities.activity_id` +
		` WHERE activities.activity_id = $3 GROUP BY 
	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
_activities_project_id.project_id,
      _activities_project_id.project_id,
	activities.activity_id, 

	activities.activity_id,
	activities.description,
	activities.is_productive,
	activities.name,
	activities.project_id,
joined_time_entries.time_entries, activities.activity_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, activityID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Project, c.joins.TimeEntries, activityID)
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByActivityID/db.Query: %w", err))
	}
	a, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Activity])
	if err != nil {
		return nil, logerror(fmt.Errorf("activities/ActivityByActivityID/pgx.CollectOneRow: %w", err))
	}

	return &a, nil
}

// FKProject_ProjectID returns the Project associated with the Activity's (ProjectID).
//
// Generated from foreign key 'activities_project_id_fkey'.
func (a *Activity) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, a.ProjectID)
}
