package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *attendeeResolver) User(ctx context.Context, obj *model.Attendee) ([]*model.User, error) {
	var Attendee []*model.User

	if err := r.DB.Model(&Attendee).Where("id = ?", obj.UserID).Order("id").Select(); err != nil {
		return nil, err
	}

	return Attendee, nil
}

func (r *attendeeResolver) Event(ctx context.Context, obj *model.Attendee) ([]*model.Event, error) {
	var Event []*model.Event

	if err := r.DB.Model(&Event).Where("id = ?mu", obj.EventID).Order("id").Select(); err != nil {
		return nil, err
	}

	return Event, nil
}

// Attendee returns generated.AttendeeResolver implementation.
func (r *Resolver) Attendee() generated.AttendeeResolver { return &attendeeResolver{r} }

type attendeeResolver struct{ *Resolver }
