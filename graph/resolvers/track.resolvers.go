package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *tracksResolver) Talks(ctx context.Context, obj *model.Tracks) ([]*model.Talk, error) {
	var Talk []*model.Talk

	if err := r.DB.Model(&Talk).Order("id").Select(); err != nil {
		return nil, err
	}

	return Talk, nil
}

func (r *tracksResolver) CreatedBy(ctx context.Context, obj *model.Tracks) ([]*model.Event, error) {
	var Event []*model.Event

	if err := r.DB.Model(&Event).Order("id").Select(); err != nil {
		return nil, err
	}

	return Event, nil
}

// Tracks returns generated.TracksResolver implementation.
func (r *Resolver) Tracks() generated.TracksResolver { return &tracksResolver{r} }

type tracksResolver struct{ *Resolver }
