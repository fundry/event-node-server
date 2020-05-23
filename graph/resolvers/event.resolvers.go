package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *eventResolver) CreatedBy(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	var createdBy []*model.User
	err := r.DB.Model(&createdBy).Where("id = ?", obj.AuthorID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return createdBy, nil
}

func (r *eventResolver) Attendees(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	var attendees []*model.User

	// this currently returns all users in the DB
	if err := r.DB.Model(&attendees).Order("id").Select(); err != nil {
		return nil, err
	}

	return attendees, nil
}

func (r *eventResolver) Tracks(ctx context.Context, obj *model.Event) ([]*model.Tracks, error) {
	var tracks []*model.Tracks

	if err := r.DB.Model(&tracks).Where("event_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}

	return tracks, nil
}

func (r *eventResolver) Teams(ctx context.Context, obj *model.Event) ([]*model.Team, error) {
	var teams []*model.Team

	if err := r.DB.Model(&teams).Where("id = ?", obj.AuthorID).Order("id").Select(); err != nil {
		return nil, err
	}

	return teams, nil
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
