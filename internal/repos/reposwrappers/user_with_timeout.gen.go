// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// UserWithTimeout implements repos.User interface instrumented with timeouts
type UserWithTimeout struct {
	repos.User
	config UserWithTimeoutConfig
}

type UserWithTimeoutConfig struct {
	ByAPIKeyTimeout time.Duration

	ByEmailTimeout time.Duration

	ByExternalIDTimeout time.Duration

	ByIDTimeout time.Duration

	ByProjectTimeout time.Duration

	ByTeamTimeout time.Duration

	ByUsernameTimeout time.Duration

	CreateTimeout time.Duration

	CreateAPIKeyTimeout time.Duration

	DeleteTimeout time.Duration

	DeleteAPIKeyTimeout time.Duration

	UpdateTimeout time.Duration
}

// NewUserWithTimeout returns UserWithTimeout
func NewUserWithTimeout(base repos.User, config UserWithTimeoutConfig) UserWithTimeout {
	return UserWithTimeout{
		User:   base,
		config: config,
	}
}

// ByAPIKey implements repos.User
func (_d UserWithTimeout) ByAPIKey(ctx context.Context, d db.DBTX, apiKey string) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.ByAPIKeyTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByAPIKeyTimeout)
		defer cancelFunc()
	}
	return _d.User.ByAPIKey(ctx, d, apiKey)
}

// ByEmail implements repos.User
func (_d UserWithTimeout) ByEmail(ctx context.Context, d db.DBTX, email string, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.ByEmailTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByEmailTimeout)
		defer cancelFunc()
	}
	return _d.User.ByEmail(ctx, d, email, opts...)
}

// ByExternalID implements repos.User
func (_d UserWithTimeout) ByExternalID(ctx context.Context, d db.DBTX, extID string, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.ByExternalIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByExternalIDTimeout)
		defer cancelFunc()
	}
	return _d.User.ByExternalID(ctx, d, extID, opts...)
}

// ByID implements repos.User
func (_d UserWithTimeout) ByID(ctx context.Context, d db.DBTX, id db.UserID, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.User.ByID(ctx, d, id, opts...)
}

// ByProject implements repos.User
func (_d UserWithTimeout) ByProject(ctx context.Context, d db.DBTX, projectID db.ProjectID) (ua1 []db.User, err error) {
	var cancelFunc func()
	if _d.config.ByProjectTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByProjectTimeout)
		defer cancelFunc()
	}
	return _d.User.ByProject(ctx, d, projectID)
}

// ByTeam implements repos.User
func (_d UserWithTimeout) ByTeam(ctx context.Context, d db.DBTX, teamID db.TeamID) (ua1 []db.User, err error) {
	var cancelFunc func()
	if _d.config.ByTeamTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByTeamTimeout)
		defer cancelFunc()
	}
	return _d.User.ByTeam(ctx, d, teamID)
}

// ByUsername implements repos.User
func (_d UserWithTimeout) ByUsername(ctx context.Context, d db.DBTX, username string, opts ...db.UserSelectConfigOption) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.ByUsernameTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByUsernameTimeout)
		defer cancelFunc()
	}
	return _d.User.ByUsername(ctx, d, username, opts...)
}

// Create implements repos.User
func (_d UserWithTimeout) Create(ctx context.Context, d db.DBTX, params *db.UserCreateParams) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.CreateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateTimeout)
		defer cancelFunc()
	}
	return _d.User.Create(ctx, d, params)
}

// CreateAPIKey implements repos.User
func (_d UserWithTimeout) CreateAPIKey(ctx context.Context, d db.DBTX, user *db.User) (up1 *db.UserAPIKey, err error) {
	var cancelFunc func()
	if _d.config.CreateAPIKeyTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.CreateAPIKeyTimeout)
		defer cancelFunc()
	}
	return _d.User.CreateAPIKey(ctx, d, user)
}

// Delete implements repos.User
func (_d UserWithTimeout) Delete(ctx context.Context, d db.DBTX, id db.UserID) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.DeleteTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteTimeout)
		defer cancelFunc()
	}
	return _d.User.Delete(ctx, d, id)
}

// DeleteAPIKey implements repos.User
func (_d UserWithTimeout) DeleteAPIKey(ctx context.Context, d db.DBTX, apiKey string) (up1 *db.UserAPIKey, err error) {
	var cancelFunc func()
	if _d.config.DeleteAPIKeyTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.DeleteAPIKeyTimeout)
		defer cancelFunc()
	}
	return _d.User.DeleteAPIKey(ctx, d, apiKey)
}

// Update implements repos.User
func (_d UserWithTimeout) Update(ctx context.Context, d db.DBTX, id db.UserID, params *db.UserUpdateParams) (up1 *db.User, err error) {
	var cancelFunc func()
	if _d.config.UpdateTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.UpdateTimeout)
		defer cancelFunc()
	}
	return _d.User.Update(ctx, d, id, params)
}
