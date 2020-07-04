package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *reminderResolver) User(ctx context.Context, obj *model.Reminder) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Reminder returns generated.ReminderResolver implementation.
func (r *Resolver) Reminder() generated.ReminderResolver { return &reminderResolver{r} }

type reminderResolver struct{ *Resolver }
