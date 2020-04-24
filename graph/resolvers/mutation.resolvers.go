package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/vickywane/event-server/graph/generated"
	InternalMiddleware "github.com/vickywane/event-server/graph/middlewares"
	"github.com/vickywane/event-server/graph/model"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEvent) (*model.Event, error) {
	event := model.Event{
		ID:          rand.Int(),
		Name:        input.Name,
		Description: input.Description,
		Alias:       input.Alias,
		Summary:     input.Summary,
		Email:       input.Email,
		EventType:   input.EventType,
		Website:     input.Website,
		Venue:       input.Venue,
		Date:        input.Date,
		BucketLink:  input.BucketLink,
		IsArchived:  false,
		IsLocked:    false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := r.DB.Insert(&event); err != nil {
		fmt.Println(err)
	}

	return &event, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id *int, input model.UpdateEvent) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	hashedPassword := HashPassword(input.Password)

	user := model.User{
		ID:         rand.Int(),
		Name:       input.Name,
		Email:      input.Email,
		CreatedAt:  time.Now(),
		Role:       input.Role,
		Password:   hashedPassword,
		BucketLink: input.BucketLink,
	}

	if err := r.DB.Insert(&user); err != nil {
		fmt.Println(err)
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *int, input model.UpdateUser) (*model.User, error) {
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

func (r *mutationResolver) UpdatePreference(ctx context.Context, id *int, input model.UpdatePreference) (*model.Preference, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFile(ctx context.Context, input model.CreateFile) (*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFile(ctx context.Context, id *int, input model.DeleteFile) (*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.CreateTeam) (*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, id *int, input model.UpdateTeam) (*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSponsor(ctx context.Context, input *model.CreateSponsor) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSponsor(ctx context.Context, id *int, input *model.UpdateSponsor) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthResponse, error) {
	response := model.AuthResponse{}

	auth := InternalMiddleware.AuthMiddleware.LoginResponse

	// Todo convert the result here from JSON to string
	if auth != nil {
		fmt.Println(json.Marshal(auth))
	}

	return &response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
