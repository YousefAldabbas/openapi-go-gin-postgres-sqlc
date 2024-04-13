package got

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

// XoTestsBook represents a row from 'xo_tests.books'.
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
type XoTestsBook struct {
	BookID XoTestsBookID `json:"bookID" db:"book_id" required:"true" nullable:"false"` // book_id
	Name   string        `json:"name" db:"name" required:"true" nullable:"false"`      // name

	AuthorsJoin     *[]XoTestsBookM2MAuthorBA   `json:"-" db:"book_authors_authors" openapi-go:"ignore"`               // M2M book_authors
	AuthorsBASKJoin *[]XoTestsBookM2MAuthorBASK `json:"-" db:"book_authors_surrogate_key_authors" openapi-go:"ignore"` // M2M book_authors_surrogate_key
	BookReviewsJoin *[]XoTestsBookReview        `json:"-" db:"book_reviews" openapi-go:"ignore"`                       // M2O books
	SellersJoin     *[]XoTestsUser              `json:"-" db:"book_sellers_sellers" openapi-go:"ignore"`               // M2M book_sellers
}

// XoTestsBookCreateParams represents insert params for 'xo_tests.books'.
type XoTestsBookCreateParams struct {
	Name string `json:"name" required:"true" nullable:"false"` // name
}

// XoTestsBookParams represents common params for both insert and update of 'xo_tests.books'.
type XoTestsBookParams interface {
	GetName() *string
}

func (p XoTestsBookCreateParams) GetName() *string {
	x := p.Name
	return &x
}

func (p XoTestsBookUpdateParams) GetName() *string {
	return p.Name
}

type XoTestsBookID int

// CreateXoTestsBook creates a new XoTestsBook in the database with the given params.
func CreateXoTestsBook(ctx context.Context, db DB, params *XoTestsBookCreateParams) (*XoTestsBook, error) {
	xtb := &XoTestsBook{
		Name: params.Name,
	}

	return xtb.Insert(ctx, db)
}

type XoTestsBookSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   XoTestsBookJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsBookSelectConfigOption func(*XoTestsBookSelectConfig)

// WithXoTestsBookLimit limits row selection.
func WithXoTestsBookLimit(limit int) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithXoTestsBookOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithXoTestsBookOrderBy(rows map[string]*models.Direction) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		te := XoTestsEntityFields[XoTestsTableEntityXoTestsBook]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}

type XoTestsBookJoins struct {
	Authors     bool `json:"authors" required:"true" nullable:"false"`     // M2M book_authors
	AuthorsBASK bool `json:"authorsBASK" required:"true" nullable:"false"` // M2M book_authors_surrogate_key
	BookReviews bool `json:"bookReviews" required:"true" nullable:"false"` // M2O book_reviews
	Sellers     bool `json:"sellers" required:"true" nullable:"false"`     // M2M book_sellers
}

// WithXoTestsBookJoin joins with the given tables.
func WithXoTestsBookJoin(joins XoTestsBookJoins) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		s.joins = XoTestsBookJoins{
			Authors:     s.joins.Authors || joins.Authors,
			AuthorsBASK: s.joins.AuthorsBASK || joins.AuthorsBASK,
			BookReviews: s.joins.BookReviews || joins.BookReviews,
			Sellers:     s.joins.Sellers || joins.Sellers,
		}
	}
}

// XoTestsBookM2MAuthorBA represents a M2M join against "xo_tests.book_authors"
type XoTestsBookM2MAuthorBA struct {
	User      XoTestsUser `json:"user" db:"users" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// XoTestsBookM2MAuthorBASK represents a M2M join against "xo_tests.book_authors_surrogate_key"
type XoTestsBookM2MAuthorBASK struct {
	User      XoTestsUser `json:"user" db:"users" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithXoTestsBookFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsBookFilters(filters map[string][]any) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsBookHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
//	}
func WithXoTestsBookHavingClause(conditions map[string][]any) XoTestsBookSelectConfigOption {
	return func(s *XoTestsBookSelectConfig) {
		s.having = conditions
	}
}

const xoTestsBookTableAuthorsJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
left join (
	select
		book_authors.book_id as book_authors_book_id
		, book_authors.pseudonym as pseudonym
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.book_authors
	join xo_tests.users on users.user_id = book_authors.author_id
	group by
		book_authors_book_id
		, users.user_id
		, pseudonym
) as xo_join_book_authors_authors on xo_join_book_authors_authors.book_authors_book_id = books.book_id
`

const xoTestsBookTableAuthorsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_authors.__users
		, xo_join_book_authors_authors.pseudonym
		)) filter (where xo_join_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const xoTestsBookTableAuthorsGroupBySQL = `books.book_id, books.book_id`

const xoTestsBookTableAuthorsBASKJoinSQL = `-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
left join (
	select
		book_authors_surrogate_key.book_id as book_authors_surrogate_key_book_id
		, book_authors_surrogate_key.pseudonym as pseudonym
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.book_authors_surrogate_key
	join xo_tests.users on users.user_id = book_authors_surrogate_key.author_id
	group by
		book_authors_surrogate_key_book_id
		, users.user_id
		, pseudonym
) as xo_join_book_authors_surrogate_key_authors on xo_join_book_authors_surrogate_key_authors.book_authors_surrogate_key_book_id = books.book_id
`

const xoTestsBookTableAuthorsBASKSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_surrogate_key_authors.__users
		, xo_join_book_authors_surrogate_key_authors.pseudonym
		)) filter (where xo_join_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') as book_authors_surrogate_key_authors`

const xoTestsBookTableAuthorsBASKGroupBySQL = `books.book_id, books.book_id`

const xoTestsBookTableBookReviewsJoinSQL = `-- M2O join generated from "book_reviews_book_id_fkey"
left join (
  select
  book_id as book_reviews_book_id
    , row(book_reviews.*) as __book_reviews
  from
    xo_tests.book_reviews
  group by
	  book_reviews_book_id, xo_tests.book_reviews.book_review_id
) as xo_join_book_reviews on xo_join_book_reviews.book_reviews_book_id = books.book_id
`

const xoTestsBookTableBookReviewsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_book_reviews.__book_reviews)) filter (where xo_join_book_reviews.book_reviews_book_id is not null), '{}') as book_reviews`

const xoTestsBookTableBookReviewsGroupBySQL = `books.book_id`

const xoTestsBookTableSellersJoinSQL = `-- M2M join generated from "book_sellers_seller_fkey"
left join (
	select
		book_sellers.book_id as book_sellers_book_id
		, users.user_id as __users_user_id
		, row(users.*) as __users
	from
		xo_tests.book_sellers
	join xo_tests.users on users.user_id = book_sellers.seller
	group by
		book_sellers_book_id
		, users.user_id
) as xo_join_book_sellers_sellers on xo_join_book_sellers_sellers.book_sellers_book_id = books.book_id
`

const xoTestsBookTableSellersSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_sellers_sellers.__users
		)) filter (where xo_join_book_sellers_sellers.__users_user_id is not null), '{}') as book_sellers_sellers`

const xoTestsBookTableSellersGroupBySQL = `books.book_id, books.book_id`

// XoTestsBookUpdateParams represents update params for 'xo_tests.books'.
type XoTestsBookUpdateParams struct {
	Name *string `json:"name" nullable:"false"` // name
}

// SetUpdateParams updates xo_tests.books struct fields with the specified params.
func (xtb *XoTestsBook) SetUpdateParams(params *XoTestsBookUpdateParams) {
	if params.Name != nil {
		xtb.Name = *params.Name
	}
}

// Insert inserts the XoTestsBook to the database.
func (xtb *XoTestsBook) Insert(ctx context.Context, db DB) (*XoTestsBook, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.books (
	name
	) VALUES (
	$1
	) RETURNING * `
	// run
	logf(sqlstr, xtb.Name)

	rows, err := db.Query(ctx, sqlstr, xtb.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Insert/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	newxtb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}

	*xtb = newxtb

	return xtb, nil
}

// Update updates a XoTestsBook in the database.
func (xtb *XoTestsBook) Update(ctx context.Context, db DB) (*XoTestsBook, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.books SET 
	name = $1 
	WHERE book_id = $2 
	RETURNING * `
	// run
	logf(sqlstr, xtb.Name, xtb.BookID)

	rows, err := db.Query(ctx, sqlstr, xtb.Name, xtb.BookID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Update/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	newxtb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}
	*xtb = newxtb

	return xtb, nil
}

// Upsert upserts a XoTestsBook in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtb *XoTestsBook) Upsert(ctx context.Context, db DB, params *XoTestsBookCreateParams) (*XoTestsBook, error) {
	var err error

	xtb.Name = params.Name

	xtb, err = xtb.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertXoTestsBook/Insert: %w", &XoError{Entity: "Book", Err: err})
			}
			xtb, err = xtb.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertXoTestsBook/Update: %w", &XoError{Entity: "Book", Err: err})
			}
		}
	}

	return xtb, err
}

// Delete deletes the XoTestsBook from the database.
func (xtb *XoTestsBook) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.books 
	WHERE book_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtb.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsBookPaginated returns a cursor-paginated list of XoTestsBook.
// At least one cursor is required.
func XoTestsBookPaginated(ctx context.Context, db DB, cursor models.PaginationCursor, opts ...XoTestsBookSelectConfigOption) ([]XoTestsBook, error) {
	c := &XoTestsBookSelectConfig{
		joins:   XoTestsBookJoins{},
		filters: make(map[string][]any),
		having:  make(map[string][]any),
		orderBy: make(map[string]models.Direction),
	}

	for _, o := range opts {
		o(c)
	}

	if cursor.Value == nil {
		return nil, logerror(fmt.Errorf("XoTestsUser/Paginated/cursorValue: %w", &XoError{Entity: "User", Err: fmt.Errorf("no cursor value for column: %s", cursor.Column)}))
	}
	field, ok := XoTestsEntityFields[XoTestsTableEntityXoTestsBook][cursor.Column]
	if !ok {
		return nil, logerror(fmt.Errorf("XoTestsBook/Paginated/cursor: %w", &XoError{Entity: "Book", Err: fmt.Errorf("invalid cursor column: %s", cursor.Column)}))
	}

	op := "<"
	if cursor.Direction == models.DirectionAsc {
		op = ">"
	}
	c.filters[fmt.Sprintf("books.%s %s $i", field.Db, op)] = []any{*cursor.Value}
	c.orderBy[field.Db] = cursor.Direction // no need to duplicate opts

	paramStart := 0 // all filters will come from the user
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
		filters += " where " + strings.Join(filterClauses, " AND ") + " "
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

	orderByClause := ""
	if len(c.orderBy) > 0 {
		orderByClause += " order by "
	} else {
		return nil, logerror(fmt.Errorf("XoTestsBook/Paginated/orderBy: %w", &XoError{Entity: "Book", Err: fmt.Errorf("at least one sorted column is required")}))
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderByClause += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsGroupBySQL)
	}

	if c.joins.AuthorsBASK {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsBASKSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsBASKJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsBASKGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsBookTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableBookReviewsGroupBySQL)
	}

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableSellersGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	books.book_id,
	books.name %s 
	 FROM xo_tests.books %s 
	 %s  %s %s %s`, selects, joins, filters, groupByClause, havingClause, orderByClause)
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookPaginated */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append(filterParams, havingParams...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Paginated/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBook/Paginated/pgx.CollectRows: %w", &XoError{Entity: "Book", Err: err}))
	}
	return res, nil
}

// XoTestsBookByBookID retrieves a row from 'xo_tests.books' as a XoTestsBook.
//
// Generated from index 'books_pkey'.
func XoTestsBookByBookID(ctx context.Context, db DB, bookID XoTestsBookID, opts ...XoTestsBookSelectConfigOption) (*XoTestsBook, error) {
	c := &XoTestsBookSelectConfig{joins: XoTestsBookJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsGroupBySQL)
	}

	if c.joins.AuthorsBASK {
		selectClauses = append(selectClauses, xoTestsBookTableAuthorsBASKSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableAuthorsBASKJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableAuthorsBASKGroupBySQL)
	}

	if c.joins.BookReviews {
		selectClauses = append(selectClauses, xoTestsBookTableBookReviewsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableBookReviewsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableBookReviewsGroupBySQL)
	}

	if c.joins.Sellers {
		selectClauses = append(selectClauses, xoTestsBookTableSellersSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookTableSellersJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookTableSellersGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupByClause := ""
	if len(groupByClauses) > 0 {
		groupByClause = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	books.book_id,
	books.name %s 
	 FROM xo_tests.books %s 
	 WHERE books.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupByClause, havingClause)
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/db.Query: %w", &XoError{Entity: "Book", Err: err}))
	}
	xtb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBook])
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/pgx.CollectOneRow: %w", &XoError{Entity: "Book", Err: err}))
	}

	return &xtb, nil
}
