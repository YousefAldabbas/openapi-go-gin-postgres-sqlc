// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// TimeEntryWithTimeout implements repos.TimeEntry interface instrumented with timeouts
type TimeEntryWithTimeout struct {
	repos.TimeEntry
	config TimeEntryWithTimeoutConfig
}

type TimeEntryWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	CreateTimeout time.Duration

	DeleteTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewTimeEntryWithTimeout returns TimeEntryWithTimeout
func NewTimeEntryWithTimeout(base repos.TimeEntry, config TimeEntryWithTimeoutConfig) TimeEntryWithTimeout {
	return TimeEntryWithTimeout{
		TimeEntry: base,
		config:    config,
	}
}

// ByID implements repos.TimeEntry
func (_d TimeEntryWithTimeout) ByID(ctx context.Context, d db.DBTX, id int64) (tp1 *db.TimeEntry, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.TimeEntry.ByID(ctx, d, id)
}

// Create implements repos.TimeEntry
func (_d TimeEntryWithTimeout) Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (tp1 *db.TimeEntry, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.TimeEntry.Create(ctx, d, params)
}

// Delete implements repos.TimeEntry
func (_d TimeEntryWithTimeout) Delete(ctx context.Context, d db.DBTX, id int64) (tp1 *db.TimeEntry, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.TimeEntry.Delete(ctx, d, id)
}

// Update implements repos.TimeEntry
func (_d TimeEntryWithTimeout) Update(ctx context.Context, d db.DBTX, id int64, params *db.TimeEntryUpdateParams) (tp1 *db.TimeEntry, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.TimeEntry.Update(ctx, d, id, params)
}
