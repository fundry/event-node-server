package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *teamResolver) Members(ctx context.Context, obj *model.Team) ([]*model.User, error) {
	var User []*model.User
	err := r.DB.Model(&User).Order("id").Select()

	if err != nil {
		return nil, err
	}
	return User, nil
}

func (r *teamResolver) CreatedBy(ctx context.Context, obj *model.Team) ([]*model.Event, error) {
	var createdBy []*model.Event
	err := r.DB.Model(&createdBy).Where("id = ?", obj.EventID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return createdBy, nil
}

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
