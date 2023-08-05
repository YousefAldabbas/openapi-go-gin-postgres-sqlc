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

// WorkItemTypeWithTimeout implements repos.WorkItemType interface instrumented with timeouts
type WorkItemTypeWithTimeout struct {
	repos.WorkItemType
	config WorkItemTypeWithTimeoutConfig
}

type WorkItemTypeWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	ByNameTimeout time.Duration
}

// NewWorkItemTypeWithTimeout returns WorkItemTypeWithTimeout
func NewWorkItemTypeWithTimeout(base repos.WorkItemType, config WorkItemTypeWithTimeoutConfig) WorkItemTypeWithTimeout {
	return WorkItemTypeWithTimeout{
		WorkItemType: base,
		config:       config,
	}
}

// ByID implements repos.WorkItemType
func (_d WorkItemTypeWithTimeout) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.WorkItemTypeSelectConfigOption) (wp1 *db.WorkItemType, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.WorkItemType.ByID(ctx, d, id, opts...)
}

// ByName implements repos.WorkItemType
func (_d WorkItemTypeWithTimeout) ByName(ctx context.Context, d db.DBTX, name string, projectID int, opts ...db.WorkItemTypeSelectConfigOption) (wp1 *db.WorkItemType, err error) {
	var cancelFunc func()
	if _d.config.ByNameTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByNameTimeout)
		defer cancelFunc()
	}
	return _d.WorkItemType.ByName(ctx, d, name, projectID, opts...)
}
