// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/prometheus.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// NotificationWithPrometheus implements repos.Notification interface with all methods wrapped
// with Prometheus metrics
type NotificationWithPrometheus struct {
	base         repos.Notification
	instanceName string
}

var notificationDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "notification_duration_seconds",
		Help:       "notification runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewNotificationWithPrometheus returns an instance of the repos.Notification decorated with prometheus summary metric
func NewNotificationWithPrometheus(base repos.Notification, instanceName string) NotificationWithPrometheus {
	return NotificationWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// Create implements repos.Notification
func (_d NotificationWithPrometheus) Create(ctx context.Context, d db.DBTX, params repos.NotificationCreateParams) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		notificationDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, d, params)
}

// Delete implements repos.Notification
func (_d NotificationWithPrometheus) Delete(ctx context.Context, d db.DBTX, notificationID int32) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		notificationDurationSummaryVec.WithLabelValues(_d.instanceName, "Delete", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Delete(ctx, d, notificationID)
}

// LatestUserNotifications implements repos.Notification
func (_d NotificationWithPrometheus) LatestUserNotifications(ctx context.Context, d db.DBTX, params repos.GetUserNotificationsParams) (ga1 []db.GetUserNotificationsRow, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		notificationDurationSummaryVec.WithLabelValues(_d.instanceName, "LatestUserNotifications", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.LatestUserNotifications(ctx, d, params)
}
