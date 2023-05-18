package entpoc

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/dsykes16/entpoc/ent"
	"github.com/dsykes16/entpoc/ent/todo"
	"github.com/rs/xid"
)

func (r *queryResolver) Node(ctx context.Context, id xid.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithFixedNodeType(todo.Table))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []xid.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithFixedNodeType(todo.Table))
}

func (r *queryResolver) Todos(ctx context.Context, after *entgql.Cursor[xid.ID], first *int, before *entgql.Cursor[xid.ID], last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// CreateTodoInput returns CreateTodoInputResolver implementation.
func (r *Resolver) CreateTodoInput() CreateTodoInputResolver { return &createTodoInputResolver{r} }

type queryResolver struct{ *Resolver }
type createTodoInputResolver struct{ *Resolver }
