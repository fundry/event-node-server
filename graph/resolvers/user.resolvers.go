package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *userResolver) Talks(ctx context.Context, obj *model.User) ([]*model.Talk, error) {
	var talks []*model.Talk

	if err := r.DB.Model(&talks).Where("speaker_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}

	return talks, nil
}

func (r *userResolver) Events(ctx context.Context, obj *model.User) ([]*model.Event, error) {
	var events []*model.Event

	err := r.DB.Model(&events).Where("author_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *userResolver) Streams(ctx context.Context, obj *model.User) ([]*model.Stream, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Reminders(ctx context.Context, obj *model.User) ([]*model.Reminder, error) {
	var reminder []*model.Reminder
	if err := r.DB.Model(&reminder).Where("user_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}
	return reminder, nil
}

func (r *userResolver) Files(ctx context.Context, obj *model.User) ([]*model.UserFile, error) {
	var file []*model.UserFile

	err := r.DB.Model(&file).Where("user_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, err
	}
	return file, nil
}

func (r *userResolver) Volunteering(ctx context.Context, obj *model.User) ([]*model.Volunteer, error) {
	var volunteer []*model.Volunteer

	err := r.DB.Model(&volunteer).Where("user_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, err
	}
	return volunteer, nil
}

func (r *userResolver) Attending(ctx context.Context, obj *model.User) ([]*model.Attendee, error) {
	var attendee []*model.Attendee

	err := r.DB.Model(&attendee).Where("user_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, err
	}
	return attendee, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
