// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// ProjectWithTimeout implements repos.Project interface instrumented with timeouts
type ProjectWithTimeout struct {
	repos.Project
	config ProjectWithTimeoutConfig
}

type ProjectWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	ByNameTimeout time.Duration
}

// NewProjectWithTimeout returns ProjectWithTimeout
func NewProjectWithTimeout(base repos.Project, config ProjectWithTimeoutConfig) ProjectWithTimeout {
	return ProjectWithTimeout{
		Project: base,
		config:  config,
	}
}

// ByID implements repos.Project
func (_d ProjectWithTimeout) ByID(ctx context.Context, d db.DBTX, id int) (pp1 *db.Project, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.Project.ByID(ctx, d, id)
}

// ByName implements repos.Project
func (_d ProjectWithTimeout) ByName(ctx context.Context, d db.DBTX, name internalmodels.Project) (pp1 *db.Project, err error) {
	var cancelFunc func()
	if _d.config.ByNameTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByNameTimeout)
		defer cancelFunc()
	}
	return _d.Project.ByName(ctx, d, name)
}
