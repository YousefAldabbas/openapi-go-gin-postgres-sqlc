// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	repomodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// ProjectBoardWithRetry implements repos.ProjectBoard interface instrumented with retries
type ProjectBoardWithRetry struct {
	repos.ProjectBoard
	_retryCount    int
	_retryInterval time.Duration
}

// NewProjectBoardWithRetry returns ProjectBoardWithRetry
func NewProjectBoardWithRetry(base repos.ProjectBoard, retryCount int, retryInterval time.Duration) ProjectBoardWithRetry {
	return ProjectBoardWithRetry{
		ProjectBoard:   base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
	}
}

// ByID implements repos.ProjectBoard
func (_d ProjectBoardWithRetry) ByID(ctx context.Context, d db.DBTX, projectID int) (pp1 *repomodels.ProjectBoard, err error) {
	pp1, err = _d.ProjectBoard.ByID(ctx, d, projectID)
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
		pp1, err = _d.ProjectBoard.ByID(ctx, d, projectID)
	}
	return
}

// Create implements repos.ProjectBoard
func (_d ProjectBoardWithRetry) Create(ctx context.Context, d db.DBTX, params repos.ProjectBoardCreateParams) (pp1 *repomodels.ProjectBoard, err error) {
	pp1, err = _d.ProjectBoard.Create(ctx, d, params)
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
		pp1, err = _d.ProjectBoard.Create(ctx, d, params)
	}
	return
}
