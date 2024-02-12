//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var CacheDemoTwoWorkItems = newCacheDemoTwoWorkItemsTable("public", "cache__demo_two_work_items", "")

type cacheDemoTwoWorkItemsTable struct {
	postgres.Table

	// Columns
	CustomDateForProject2 postgres.ColumnTimestampz
	WorkItemID            postgres.ColumnInteger
	Title                 postgres.ColumnString
	Description           postgres.ColumnString
	WorkItemTypeID        postgres.ColumnInteger
	Metadata              postgres.ColumnString
	TeamID                postgres.ColumnInteger
	KanbanStepID          postgres.ColumnInteger
	ClosedAt              postgres.ColumnTimestampz
	TargetDate            postgres.ColumnTimestampz
	CreatedAt             postgres.ColumnTimestampz
	UpdatedAt             postgres.ColumnTimestampz
	DeletedAt             postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CacheDemoTwoWorkItemsTable struct {
	cacheDemoTwoWorkItemsTable

	EXCLUDED cacheDemoTwoWorkItemsTable
}

// AS creates new CacheDemoTwoWorkItemsTable with assigned alias
func (a CacheDemoTwoWorkItemsTable) AS(alias string) *CacheDemoTwoWorkItemsTable {
	return newCacheDemoTwoWorkItemsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CacheDemoTwoWorkItemsTable with assigned schema name
func (a CacheDemoTwoWorkItemsTable) FromSchema(schemaName string) *CacheDemoTwoWorkItemsTable {
	return newCacheDemoTwoWorkItemsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CacheDemoTwoWorkItemsTable with assigned table prefix
func (a CacheDemoTwoWorkItemsTable) WithPrefix(prefix string) *CacheDemoTwoWorkItemsTable {
	return newCacheDemoTwoWorkItemsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CacheDemoTwoWorkItemsTable with assigned table suffix
func (a CacheDemoTwoWorkItemsTable) WithSuffix(suffix string) *CacheDemoTwoWorkItemsTable {
	return newCacheDemoTwoWorkItemsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCacheDemoTwoWorkItemsTable(schemaName, tableName, alias string) *CacheDemoTwoWorkItemsTable {
	return &CacheDemoTwoWorkItemsTable{
		cacheDemoTwoWorkItemsTable: newCacheDemoTwoWorkItemsTableImpl(schemaName, tableName, alias),
		EXCLUDED:                   newCacheDemoTwoWorkItemsTableImpl("", "excluded", ""),
	}
}

func newCacheDemoTwoWorkItemsTableImpl(schemaName, tableName, alias string) cacheDemoTwoWorkItemsTable {
	var (
		CustomDateForProject2Column = postgres.TimestampzColumn("custom_date_for_project_2")
		WorkItemIDColumn            = postgres.IntegerColumn("work_item_id")
		TitleColumn                 = postgres.StringColumn("title")
		DescriptionColumn           = postgres.StringColumn("description")
		WorkItemTypeIDColumn        = postgres.IntegerColumn("work_item_type_id")
		MetadataColumn              = postgres.StringColumn("metadata")
		TeamIDColumn                = postgres.IntegerColumn("team_id")
		KanbanStepIDColumn          = postgres.IntegerColumn("kanban_step_id")
		ClosedAtColumn              = postgres.TimestampzColumn("closed_at")
		TargetDateColumn            = postgres.TimestampzColumn("target_date")
		CreatedAtColumn             = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn             = postgres.TimestampzColumn("updated_at")
		DeletedAtColumn             = postgres.TimestampzColumn("deleted_at")
		allColumns                  = postgres.ColumnList{CustomDateForProject2Column, WorkItemIDColumn, TitleColumn, DescriptionColumn, WorkItemTypeIDColumn, MetadataColumn, TeamIDColumn, KanbanStepIDColumn, ClosedAtColumn, TargetDateColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
		mutableColumns              = postgres.ColumnList{CustomDateForProject2Column, TitleColumn, DescriptionColumn, WorkItemTypeIDColumn, MetadataColumn, TeamIDColumn, KanbanStepIDColumn, ClosedAtColumn, TargetDateColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
	)

	return cacheDemoTwoWorkItemsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CustomDateForProject2: CustomDateForProject2Column,
		WorkItemID:            WorkItemIDColumn,
		Title:                 TitleColumn,
		Description:           DescriptionColumn,
		WorkItemTypeID:        WorkItemTypeIDColumn,
		Metadata:              MetadataColumn,
		TeamID:                TeamIDColumn,
		KanbanStepID:          KanbanStepIDColumn,
		ClosedAt:              ClosedAtColumn,
		TargetDate:            TargetDateColumn,
		CreatedAt:             CreatedAtColumn,
		UpdatedAt:             UpdatedAtColumn,
		DeletedAt:             DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
