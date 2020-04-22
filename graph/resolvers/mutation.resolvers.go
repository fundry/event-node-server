package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEvent) (*model.Event, error) {
	event := model.Event{
		ID:          rand.Int(),
		Name:        input.Name,
		Description: input.Description,
		Email:       input.Email,
		CreatedAt:   time.Now(),
	}

	if err := r.DB.Insert(&event); err != nil {
		fmt.Println(err)
	}

	return &event, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id *int, input *model.UpdateEvent) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	user := model.User{
		ID:        rand.Int(),
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: time.Now(),
		Role:      input.Role,
	}

	if err := r.DB.Insert(&user); err != nil {
		fmt.Println(err)
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *int, input *model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePreference(ctx context.Context, input model.CreatePreference) (*model.Preference, error) {
	preference := model.Preference{
		ID:        rand.Int(),
		Name:      input.Name,
		CreatedAt: time.Now(),
	}

	if err := r.DB.Insert(&preference); err != nil {
		fmt.Println(err)
	}

	return &preference, nil
}

func (r *mutationResolver) UpdatePreference(ctx context.Context, id *int, input *model.UpdatePreference) (*model.Preference, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
