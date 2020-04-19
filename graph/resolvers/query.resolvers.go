package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *queryResolver) Event(ctx context.Context, id int) (*model.Event, error) {
	event := model.Event{ID: id}

	if err := r.DB.Select(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	var events []*model.Event

	err := r.DB.Model(&events).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return events, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	User := model.User{ID: id}

	if err := r.DB.Select(&User); err != nil {
		return nil, err
	}

	return &User, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var Users []*model.User

	err := r.DB.Model(&Users).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Users, nil
}

func (r *queryResolver) Preference(ctx context.Context, id int) (*model.Preference, error) {
	Preference := model.Preference{ID: id}

	if err := r.DB.Select(&Preference); err != nil {
		return nil, err
	}

	return &Preference, nil
}

func (r *queryResolver) Preferences(ctx context.Context) ([]*model.Preference, error) {
	var Preferences []*model.Preference

	err := r.DB.Model(&Preferences).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Preferences, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
