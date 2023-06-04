package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// WorkItemTag represents a row from 'public.work_item_tags'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type WorkItemTag struct {
	WorkItemTagID int    `json:"workItemTagID" db:"work_item_tag_id" required:"true"` // work_item_tag_id
	ProjectID     int    `json:"projectID" db:"project_id" required:"true"`           // project_id
	Name          string `json:"name" db:"name" required:"true"`                      // name
	Description   string `json:"description" db:"description" required:"true"`        // description
	Color         string `json:"color" db:"color" required:"true"`                    // color

	ProjectJoin              *Project    `json:"-" db:"project_project_id" openapi-go:"ignore"`                 // O2O projects (generated from M2O)
	WorkItemTagWorkItemsJoin *[]WorkItem `json:"-" db:"work_item_work_item_tag_work_items" openapi-go:"ignore"` // M2M work_item_work_item_tag

}

// WorkItemTagCreateParams represents insert params for 'public.work_item_tags'.
type WorkItemTagCreateParams struct {
	ProjectID   int    `json:"projectID" required:"true"`   // project_id
	Name        string `json:"name" required:"true"`        // name
	Description string `json:"description" required:"true"` // description
	Color       string `json:"color" required:"true"`       // color
}

// CreateWorkItemTag creates a new WorkItemTag in the database with the given params.
func CreateWorkItemTag(ctx context.Context, db DB, params *WorkItemTagCreateParams) (*WorkItemTag, error) {
	wit := &WorkItemTag{
		ProjectID:   params.ProjectID,
		Name:        params.Name,
		Description: params.Description,
		Color:       params.Color,
	}

	return wit.Insert(ctx, db)
}

// WorkItemTagUpdateParams represents update params for 'public.work_item_tags'.
type WorkItemTagUpdateParams struct {
	ProjectID   *int    `json:"projectID" required:"true"`   // project_id
	Name        *string `json:"name" required:"true"`        // name
	Description *string `json:"description" required:"true"` // description
	Color       *string `json:"color" required:"true"`       // color
}

// SetUpdateParams updates public.work_item_tags struct fields with the specified params.
func (wit *WorkItemTag) SetUpdateParams(params *WorkItemTagUpdateParams) {
	if params.ProjectID != nil {
		wit.ProjectID = *params.ProjectID
	}
	if params.Name != nil {
		wit.Name = *params.Name
	}
	if params.Description != nil {
		wit.Description = *params.Description
	}
	if params.Color != nil {
		wit.Color = *params.Color
	}
}

type WorkItemTagSelectConfig struct {
	limit   string
	orderBy string
	joins   WorkItemTagJoins
	filters map[string][]any
}
type WorkItemTagSelectConfigOption func(*WorkItemTagSelectConfig)

// WithWorkItemTagLimit limits row selection.
func WithWorkItemTagLimit(limit int) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type WorkItemTagOrderBy string

const ()

type WorkItemTagJoins struct {
	Project              bool // O2O projects
	WorkItemsWorkItemTag bool // M2M work_item_work_item_tag
}

// WithWorkItemTagJoin joins with the given tables.
func WithWorkItemTagJoin(joins WorkItemTagJoins) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.joins = WorkItemTagJoins{
			Project:              s.joins.Project || joins.Project,
			WorkItemsWorkItemTag: s.joins.WorkItemsWorkItemTag || joins.WorkItemsWorkItemTag,
		}
	}
}

// WithWorkItemTagFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithWorkItemTagFilters(filters map[string][]any) WorkItemTagSelectConfigOption {
	return func(s *WorkItemTagSelectConfig) {
		s.filters = filters
	}
}

const workItemTagTableProjectJoinSQL = `-- O2O join generated from "work_item_tags_project_id_fkey (Generated from M2O)"
left join projects as _work_item_tags_project_id on _work_item_tags_project_id.project_id = work_item_tags.project_id
`

const workItemTagTableProjectSelectSQL = `(case when _work_item_tags_project_id.project_id is not null then row(_work_item_tags_project_id.*) end) as project_project_id`

const workItemTagTableProjectGroupBySQL = `_work_item_tags_project_id.project_id,
      _work_item_tags_project_id.project_id,
	work_item_tags.work_item_tag_id`

const workItemTagTableWorkItemsWorkItemTagJoinSQL = `-- M2M join generated from "work_item_work_item_tag_work_item_id_fkey"
left join (
	select
		work_item_work_item_tag.work_item_tag_id as work_item_work_item_tag_work_item_tag_id
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		work_item_work_item_tag
	join work_items on work_items.work_item_id = work_item_work_item_tag.work_item_id
	group by
		work_item_work_item_tag_work_item_tag_id
		, work_items.work_item_id
) as joined_work_item_work_item_tag_work_items on joined_work_item_work_item_tag_work_items.work_item_work_item_tag_work_item_tag_id = work_item_tags.work_item_tag_id
`

const workItemTagTableWorkItemsWorkItemTagSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_work_item_tag_work_items.__work_items
		)) filter (where joined_work_item_work_item_tag_work_items.__work_items_work_item_id is not null), '{}') as work_item_work_item_tag_work_items`

const workItemTagTableWorkItemsWorkItemTagGroupBySQL = `work_item_tags.work_item_tag_id, work_item_tags.work_item_tag_id`

// Insert inserts the WorkItemTag to the database.
func (wit *WorkItemTag) Insert(ctx context.Context, db DB) (*WorkItemTag, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.work_item_tags (
	project_id, name, description, color
	) VALUES (
	$1, $2, $3, $4
	) RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Insert/db.Query: %w", err))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Insert/pgx.CollectOneRow: %w", err))
	}

	*wit = newwit

	return wit, nil
}

// Update updates a WorkItemTag in the database.
func (wit *WorkItemTag) Update(ctx context.Context, db DB) (*WorkItemTag, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.work_item_tags SET 
	project_id = $1, name = $2, description = $3, color = $4 
	WHERE work_item_tag_id = $5 
	RETURNING * `
	// run
	logf(sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)

	rows, err := db.Query(ctx, sqlstr, wit.ProjectID, wit.Name, wit.Description, wit.Color, wit.WorkItemTagID)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Update/db.Query: %w", err))
	}
	newwit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Update/pgx.CollectOneRow: %w", err))
	}
	*wit = newwit

	return wit, nil
}

// Upsert upserts a WorkItemTag in the database.
// Requires appropiate PK(s) to be set beforehand.
func (wit *WorkItemTag) Upsert(ctx context.Context, db DB, params *WorkItemTagCreateParams) (*WorkItemTag, error) {
	var err error

	wit.ProjectID = params.ProjectID
	wit.Name = params.Name
	wit.Description = params.Description
	wit.Color = params.Color

	wit, err = wit.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			wit, err = wit.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return wit, err
}

// Delete deletes the WorkItemTag from the database.
func (wit *WorkItemTag) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.work_item_tags 
	WHERE work_item_tag_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, wit.WorkItemTagID); err != nil {
		return logerror(err)
	}
	return nil
}

// WorkItemTagPaginatedByWorkItemTagIDAsc returns a cursor-paginated list of WorkItemTag in Asc order.
func WorkItemTagPaginatedByWorkItemTagIDAsc(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.work_item_tag_id > $1
	 %s   %s 
  ORDER BY 
		work_item_tag_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagPaginatedByWorkItemTagIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagPaginatedByProjectIDAsc returns a cursor-paginated list of WorkItemTag in Asc order.
func WorkItemTagPaginatedByProjectIDAsc(ctx context.Context, db DB, projectID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.project_id > $1
	 %s   %s 
  ORDER BY 
		project_id Asc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagPaginatedByProjectIDAsc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagPaginatedByWorkItemTagIDDesc returns a cursor-paginated list of WorkItemTag in Desc order.
func WorkItemTagPaginatedByWorkItemTagIDDesc(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.work_item_tag_id < $1
	 %s   %s 
  ORDER BY 
		work_item_tag_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagPaginatedByWorkItemTagIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagPaginatedByProjectIDDesc returns a cursor-paginated list of WorkItemTag in Desc order.
func WorkItemTagPaginatedByProjectIDDesc(ctx context.Context, db DB, projectID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.project_id < $1
	 %s   %s 
  ORDER BY 
		project_id Desc`, selects, joins, filters, groupbys)
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagPaginatedByProjectIDDesc */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagByNameProjectID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagByNameProjectID(ctx context.Context, db DB, name string, projectID int, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 2
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.name = $1 AND work_item_tags.project_id = $2
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagByNameProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, name, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{name, projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByNameProjectID/db.Query: %w", err))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByNameProjectID/pgx.CollectOneRow: %w", err))
	}

	return &wit, nil
}

// WorkItemTagsByName retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagsByName(ctx context.Context, db DB, name string, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.name = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagsByName */\n" + sqlstr

	// run
	// logf(sqlstr, name)
	rows, err := db.Query(ctx, sqlstr, append([]any{name}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagsByProjectID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_name_project_id_key'.
func WorkItemTagsByProjectID(ctx context.Context, db DB, projectID int, opts ...WorkItemTagSelectConfigOption) ([]WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.project_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagsByProjectID */\n" + sqlstr

	// run
	// logf(sqlstr, projectID)
	rows, err := db.Query(ctx, sqlstr, append([]any{projectID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/Query: %w", err))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("WorkItemTag/WorkItemTagByNameProjectID/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// WorkItemTagByWorkItemTagID retrieves a row from 'public.work_item_tags' as a WorkItemTag.
//
// Generated from index 'work_item_tags_pkey'.
func WorkItemTagByWorkItemTagID(ctx context.Context, db DB, workItemTagID int, opts ...WorkItemTagSelectConfigOption) (*WorkItemTag, error) {
	c := &WorkItemTagSelectConfig{joins: WorkItemTagJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Project {
		selectClauses = append(selectClauses, workItemTagTableProjectSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableProjectJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableProjectGroupBySQL)
	}

	if c.joins.WorkItemsWorkItemTag {
		selectClauses = append(selectClauses, workItemTagTableWorkItemsWorkItemTagSelectSQL)
		joinClauses = append(joinClauses, workItemTagTableWorkItemsWorkItemTagJoinSQL)
		groupByClauses = append(groupByClauses, workItemTagTableWorkItemsWorkItemTagGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	work_item_tags.work_item_tag_id,
	work_item_tags.project_id,
	work_item_tags.name,
	work_item_tags.description,
	work_item_tags.color %s 
	 FROM public.work_item_tags %s 
	 WHERE work_item_tags.work_item_tag_id = $1
	 %s   %s 
`, selects, joins, filters, groupbys)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* WorkItemTagByWorkItemTagID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemTagID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemTagID}, filterParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/db.Query: %w", err))
	}
	wit, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[WorkItemTag])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_tags/WorkItemTagByWorkItemTagID/pgx.CollectOneRow: %w", err))
	}

	return &wit, nil
}

// FKProject_ProjectID returns the Project associated with the WorkItemTag's (ProjectID).
//
// Generated from foreign key 'work_item_tags_project_id_fkey'.
func (wit *WorkItemTag) FKProject_ProjectID(ctx context.Context, db DB) (*Project, error) {
	return ProjectByProjectID(ctx, db, wit.ProjectID)
}
