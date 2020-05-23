package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *volunteerResolver) Team(ctx context.Context, obj *model.Volunteer) (*model.Team, error) {
	var Team *model.Team

	if err := r.DB.Model(&Team).Where("team_id = ?", obj.TeamID).Order("id").Select(); err != nil {
		return nil, err
	}

	return Team, nil
}

func (r *volunteerResolver) Event(ctx context.Context, obj *model.Volunteer) (*model.Event, error) {
	var Event *model.Event

	if err := r.DB.Model(&Event).Where("event_id = ?", obj.EventID).Order("id").Select(); err != nil {
		return nil, err
	}

	return Event, nil
}

func (r *volunteerResolver) User(ctx context.Context, obj *model.Volunteer) (*model.User, error) {
	var User *model.User

	if err := r.DB.Model(&User).Where("user_id = ?", obj.UserID).Order("id").Select(); err != nil {
		return nil, err
	}

	return User, nil
}

// Volunteer returns generated.VolunteerResolver implementation.
func (r *Resolver) Volunteer() generated.VolunteerResolver { return &volunteerResolver{r} }

type volunteerResolver struct{ *Resolver }
