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

// UserWithRetry implements repos.User interface instrumented with retries
type UserWithRetry struct {
	repos.User
	_retryCount    int
	_retryInterval time.Duration
	logger         *zap.SugaredLogger
}

// NewUserWithRetry returns UserWithRetry
func NewUserWithRetry(base repos.User, logger *zap.SugaredLogger, retryCount int, retryInterval time.Duration) UserWithRetry {
	return UserWithRetry{
		User:           base,
		_retryCount:    retryCount,
		_retryInterval: retryInterval,
		logger:         logger,
	}
}

// ByAPIKey implements repos.User
func (_d UserWithRetry) ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryByAPIKey")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.ByAPIKey(ctx, d, apiKey)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByAPIKey")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryByAPIKey"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.ByAPIKey(ctx, d, apiKey)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByAPIKey")
	}
	return
}

// ByEmail implements repos.User
func (_d UserWithRetry) ByEmail(ctx context.Context, d db.DBTX, email string, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryByEmail")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.ByEmail(ctx, d, email, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByEmail")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryByEmail"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.ByEmail(ctx, d, email, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByEmail")
	}
	return
}

// ByExternalID implements repos.User
func (_d UserWithRetry) ByExternalID(ctx context.Context, d db.DBTX, extID string, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryByExternalID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.ByExternalID(ctx, d, extID, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByExternalID")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryByExternalID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.ByExternalID(ctx, d, extID, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByExternalID")
	}
	return
}

// ByID implements repos.User
func (_d UserWithRetry) ByID(ctx context.Context, d db.DBTX, id db.UserID, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryByID")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.ByID(ctx, d, id, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByID")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryByID"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.ByID(ctx, d, id, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByID")
	}
	return
}

// ByProject implements repos.User
func (_d UserWithRetry) ByProject(ctx context.Context, d db.DBTX, projectID db.ProjectID) (ua1 []db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryByProject")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	ua1, err = _d.User.ByProject(ctx, d, projectID)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByProject")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryByProject"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		ua1, err = _d.User.ByProject(ctx, d, projectID)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByProject")
	}
	return
}

// ByTeam implements repos.User
func (_d UserWithRetry) ByTeam(ctx context.Context, d db.DBTX, teamID db.TeamID) (ua1 []db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryByTeam")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	ua1, err = _d.User.ByTeam(ctx, d, teamID)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByTeam")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryByTeam"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		ua1, err = _d.User.ByTeam(ctx, d, teamID)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByTeam")
	}
	return
}

// ByUsername implements repos.User
func (_d UserWithRetry) ByUsername(ctx context.Context, d db.DBTX, username string, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryByUsername")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.ByUsername(ctx, d, username, opts...)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByUsername")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryByUsername"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.ByUsername(ctx, d, username, opts...)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryByUsername")
	}
	return
}

// Create implements repos.User
func (_d UserWithRetry) Create(ctx context.Context, d db.DBTX, params *db.UserCreateParams) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryCreate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.Create(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryCreate")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryCreate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.Create(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryCreate")
	}
	return
}

// CreateAPIKey implements repos.User
func (_d UserWithRetry) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (up1 *db.UserAPIKey, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryCreateAPIKey")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.CreateAPIKey(ctx, d, user)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryCreateAPIKey")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryCreateAPIKey"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.CreateAPIKey(ctx, d, user)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryCreateAPIKey")
	}
	return
}

// Delete implements repos.User
func (_d UserWithRetry) Delete(ctx context.Context, d db.DBTX, id db.UserID) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryDelete")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.Delete(ctx, d, id)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryDelete")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryDelete"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.Delete(ctx, d, id)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryDelete")
	}
	return
}

// DeleteAPIKey implements repos.User
func (_d UserWithRetry) DeleteAPIKey(ctx context.Context, d db.DBTX, apiKey string) (up1 *db.UserAPIKey, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryDeleteAPIKey")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.DeleteAPIKey(ctx, d, apiKey)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryDeleteAPIKey")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryDeleteAPIKey"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.DeleteAPIKey(ctx, d, apiKey)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryDeleteAPIKey")
	}
	return
}

// Paginated implements repos.User
func (_d UserWithRetry) Paginated(ctx context.Context, d db.DBTX, params repos.GetPaginatedUsersParams) (ua1 []db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryPaginated")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	ua1, err = _d.User.Paginated(ctx, d, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryPaginated")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryPaginated"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		ua1, err = _d.User.Paginated(ctx, d, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryPaginated")
	}
	return
}

// Update implements repos.User
func (_d UserWithRetry) Update(ctx context.Context, d db.DBTX, id db.UserID, params *db.UserUpdateParams) (up1 *db.User, err error) {
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "SAVEPOINT UserWithRetryUpdate")
		if err != nil {
			err = fmt.Errorf("could not store savepoint: %w", err)
			return
		}
	}
	up1, err = _d.User.Update(ctx, d, id, params)
	if err == nil || _d._retryCount < 1 {
		if tx, ok := d.(pgx.Tx); ok {
			_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryUpdate")
		}
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
			if _, err = tx.Exec(ctx, "ROLLBACK to UserWithRetryUpdate"); err != nil {
				err = fmt.Errorf("could not rollback to savepoint: %w", err)
				return
			}
		}

		up1, err = _d.User.Update(ctx, d, id, params)
	}
	if tx, ok := d.(pgx.Tx); ok {
		_, err = tx.Exec(ctx, "RELEASE SAVEPOINT UserWithRetryUpdate")
	}
	return
}
