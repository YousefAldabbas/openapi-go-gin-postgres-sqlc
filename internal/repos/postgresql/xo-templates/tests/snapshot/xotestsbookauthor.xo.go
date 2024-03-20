package got

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

// XoTestsBookAuthor represents a row from 'xo_tests.book_authors'.
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
type XoTestsBookAuthor struct {
	BookID    XoTestsBookID `json:"bookID" db:"book_id" required:"true" nullable:"false"`     // book_id
	AuthorID  XoTestsUserID `json:"authorID" db:"author_id" required:"true" nullable:"false"` // author_id
	Pseudonym *string       `json:"pseudonym" db:"pseudonym"`                                 // pseudonym

	BooksJoin   *[]XoTestsBookAuthorM2MBookBA   `json:"-" db:"book_authors_books" openapi-go:"ignore"`   // M2M book_authors
	AuthorsJoin *[]XoTestsBookAuthorM2MAuthorBA `json:"-" db:"book_authors_authors" openapi-go:"ignore"` // M2M book_authors
}

// XoTestsBookAuthorCreateParams represents insert params for 'xo_tests.book_authors'.
type XoTestsBookAuthorCreateParams struct {
	AuthorID  XoTestsUserID `json:"authorID" required:"true" nullable:"false"` // author_id
	BookID    XoTestsBookID `json:"bookID" required:"true" nullable:"false"`   // book_id
	Pseudonym *string       `json:"pseudonym"`                                 // pseudonym
}

// XoTestsBookAuthorParams represents common params for both insert and update of 'xo_tests.book_authors'.
type XoTestsBookAuthorParams interface {
	GetAuthorID() *XoTestsUserID
	GetBookID() *XoTestsBookID
	GetPseudonym() *string
}

func (p XoTestsBookAuthorCreateParams) GetAuthorID() *XoTestsUserID {
	x := p.AuthorID
	return &x
}

func (p XoTestsBookAuthorUpdateParams) GetAuthorID() *XoTestsUserID {
	return p.AuthorID
}

func (p XoTestsBookAuthorCreateParams) GetBookID() *XoTestsBookID {
	x := p.BookID
	return &x
}

func (p XoTestsBookAuthorUpdateParams) GetBookID() *XoTestsBookID {
	return p.BookID
}

func (p XoTestsBookAuthorCreateParams) GetPseudonym() *string {
	return p.Pseudonym
}

func (p XoTestsBookAuthorUpdateParams) GetPseudonym() *string {
	if p.Pseudonym != nil {
		return *p.Pseudonym
	}
	return nil
}

// CreateXoTestsBookAuthor creates a new XoTestsBookAuthor in the database with the given params.
func CreateXoTestsBookAuthor(ctx context.Context, db DB, params *XoTestsBookAuthorCreateParams) (*XoTestsBookAuthor, error) {
	xtba := &XoTestsBookAuthor{
		AuthorID:  params.AuthorID,
		BookID:    params.BookID,
		Pseudonym: params.Pseudonym,
	}

	return xtba.Insert(ctx, db)
}

type XoTestsBookAuthorSelectConfig struct {
	limit   string
	orderBy string
	joins   XoTestsBookAuthorJoins
	filters map[string][]any
	having  map[string][]any
}
type XoTestsBookAuthorSelectConfigOption func(*XoTestsBookAuthorSelectConfig)

// WithXoTestsBookAuthorLimit limits row selection.
func WithXoTestsBookAuthorLimit(limit int) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type XoTestsBookAuthorOrderBy string

type XoTestsBookAuthorJoins struct {
	Books   bool `json:"books" required:"true" nullable:"false"`   // M2M book_authors
	Authors bool `json:"authors" required:"true" nullable:"false"` // M2M book_authors
}

// WithXoTestsBookAuthorJoin joins with the given tables.
func WithXoTestsBookAuthorJoin(joins XoTestsBookAuthorJoins) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		s.joins = XoTestsBookAuthorJoins{
			Books:   s.joins.Books || joins.Books,
			Authors: s.joins.Authors || joins.Authors,
		}
	}
}

// XoTestsBookAuthorM2MBookBA represents a M2M join against "xo_tests.book_authors"
type XoTestsBookAuthorM2MBookBA struct {
	Book      XoTestsBook `json:"book" db:"books" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// XoTestsBookAuthorM2MAuthorBA represents a M2M join against "xo_tests.book_authors"
type XoTestsBookAuthorM2MAuthorBA struct {
	User      XoTestsUser `json:"user" db:"users" required:"true"`
	Pseudonym *string     `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithXoTestsBookAuthorFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithXoTestsBookAuthorFilters(filters map[string][]any) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		s.filters = filters
	}
}

// WithXoTestsBookAuthorHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
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
func WithXoTestsBookAuthorHavingClause(conditions map[string][]any) XoTestsBookAuthorSelectConfigOption {
	return func(s *XoTestsBookAuthorSelectConfig) {
		s.having = conditions
	}
}

const xoTestsBookAuthorTableBooksJoinSQL = `-- M2M join generated from "book_authors_book_id_fkey"
left join (
	select
		book_authors.author_id as book_authors_author_id
		, book_authors.pseudonym as pseudonym
		, books.book_id as __books_book_id
		, row(books.*) as __books
	from
		xo_tests.book_authors
	join xo_tests.books on books.book_id = book_authors.book_id
	group by
		book_authors_author_id
		, books.book_id
		, pseudonym
) as xo_join_book_authors_books on xo_join_book_authors_books.book_authors_author_id = book_authors.author_id
`

const xoTestsBookAuthorTableBooksSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_books.__books
		, xo_join_book_authors_books.pseudonym
		)) filter (where xo_join_book_authors_books.__books_book_id is not null), '{}') as book_authors_books`

const xoTestsBookAuthorTableBooksGroupBySQL = `book_authors.author_id, book_authors.book_id, book_authors.author_id`

const xoTestsBookAuthorTableAuthorsJoinSQL = `-- M2M join generated from "book_authors_author_id_fkey"
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
) as xo_join_book_authors_authors on xo_join_book_authors_authors.book_authors_book_id = book_authors.book_id
`

const xoTestsBookAuthorTableAuthorsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_book_authors_authors.__users
		, xo_join_book_authors_authors.pseudonym
		)) filter (where xo_join_book_authors_authors.__users_user_id is not null), '{}') as book_authors_authors`

const xoTestsBookAuthorTableAuthorsGroupBySQL = `book_authors.book_id, book_authors.book_id, book_authors.author_id`

// XoTestsBookAuthorUpdateParams represents update params for 'xo_tests.book_authors'.
type XoTestsBookAuthorUpdateParams struct {
	AuthorID  *XoTestsUserID `json:"authorID" nullable:"false"` // author_id
	BookID    *XoTestsBookID `json:"bookID" nullable:"false"`   // book_id
	Pseudonym **string       `json:"pseudonym"`                 // pseudonym
}

// SetUpdateParams updates xo_tests.book_authors struct fields with the specified params.
func (xtba *XoTestsBookAuthor) SetUpdateParams(params *XoTestsBookAuthorUpdateParams) {
	if params.AuthorID != nil {
		xtba.AuthorID = *params.AuthorID
	}
	if params.BookID != nil {
		xtba.BookID = *params.BookID
	}
	if params.Pseudonym != nil {
		xtba.Pseudonym = *params.Pseudonym
	}
}

// Insert inserts the XoTestsBookAuthor to the database.
func (xtba *XoTestsBookAuthor) Insert(ctx context.Context, db DB) (*XoTestsBookAuthor, error) {
	// insert (manual)
	sqlstr := `INSERT INTO xo_tests.book_authors (
	author_id, book_id, pseudonym
	) VALUES (
	$1, $2, $3
	)
	 RETURNING * `
	// run
	logf(sqlstr, xtba.AuthorID, xtba.BookID, xtba.Pseudonym)
	rows, err := db.Query(ctx, sqlstr, xtba.AuthorID, xtba.BookID, xtba.Pseudonym)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Insert/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newxtba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*xtba = newxtba

	return xtba, nil
}

// Update updates a XoTestsBookAuthor in the database.
func (xtba *XoTestsBookAuthor) Update(ctx context.Context, db DB) (*XoTestsBookAuthor, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.book_authors SET 
	pseudonym = $1 
	WHERE book_id = $2  AND author_id = $3 
	RETURNING * `
	// run
	logf(sqlstr, xtba.Pseudonym, xtba.BookID, xtba.AuthorID)

	rows, err := db.Query(ctx, sqlstr, xtba.Pseudonym, xtba.BookID, xtba.AuthorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Update/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	newxtba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/Update/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}
	*xtba = newxtba

	return xtba, nil
}

// Upsert upserts a XoTestsBookAuthor in the database.
// Requires appropriate PK(s) to be set beforehand.
func (xtba *XoTestsBookAuthor) Upsert(ctx context.Context, db DB, params *XoTestsBookAuthorCreateParams) (*XoTestsBookAuthor, error) {
	var err error

	xtba.AuthorID = params.AuthorID
	xtba.BookID = params.BookID
	xtba.Pseudonym = params.Pseudonym

	xtba, err = xtba.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "Book author", Err: err})
			}
			xtba, err = xtba.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "Book author", Err: err})
			}
		}
	}

	return xtba, err
}

// Delete deletes the XoTestsBookAuthor from the database.
func (xtba *XoTestsBookAuthor) Delete(ctx context.Context, db DB) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM xo_tests.book_authors 
	WHERE book_id = $1 AND author_id = $2 `
	// run
	if _, err := db.Exec(ctx, sqlstr, xtba.BookID, xtba.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// XoTestsBookAuthorByBookIDAuthorID retrieves a row from 'xo_tests.book_authors' as a XoTestsBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func XoTestsBookAuthorByBookIDAuthorID(ctx context.Context, db DB, bookID XoTestsBookID, authorID XoTestsUserID, opts ...XoTestsBookAuthorSelectConfigOption) (*XoTestsBookAuthor, error) {
	c := &XoTestsBookAuthorSelectConfig{joins: XoTestsBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableAuthorsGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.book_id = $1 AND book_authors.author_id = $2
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorByBookIDAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID, authorID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/db.Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	xtba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/pgx.CollectOneRow: %w", &XoError{Entity: "Book author", Err: err}))
	}

	return &xtba, nil
}

// XoTestsBookAuthorsByBookID retrieves a row from 'xo_tests.book_authors' as a XoTestsBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func XoTestsBookAuthorsByBookID(ctx context.Context, db DB, bookID XoTestsBookID, opts ...XoTestsBookAuthorSelectConfigOption) ([]XoTestsBookAuthor, error) {
	c := &XoTestsBookAuthorSelectConfig{joins: XoTestsBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableAuthorsGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.book_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsByBookID */\n" + sqlstr

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{bookID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// XoTestsBookAuthorsByAuthorID retrieves a row from 'xo_tests.book_authors' as a XoTestsBookAuthor.
//
// Generated from index 'book_authors_pkey'.
func XoTestsBookAuthorsByAuthorID(ctx context.Context, db DB, authorID XoTestsUserID, opts ...XoTestsBookAuthorSelectConfigOption) ([]XoTestsBookAuthor, error) {
	c := &XoTestsBookAuthorSelectConfig{joins: XoTestsBookAuthorJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

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

	if c.joins.Books {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableBooksSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableBooksJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableBooksGroupBySQL)
	}

	if c.joins.Authors {
		selectClauses = append(selectClauses, xoTestsBookAuthorTableAuthorsSelectSQL)
		joinClauses = append(joinClauses, xoTestsBookAuthorTableAuthorsJoinSQL)
		groupByClauses = append(groupByClauses, xoTestsBookAuthorTableAuthorsGroupBySQL)
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
	book_authors.author_id,
	book_authors.book_id,
	book_authors.pseudonym %s 
	 FROM xo_tests.book_authors %s 
	 WHERE book_authors.author_id = $1
	 %s   %s 
  %s 
`, selects, joins, filters, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* XoTestsBookAuthorsByAuthorID */\n" + sqlstr

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, append([]any{authorID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/Query: %w", &XoError{Entity: "Book author", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[XoTestsBookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("XoTestsBookAuthor/BookAuthorByBookIDAuthorID/pgx.CollectRows: %w", &XoError{Entity: "Book author", Err: err}))
	}
	return res, nil
}

// FKUser_AuthorID returns the User associated with the XoTestsBookAuthor's (AuthorID).
//
// Generated from foreign key 'book_authors_author_id_fkey'.
func (xtba *XoTestsBookAuthor) FKUser_AuthorID(ctx context.Context, db DB) (*XoTestsUser, error) {
	return XoTestsUserByUserID(ctx, db, xtba.AuthorID)
}

// FKBook_BookID returns the Book associated with the XoTestsBookAuthor's (BookID).
//
// Generated from foreign key 'book_authors_book_id_fkey'.
func (xtba *XoTestsBookAuthor) FKBook_BookID(ctx context.Context, db DB) (*XoTestsBook, error) {
	return XoTestsBookByBookID(ctx, db, xtba.BookID)
}
