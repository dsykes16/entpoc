// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/dsykes16/entpoc/ent/schema"
	"github.com/dsykes16/entpoc/ent/todo"
	"github.com/rs/xid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoMixin := schema.Todo{}.Mixin()
	todoMixinFields0 := todoMixin[0].Fields()
	_ = todoMixinFields0
	todoMixinFields1 := todoMixin[1].Fields()
	_ = todoMixinFields1
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoMixinFields1[0].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoMixinFields1[1].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	// todo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	todo.UpdateDefaultUpdatedAt = todoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// todoDescText is the schema descriptor for text field.
	todoDescText := todoFields[0].Descriptor()
	// todo.TextValidator is a validator for the "text" field. It is called by the builders before save.
	todo.TextValidator = todoDescText.Validators[0].(func(string) error)
	// todoDescPriority is the schema descriptor for priority field.
	todoDescPriority := todoFields[2].Descriptor()
	// todo.DefaultPriority holds the default value on creation for the priority field.
	todo.DefaultPriority = todoDescPriority.Default.(int)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoMixinFields0[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() xid.ID)
}
