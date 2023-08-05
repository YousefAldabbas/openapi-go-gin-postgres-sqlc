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

// KanbanStepWithTracing implements repos.KanbanStep interface instrumented with opentracing spans
type KanbanStepWithTracing struct {
	repos.KanbanStep
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewKanbanStepWithTracing returns KanbanStepWithTracing
func NewKanbanStepWithTracing(base repos.KanbanStep, instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) KanbanStepWithTracing {
	d := KanbanStepWithTracing{
		KanbanStep: base,
		_instance:  instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ByID implements repos.KanbanStep
func (_d KanbanStepWithTracing) ByID(ctx context.Context, d db.DBTX, id int, opts ...db.KanbanStepSelectConfigOption) (kp1 *db.KanbanStep, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.KanbanStep.ByID")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":  ctx,
				"d":    d,
				"id":   id,
				"opts": opts}, map[string]interface{}{
				"kp1": kp1,
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
	return _d.KanbanStep.ByID(ctx, d, id, opts...)
}

// ByProject implements repos.KanbanStep
func (_d KanbanStepWithTracing) ByProject(ctx context.Context, d db.DBTX, projectID int, opts ...db.KanbanStepSelectConfigOption) (ka1 []db.KanbanStep, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "repos.KanbanStep.ByProject")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx":       ctx,
				"d":         d,
				"projectID": projectID,
				"opts":      opts}, map[string]interface{}{
				"ka1": ka1,
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
	return _d.KanbanStep.ByProject(ctx, d, projectID, opts...)
}
