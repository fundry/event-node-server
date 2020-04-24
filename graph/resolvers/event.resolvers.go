package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *eventResolver) CreatedBy(ctx context.Context, obj *model.Event) (*model.User, error) {
	var createdBy *model.User

	err := r.DB.Model(&createdBy).Where("user_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return createdBy, nil
}

func (r *eventResolver) Attendees(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	var attendees []*model.User

	err := r.DB.Model(&attendees).Where("user_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return attendees, nil
}

func (r *eventResolver) Teams(ctx context.Context, obj *model.Event) ([]*model.Team, error) {
	var teams []*model.Team

	err := r.DB.Model(&teams).Where("team_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return teams, nil
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
