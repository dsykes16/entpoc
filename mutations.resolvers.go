package entpoc

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dsykes16/entpoc/ent"
	"github.com/rs/xid"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	return ent.FromContext(ctx).Todo.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id xid.ID, input ent.UpdateTodoInput) (*ent.Todo, error) {
	return ent.FromContext(ctx).Todo.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
