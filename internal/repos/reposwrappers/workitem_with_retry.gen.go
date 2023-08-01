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

// WorkItemWithRetry implements repos.WorkItem interface instrumented with retries
type WorkItemWithRetry struct {
	repos.WorkItem
	_retryCount    int
	_retryInterval time.Duration
}

// NewWorkItemWithRetry returns WorkItemWithRetry
func NewWorkItemWithRetry(base repos.WorkItem, retryCount int, retryInterval time.Duration) WorkItemWithRetry {
	return WorkItemWithRetry{
		WorkItem:       base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.WorkItem
func (_d WorkItemWithRetry) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	wp1, err = _d.WorkItem.ByID(ctx, d, id, opts...)
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
		wp1, err = _d.WorkItem.ByID(ctx, d, id, opts...)
	}
	return
}
