// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// ProjectWithRetry implements repos.Project interface instrumented with retries
type ProjectWithRetry struct {
	repos.Project
	_retryCount    int
	_retryInterval time.Duration
}

// NewProjectWithRetry returns ProjectWithRetry
func NewProjectWithRetry(base repos.Project, retryCount int, retryInterval time.Duration) ProjectWithRetry {
	return ProjectWithRetry{
		Project:        base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.Project
func (_d ProjectWithRetry) ByID(ctx context.Context, d db.DBTX, id db.ProjectID, opts ...db.ProjectSelectConfigOption) (pp1 *db.Project, err error) {
	pp1, err = _d.Project.ByID(ctx, d, id, opts...)
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
		pp1, err = _d.Project.ByID(ctx, d, id, opts...)
	}
	return
}

// ByName implements repos.Project
func (_d ProjectWithRetry) ByName(ctx context.Context, d db.DBTX, name models.Project, opts ...db.ProjectSelectConfigOption) (pp1 *db.Project, err error) {
	pp1, err = _d.Project.ByName(ctx, d, name, opts...)
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
		pp1, err = _d.Project.ByName(ctx, d, name, opts...)
	}
	return
}
