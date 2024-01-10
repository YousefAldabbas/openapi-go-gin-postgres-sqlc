package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// ExtraSchemaBook represents a row from 'extra_schema.books'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private to exclude a field from JSON.
//     -- not-required to make a schema field not required.
//     -- hidden to exclude field from OpenAPI generation.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type ExtraSchemaBook struct {
	BookID ExtraSchemaBookID `json:"bookID" db:"book_id" required:"true" nullable:"false"` // book_id
	Name   string            `json:"name" db:"name" required:"true" nullable:"false"`      // name

	BookAuthorsJoin     *[]User__BA_ExtraSchemaBook   `json:"-" db:"book_authors_authors" openapi-go:"ignore"`               // M2M book_authors
	BookAuthorsJoinBASK *[]User__BASK_ExtraSchemaBook `json:"-" db:"book_authors_surrogate_key_authors" openapi-go:"ignore"` // M2M book_authors_surrogate_key
	BookBookReviewsJoin *[]ExtraSchemaBookReview      `json:"-" db:"book_reviews" openapi-go:"ignore"`                       // M2O books
	BookSellersJoin     *[]ExtraSchemaUser            `json:"-" db:"book_sellers_sellers" openapi-go:"ignore"`               // M2M book_sellers

}

// ExtraSchemaBookCreateParams represents insert params for 'extra_schema.books'.
type ExtraSchemaBookCreateParams struct {
	Name string `json:"name" required:"true" nullable:"false"` // name
}

type ExtraSchemaBookID int

// CreateExtraSchemaBook creates a new ExtraSchemaBook in the database with the given params.
func CreateExtraSchemaBook(ctx context.Context, db DB, params *ExtraSchemaBookCreateParams) (*ExtraSchemaBook, error) {
	esb := &ExtraSchemaBook{
		Name: params.Name,
	}

	return esb.Insert(ctx, db)
}

// ExtraSchemaBookUpdateParams represents update params for 'extra_schema.books'.
type ExtraSchemaBookUpdateParams struct {
	Name *string `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates extra_schema.books struct fields with the specified params.
func (esb *ExtraSchemaBook) SetUpdateParams(params *ExtraSchemaBookUpdateParams) {
	if params.Name != nil {
		esb.Name = *params.Name
	}
}

type ExtraSchemaBookSelectConfig struct {
	limit   string
	orderBy string
	joins   ExtraSchemaBookJoins
	filters map[string][]any
	having  map[string][]any
}
type ExtraSchemaBookSelectConfigOption func(*ExtraSchemaBookSelectConfig)

// WithExtraSchemaBookLimit limits row selection.
func WithExtraSchemaBookLimit(limit int) ExtraSchemaBookSelectConfigOption {
	return func(s *ExtraSchemaBookSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type ExtraSchemaBookOrderBy string

const ()

type ExtraSchemaBookJoins struct {
	AuthorsBook      bool // M2M book_authors
	AuthorsBookUsers bool // M2M book_authors_surrogate_key
	BookReviews      bool // M2O book_reviews
	Sellers          bool // M2M book_sellers
}

// WithExtraSchemaBookJoin joins with the given tables.
func WithExtraSchemaBookJoin(joins ExtraSchemaBookJoins) ExtraSchemaBookSelectConfigOption {
	return func(s *ExtraSchemaBookSelectConfig) {
		s.joins = ExtraSchemaBookJoins{
			AuthorsBook:      s.joins.AuthorsBook || joins.AuthorsBook,
			AuthorsBookUsers: s.joins.AuthorsBookUsers || joins.AuthorsBookUsers,
			BookReviews:      s.joins.BookReviews || joins.BookReviews,
			Sellers:          s.joins.Sellers || joins.Sellers,
		}
	}
}

// User__BA_ExtraSchemaBook represents a M2M join against "extra_schema.book_authors"
type User__BA_ExtraSchemaBook struct {
	User      ExtraSchemaUser `json:"user" db:"users" required:"true"`
	Pseudonym *string         `json:"pseudonym" db:"pseudonym" required:"true" `
}

// User__BASK_ExtraSchemaBook represents a M2M join against "extra_schema.book_authors_surrogate_key"
type User__BASK_ExtraSchemaBook struct {
	User      ExtraSchemaUser `json:"user" db:"users" required:"true"`
	Pseudonym *string         `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithExtraSchemaBookFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithExtraSchemaBookFilters(filters map[string][]any) ExtraSchemaBookSelectConfigOption {
	return func(s *ExtraSchemaBookSelectConfig) {
		s.filters = filters
	}
}

// WithExtraSchemaBookHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(assigned_users_join.user_id))": {userId},
//	}
func WithExtraSchemaBookHavingClause(conditions map[string][]any) ExtraSchemaBookSelectConfigOption {
	return func(s *ExtraSchemaBookSelectConfig) {
		s.having = conditions
	}
}

const extraSchemaBookTableAuthorsBookJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
		book_authors.book_id as book_authors_book_id
		, book_authors.pseudonym as pseudonym
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.book_authors
	join extra_schema.users on users.user_id = book_authors.author_id
	group by
		book_authors_book_id
		, users.user_id
		, pseudonym
) as joined_book_authors_authors on joined_book_authors_authors.book_authors_book_id = books.book_id
`

const extraSchemaBookTableAuthorsBookSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_authors.__users
		, joined_book_authors_authors.pseudonym
		)) filter (where joined_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const extraSchemaBookTableAuthorsBookGroupBySQL = `books.book_id, books.book_id`

const extraSchemaBookTableAuthorsBookUsersJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
left join (
	select
		book_authors_surrogate_key.book_id as book_authors_surrogate_key_book_id
		, book_authors_surrogate_key.pseudonym as pseudonym
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.book_authors_surrogate_key
	join extra_schema.users on users.user_id = book_authors_surrogate_key.author_id
	group by
		book_authors_surrogate_key_book_id
		, users.user_id
		, pseudonym
) as joined_book_authors_surrogate_key_authors on joined_book_authors_surrogate_key_authors.book_authors_surrogate_key_book_id = books.book_id
`

const extraSchemaBookTableAuthorsBookUsersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_authors.__users
		, joined_book_authors_surrogate_key_authors.pseudonym
		)) filter (where joined_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') as book_authors_surrogate_key_authors`

const extraSchemaBookTableAuthorsBookUsersGroupBySQL = `books.book_id, books.book_id`

const extraSchemaBookTableBookReviewsJoinSQL = `-- M2O join generated from "book_reviews_book_id_fkey"
left join (
  select
  book_id as book_reviews_book_id
    , array_agg(book_reviews.*) as book_reviews
  from
    extra_schema.book_reviews
  group by
        book_id
) as joined_book_reviews on joined_book_reviews.book_reviews_book_id = books.book_id
`

const extraSchemaBookTableBookReviewsSelectSQL = `COALESCE(joined_book_reviews.book_reviews, '{}') as book_reviews`

const extraSchemaBookTableBookReviewsGroupBySQL = `joined_book_reviews.book_reviews, books.book_id`

const extraSchemaBookTableSellersJoinSQL = `-- M2M join generated from "book_sellers_seller_fkey"
left join (
	select
		book_sellers.book_id as book_sellers_book_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		extra_schema.book_sellers
	join extra_schema.users on users.user_id = book_sellers.seller
	group by
		book_sellers_book_id
		, users.user_id
) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = books.book_id
`

const extraSchemaBookTableSellersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users_user_id is not null), '{}') as book_sellers_sellers`

const extraSchemaBookTableSellersGroupBySQL = `books.book_id, books.book_id`

// Insert inserts the ExtraSchemaBook to the database.
func (esb *ExtraSchemaBook) Insert(ctx context.Context, db DB) (*ExtraSchemaBook, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO extra_schema.books (
	name
	) VALUES (
	$1
	) RETURNING * `
	// run
	logf(sqlstr, esb.Name)

	rows, err := db.Query(ctx, sqlstr, esb.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBook/Insert/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	newesb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBook/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}

	*esb = newesb

	return esb, nil
}

// Update updates a ExtraSchemaBook in the database.
func (esb *ExtraSchemaBook) Update(ctx context.Context, db DB) (*ExtraSchemaBook, error) {
	// update with composite primary key
	sqlstr := `UPDATE extra_schema.books SET 
	name = $1 
	WHERE book_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, esb.Name, esb.BookID)

	rows, err := db.Query(ctx, sqlstr, esb.Name, esb.BookID)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBook/Update/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	newesb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBook/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}
	*esb = newesb

	return esb, nil
}

// Upsert upserts a ExtraSchemaBook in the database.
// Requires appropriate PK(s) to be set beforehand.
func (esb *ExtraSchemaBook) Upsert(ctx context.Context, db DB, params *ExtraSchemaBookCreateParams) (*ExtraSchemaBook, error) {
	var err error

	esb.Name = params.Name

	esb, err = esb.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Book", Err: err})
			}
			esb, err = esb.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Book", Err: err})
			}
		}
	}

	return esb, err
}

// Delete deletes the ExtraSchemaBook from the database.
func (esb *ExtraSchemaBook) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM extra_schema.books 
	WHERE book_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, esb.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// ExtraSchemaBookPaginatedByBookID returns a cursor-paginated list of ExtraSchemaBook.
func ExtraSchemaBookPaginatedByBookID(ctx context.Context, db DB, bookID ExtraSchemaBookID, direction models.Direction, opts ...ExtraSchemaBookSelectConfigOption) ([]ExtraSchemaBook, error) {
	c := &ExtraSchemaBookSelectConfig{joins: ExtraSchemaBookJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, extraSchemaBookTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableAuthorsBookGroupBySQL)
	}

	if c.joins.AuthorsBookUsers {
		selectClauses = append(selectClauses, extraSchemaBookTableAuthorsBookUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableAuthorsBookUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableAuthorsBookUsersGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, extraSchemaBookTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableBookReviewsGroupBySQL)
	}

	if c.joins.Sellers {
		selectClauses = append(selectClauses, extraSchemaBookTableSellersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableSellersGroupBySQL)
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

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	books.book_id,
	books.name %s 
	 FROM extra_schema.books %s 
	 WHERE books.book_id %s $1
	 %s   %s 
  %s 
  ORDER BY 
		book_id %s `, selects, joins, operator, filters, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookPaginatedByBookID */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBook/Paginated/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[ExtraSchemaBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("ExtraSchemaBook/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book", Err: err}))
	}
	return res, nil
}

// ExtraSchemaBookByBookID retrieves a row from 'extra_schema.books' as a ExtraSchemaBook.
//
// Generated from index 'books_pkey'.
func ExtraSchemaBookByBookID(ctx context.Context, db DB, bookID ExtraSchemaBookID, opts ...ExtraSchemaBookSelectConfigOption) (*ExtraSchemaBook, error) {
	c := &ExtraSchemaBookSelectConfig{joins: ExtraSchemaBookJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.AuthorsBook {
		selectClauses = append(selectClauses, extraSchemaBookTableAuthorsBookSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableAuthorsBookJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableAuthorsBookGroupBySQL)
	}

	if c.joins.AuthorsBookUsers {
		selectClauses = append(selectClauses, extraSchemaBookTableAuthorsBookUsersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableAuthorsBookUsersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableAuthorsBookUsersGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, extraSchemaBookTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableBookReviewsGroupBySQL)
	}

	if c.joins.Sellers {
		selectClauses = append(selectClauses, extraSchemaBookTableSellersSelectSQL)
		joinClauses = append(joinClauses, extraSchemaBookTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, extraSchemaBookTableSellersGroupBySQL)
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
	books.book_id,
	books.name %s 
	 FROM extra_schema.books %s 
	 WHERE books.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* ExtraSchemaBookByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	esb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[ExtraSchemaBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}

	return &esb, nil
}
