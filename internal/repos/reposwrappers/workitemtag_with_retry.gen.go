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

// WorkItemTagWithRetry implements repos.WorkItemTag interface instrumented with retries
type WorkItemTagWithRetry struct {
	repos.WorkItemTag
	_retryCount    int
	_retryInterval time.Duration
}

// NewWorkItemTagWithRetry returns WorkItemTagWithRetry
func NewWorkItemTagWithRetry(base repos.WorkItemTag, retryCount int, retryInterval time.Duration) WorkItemTagWithRetry {
	return WorkItemTagWithRetry{
		WorkItemTag:    base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) ByID(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemTag, err error) {
	wp1, err = _d.WorkItemTag.ByID(ctx, d, id)
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
		wp1, err = _d.WorkItemTag.ByID(ctx, d, id)
	}
	return
}

// ByName implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (wp1 *db.WorkItemTag, err error) {
	wp1, err = _d.WorkItemTag.ByName(ctx, d, name, projectID)
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
		wp1, err = _d.WorkItemTag.ByName(ctx, d, name, projectID)
	}
	return
}

// Create implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) Create(ctx context.Context, d db.DBTX, params *db.WorkItemTagCreateParams) (wp1 *db.WorkItemTag, err error) {
	wp1, err = _d.WorkItemTag.Create(ctx, d, params)
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
		wp1, err = _d.WorkItemTag.Create(ctx, d, params)
	}
	return
}

// Delete implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) Delete(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemTag, err error) {
	wp1, err = _d.WorkItemTag.Delete(ctx, d, id)
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
		wp1, err = _d.WorkItemTag.Delete(ctx, d, id)
	}
	return
}

// Update implements repos.WorkItemTag
func (_d WorkItemTagWithRetry) Update(ctx context.Context, d db.DBTX, id int, params *db.WorkItemTagUpdateParams) (wp1 *db.WorkItemTag, err error) {
	wp1, err = _d.WorkItemTag.Update(ctx, d, id, params)
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
		wp1, err = _d.WorkItemTag.Update(ctx, d, id, params)
	}
	return
}
