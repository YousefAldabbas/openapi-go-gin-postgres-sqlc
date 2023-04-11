// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// TeamWithRetry implements repos.Team interface instrumented with retries
type TeamWithRetry struct {
	repos.Team
	_retryCount    int
	_retryInterval time.Duration
}

// NewTeamWithRetry returns TeamWithRetry
func NewTeamWithRetry(base repos.Team, retryCount int, retryInterval time.Duration) TeamWithRetry {
	return TeamWithRetry{
		Team:           base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.Team
func (_d TeamWithRetry) ByID(ctx context.Context, d db.DBTX, id int) (tp1 *db.Team, err error) {
	tp1, err = _d.Team.ByID(ctx, d, id)
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
		tp1, err = _d.Team.ByID(ctx, d, id)
	}
	return
}

// ByName implements repos.Team
func (_d TeamWithRetry) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (tp1 *db.Team, err error) {
	tp1, err = _d.Team.ByName(ctx, d, name, projectID)
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
		tp1, err = _d.Team.ByName(ctx, d, name, projectID)
	}
	return
}

// Create implements repos.Team
func (_d TeamWithRetry) Create(ctx context.Context, d db.DBTX, params repos.TeamCreateParams) (tp1 *db.Team, err error) {
	tp1, err = _d.Team.Create(ctx, d, params)
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
		tp1, err = _d.Team.Create(ctx, d, params)
	}
	return
}

// Delete implements repos.Team
func (_d TeamWithRetry) Delete(ctx context.Context, d db.DBTX, id int) (tp1 *db.Team, err error) {
	tp1, err = _d.Team.Delete(ctx, d, id)
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
		tp1, err = _d.Team.Delete(ctx, d, id)
	}
	return
}

// Update implements repos.Team
func (_d TeamWithRetry) Update(ctx context.Context, d db.DBTX, id int, params repos.TeamUpdateParams) (tp1 *db.Team, err error) {
	tp1, err = _d.Team.Update(ctx, d, id, params)
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
		tp1, err = _d.Team.Update(ctx, d, id, params)
	}
	return
}
