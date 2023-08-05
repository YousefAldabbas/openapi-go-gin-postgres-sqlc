// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// TimeEntryWithRetry implements repos.TimeEntry interface instrumented with retries
type TimeEntryWithRetry struct {
	repos.TimeEntry
	_retryCount    int
	_retryInterval time.Duration
}

// NewTimeEntryWithRetry returns TimeEntryWithRetry
func NewTimeEntryWithRetry(base repos.TimeEntry, retryCount int, retryInterval time.Duration) TimeEntryWithRetry {
	return TimeEntryWithRetry{
		TimeEntry:      base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.TimeEntry
func (_d TimeEntryWithRetry) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.TimeEntrySelectConfigOption) (tp1 *db.TimeEntry, err error) {
	tp1, err = _d.TimeEntry.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		tp1, err = _d.TimeEntry.ByID(ctx, d, id, opts...)
	}
	return
}

// Create implements repos.TimeEntry
func (_d TimeEntryWithRetry) Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (tp1 *db.TimeEntry, err error) {
	tp1, err = _d.TimeEntry.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		tp1, err = _d.TimeEntry.Create(ctx, d, params)
	}
	return
}

// Delete implements repos.TimeEntry
func (_d TimeEntryWithRetry) Delete(ctx context.Context, d db.DBTX, id int) (tp1 *db.TimeEntry, err error) {
	tp1, err = _d.TimeEntry.Delete(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		tp1, err = _d.TimeEntry.Delete(ctx, d, id)
	}
	return
}

// Update implements repos.TimeEntry
func (_d TimeEntryWithRetry) Update(ctx context.Context, d db.DBTX, id int, params *db.TimeEntryUpdateParams) (tp1 *db.TimeEntry, err error) {
	tp1, err = _d.TimeEntry.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		tp1, err = _d.TimeEntry.Update(ctx, d, id, params)
	}
	return
}
