// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// KanbanStepWithTimeout implements repos.KanbanStep interface instrumented with timeouts
type KanbanStepWithTimeout struct {
	repos.KanbanStep
	config KanbanStepWithTimeoutConfig
}

type KanbanStepWithTimeoutConfig struct {
	ByIDTimeout time.Duration

	ByProjectTimeout time.Duration
}

// NewKanbanStepWithTimeout returns KanbanStepWithTimeout
func NewKanbanStepWithTimeout(base repos.KanbanStep, config KanbanStepWithTimeoutConfig) KanbanStepWithTimeout {
	return KanbanStepWithTimeout{
		KanbanStep: base,
		config:     config,
	}
}

// ByID implements repos.KanbanStep
func (_d KanbanStepWithTimeout) ByID(ctx context.Context, d models.DBTX, id models.KanbanStepID, opts ...models.KanbanStepSelectConfigOption) (kp1 *models.KanbanStep, err error) {
	var cancelFunc func()
	if _d.config.ByIDTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByIDTimeout)
		defer cancelFunc()
	}
	return _d.KanbanStep.ByID(ctx, d, id, opts...)
}

// ByProject implements repos.KanbanStep
func (_d KanbanStepWithTimeout) ByProject(ctx context.Context, d models.DBTX, projectID models.ProjectID, opts ...models.KanbanStepSelectConfigOption) (ka1 []models.KanbanStep, err error) {
	var cancelFunc func()
	if _d.config.ByProjectTimeout > 0 {
		ctx, cancelFunc = context.WithTimeout(ctx, _d.config.ByProjectTimeout)
		defer cancelFunc()
	}
	return _d.KanbanStep.ByProject(ctx, d, projectID, opts...)
}
