// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}},
		{Name: "priority", Type: field.TypeInt, Default: 0},
		{Name: "todo_children", Type: field.TypeString, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_todos_children",
				Columns:    []*schema.Column{TodosColumns[6]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = TodosTable
}
