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

// Book represents a row from 'xo_tests.books'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":private to exclude a field from JSON.
//   - "type":<pkg.type> to override the type annotation.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type Book struct {
	BookID int    `json:"bookID" db:"book_id" required:"true"` // book_id
	Name   string `json:"name" db:"name" required:"true"`      // name

	BookAuthorsJoin     *[]User__BA_Book   `json:"-" db:"book_authors_authors" openapi-go:"ignore"`               // M2M book_authors
	BookAuthorsJoinBASK *[]User__BASK_Book `json:"-" db:"book_authors_surrogate_key_authors" openapi-go:"ignore"` // M2M book_authors_surrogate_key
	BookBookReviewsJoin *[]BookReview      `json:"-" db:"book_reviews" openapi-go:"ignore"`                       // M2O books
	BookSellersJoin     *[]User            `json:"-" db:"book_sellers_sellers" openapi-go:"ignore"`               // M2M book_sellers
}

// BookCreateParams represents insert params for 'xo_tests.books'.
type BookCreateParams struct {
	Name string `json:"name" required:"true"` // name
}

// CreateBook creates a new Book in the database with the given params.
func CreateBook(ctx context.Context, db DB, params *BookCreateParams) (*Book, error) {
	b := &Book{
		Name: params.Name,
	}

	return b.Insert(ctx, db)
}

// BookUpdateParams represents update params for 'xo_tests.books'.
type BookUpdateParams struct {
	Name *string `json:"name" required:"true"` // name
}

// SetUpdateParams updates xo_tests.books struct fields with the specified params.
func (b *Book) SetUpdateParams(params *BookUpdateParams) {
	if params.Name != nil {
		b.Name = *params.Name
	}
}

type BookSelectConfig struct {
	limit   string
	orderBy string
	joins   BookJoins
	filters map[string][]any
}
type BookSelectConfigOption func(*BookSelectConfig)

// WithBookLimit limits row selection.
func WithBookLimit(limit int) BookSelectConfigOption {
	return func(s *BookSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type BookOrderBy string

type BookJoins struct {
	AuthorsBook      bool // M2M book_authors
	AuthorsBookUsers bool // M2M book_authors_surrogate_key
	BookReviews      bool // M2O book_reviews
	Sellers          bool // M2M book_sellers
}

// WithBookJoin joins with the given tables.
func WithBookJoin(joins BookJoins) BookSelectConfigOption {
	return func(s *BookSelectConfig) {
		s.joins = BookJoins{
			AuthorsBook:      s.joins.AuthorsBook || joins.AuthorsBook,
			AuthorsBookUsers: s.joins.AuthorsBookUsers || joins.AuthorsBookUsers,
			BookReviews:      s.joins.BookReviews || joins.BookReviews,
			Sellers:          s.joins.Sellers || joins.Sellers,
		}
	}
}

// User__BA_Book represents a M2M join against "xo_tests.book_authors"
type User__BA_Book struct {
	User      User    `json:"user" db:"users" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// User__BASK_Book represents a M2M join against "xo_tests.book_authors_surrogate_key"
type User__BASK_Book struct {
	User      User    `json:"user" db:"users" required:"true"`
	Pseudonym *string `json:"pseudonym" db:"pseudonym" required:"true" `
}

// WithBookFilters adds the given filters, which may be parameterized with $i.
// Filters are joined with AND.
// NOTE: SQL injection prone.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithBookFilters(filters map[string][]any) BookSelectConfigOption {
	return func(s *BookSelectConfig) {
		s.filters = filters
	}
}

// Insert inserts the Book to the database.
func (b *Book) Insert(ctx context.Context, db DB) (*Book, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO xo_tests.books (` +
		`name` +
		`) VALUES (` +
		`$1` +
		`) RETURNING * `
	// run
	logf(sqlstr, b.Name)

	rows, err := db.Query(ctx, sqlstr, b.Name)
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Insert/db.Query: %w", err))
	}
	newb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Book])
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Insert/pgx.CollectOneRow: %w", err))
	}

	*b = newb

	return b, nil
}

// Update updates a Book in the database.
func (b *Book) Update(ctx context.Context, db DB) (*Book, error) {
	// update with composite primary key
	sqlstr := `UPDATE xo_tests.books SET ` +
		`name = $1 ` +
		`WHERE book_id = $2 ` +
		`RETURNING * `
	// run
	logf(sqlstr, b.Name, b.BookID)

	rows, err := db.Query(ctx, sqlstr, b.Name, b.BookID)
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Update/db.Query: %w", err))
	}
	newb, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Book])
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Update/pgx.CollectOneRow: %w", err))
	}
	*b = newb

	return b, nil
}

// Upsert upserts a Book in the database.
// Requires appropiate PK(s) to be set beforehand.
func (b *Book) Upsert(ctx context.Context, db DB, params *BookCreateParams) (*Book, error) {
	var err error

	b.Name = params.Name

	b, err = b.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			b, err = b.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return b, err
}

// Delete deletes the Book from the database.
func (b *Book) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM xo_tests.books ` +
		`WHERE book_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, b.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// BookPaginatedByBookIDAsc returns a cursor-paginated list of Book in Asc order.
func BookPaginatedByBookIDAsc(ctx context.Context, db DB, bookID int, opts ...BookSelectConfigOption) ([]Book, error) {
	c := &BookSelectConfig{joins: BookJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 5
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`books.book_id,
books.name,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_authors.__users
		, joined_book_authors_authors.pseudonym
		)) filter (where joined_book_authors_authors.__users_user_id is not null), '{}') end) as book_authors_authors,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_authors.__users
		, joined_book_authors_surrogate_key_authors.pseudonym
		)) filter (where joined_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') end) as book_authors_surrogate_key_authors,
(case when $3::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews,
(case when $4::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users_user_id is not null), '{}') end) as book_sellers_sellers `+
		`FROM xo_tests.books `+
		`-- M2M join generated from "book_authors_author_id_fkey"
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
  ) as joined_book_authors_authors on joined_book_authors_authors.book_authors_book_id = books.book_id

-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
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
  ) as joined_book_authors_surrogate_key_authors on joined_book_authors_surrogate_key_authors.book_authors_surrogate_key_book_id = books.book_id

-- M2O join generated from "book_reviews_book_id_fkey"
left join (
  select
  book_id as book_reviews_book_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        book_id) joined_book_reviews on joined_book_reviews.book_reviews_book_id = books.book_id
-- M2M join generated from "book_sellers_seller_fkey"
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
  ) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = books.book_id
`+
		` WHERE books.book_id > $5`+
		` %s  GROUP BY books.book_id, 
books.name, 
books.book_id, books.book_id, 
books.book_id, books.book_id, 
joined_book_reviews.book_reviews, books.book_id, 
books.book_id, books.book_id ORDER BY 
		book_id Asc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.AuthorsBook, c.joins.AuthorsBookUsers, c.joins.BookReviews, c.joins.Sellers, bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Paginated/Asc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Book])
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Paginated/Asc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookPaginatedByBookIDDesc returns a cursor-paginated list of Book in Desc order.
func BookPaginatedByBookIDDesc(ctx context.Context, db DB, bookID int, opts ...BookSelectConfigOption) ([]Book, error) {
	c := &BookSelectConfig{joins: BookJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 5
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`books.book_id,
books.name,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_authors.__users
		, joined_book_authors_authors.pseudonym
		)) filter (where joined_book_authors_authors.__users_user_id is not null), '{}') end) as book_authors_authors,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_authors.__users
		, joined_book_authors_surrogate_key_authors.pseudonym
		)) filter (where joined_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') end) as book_authors_surrogate_key_authors,
(case when $3::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews,
(case when $4::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users_user_id is not null), '{}') end) as book_sellers_sellers `+
		`FROM xo_tests.books `+
		`-- M2M join generated from "book_authors_author_id_fkey"
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
  ) as joined_book_authors_authors on joined_book_authors_authors.book_authors_book_id = books.book_id

-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
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
  ) as joined_book_authors_surrogate_key_authors on joined_book_authors_surrogate_key_authors.book_authors_surrogate_key_book_id = books.book_id

-- M2O join generated from "book_reviews_book_id_fkey"
left join (
  select
  book_id as book_reviews_book_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        book_id) joined_book_reviews on joined_book_reviews.book_reviews_book_id = books.book_id
-- M2M join generated from "book_sellers_seller_fkey"
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
  ) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = books.book_id
`+
		` WHERE books.book_id < $5`+
		` %s  GROUP BY books.book_id, 
books.name, 
books.book_id, books.book_id, 
books.book_id, books.book_id, 
joined_book_reviews.book_reviews, books.book_id, 
books.book_id, books.book_id ORDER BY 
		book_id Desc `, filters)
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.AuthorsBook, c.joins.AuthorsBookUsers, c.joins.BookReviews, c.joins.Sellers, bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Paginated/Desc/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[Book])
	if err != nil {
		return nil, logerror(fmt.Errorf("Book/Paginated/Desc/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookByBookID retrieves a row from 'xo_tests.books' as a Book.
//
// Generated from index 'books_pkey'.
func BookByBookID(ctx context.Context, db DB, bookID int, opts ...BookSelectConfigOption) (*Book, error) {
	c := &BookSelectConfig{joins: BookJoins{}, filters: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 5
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterValues []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterValues = append(filterValues, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT `+
		`books.book_id,
books.name,
(case when $1::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_authors.__users
		, joined_book_authors_authors.pseudonym
		)) filter (where joined_book_authors_authors.__users_user_id is not null), '{}') end) as book_authors_authors,
(case when $2::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_authors_surrogate_key_authors.__users
		, joined_book_authors_surrogate_key_authors.pseudonym
		)) filter (where joined_book_authors_surrogate_key_authors.__users_user_id is not null), '{}') end) as book_authors_surrogate_key_authors,
(case when $3::boolean = true then COALESCE(joined_book_reviews.book_reviews, '{}') end) as book_reviews,
(case when $4::boolean = true then COALESCE(
		ARRAY_AGG( DISTINCT (
		joined_book_sellers_sellers.__users
		)) filter (where joined_book_sellers_sellers.__users_user_id is not null), '{}') end) as book_sellers_sellers `+
		`FROM xo_tests.books `+
		`-- M2M join generated from "book_authors_author_id_fkey"
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
  ) as joined_book_authors_authors on joined_book_authors_authors.book_authors_book_id = books.book_id

-- M2M join generated from "book_authors_surrogate_key_author_id_fkey"
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
  ) as joined_book_authors_surrogate_key_authors on joined_book_authors_surrogate_key_authors.book_authors_surrogate_key_book_id = books.book_id

-- M2O join generated from "book_reviews_book_id_fkey"
left join (
  select
  book_id as book_reviews_book_id
    , array_agg(book_reviews.*) as book_reviews
  from
    xo_tests.book_reviews
  group by
        book_id) joined_book_reviews on joined_book_reviews.book_reviews_book_id = books.book_id
-- M2M join generated from "book_sellers_seller_fkey"
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
  ) as joined_book_sellers_sellers on joined_book_sellers_sellers.book_sellers_book_id = books.book_id
`+
		` WHERE books.book_id = $5`+
		` %s  GROUP BY 
books.book_id, books.book_id, 
books.book_id, books.book_id, 
joined_book_reviews.book_reviews, books.book_id, 
books.book_id, books.book_id `, filters)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, append([]any{c.joins.AuthorsBook, c.joins.AuthorsBookUsers, c.joins.BookReviews, c.joins.Sellers, bookID}, filterValues...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/db.Query: %w", err))
	}
	b, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[Book])
	if err != nil {
		return nil, logerror(fmt.Errorf("books/BookByBookID/pgx.CollectOneRow: %w", err))
	}

	return &b, nil
}
