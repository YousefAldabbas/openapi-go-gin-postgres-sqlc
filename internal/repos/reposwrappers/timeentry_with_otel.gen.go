// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/otel.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TimeEntryWithTracing implements repos.TimeEntry interface instrumented with opentracing spans
type TimeEntryWithTracing struct {
	repos.TimeEntry
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewTimeEntryWithTracing returns TimeEntryWithTracing
func NewTimeEntryWithTracing(base repos.TimeEntry, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) TimeEntryWithTracing {
	d := TimeEntryWithTracing{
		TimeEntry: base,
		_instance: instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements repos.TimeEntry
func (_d TimeEntryWithTracing) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.TimeEntrySelectConfigOption) (tp1 *db.TimeEntry, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.TimeEntry.ByID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"d":    d,
				"id":   id,
				"opts": opts}, map[string]interface{}{
				"tp1": tp1,
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
	return _d.TimeEntry.ByID(ctx, d, id, opts...)
}

// Create implements repos.TimeEntry
func (_d TimeEntryWithTracing) Create(ctx context.Context, d db.DBTX, params *db.TimeEntryCreateParams) (tp1 *db.TimeEntry, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.TimeEntry.Create")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"tp1": tp1,
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
	return _d.TimeEntry.Create(ctx, d, params)
}

// Delete implements repos.TimeEntry
func (_d TimeEntryWithTracing) Delete(ctx context.Context, d db.DBTX, id int) (tp1 *db.TimeEntry, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.TimeEntry.Delete")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
				"tp1": tp1,
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
	return _d.TimeEntry.Delete(ctx, d, id)
}

// Update implements repos.TimeEntry
func (_d TimeEntryWithTracing) Update(ctx context.Context, d db.DBTX, id int, params *db.TimeEntryUpdateParams) (tp1 *db.TimeEntry, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.TimeEntry.Update")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"id":     id,
				"params": params}, map[string]interface{}{
				"tp1": tp1,
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
	return _d.TimeEntry.Update(ctx, d, id, params)
}
