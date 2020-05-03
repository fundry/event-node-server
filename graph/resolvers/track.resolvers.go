package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *trackResolver) Talks(ctx context.Context, obj *model.Track) ([]*model.Talk, error) {
	// var Talks []*model.Talk
	//
	// //Todo: find a way to get user id that matches event && user
	// err := r.DB.Model(&Talks).Order("id").Select()
	//
	// /*	err := r.DB.Model(&Talks).Where("event.id = ?",
	// 	obj.ID).Order("id").Select()
	// */
	//
	// if err != nil {
	// 	return nil, err
	// }
	// return Talks, nil
	//
	panic("not done")
}

func (r *trackResolver) CreatedBy(ctx context.Context, obj *model.Track) ([]*model.Event, error) {
	var Event []*model.Event

	if err := r.DB.Model(&Event).Order("id").Select(); err != nil {
		return nil, err
	}

	return Event, nil
}

// Track returns generated.TrackResolver implementation.
func (r *Resolver) Track() generated.TrackResolver { return &trackResolver{r} }

type trackResolver struct{ *Resolver }
