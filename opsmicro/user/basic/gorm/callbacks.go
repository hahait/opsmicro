package gorm

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"gorm.io/gorm"
	"github.com/opentracing/opentracing-go/log"
)

const (
	spanKey = "spankey"
	operName = "db.gorm"
)

func GormTraceBefore(db *gorm.DB) {
	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, operName)
	db.InstanceSet(spanKey, span)
}

func GormTraceAfter(db *gorm.DB) {
	spanIn, y := db.InstanceGet(spanKey)
	if !y {
		return
	}
	span := spanIn.(opentracing.Span)
	defer span.Finish()
	ext.DBType.Set(span, db.Statement.Dialector.Name())
	span.SetTag("db.table", db.Statement.Table)
	span.SetTag("db.SQL", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...))

	if err := db.Error; err != nil {
		span.LogFields(log.String("Error", err.Error()))
	}
}

type GormTracePlugin struct {}

func (g *GormTracePlugin) Name() string {
	return "GormTracePlugin"
}

func (g *GormTracePlugin) Initialize(db *gorm.DB) error {
	createCallback := db.Callback().Create()
	createCallback.Before("gorm:create").Register("gorm:trace_before_create", GormTraceBefore)
	createCallback.After("gorm:create").Register("gorm:trace_after_create", GormTraceAfter)

	queryCallback := db.Callback().Query()
	queryCallback.Before("gorm:query").Register("gorm:trace_before_query", GormTraceBefore)
	queryCallback.After("gorm:query").Register("gorm:trace_after_query", GormTraceAfter)

	deleteCallback := db.Callback().Delete()
	deleteCallback.Before("gorm:delete").Register("gorm:trace_before_delete", GormTraceBefore)
	deleteCallback.After("gorm:delete").Register("gorm:trace_after_delete", GormTraceAfter)

	updateCallback := db.Callback().Update()
	updateCallback.Before("gorm:update").Register("gorm:trace_before_update", GormTraceBefore)
	updateCallback.After("gorm:update").Register("gorm:trace_after_update", GormTraceAfter)

	return nil
}