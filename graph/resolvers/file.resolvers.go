package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *eventFileResolver) Event(ctx context.Context, obj *model.EventFile) ([]*model.Event, error) {
	var event []*model.Event
	if err := r.DB.Model(&event).Where("id = ?", obj.EventID).Order("id").Select(); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventFileResolver) UploadedBy(ctx context.Context, obj *model.EventFile) ([]*model.User, error) {
	var uploadedBy []*model.User
	if err := r.DB.Model(&uploadedBy).Where("id = ?", obj.UserID).Order("id").Select(); err != nil {
		return nil, err
	}

	return uploadedBy, nil
}

func (r *userFileResolver) User(ctx context.Context, obj *model.UserFile) ([]*model.User, error) {
	var user []*model.User
	if err := r.DB.Model(&user).Where("id = ?", obj.UserID).Order("id").Select(); err != nil {
		return nil, err
	}

	return user, nil
}

// EventFile returns generated.EventFileResolver implementation.
func (r *Resolver) EventFile() generated.EventFileResolver { return &eventFileResolver{r} }

// UserFile returns generated.UserFileResolver implementation.
func (r *Resolver) UserFile() generated.UserFileResolver { return &userFileResolver{r} }

type eventFileResolver struct{ *Resolver }
type userFileResolver struct{ *Resolver }
