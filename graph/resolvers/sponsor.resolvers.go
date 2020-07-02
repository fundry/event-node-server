package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *sponsorResolver) Event(ctx context.Context, obj *model.Sponsor) (*model.Event, error) {
	var event *model.Event

	err := r.DB.Model(event).Where("sponsor_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return event, nil
}

// Sponsor returns generated.SponsorResolver implementation.
func (r *Resolver) Sponsor() generated.SponsorResolver { return &sponsorResolver{r} }

type sponsorResolver struct{ *Resolver }
