// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap-templates/otel.tmpl
// gowrap: http://github.com/hexdigest/gowrap

package reposwrappers

import (
	"context"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// WorkItemTagWithTracing implements repos.WorkItemTag interface instrumented with opentracing spans
type WorkItemTagWithTracing struct {
	repos.WorkItemTag
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewWorkItemTagWithTracing returns WorkItemTagWithTracing
func NewWorkItemTagWithTracing(base repos.WorkItemTag, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) WorkItemTagWithTracing {
	d := WorkItemTagWithTracing{
		WorkItemTag: base,
		_instance:   instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements repos.WorkItemTag
func (_d WorkItemTagWithTracing) ByID(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemTag, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemTag.ByID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
				"wp1": wp1,
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
	return _d.WorkItemTag.ByID(ctx, d, id)
}

// ByName implements repos.WorkItemTag
func (_d WorkItemTagWithTracing) ByName(ctx context.Context, d db.DBTX, name string, projectID int) (wp1 *db.WorkItemTag, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemTag.ByName")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":       ctx,
				"d":         d,
				"name":      name,
				"projectID": projectID}, map[string]interface{}{
				"wp1": wp1,
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
	return _d.WorkItemTag.ByName(ctx, d, name, projectID)
}

// Create implements repos.WorkItemTag
func (_d WorkItemTagWithTracing) Create(ctx context.Context, d db.DBTX, params repos.WorkItemTagCreateParams) (wp1 *db.WorkItemTag, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemTag.Create")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"params": params}, map[string]interface{}{
				"wp1": wp1,
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
	return _d.WorkItemTag.Create(ctx, d, params)
}

// Delete implements repos.WorkItemTag
func (_d WorkItemTagWithTracing) Delete(ctx context.Context, d db.DBTX, id int) (wp1 *db.WorkItemTag, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemTag.Delete")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"d":   d,
				"id":  id}, map[string]interface{}{
				"wp1": wp1,
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
	return _d.WorkItemTag.Delete(ctx, d, id)
}

// Update implements repos.WorkItemTag
func (_d WorkItemTagWithTracing) Update(ctx context.Context, d db.DBTX, id int, params repos.WorkItemTagUpdateParams) (wp1 *db.WorkItemTag, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.WorkItemTag.Update")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":    ctx,
				"d":      d,
				"id":     id,
				"params": params}, map[string]interface{}{
				"wp1": wp1,
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
	return _d.WorkItemTag.Update(ctx, d, id, params)
}
