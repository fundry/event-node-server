package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *eventResolver) Talk(ctx context.Context, obj *model.Event) ([]*model.EventTalk, error) {
	var eventTalk []*model.EventTalk
	err := r.DB.Model(&eventTalk).Where("event_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return eventTalk, nil
}

func (r *eventResolver) MeetupGroups(ctx context.Context, obj *model.Event) ([]*model.MeetupGroups, error) {
	var meetupGroups []*model.MeetupGroups
	err := r.DB.Model(&meetupGroups).Where("event_id = ?", obj.ID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return meetupGroups, nil
}

func (r *eventResolver) CreatedBy(ctx context.Context, obj *model.Event) ([]*model.User, error) {
	var createdBy []*model.User
	err := r.DB.Model(&createdBy).Where("id = ?", obj.AuthorID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return createdBy, nil
}

func (r *eventResolver) Attendees(ctx context.Context, obj *model.Event) ([]*model.Attendee, error) {
	var attendees []*model.Attendee

	// this currently returns all users in the DB
	if err := r.DB.Model(&attendees).Where("event_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}

	return attendees, nil
}

func (r *eventResolver) Tracks(ctx context.Context, obj *model.Event) ([]*model.Tracks, error) {
	var tracks []*model.Tracks

	if err := r.DB.Model(&tracks).Where("event_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}
	return tracks, nil
}

func (r *eventResolver) CartItemsCategory(ctx context.Context, obj *model.Event) ([]*model.Category, error) {
	var categories []*model.Category

	if err := r.DB.Model(&categories).Where("event_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *eventResolver) Teams(ctx context.Context, obj *model.Event) ([]*model.Team, error) {
	var teams []*model.Team

	if err := r.DB.Model(&teams).Where("event_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}

	return teams, nil
}

func (r *eventResolver) Volunteer(ctx context.Context, obj *model.Event) ([]*model.Volunteer, error) {
	var volunteer []*model.Volunteer

	if err := r.DB.Model(&volunteer).Where("event_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}

	return volunteer, nil
}

func (r *eventTalkResolver) Event(ctx context.Context, obj *model.EventTalk) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *eventTalkResolver) Draft(ctx context.Context, obj *model.EventTalk) ([]*model.Talk, error) {
	var talk []*model.Talk

	if err := r.DB.Model(&talk).Where("id = ?", obj.DraftID).Order("id").Select(); err != nil {
		return nil, err
	}

	return talk, nil
}

func (r *meetupGroupsResolver) Event(ctx context.Context, obj *model.MeetupGroups) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *meetupGroupsResolver) Lead(ctx context.Context, obj *model.MeetupGroups) ([]*model.User, error) {
	var meetupGroupLeader []*model.User

	if err := r.DB.Model(&meetupGroupLeader).Where("id = ?", obj.LeadID).Order("id").Select(); err != nil {
		return nil, err
	}

	return meetupGroupLeader, nil
}

func (r *meetupGroupsResolver) Members(ctx context.Context, obj *model.MeetupGroups) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

// EventTalk returns generated.EventTalkResolver implementation.
func (r *Resolver) EventTalk() generated.EventTalkResolver { return &eventTalkResolver{r} }

// MeetupGroups returns generated.MeetupGroupsResolver implementation.
func (r *Resolver) MeetupGroups() generated.MeetupGroupsResolver { return &meetupGroupsResolver{r} }

type eventResolver struct{ *Resolver }
type eventTalkResolver struct{ *Resolver }
type meetupGroupsResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *eventTalkResolver) Author(ctx context.Context, obj *model.EventTalk) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
