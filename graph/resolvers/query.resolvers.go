package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	InternalMiddlewares "github.com/vickywane/event-server/graph/middlewares"
	"github.com/vickywane/event-server/graph/model"
)

func (r *queryResolver) Event(ctx context.Context, id *int, name string) (*model.Event, error) {
	event := model.Event{ID: *id, Name: name}

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

func (r *queryResolver) User(ctx context.Context, id *int, name string) (*model.User, error) {
	User := model.User{ID: *id, Name: name}

	if err := r.DB.Select(&User); err != nil {
		return nil, err
	}

	return &User, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	gc, CtxErr := InternalMiddlewares.GinContextFromContext(ctx)

	if CtxErr != nil {
		fmt.Println(gc, "context resolver")
		fmt.Println(CtxErr, "context resolver")
	} else {
		fmt.Println(CtxErr, "context")
	}

	var Users []*model.User

	err := r.DB.Model(&Users).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Users, nil
}

func (r *queryResolver) Preference(ctx context.Context, id *int, name string) (*model.Preference, error) {
	Preference := model.Preference{ID: *id, Name: name}

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

func (r *queryResolver) File(ctx context.Context, id *int, name string) (*model.File, error) {
	File := model.File{ID: *id, Filename: name}

	if err := r.DB.Select(&File); err != nil {
		return nil, err
	}

	return &File, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*model.File, error) {
	var Files []*model.File

	err := r.DB.Model(&Files).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Files, nil
}

func (r *queryResolver) Team(ctx context.Context, id *int, name string) (*model.Team, error) {
	Team := model.Team{ID: *id, Name: name}

	if err := r.DB.Select(&Team); err != nil {
		return nil, err
	}

	return &Team, nil
}

func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	var Teams []*model.Team

	err := r.DB.Model(&Teams).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Teams, nil
}

func (r *queryResolver) Sponsor(ctx context.Context, id *int, name *string) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sponsors(ctx context.Context) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
