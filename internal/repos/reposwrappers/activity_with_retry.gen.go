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

// ActivityWithRetry implements repos.Activity interface instrumented with retries
type ActivityWithRetry struct {
	repos.Activity
	_retryCount    int
	_retryInterval time.Duration
}

// NewActivityWithRetry returns ActivityWithRetry
func NewActivityWithRetry(base repos.Activity, retryCount int, retryInterval time.Duration) ActivityWithRetry {
	return ActivityWithRetry{
		Activity:       base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.Activity
func (_d ActivityWithRetry) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.ActivitySelectConfigOption) (ap1 *db.Activity, err error) {
	ap1, err = _d.Activity.ByID(ctx, d, id, opts...)
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
		ap1, err = _d.Activity.ByID(ctx, d, id, opts...)
	}
	return
}

// ByName implements repos.Activity
func (_d ActivityWithRetry) ByName(ctx context.Context, d db.DBTX, name string, projectID int, opts ...db.ActivitySelectConfigOption) (ap1 *db.Activity, err error) {
	ap1, err = _d.Activity.ByName(ctx, d, name, projectID, opts...)
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
		ap1, err = _d.Activity.ByName(ctx, d, name, projectID, opts...)
	}
	return
}

// ByProjectID implements repos.Activity
func (_d ActivityWithRetry) ByProjectID(ctx context.Context, d db.DBTX, projectID int, opts ...db.ActivitySelectConfigOption) (aa1 []db.Activity, err error) {
	aa1, err = _d.Activity.ByProjectID(ctx, d, projectID, opts...)
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
		aa1, err = _d.Activity.ByProjectID(ctx, d, projectID, opts...)
	}
	return
}

// Create implements repos.Activity
func (_d ActivityWithRetry) Create(ctx context.Context, d db.DBTX, params *db.ActivityCreateParams) (ap1 *db.Activity, err error) {
	ap1, err = _d.Activity.Create(ctx, d, params)
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
		ap1, err = _d.Activity.Create(ctx, d, params)
	}
	return
}

// Delete implements repos.Activity
func (_d ActivityWithRetry) Delete(ctx context.Context, d db.DBTX, id int) (ap1 *db.Activity, err error) {
	ap1, err = _d.Activity.Delete(ctx, d, id)
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
		ap1, err = _d.Activity.Delete(ctx, d, id)
	}
	return
}

// Update implements repos.Activity
func (_d ActivityWithRetry) Update(ctx context.Context, d db.DBTX, id int, params *db.ActivityUpdateParams) (ap1 *db.Activity, err error) {
	ap1, err = _d.Activity.Update(ctx, d, id, params)
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
		ap1, err = _d.Activity.Update(ctx, d, id, params)
	}
	return
}
