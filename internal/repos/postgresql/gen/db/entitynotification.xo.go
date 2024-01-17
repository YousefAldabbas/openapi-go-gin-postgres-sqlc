package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// EntityNotification represents a row from 'public.entity_notifications'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type EntityNotification struct {
	EntityNotificationID EntityNotificationID `json:"entityNotificationID" db:"entity_notification_id" required:"true" nullable:"false"`   // entity_notification_id
	ID                   string               `json:"id" db:"id" required:"true" nullable:"false"`                                         // id
	Message              string               `json:"message" db:"message" required:"true" nullable:"false"`                               // message
	Topic                models.Topics        `json:"topic" db:"topic" required:"true" nullable:"false" ref:"#/components/schemas/Topics"` // topic
	CreatedAt            time.Time            `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                          // created_at

}

// EntityNotificationCreateParams represents insert params for 'public.entity_notifications'.
type EntityNotificationCreateParams struct {
	ID      string        `json:"id" required:"true" nullable:"false"`                                      // id
	Message string        `json:"message" required:"true" nullable:"false"`                                 // message
	Topic   models.Topics `json:"topic" required:"true" nullable:"false" ref:"#/components/schemas/Topics"` // topic
}

type EntityNotificationID int

// CreateEntityNotification creates a new EntityNotification in the database with the given params.
func CreateEntityNotification(ctx context.Context, db DB, params *EntityNotificationCreateParams) (*EntityNotification, error) {
	en := &EntityNotification{
		ID:      params.ID,
		Message: params.Message,
		Topic:   params.Topic,
	}

	return en.Insert(ctx, db)
}

type EntityNotificationSelectConfig struct {
	limit   string
	orderBy string
	joins   EntityNotificationJoins
	filters map[string][]any
	having  map[string][]any
}
type EntityNotificationSelectConfigOption func(*EntityNotificationSelectConfig)

// WithEntityNotificationLimit limits row selection.
func WithEntityNotificationLimit(limit int) EntityNotificationSelectConfigOption {
	return func(s *EntityNotificationSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type EntityNotificationOrderBy string

const (
	EntityNotificationCreatedAtDescNullsFirst EntityNotificationOrderBy = " created_at DESC NULLS FIRST "
	EntityNotificationCreatedAtDescNullsLast  EntityNotificationOrderBy = " created_at DESC NULLS LAST "
	EntityNotificationCreatedAtAscNullsFirst  EntityNotificationOrderBy = " created_at ASC NULLS FIRST "
	EntityNotificationCreatedAtAscNullsLast   EntityNotificationOrderBy = " created_at ASC NULLS LAST "
)

// WithEntityNotificationOrderBy orders results by the given columns.
func WithEntityNotificationOrderBy(rows ...EntityNotificationOrderBy) EntityNotificationSelectConfigOption {
	return func(s *EntityNotificationSelectConfig) {
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

type EntityNotificationJoins struct {
}

// WithEntityNotificationJoin joins with the given tables.
func WithEntityNotificationJoin(joins EntityNotificationJoins) EntityNotificationSelectConfigOption {
	return func(s *EntityNotificationSelectConfig) {
		s.joins = EntityNotificationJoins{}
	}
}

// WithEntityNotificationFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithEntityNotificationFilters(filters map[string][]any) EntityNotificationSelectConfigOption {
	return func(s *EntityNotificationSelectConfig) {
		s.filters = filters
	}
}

// WithEntityNotificationHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithEntityNotificationHavingClause(conditions map[string][]any) EntityNotificationSelectConfigOption {
	return func(s *EntityNotificationSelectConfig) {
		s.having = conditions
	}
}

// EntityNotificationUpdateParams represents update params for 'public.entity_notifications'.
type EntityNotificationUpdateParams struct {
	ID      *string        `json:"id" nullable:"false"`                                      // id
	Message *string        `json:"message" nullable:"false"`                                 // message
	Topic   *models.Topics `json:"topic" nullable:"false" ref:"#/components/schemas/Topics"` // topic
}

// SetUpdateParams updates public.entity_notifications struct fields with the specified params.
func (en *EntityNotification) SetUpdateParams(params *EntityNotificationUpdateParams) {
	if params.ID != nil {
		en.ID = *params.ID
	}
	if params.Message != nil {
		en.Message = *params.Message
	}
	if params.Topic != nil {
		en.Topic = *params.Topic
	}
}

// Insert inserts the EntityNotification to the database.
func (en *EntityNotification) Insert(ctx context.Context, db DB) (*EntityNotification, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.entity_notifications (
	id, message, topic
	) VALUES (
	$1, $2, $3
	) RETURNING * `
	// run
	logf(sqlstr, en.ID, en.Message, en.Topic)

	rows, err := db.Query(ctx, sqlstr, en.ID, en.Message, en.Topic)
	if err != nil {
		return nil, logerror(fmt.Errorf("EntityNotification/Insert/db.Query: %w", &XoError{Entity: "Entity notification", Err: err}))
	}
	newen, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[EntityNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("EntityNotification/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Entity notification", Err: err}))
	}

	*en = newen

	return en, nil
}

// Update updates a EntityNotification in the database.
func (en *EntityNotification) Update(ctx context.Context, db DB) (*EntityNotification, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.entity_notifications SET 
	id = $1, message = $2, topic = $3 
	WHERE entity_notification_id = $4 
	RETURNING * `
	// run
	logf(sqlstr, en.ID, en.Message, en.Topic, en.EntityNotificationID)

	rows, err := db.Query(ctx, sqlstr, en.ID, en.Message, en.Topic, en.EntityNotificationID)
	if err != nil {
		return nil, logerror(fmt.Errorf("EntityNotification/Update/db.Query: %w", &XoError{Entity: "Entity notification", Err: err}))
	}
	newen, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[EntityNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("EntityNotification/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Entity notification", Err: err}))
	}
	*en = newen

	return en, nil
}

// Upsert upserts a EntityNotification in the database.
// Requires appropriate PK(s) to be set beforehand.
func (en *EntityNotification) Upsert(ctx context.Context, db DB, params *EntityNotificationCreateParams) (*EntityNotification, error) {
	var err error

	en.ID = params.ID
	en.Message = params.Message
	en.Topic = params.Topic

	en, err = en.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Entity notification", Err: err})
			}
			en, err = en.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Entity notification", Err: err})
			}
		}
	}

	return en, err
}

// Delete deletes the EntityNotification from the database.
func (en *EntityNotification) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.entity_notifications 
	WHERE entity_notification_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, en.EntityNotificationID); err != nil {
		return logerror(err)
	}
	return nil
}

// EntityNotificationPaginatedByEntityNotificationID returns a cursor-paginated list of EntityNotification.
func EntityNotificationPaginatedByEntityNotificationID(ctx context.Context, db DB, entityNotificationID EntityNotificationID, direction models.Direction, opts ...EntityNotificationSelectConfigOption) ([]EntityNotification, error) {
	c := &EntityNotificationSelectConfig{joins: EntityNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
	entity_notifications.created_at,
	entity_notifications.entity_notification_id,
	entity_notifications.id,
	entity_notifications.message,
	entity_notifications.topic %s 
	 FROM public.entity_notifications %s 
	 WHERE entity_notifications.entity_notification_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		entity_notification_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* EntityNotificationPaginatedByEntityNotificationID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{entityNotificationID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("EntityNotification/Paginated/db.Query: %w", &XoError{Entity: "Entity notification", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[EntityNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("EntityNotification/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Entity notification", Err: err}))
	}
	return res, nil
}

// EntityNotificationByEntityNotificationID retrieves a row from 'public.entity_notifications' as a EntityNotification.
//
// Generated from index 'entity_notifications_pkey'.
func EntityNotificationByEntityNotificationID(ctx context.Context, db DB, entityNotificationID EntityNotificationID, opts ...EntityNotificationSelectConfigOption) (*EntityNotification, error) {
	c := &EntityNotificationSelectConfig{joins: EntityNotificationJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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
	entity_notifications.created_at,
	entity_notifications.entity_notification_id,
	entity_notifications.id,
	entity_notifications.message,
	entity_notifications.topic %s 
	 FROM public.entity_notifications %s 
	 WHERE entity_notifications.entity_notification_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* EntityNotificationByEntityNotificationID */\n" + sqlstr

	// run
	// logf(sqlstr, entityNotificationID)
	rows, err := db.Query(ctx, sqlstr, append([]any{entityNotificationID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("entity_notifications/EntityNotificationByEntityNotificationID/db.Query: %w", &XoError{Entity: "Entity notification", Err: err}))
	}
	en, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[EntityNotification])
	if err != nil {
		return nil, logerror(fmt.Errorf("entity_notifications/EntityNotificationByEntityNotificationID/pgx.CollectOneRow: %w", &XoError{Entity: "Entity notification", Err: err}))
	}

	return &en, nil
}
