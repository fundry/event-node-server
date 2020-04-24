package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *userResolver) Events(ctx context.Context, obj *model.User) ([]*model.Event, error) {
	var events []*model.Event

	err := r.DB.Model(&events).Where("event_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return events, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
