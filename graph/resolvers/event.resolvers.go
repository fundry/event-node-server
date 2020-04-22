package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *eventResolver) Attendees(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventResolver) CreatedBy(ctx context.Context, obj *model.Event) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
