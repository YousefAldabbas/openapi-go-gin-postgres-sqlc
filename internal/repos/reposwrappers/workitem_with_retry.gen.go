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

// WorkItemWithRetry implements repos.WorkItem interface instrumented with retries
type WorkItemWithRetry struct {
	repos.WorkItem
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewWorkItemWithRetry returns WorkItemWithRetry
func NewWorkItemWithRetry(base repos.WorkItem, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) WorkItemWithRetry {
	return WorkItemWithRetry{
		WorkItem:       base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// AssignTag implements repos.WorkItem
func (_d WorkItemWithRetry) AssignTag(ctx context.Context, d db.DBTX, params *db.WorkItemWorkItemTagCreateParams) (err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemWithRetryAssignTag")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	err = _d.WorkItem.AssignTag(ctx, d, params)
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemWithRetryAssignTag"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		err = _d.WorkItem.AssignTag(ctx, d, params)
	}
	return
}

// AssignUser implements repos.WorkItem
func (_d WorkItemWithRetry) AssignUser(ctx context.Context, d db.DBTX, params *db.WorkItemAssignedUserCreateParams) (err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemWithRetryAssignUser")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	err = _d.WorkItem.AssignUser(ctx, d, params)
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemWithRetryAssignUser"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		err = _d.WorkItem.AssignUser(ctx, d, params)
	}
	return
}

// ByID implements repos.WorkItem
func (_d WorkItemWithRetry) ByID(ctx context.Context, d db.DBTX, id db.WorkItemID, opts ...db.WorkItemSelectConfigOption) (wp1 *db.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.WorkItem.ByID(ctx, d, id, opts...)
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItem.ByID(ctx, d, id, opts...)
	}
	return
}

// Delete implements repos.WorkItem
func (_d WorkItemWithRetry) Delete(ctx context.Context, d db.DBTX, id db.WorkItemID) (wp1 *db.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemWithRetryDelete")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.WorkItem.Delete(ctx, d, id)
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemWithRetryDelete"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItem.Delete(ctx, d, id)
	}
	return
}

// RemoveAssignedUser implements repos.WorkItem
func (_d WorkItemWithRetry) RemoveAssignedUser(ctx context.Context, d db.DBTX, memberID db.UserID, workItemID db.WorkItemID) (err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemWithRetryRemoveAssignedUser")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	err = _d.WorkItem.RemoveAssignedUser(ctx, d, memberID, workItemID)
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemWithRetryRemoveAssignedUser"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		err = _d.WorkItem.RemoveAssignedUser(ctx, d, memberID, workItemID)
	}
	return
}

// RemoveTag implements repos.WorkItem
func (_d WorkItemWithRetry) RemoveTag(ctx context.Context, d db.DBTX, tagID db.WorkItemTagID, workItemID db.WorkItemID) (err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemWithRetryRemoveTag")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	err = _d.WorkItem.RemoveTag(ctx, d, tagID, workItemID)
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemWithRetryRemoveTag"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		err = _d.WorkItem.RemoveTag(ctx, d, tagID, workItemID)
	}
	return
}

// Restore implements repos.WorkItem
func (_d WorkItemWithRetry) Restore(ctx context.Context, d db.DBTX, id db.WorkItemID) (wp1 *db.WorkItem, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT WorkItemWithRetryRestore")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	wp1, err = _d.WorkItem.Restore(ctx, d, id)
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
			if _, err = tx.Exec(ctx, "ROLLBACK to WorkItemWithRetryRestore"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}

			if _, err = tx.Exec(ctx, "BEGIN"); err != nil {
				err = fmt.Errorf("could not begin transaction after rollback: %w", err)
				return
			}
		}

		wp1, err = _d.WorkItem.Restore(ctx, d, id)
	}
	return
}
