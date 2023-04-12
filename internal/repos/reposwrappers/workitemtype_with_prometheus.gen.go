// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/prometheus.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// WorkItemTypeWithPrometheus implements repos.WorkItemType interface with all methods wrapped
// with Prometheus metrics
type WorkItemTypeWithPrometheus struct {
	base         repos.WorkItemType
	instanceName string
}

var workitemtypeDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "workitemtype_duration_seconds",
		Help:       "workitemtype runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewWorkItemTypeWithPrometheus returns an instance of the repos.WorkItemType decorated with prometheus summary metric
func NewWorkItemTypeWithPrometheus(base repos.WorkItemType, instanceName string) WorkItemTypeWithPrometheus {
	return WorkItemTypeWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// ByID implements repos.WorkItemType
func (_d WorkItemTypeWithPrometheus) ByID(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemType, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtypeDurationSummaryVec.WithLabelValues(_d.instanceName, "ByID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByID(ctx, d, id)
}

// ByName implements repos.WorkItemType
func (_d WorkItemTypeWithPrometheus) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (wp1 *db.WorkItemType, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtypeDurationSummaryVec.WithLabelValues(_d.instanceName, "ByName", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByName(ctx, d, name, projectID)
}

// Create implements repos.WorkItemType
func (_d WorkItemTypeWithPrometheus) Create(ctx context.Context, d db.DBTX, params db.WorkItemTypeCreateParams) (wp1 *db.WorkItemType, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtypeDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, d, params)
}

// Delete implements repos.WorkItemType
func (_d WorkItemTypeWithPrometheus) Delete(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemType, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtypeDurationSummaryVec.WithLabelValues(_d.instanceName, "Delete", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Delete(ctx, d, id)
}

// Update implements repos.WorkItemType
func (_d WorkItemTypeWithPrometheus) Update(ctx context.Context, d db.DBTX, id int, params db.WorkItemTypeUpdateParams) (wp1 *db.WorkItemType, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		workitemtypeDurationSummaryVec.WithLabelValues(_d.instanceName, "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Update(ctx, d, id, params)
}
