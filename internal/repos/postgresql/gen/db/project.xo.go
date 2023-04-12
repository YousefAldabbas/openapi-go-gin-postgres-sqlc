package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v5"
)

// Project represents a row from 'public.projects'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type Project struct {
	ProjectID          int            `json:"projectID" db:"project_id" required:"true"`                         // project_id
	Name               models.Project `json:"name" db:"name" required:"true" ref:"#/components/schemas/Project"` // name
	Description        string         `json:"description" db:"description" required:"true"`                      // description
	WorkItemsTableName string         `json:"-" db:"work_items_table_name"`                                      // work_items_table_name
	Initialized        bool           `json:"initialized" db:"initialized" required:"true"`                      // initialized
	BoardConfig        pgtype.JSONB   `json:"-" db:"board_config"`                                               // board_config
	CreatedAt          time.Time      `json:"createdAt" db:"created_at" required:"true"`                         // created_at
	UpdatedAt          time.Time      `json:"updatedAt" db:"updated_at" required:"true"`                         // updated_at

	Activities    *[]Activity     `json:"activities" db:"activities"`         // O2M
	KanbanSteps   *[]KanbanStep   `json:"kanbanSteps" db:"kanban_steps"`      // O2M
	Teams         *[]Team         `json:"teams" db:"teams"`                   // O2M
	WorkItemTags  *[]WorkItemTag  `json:"workItemTags" db:"work_item_tags"`   // O2M
	WorkItemTypes *[]WorkItemType `json:"workItemTypes" db:"work_item_types"` // O2M
	// xo fields
	_exists, _deleted bool
}

// ProjectCreateParams represents insert params for 'public.projects'
type ProjectCreateParams struct {
	Name               models.Project `json:"name"`        // name
	Description        string         `json:"description"` // description
	WorkItemsTableName string         `json:"-"`           // work_items_table_name
	Initialized        bool           `json:"initialized"` // initialized
	BoardConfig        pgtype.JSONB   `json:"-"`           // board_config
}

// ProjectUpdateParams represents update params for 'public.projects'
type ProjectUpdateParams struct {
	Name               *models.Project `json:"name"`        // name
	Description        *string         `json:"description"` // description
	WorkItemsTableName *string         `json:"-"`           // work_items_table_name
	Initialized        *bool           `json:"initialized"` // initialized
	BoardConfig        *pgtype.JSONB   `json:"-"`           // board_config
}

type ProjectSelectConfig struct {
	limit   string
	orderBy string
	joins   ProjectJoins
}
type ProjectSelectConfigOption func(*ProjectSelectConfig)

// WithProjectLimit limits row selection.
func WithProjectLimit(limit int) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type ProjectOrderBy = string

const (
	ProjectCreatedAtDescNullsFirst ProjectOrderBy = " created_at DESC NULLS FIRST "
	ProjectCreatedAtDescNullsLast  ProjectOrderBy = " created_at DESC NULLS LAST "
	ProjectCreatedAtAscNullsFirst  ProjectOrderBy = " created_at ASC NULLS FIRST "
	ProjectCreatedAtAscNullsLast   ProjectOrderBy = " created_at ASC NULLS LAST "
	ProjectUpdatedAtDescNullsFirst ProjectOrderBy = " updated_at DESC NULLS FIRST "
	ProjectUpdatedAtDescNullsLast  ProjectOrderBy = " updated_at DESC NULLS LAST "
	ProjectUpdatedAtAscNullsFirst  ProjectOrderBy = " updated_at ASC NULLS FIRST "
	ProjectUpdatedAtAscNullsLast   ProjectOrderBy = " updated_at ASC NULLS LAST "
)

// WithProjectOrderBy orders results by the given columns.
func WithProjectOrderBy(rows ...ProjectOrderBy) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type ProjectJoins struct {
	Activities    bool
	KanbanSteps   bool
	Teams         bool
	WorkItemTags  bool
	WorkItemTypes bool
}

// WithProjectJoin joins with the given tables.
func WithProjectJoin(joins ProjectJoins) ProjectSelectConfigOption {
	return func(s *ProjectSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the Project exists in the database.
func (p *Project) Exists() bool {
	return p._exists
}

// Deleted returns true when the Project has been marked for deletion from
// the database.
func (p *Project) Deleted() bool {
	return p._deleted
}

// Insert inserts the Project to the database.
func (p *Project) Insert(ctx context.Context, db DB) (*Project, error) {
	switch {
	case p._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.projects (` +
		`name, description, work_items_table_name, initialized, board_config` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING * `
	// run
	logf(sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.Initialized, p.BoardConfig)

	rows, err := db.Query(ctx, sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.Initialized, p.BoardConfig)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Insert/db.Query: %w", err))
	}
	newp, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Insert/pgx.CollectOneRow: %w", err))
	}
	newp._exists = true
	*p = newp

	return p, nil
}

// Update updates a Project in the database.
func (p *Project) Update(ctx context.Context, db DB) (*Project, error) {
	switch {
	case !p._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.projects SET ` +
		`name = $1, description = $2, work_items_table_name = $3, initialized = $4, board_config = $5 ` +
		`WHERE project_id = $6 ` +
		`RETURNING * `
	// run
	logf(sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.Initialized, p.BoardConfig, p.CreatedAt, p.UpdatedAt, p.ProjectID)

	rows, err := db.Query(ctx, sqlstr, p.Name, p.Description, p.WorkItemsTableName, p.Initialized, p.BoardConfig, p.ProjectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Update/db.Query: %w", err))
	}
	newp, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("Project/Update/pgx.CollectOneRow: %w", err))
	}
	newp._exists = true
	*p = newp

	return p, nil
}

// Save saves the Project to the database.
func (p *Project) Save(ctx context.Context, db DB) (*Project, error) {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for Project.
func (p *Project) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.projects (` +
		`project_id, name, description, work_items_table_name, initialized, board_config` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)` +
		` ON CONFLICT (project_id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, description = EXCLUDED.description, work_items_table_name = EXCLUDED.work_items_table_name, initialized = EXCLUDED.initialized, board_config = EXCLUDED.board_config  `
	// run
	logf(sqlstr, p.ProjectID, p.Name, p.Description, p.WorkItemsTableName, p.Initialized, p.BoardConfig)
	if _, err := db.Exec(ctx, sqlstr, p.ProjectID, p.Name, p.Description, p.WorkItemsTableName, p.Initialized, p.BoardConfig); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the Project from the database.
func (p *Project) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.projects ` +
		`WHERE project_id = $1 `
	// run
	logf(sqlstr, p.ProjectID)
	if _, err := db.Exec(ctx, sqlstr, p.ProjectID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// ProjectByName retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_name_key'.
func ProjectByName(ctx context.Context, db DB, name models.Project, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`projects.project_id,
projects.name,
projects.description,
projects.work_items_table_name,
projects.initialized,
projects.board_config,
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then joined_activities.activities end) as activities,
(case when $2::boolean = true then joined_kanban_steps.kanban_steps end) as kanban_steps,
(case when $3::boolean = true then joined_teams.teams end) as teams,
(case when $4::boolean = true then joined_work_item_tags.work_item_tags end) as work_item_tags,
(case when $5::boolean = true then joined_work_item_types.work_item_types end) as work_item_types ` +
		`FROM public.projects ` +
		`-- O2M join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , array_agg(activities.*) as activities
  from
    activities
   group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- O2M join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , array_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
   group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- O2M join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , array_agg(teams.*) as teams
  from
    teams
   group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- O2M join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
   group by
        project_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = projects.project_id
-- O2M join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , array_agg(work_item_types.*) as work_item_types
  from
    work_item_types
   group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.name = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activities, c.joins.KanbanSteps, c.joins.Teams, c.joins.WorkItemTags, c.joins.WorkItemTypes, name)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByName/db.Query: %w", err))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByName/pgx.CollectOneRow: %w", err))
	}
	p._exists = true
	return &p, nil
}

// ProjectByProjectID retrieves a row from 'public.projects' as a Project.
//
// Generated from index 'projects_pkey'.
func ProjectByProjectID(ctx context.Context, db DB, projectID int, opts ...ProjectSelectConfigOption) (*Project, error) {
	c := &ProjectSelectConfig{joins: ProjectJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`projects.project_id,
projects.name,
projects.description,
projects.work_items_table_name,
projects.initialized,
projects.board_config,
projects.created_at,
projects.updated_at,
(case when $1::boolean = true then joined_activities.activities end) as activities,
(case when $2::boolean = true then joined_kanban_steps.kanban_steps end) as kanban_steps,
(case when $3::boolean = true then joined_teams.teams end) as teams,
(case when $4::boolean = true then joined_work_item_tags.work_item_tags end) as work_item_tags,
(case when $5::boolean = true then joined_work_item_types.work_item_types end) as work_item_types ` +
		`FROM public.projects ` +
		`-- O2M join generated from "activities_project_id_fkey"
left join (
  select
  project_id as activities_project_id
    , array_agg(activities.*) as activities
  from
    activities
   group by
        project_id) joined_activities on joined_activities.activities_project_id = projects.project_id
-- O2M join generated from "kanban_steps_project_id_fkey"
left join (
  select
  project_id as kanban_steps_project_id
    , array_agg(kanban_steps.*) as kanban_steps
  from
    kanban_steps
   group by
        project_id) joined_kanban_steps on joined_kanban_steps.kanban_steps_project_id = projects.project_id
-- O2M join generated from "teams_project_id_fkey"
left join (
  select
  project_id as teams_project_id
    , array_agg(teams.*) as teams
  from
    teams
   group by
        project_id) joined_teams on joined_teams.teams_project_id = projects.project_id
-- O2M join generated from "work_item_tags_project_id_fkey"
left join (
  select
  project_id as work_item_tags_project_id
    , array_agg(work_item_tags.*) as work_item_tags
  from
    work_item_tags
   group by
        project_id) joined_work_item_tags on joined_work_item_tags.work_item_tags_project_id = projects.project_id
-- O2M join generated from "work_item_types_project_id_fkey"
left join (
  select
  project_id as work_item_types_project_id
    , array_agg(work_item_types.*) as work_item_types
  from
    work_item_types
   group by
        project_id) joined_work_item_types on joined_work_item_types.work_item_types_project_id = projects.project_id` +
		` WHERE projects.project_id = $6 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, c.joins.Activities, c.joins.KanbanSteps, c.joins.Teams, c.joins.WorkItemTags, c.joins.WorkItemTypes, projectID)
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByProjectID/db.Query: %w", err))
	}
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Project])
	if err != nil {
		return nil, logerror(fmt.Errorf("projects/ProjectByProjectID/pgx.CollectOneRow: %w", err))
	}
	p._exists = true
	return &p, nil
}
