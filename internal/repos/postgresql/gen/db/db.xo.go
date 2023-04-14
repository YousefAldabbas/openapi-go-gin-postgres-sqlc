package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"io"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	// logf is used by generated code to log SQL queries.
	logf = func(string, ...any) {}
	// errf is used by generated code to log SQL errors.
	errf = func(string, ...any) {}
)

type Trigger struct{}

// logerror logs the error and returns it.
func logerror(err error) error {
	errf("ERROR: %v", err)
	return err
}

// Logf logs a message using the package logger.
func Logf(s string, v ...any) {
	logf(s, v...)
}

// SetLogger sets the package logger. Valid logger types:
//
//	io.Writer
//	func(string, ...any) (int, error) // fmt.Printf
//	func(string, ...any) // log.Printf
func SetLogger(logger any) {
	logf = convLogger(logger)
}

// Errorf logs an error message using the package error logger.
func Errorf(s string, v ...any) {
	errf(s, v...)
}

// SetErrorLogger sets the package error logger. Valid logger types:
//
//	io.Writer
//	func(string, ...any) (int, error) // fmt.Printf
//	func(string, ...any) // log.Printf
func SetErrorLogger(logger any) {
	errf = convLogger(logger)
}

// convLogger converts logger to the standard logger interface.
func convLogger(logger any) func(string, ...any) {
	switch z := logger.(type) {
	case io.Writer:
		return func(s string, v ...any) {
			fmt.Fprintf(z, s, v...)
		}
	case func(string, ...any) (int, error): // fmt.Printf
		return func(s string, v ...any) {
			_, _ = z(s, v...)
		}
	case func(string, ...any): // log.Printf
		return z
	}
	panic(fmt.Sprintf("unsupported logger type %T", logger))
}

// DB is the common interface for database operations that can be used with
// types from schema 'public'.
type DB interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
}

// Error is an error.
type Error string

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// Error values.
const (
	// ErrAlreadyExists is the already exists error.
	ErrAlreadyExists Error = "already exists"
	// ErrDoesNotExist is the does not exist error.
	ErrDoesNotExist Error = "does not exist"
	// ErrMarkedForDeletion is the marked for deletion error.
	ErrMarkedForDeletion Error = "marked for deletion"
)

// ErrInsertFailed is the insert failed error.
type ErrInsertFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrInsertFailed) Error() string {
	return fmt.Sprintf("insert failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrInsertFailed) Unwrap() error {
	return err.Err
}

// ErrUpdateFailed is the update failed error.
type ErrUpdateFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrUpdateFailed) Error() string {
	return fmt.Sprintf("update failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrUpdateFailed) Unwrap() error {
	return err.Err
}

// ErrUpsertFailed is the upsert failed error.
type ErrUpsertFailed struct {
	Err error
}

// Error satisfies the error interface.
func (err *ErrUpsertFailed) Error() string {
	return fmt.Sprintf("upsert failed: %v", err.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *ErrUpsertFailed) Unwrap() error {
	return err.Err
}
