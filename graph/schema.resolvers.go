// Package graph ...
// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/masaok/gqlgen-todos/graph/generated"
	"github.com/masaok/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	val, _ := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	todo := &model.Todo{
		Text: input.Text,
		// ID:   fmt.Sprintf("T%d", rand.Int()),
		// ID:   fmt.Sprintf("T%d", rand.Int(rand.Reader, big.NewInt(27))),
		ID:   fmt.Sprintf("T%d", val),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Mutation ...
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query ...
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
