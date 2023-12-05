// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/otel.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// NotificationWithTracing implements repos.Notification interface instrumented with opentracing spans
type NotificationWithTracing struct {
	repos.Notification
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewNotificationWithTracing returns NotificationWithTracing
func NewNotificationWithTracing(base repos.Notification, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) NotificationWithTracing {
	d := NotificationWithTracing{
		Notification: base,
		_instance:    instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// Create implements repos.Notification
func (_d NotificationWithTracing) Create(ctx context.Context, d db.DBTX, params *db.NotificationCreateParams) (up1 *db.UserNotification, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.Notification.Create")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"up1": up1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Notification.Create(ctx, d, params)
}

// Delete implements repos.Notification
func (_d NotificationWithTracing) Delete(ctx context.Context, d db.DBTX, id db.NotificationID) (np1 *db.Notification, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.Notification.Delete")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
				"np1": np1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Notification.Delete(ctx, d, id)
}

// LatestNotifications implements repos.Notification
func (_d NotificationWithTracing) LatestNotifications(ctx context.Context, d db.DBTX, params *db.GetUserNotificationsParams) (ga1 []db.GetUserNotificationsRow, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.Notification.LatestNotifications")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"ga1": ga1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Notification.LatestNotifications(ctx, d, params)
}

// PaginatedNotifications implements repos.Notification
func (_d NotificationWithTracing) PaginatedNotifications(ctx context.Context, d db.DBTX, userID db.UserID, params models.GetPaginatedNotificationsParams) (ua1 []db.UserNotification, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.Notification.PaginatedNotifications")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"userID": userID,
				"params": params}, map[string]interface{}{
				"ua1": ua1,
				"err": err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.Notification.PaginatedNotifications(ctx, d, userID, params)
}
