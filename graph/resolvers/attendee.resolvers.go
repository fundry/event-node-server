package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *attendeeResolver) User(ctx context.Context, obj *model.Attendee) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *attendeeResolver) Event(ctx context.Context, obj *model.Attendee) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Attendee returns generated.AttendeeResolver implementation.
func (r *Resolver) Attendee() generated.AttendeeResolver { return &attendeeResolver{r} }

type attendeeResolver struct{ *Resolver }
