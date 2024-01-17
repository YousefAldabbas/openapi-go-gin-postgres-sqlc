// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/retry-repo.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// WorkItemTypeWithRetry implements repos.WorkItemType interface instrumented with retries
type WorkItemTypeWithRetry struct {
	repos.WorkItemType
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewWorkItemTypeWithRetry returns WorkItemTypeWithRetry
func NewWorkItemTypeWithRetry(base repos.WorkItemType, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) WorkItemTypeWithRetry {
	return WorkItemTypeWithRetry{
		WorkItemType:   base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// ByID implements repos.WorkItemType
func (_d WorkItemTypeWithRetry) ByID(ctx context.Context, d db.DBTX, id db.WorkItemTypeID, opts ...db.WorkItemTypeSelectConfigOption) (wp1 *db.WorkItemType, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemTypeWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.WorkItemType.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		_d.logger.Debugf("retry %d/%d: %s", _i+1, _d._retryCount, err)
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		if tx, ok := d.(pgx.Tx); ok {
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemTypeWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemType.ByID(ctx, d, id, opts...)
	}
	return
}

// ByName implements repos.WorkItemType
func (_d WorkItemTypeWithRetry) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID, opts ...db.WorkItemTypeSelectConfigOption) (wp1 *db.WorkItemType, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemTypeWithRetryByName")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.WorkItemType.ByName(ctx, d, name, projectID, opts...)
	if err == nil || _d._retryCount < 1 {
		return
	}
	_ticker := time.NewTicker(_d._retryInterval)
	defer _ticker.Stop()
	for _i := 0; _i < _d._retryCount && err != nil; _i++ {
		_d.logger.Debugf("retry %d/%d: %s", _i+1, _d._retryCount, err)
		select {
		case <-ctx.Done():
			return
		case <-_ticker.C:
		}
		if tx, ok := d.(pgx.Tx); ok {
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemTypeWithRetryByName"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItemType.ByName(ctx, d, name, projectID, opts...)
	}
	return
}
