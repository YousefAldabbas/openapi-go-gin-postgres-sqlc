package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgx/v5"
)

// CacheDemoWorkItem represents a row from 'cache.demo_work_items'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type CacheDemoWorkItem struct {
	Ref            string         `json:"ref" db:"ref" required:"true" nullable:"false"`                          // ref
	Line           string         `json:"line" db:"line" required:"true" nullable:"false"`                        // line
	LastMessageAt  time.Time      `json:"lastMessageAt" db:"last_message_at" required:"true" nullable:"false"`    // last_message_at
	Reopened       bool           `json:"reopened" db:"reopened" required:"true" nullable:"false"`                // reopened
	WorkItemID     int            `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`          // work_item_id
	Title          string         `json:"title" db:"title" required:"true" nullable:"false"`                      // title
	Description    string         `json:"description" db:"description" required:"true" nullable:"false"`          // description
	WorkItemTypeID int            `json:"workItemTypeID" db:"work_item_type_id" required:"true" nullable:"false"` // work_item_type_id
	Metadata       map[string]any `json:"metadata" db:"metadata" required:"true" nullable:"false"`                // metadata
	TeamID         int            `json:"teamID" db:"team_id" required:"true" nullable:"false"`                   // team_id
	KanbanStepID   int            `json:"kanbanStepID" db:"kanban_step_id" required:"true" nullable:"false"`      // kanban_step_id
	ClosedAt       *time.Time     `json:"closedAt" db:"closed_at"`                                                // closed_at
	TargetDate     time.Time      `json:"targetDate" db:"target_date" required:"true" nullable:"false"`           // target_date
	CreatedAt      time.Time      `json:"createdAt" db:"created_at" required:"true" nullable:"false"`             // created_at
	UpdatedAt      time.Time      `json:"updatedAt" db:"updated_at" required:"true" nullable:"false"`             // updated_at
	DeletedAt      *time.Time     `json:"deletedAt" db:"deleted_at"`                                              // deleted_at
}

type CacheDemoWorkItemSelectConfig struct {
	limit   string
	orderBy string
	joins   CacheDemoWorkItemJoins
	filters map[string][]any
	having  map[string][]any

	deletedAt string
}
type CacheDemoWorkItemSelectConfigOption func(*CacheDemoWorkItemSelectConfig)

// WithCacheDemoWorkItemLimit limits row selection.
func WithCacheDemoWorkItemLimit(limit int) CacheDemoWorkItemSelectConfigOption {
	return func(s *CacheDemoWorkItemSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDeletedCacheDemoWorkItemOnly limits result to records marked as deleted.
func WithDeletedCacheDemoWorkItemOnly() CacheDemoWorkItemSelectConfigOption {
	return func(s *CacheDemoWorkItemSelectConfig) {
		s.deletedAt = " not null "
	}
}

type CacheDemoWorkItemOrderBy string

const (
	CacheDemoWorkItemLastMessageAtDescNullsFirst CacheDemoWorkItemOrderBy = " last_message_at DESC NULLS FIRST "
	CacheDemoWorkItemLastMessageAtDescNullsLast  CacheDemoWorkItemOrderBy = " last_message_at DESC NULLS LAST "
	CacheDemoWorkItemLastMessageAtAscNullsFirst  CacheDemoWorkItemOrderBy = " last_message_at ASC NULLS FIRST "
	CacheDemoWorkItemLastMessageAtAscNullsLast   CacheDemoWorkItemOrderBy = " last_message_at ASC NULLS LAST "
	CacheDemoWorkItemClosedAtDescNullsFirst      CacheDemoWorkItemOrderBy = " closed_at DESC NULLS FIRST "
	CacheDemoWorkItemClosedAtDescNullsLast       CacheDemoWorkItemOrderBy = " closed_at DESC NULLS LAST "
	CacheDemoWorkItemClosedAtAscNullsFirst       CacheDemoWorkItemOrderBy = " closed_at ASC NULLS FIRST "
	CacheDemoWorkItemClosedAtAscNullsLast        CacheDemoWorkItemOrderBy = " closed_at ASC NULLS LAST "
	CacheDemoWorkItemTargetDateDescNullsFirst    CacheDemoWorkItemOrderBy = " target_date DESC NULLS FIRST "
	CacheDemoWorkItemTargetDateDescNullsLast     CacheDemoWorkItemOrderBy = " target_date DESC NULLS LAST "
	CacheDemoWorkItemTargetDateAscNullsFirst     CacheDemoWorkItemOrderBy = " target_date ASC NULLS FIRST "
	CacheDemoWorkItemTargetDateAscNullsLast      CacheDemoWorkItemOrderBy = " target_date ASC NULLS LAST "
	CacheDemoWorkItemCreatedAtDescNullsFirst     CacheDemoWorkItemOrderBy = " created_at DESC NULLS FIRST "
	CacheDemoWorkItemCreatedAtDescNullsLast      CacheDemoWorkItemOrderBy = " created_at DESC NULLS LAST "
	CacheDemoWorkItemCreatedAtAscNullsFirst      CacheDemoWorkItemOrderBy = " created_at ASC NULLS FIRST "
	CacheDemoWorkItemCreatedAtAscNullsLast       CacheDemoWorkItemOrderBy = " created_at ASC NULLS LAST "
	CacheDemoWorkItemUpdatedAtDescNullsFirst     CacheDemoWorkItemOrderBy = " updated_at DESC NULLS FIRST "
	CacheDemoWorkItemUpdatedAtDescNullsLast      CacheDemoWorkItemOrderBy = " updated_at DESC NULLS LAST "
	CacheDemoWorkItemUpdatedAtAscNullsFirst      CacheDemoWorkItemOrderBy = " updated_at ASC NULLS FIRST "
	CacheDemoWorkItemUpdatedAtAscNullsLast       CacheDemoWorkItemOrderBy = " updated_at ASC NULLS LAST "
	CacheDemoWorkItemDeletedAtDescNullsFirst     CacheDemoWorkItemOrderBy = " deleted_at DESC NULLS FIRST "
	CacheDemoWorkItemDeletedAtDescNullsLast      CacheDemoWorkItemOrderBy = " deleted_at DESC NULLS LAST "
	CacheDemoWorkItemDeletedAtAscNullsFirst      CacheDemoWorkItemOrderBy = " deleted_at ASC NULLS FIRST "
	CacheDemoWorkItemDeletedAtAscNullsLast       CacheDemoWorkItemOrderBy = " deleted_at ASC NULLS LAST "
)

// WithCacheDemoWorkItemOrderBy orders results by the given columns.
func WithCacheDemoWorkItemOrderBy(rows ...CacheDemoWorkItemOrderBy) CacheDemoWorkItemSelectConfigOption {
	return func(s *CacheDemoWorkItemSelectConfig) {
		if len(rows) > 0 {
			orderStrings := make([]string, len(rows))
			for i, row := range rows {
				orderStrings[i] = string(row)
			}
			s.orderBy = " order by "
			s.orderBy += strings.Join(orderStrings, ", ")
		}
	}
}

type CacheDemoWorkItemJoins struct{}

// WithCacheDemoWorkItemJoin joins with the given tables.
func WithCacheDemoWorkItemJoin(joins CacheDemoWorkItemJoins) CacheDemoWorkItemSelectConfigOption {
	return func(s *CacheDemoWorkItemSelectConfig) {
		s.joins = CacheDemoWorkItemJoins{}
	}
}

// WithCacheDemoWorkItemFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithCacheDemoWorkItemFilters(filters map[string][]any) CacheDemoWorkItemSelectConfigOption {
	return func(s *CacheDemoWorkItemSelectConfig) {
		s.filters = filters
	}
}

// WithCacheDemoWorkItemHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithCacheDemoWorkItemHavingClause(conditions map[string][]any) CacheDemoWorkItemSelectConfigOption {
	return func(s *CacheDemoWorkItemSelectConfig) {
		s.having = conditions
	}
}

// CacheDemoWorkItemPaginatedByWorkItemID returns a cursor-paginated list of CacheDemoWorkItem.
func CacheDemoWorkItemPaginatedByWorkItemID(ctx context.Context, db DB, workItemID int, direction models.Direction, opts ...CacheDemoWorkItemSelectConfigOption) ([]CacheDemoWorkItem, error) {
	c := &CacheDemoWorkItemSelectConfig{deletedAt: " null ", joins: CacheDemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	demo_work_items.ref,
	demo_work_items.line,
	demo_work_items.last_message_at,
	demo_work_items.reopened,
	demo_work_items.work_item_id,
	demo_work_items.title,
	demo_work_items.description,
	demo_work_items.work_item_type_id,
	demo_work_items.metadata,
	demo_work_items.team_id,
	demo_work_items.kanban_step_id,
	demo_work_items.closed_at,
	demo_work_items.target_date,
	demo_work_items.created_at,
	demo_work_items.updated_at,
	demo_work_items.deleted_at %s 
	 FROM cache.demo_work_items %s 
	 WHERE demo_work_items.work_item_id %s $1
	 %s   AND demo_work_items.deleted_at is %s  %s 
  %s 
  ORDER BY 
		work_item_id %s `, selects, joins, operator, filters, c.deletedAt, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* CacheDemoWorkItemPaginatedByWorkItemID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoWorkItem/Paginated/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[CacheDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("CacheDemoWorkItem/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	return res, nil
}

// CacheDemoWorkItemByWorkItemID retrieves a row from 'cache.demo_work_items' as a CacheDemoWorkItem.
//
// Generated from index 'cache_demo_work_items_work_item_id_unique'.
func CacheDemoWorkItemByWorkItemID(ctx context.Context, db DB, workItemID int, opts ...CacheDemoWorkItemSelectConfigOption) (*CacheDemoWorkItem, error) {
	c := &CacheDemoWorkItemSelectConfig{deletedAt: " null ", joins: CacheDemoWorkItemJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

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
	demo_work_items.ref,
	demo_work_items.line,
	demo_work_items.last_message_at,
	demo_work_items.reopened,
	demo_work_items.work_item_id,
	demo_work_items.title,
	demo_work_items.description,
	demo_work_items.work_item_type_id,
	demo_work_items.metadata,
	demo_work_items.team_id,
	demo_work_items.kanban_step_id,
	demo_work_items.closed_at,
	demo_work_items.target_date,
	demo_work_items.created_at,
	demo_work_items.updated_at,
	demo_work_items.deleted_at %s 
	 FROM cache.demo_work_items %s 
	 WHERE demo_work_items.work_item_id = $1
	 %s   AND demo_work_items.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* CacheDemoWorkItemByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/db.Query: %w", &XoError{Entity: "Demo work item", Err: err}))
	}
	cdwi, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[CacheDemoWorkItem])
	if err != nil {
		return nil, logerror(fmt.Errorf("demo_work_items/DemoWorkItemByWorkItemID/pgx.CollectOneRow: %w", &XoError{Entity: "Demo work item", Err: err}))
	}

	return &cdwi, nil
}
