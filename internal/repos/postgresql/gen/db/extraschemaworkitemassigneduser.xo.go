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

// ExtraSchemaWorkItemAssignedUser represents a row from 'extra_schema.work_item_assigned_user'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaWorkItemAssignedUser struct {
	WorkItemID      ExtraSchemaWorkItemID    `json:"workItemID" db:"work_item_id" required:"true" nullable:"false"`                           // work_item_id
	AssignedUser    ExtraSchemaUserID        `json:"assignedUser" db:"assigned_user" required:"true" nullable:"false"`                        // assigned_user
	ExtraSchemaRole *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role

	AssignedUserWorkItemsJoin *[]WorkItem__WIAU_ExtraSchemaWorkItemAssignedUser `json:"-" db:"work_item_assigned_user_work_items" openapi-go:"ignore"`     // M2M work_item_assigned_user
	WorkItemAssignedUsersJoin *[]User__WIAU_ExtraSchemaWorkItemAssignedUser     `json:"-" db:"work_item_assigned_user_assigned_users" openapi-go:"ignore"` // M2M work_item_assigned_user

}

// ExtraSchemaWorkItemAssignedUserCreateParams represents insert params for 'extra_schema.work_item_assigned_user'.
type ExtraSchemaWorkItemAssignedUserCreateParams struct {
	AssignedUser    ExtraSchemaUserID        `json:"assignedUser" required:"true" nullable:"false"`                                 // assigned_user
	ExtraSchemaRole *ExtraSchemaWorkItemRole `json:"role" required:"true" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID      ExtraSchemaWorkItemID    `json:"workItemID" required:"true" nullable:"false"`                                   // work_item_id
}

// CreateExtraSchemaWorkItemAssignedUser creates a new ExtraSchemaWorkItemAssignedUser in the database with the given params.
func CreateExtraSchemaWorkItemAssignedUser(ctx context.Context, db DB, params *ExtraSchemaWorkItemAssignedUserCreateParams) (*ExtraSchemaWorkItemAssignedUser, error) {
	eswiau := &ExtraSchemaWorkItemAssignedUser{
		AssignedUser:    params.AssignedUser,
		ExtraSchemaRole: params.ExtraSchemaRole,
		WorkItemID:      params.WorkItemID,
	}

	return eswiau.Insert(ctx, db)
}

type ExtraSchemaWorkItemAssignedUserSelectConfig struct {
	limit   string
	orderBy string
	joins   ExtraSchemaWorkItemAssignedUserJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaWorkItemAssignedUserSelectConfigOption func(*ExtraSchemaWorkItemAssignedUserSelectConfig)

// WithExtraSchemaWorkItemAssignedUserLimit limits row selection.
func WithExtraSchemaWorkItemAssignedUserLimit(limit int) ExtraSchemaWorkItemAssignedUserSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssignedUserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ExtraSchemaWorkItemAssignedUserOrderBy string

const ()

type ExtraSchemaWorkItemAssignedUserJoins struct {
	WorkItemsAssignedUser bool // M2M work_item_assigned_user
	AssignedUsers         bool // M2M work_item_assigned_user
}

// WithExtraSchemaWorkItemAssignedUserJoin joins with the given tables.
func WithExtraSchemaWorkItemAssignedUserJoin(joins ExtraSchemaWorkItemAssignedUserJoins) ExtraSchemaWorkItemAssignedUserSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssignedUserSelectConfig) {
		s.joins = ExtraSchemaWorkItemAssignedUserJoins{
			WorkItemsAssignedUser: s.joins.WorkItemsAssignedUser || joins.WorkItemsAssignedUser,
			AssignedUsers:         s.joins.AssignedUsers || joins.AssignedUsers,
		}
	}
}

// WorkItem__WIAU_ExtraSchemaWorkItemAssignedUser represents a M2M join against "extra_schema.work_item_assigned_user"
type WorkItem__WIAU_ExtraSchemaWorkItemAssignedUser struct {
	WorkItem ExtraSchemaWorkItem      `json:"workItem" db:"work_items" required:"true"`
	Role     *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// User__WIAU_ExtraSchemaWorkItemAssignedUser represents a M2M join against "extra_schema.work_item_assigned_user"
type User__WIAU_ExtraSchemaWorkItemAssignedUser struct {
	User ExtraSchemaUser          `json:"user" db:"users" required:"true"`
	Role *ExtraSchemaWorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithExtraSchemaWorkItemAssignedUserFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaWorkItemAssignedUserFilters(filters map[string][]any) ExtraSchemaWorkItemAssignedUserSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssignedUserSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaWorkItemAssignedUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithExtraSchemaWorkItemAssignedUserHavingClause(conditions map[string][]any) ExtraSchemaWorkItemAssignedUserSelectConfigOption {
	return func(s *ExtraSchemaWorkItemAssignedUserSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL = `-- M2M join generated from "work_item_assigned_user_work_item_id_fkey"
left join (
	select
		work_item_assigned_user.assigned_user as work_item_assigned_user_assigned_user
		, work_item_assigned_user.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		extra_schema.work_item_assigned_user
	join extra_schema.work_items on work_items.work_item_id = work_item_assigned_user.work_item_id
	group by
		work_item_assigned_user_assigned_user
		, work_items.work_item_id
		, role
) as joined_work_item_assigned_user_work_items on joined_work_item_assigned_user_work_items.work_item_assigned_user_assigned_user = work_item_assigned_user.assigned_user
`

const extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_work_items.__work_items
		, joined_work_item_assigned_user_work_items.role
		)) filter (where joined_work_item_assigned_user_work_items.__work_items_work_item_id is not null), '{}') as work_item_assigned_user_work_items`

const extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL = `work_item_assigned_user.assigned_user, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user`

const extraSchemaWorkItemAssignedUserTableAssignedUsersJoinSQL = `-- M2M join generated from "work_item_assigned_user_assigned_user_fkey"
left join (
	select
		work_item_assigned_user.work_item_id as work_item_assigned_user_work_item_id
		, work_item_assigned_user.role as role
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.work_item_assigned_user
	join extra_schema.users on users.user_id = work_item_assigned_user.assigned_user
	group by
		work_item_assigned_user_work_item_id
		, users.user_id
		, role
) as joined_work_item_assigned_user_assigned_users on joined_work_item_assigned_user_assigned_users.work_item_assigned_user_work_item_id = work_item_assigned_user.work_item_id
`

const extraSchemaWorkItemAssignedUserTableAssignedUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_work_item_assigned_user_assigned_users.__users
		, joined_work_item_assigned_user_assigned_users.role
		)) filter (where joined_work_item_assigned_user_assigned_users.__users_user_id is not null), '{}') as work_item_assigned_user_assigned_users`

const extraSchemaWorkItemAssignedUserTableAssignedUsersGroupBySQL = `work_item_assigned_user.work_item_id, work_item_assigned_user.work_item_id, work_item_assigned_user.assigned_user`

// ExtraSchemaWorkItemAssignedUserUpdateParams represents update params for 'extra_schema.work_item_assigned_user'.
type ExtraSchemaWorkItemAssignedUserUpdateParams struct {
	AssignedUser    *ExtraSchemaUserID        `json:"assignedUser" nullable:"false"`                                 // assigned_user
	ExtraSchemaRole **ExtraSchemaWorkItemRole `json:"role" nullable:"false" ref:"#/components/schemas/WorkItemRole"` // role
	WorkItemID      *ExtraSchemaWorkItemID    `json:"workItemID" nullable:"false"`                                   // work_item_id
}

// SetUpdateParams updates extra_schema.work_item_assigned_user struct fields with the specified params.
func (eswiau *ExtraSchemaWorkItemAssignedUser) SetUpdateParams(params *ExtraSchemaWorkItemAssignedUserUpdateParams) {
	if params.AssignedUser != nil {
		eswiau.AssignedUser = *params.AssignedUser
	}
	if params.ExtraSchemaRole != nil {
		eswiau.ExtraSchemaRole = *params.ExtraSchemaRole
	}
	if params.WorkItemID != nil {
		eswiau.WorkItemID = *params.WorkItemID
	}
}

// Insert inserts the ExtraSchemaWorkItemAssignedUser to the database.
func (eswiau *ExtraSchemaWorkItemAssignedUser) Insert(ctx context.Context, db DB) (*ExtraSchemaWorkItemAssignedUser, error) {
	// insert (manual)
	sqlstr := `INSERT INTO extra_schema.work_item_assigned_user (
	assigned_user, role, work_item_id
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, eswiau.AssignedUser, eswiau.ExtraSchemaRole, eswiau.WorkItemID)
	rows, err := db.Query(ctx, sqlstr, eswiau.AssignedUser, eswiau.ExtraSchemaRole, eswiau.WorkItemID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/Insert/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	neweswiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	*eswiau = neweswiau

	return eswiau, nil
}

// Update updates a ExtraSchemaWorkItemAssignedUser in the database.
func (eswiau *ExtraSchemaWorkItemAssignedUser) Update(ctx context.Context, db DB) (*ExtraSchemaWorkItemAssignedUser, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.work_item_assigned_user SET 
	role = $1 
	WHERE work_item_id = $2  AND assigned_user = $3 
	RETURNING * `
	// run
	logf(sqlstr, eswiau.ExtraSchemaRole, eswiau.WorkItemID, eswiau.AssignedUser)

	rows, err := db.Query(ctx, sqlstr, eswiau.ExtraSchemaRole, eswiau.WorkItemID, eswiau.AssignedUser)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/Update/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	neweswiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	*eswiau = neweswiau

	return eswiau, nil
}

// Upsert upserts a ExtraSchemaWorkItemAssignedUser in the database.
// Requires appropriate PK(s) to be set beforehand.
func (eswiau *ExtraSchemaWorkItemAssignedUser) Upsert(ctx context.Context, db DB, params *ExtraSchemaWorkItemAssignedUserCreateParams) (*ExtraSchemaWorkItemAssignedUser, error) {
	var err error

	eswiau.AssignedUser = params.AssignedUser
	eswiau.ExtraSchemaRole = params.ExtraSchemaRole
	eswiau.WorkItemID = params.WorkItemID

	eswiau, err = eswiau.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Work item assigned user", Err: err})
			}
			eswiau, err = eswiau.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Work item assigned user", Err: err})
			}
		}
	}

	return eswiau, err
}

// Delete deletes the ExtraSchemaWorkItemAssignedUser from the database.
func (eswiau *ExtraSchemaWorkItemAssignedUser) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM extra_schema.work_item_assigned_user 
	WHERE work_item_id = $1 AND assigned_user = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, eswiau.WorkItemID, eswiau.AssignedUser); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaWorkItemAssignedUsersByAssignedUserWorkItemID retrieves a row from 'extra_schema.work_item_assigned_user' as a ExtraSchemaWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_assigned_user_work_item_id_idx'.
func ExtraSchemaWorkItemAssignedUsersByAssignedUserWorkItemID(ctx context.Context, db DB, assignedUser ExtraSchemaUserID, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemAssignedUserSelectConfigOption) ([]ExtraSchemaWorkItemAssignedUser, error) {
	c := &ExtraSchemaWorkItemAssignedUserSelectConfig{joins: ExtraSchemaWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM extra_schema.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.assigned_user = $1 AND work_item_assigned_user.work_item_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssignedUsersByAssignedUserWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, assignedUser, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignedUser, workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/WorkItemAssignedUserByAssignedUserWorkItemID/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemAssignedUserByWorkItemIDAssignedUser retrieves a row from 'extra_schema.work_item_assigned_user' as a ExtraSchemaWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func ExtraSchemaWorkItemAssignedUserByWorkItemIDAssignedUser(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, assignedUser ExtraSchemaUserID, opts ...ExtraSchemaWorkItemAssignedUserSelectConfigOption) (*ExtraSchemaWorkItemAssignedUser, error) {
	c := &ExtraSchemaWorkItemAssignedUserSelectConfig{joins: ExtraSchemaWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM extra_schema.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.work_item_id = $1 AND work_item_assigned_user.assigned_user = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssignedUserByWorkItemIDAssignedUser */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID, assignedUser)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID, assignedUser}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/db.Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	eswiau, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("work_item_assigned_user/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectOneRow: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}

	return &eswiau, nil
}

// ExtraSchemaWorkItemAssignedUsersByWorkItemID retrieves a row from 'extra_schema.work_item_assigned_user' as a ExtraSchemaWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func ExtraSchemaWorkItemAssignedUsersByWorkItemID(ctx context.Context, db DB, workItemID ExtraSchemaWorkItemID, opts ...ExtraSchemaWorkItemAssignedUserSelectConfigOption) ([]ExtraSchemaWorkItemAssignedUser, error) {
	c := &ExtraSchemaWorkItemAssignedUserSelectConfig{joins: ExtraSchemaWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM extra_schema.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.work_item_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssignedUsersByWorkItemID */\n" + sqlstr

	// run
	// logf(sqlstr, workItemID)
	rows, err := db.Query(ctx, sqlstr, append([]any{workItemID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// ExtraSchemaWorkItemAssignedUsersByAssignedUser retrieves a row from 'extra_schema.work_item_assigned_user' as a ExtraSchemaWorkItemAssignedUser.
//
// Generated from index 'work_item_assigned_user_pkey'.
func ExtraSchemaWorkItemAssignedUsersByAssignedUser(ctx context.Context, db DB, assignedUser ExtraSchemaUserID, opts ...ExtraSchemaWorkItemAssignedUserSelectConfigOption) ([]ExtraSchemaWorkItemAssignedUser, error) {
	c := &ExtraSchemaWorkItemAssignedUserSelectConfig{joins: ExtraSchemaWorkItemAssignedUserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.WorkItemsAssignedUser {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableWorkItemsAssignedUserGroupBySQL)
	}

	if c.joins.AssignedUsers {
		selectClauses = append(selectClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaWorkItemAssignedUserTableAssignedUsersGroupBySQL)
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
	work_item_assigned_user.assigned_user,
	work_item_assigned_user.role,
	work_item_assigned_user.work_item_id %s 
	 FROM extra_schema.work_item_assigned_user %s 
	 WHERE work_item_assigned_user.assigned_user = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaWorkItemAssignedUsersByAssignedUser */\n" + sqlstr

	// run
	// logf(sqlstr, assignedUser)
	rows, err := db.Query(ctx, sqlstr, append([]any{assignedUser}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/Query: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaWorkItemAssignedUser])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaWorkItemAssignedUser/WorkItemAssignedUserByWorkItemIDAssignedUser/pgx.CollectRows: %w", &XoError{Entity: "Work item assigned user", Err: err}))
	}
	return res, nil
}

// FKUser_AssignedUser returns the User associated with the ExtraSchemaWorkItemAssignedUser's (AssignedUser).
//
// Generated from foreign key 'work_item_assigned_user_assigned_user_fkey'.
func (eswiau *ExtraSchemaWorkItemAssignedUser) FKUser_AssignedUser(ctx context.Context, db DB) (*ExtraSchemaUser, error) {
	return ExtraSchemaUserByUserID(ctx, db, eswiau.AssignedUser)
}

// FKWorkItem_WorkItemID returns the WorkItem associated with the ExtraSchemaWorkItemAssignedUser's (WorkItemID).
//
// Generated from foreign key 'work_item_assigned_user_work_item_id_fkey'.
func (eswiau *ExtraSchemaWorkItemAssignedUser) FKWorkItem_WorkItemID(ctx context.Context, db DB) (*ExtraSchemaWorkItem, error) {
	return ExtraSchemaWorkItemByWorkItemID(ctx, db, eswiau.WorkItemID)
}
