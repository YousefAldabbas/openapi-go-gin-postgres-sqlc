package got

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// BookAuthor represents a row from 'public.book_authors'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|O2M|M2O|M2M" to generate joins (not executed by default).
type BookAuthor struct {
	BookID   int       `json:"bookID" db:"book_id" required:"true"`     // book_id
	AuthorID uuid.UUID `json:"authorID" db:"author_id" required:"true"` // author_id
}

// BookAuthorCreateParams represents insert params for 'public.book_authors'
type BookAuthorCreateParams struct {
	BookID   int       `json:"bookID"`   // book_id
	AuthorID uuid.UUID `json:"authorID"` // author_id
}

// BookAuthorUpdateParams represents update params for 'public.book_authors'
type BookAuthorUpdateParams struct {
	BookID   *int       `json:"bookID"`   // book_id
	AuthorID *uuid.UUID `json:"authorID"` // author_id
}

type BookAuthorSelectConfig struct {
	limit   string
	orderBy string
	joins   BookAuthorJoins
}
type BookAuthorSelectConfigOption func(*BookAuthorSelectConfig)

// WithBookAuthorLimit limits row selection.
func WithBookAuthorLimit(limit int) BookAuthorSelectConfigOption {
	return func(s *BookAuthorSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type BookAuthorOrderBy = string

type BookAuthorJoins struct{}

// WithBookAuthorJoin joins with the given tables.
func WithBookAuthorJoin(joins BookAuthorJoins) BookAuthorSelectConfigOption {
	return func(s *BookAuthorSelectConfig) {
		s.joins = joins
	}
}

// BookAuthorByBookIDAuthorID retrieves a row from 'public.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_book_id_author_id_key'.
func BookAuthorByBookIDAuthorID(ctx context.Context, db DB, bookID int, authorID uuid.UUID, opts ...BookAuthorSelectConfigOption) (*BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors.book_id,
book_authors.author_id ` +
		`FROM public.book_authors ` +
		`` +
		` WHERE book_authors.book_id = $1 AND book_authors.author_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID, authorID)
	rows, err := db.Query(ctx, sqlstr, bookID, authorID)
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/db.Query: %w", err))
	}
	ba, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("book_authors/BookAuthorByBookIDAuthorID/pgx.CollectOneRow: %w", err))
	}

	return &ba, nil
}

// BookAuthorsByBookID retrieves a row from 'public.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_book_id_author_id_key'.
func BookAuthorsByBookID(ctx context.Context, db DB, bookID int, opts ...BookAuthorSelectConfigOption) ([]BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors.book_id,
book_authors.author_id ` +
		`FROM public.book_authors ` +
		`` +
		` WHERE book_authors.book_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, bookID)
	rows, err := db.Query(ctx, sqlstr, bookID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// BookAuthorsByAuthorID retrieves a row from 'public.book_authors' as a BookAuthor.
//
// Generated from index 'book_authors_book_id_author_id_key'.
func BookAuthorsByAuthorID(ctx context.Context, db DB, authorID uuid.UUID, opts ...BookAuthorSelectConfigOption) ([]BookAuthor, error) {
	c := &BookAuthorSelectConfig{joins: BookAuthorJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`book_authors.book_id,
book_authors.author_id ` +
		`FROM public.book_authors ` +
		`` +
		` WHERE book_authors.author_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, authorID)
	rows, err := db.Query(ctx, sqlstr, authorID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[BookAuthor])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}
