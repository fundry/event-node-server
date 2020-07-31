package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
	CustomResponse "github.com/vickywane/event-server/graph/validators"
)

func (r *queryResolver) Event(ctx context.Context, id *int, name string) (*model.Event, error) {
	event := model.Event{ID: *id, Name: name}
	fmt.Println(event.AuthorID)
	// if err := r.DB.Model(&event).Column("user").Relation("CreatedBy").Select(); err != nil {
	// 	return  nil, err
	// }

	if err := r.DB.Select(&event); err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *queryResolver) Events(ctx context.Context, limit *int) ([]*model.Event, error) {
	var events []*model.Event

	if limit != nil {
		QueryErr = r.DB.Model(&events).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&events).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}

	return events, nil
}

func (r *queryResolver) Attendees(ctx context.Context, limit *int, eventID *int) ([]*model.Attendee, error) {
	var attendee []*model.Attendee
	if err := r.DB.Select(attendee); err != nil {
		return nil, err
	}

	return attendee, nil
}

func (r *queryResolver) EventTalk(ctx context.Context, limit *int, talkID int) ([]*model.Talk, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MeetupGroups(ctx context.Context, limit *int) ([]*model.MeetupGroups, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMeetupGroup(ctx context.Context, id int) (*model.MeetupGroups, error) {
	Meetup := model.MeetupGroups{ID: id}
	// if err := r.DB.Model(&event).Column("user").Relation("CreatedBy").Select(); err != nil {
	// 	return  nil, err
	// }

	if err := r.DB.Select(&Meetup); err != nil {
		return nil, err
	}

	return &Meetup, nil
}

func (r *queryResolver) GetEventTalks(ctx context.Context, areApproved bool, limit *int, eventID *int) ([]*model.EventTalk, error) {
	var talks []*model.EventTalk

	if err := r.DB.Model(&model.Event{}).Where("id = ?", eventID).First(); err != nil {
		return nil, errors.New("event doesnt exist or not found ")
	}

	// err := r.DB.Model(talks).Where("is_accepted = ?", areApproved).Limit(*limit).Select(&talks)
	if err := r.DB.Model(&talks).Where("is_accepted = ?", areApproved).Limit(*limit).Select(); err != nil {
		return nil, err
	}

	return talks, nil
}

func (r *queryResolver) Stream(ctx context.Context, id int) (*model.Stream, error) {
	stream := &model.Stream{ID: id}

	if err := r.DB.Select(stream); err != nil {
		return nil, err
	}

	return stream, nil
}

func (r *queryResolver) Streams(ctx context.Context, limit *int) ([]*model.Stream, error) {
	var streams []*model.Stream

	if limit != nil {
		QueryErr = r.DB.Model(&streams).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&streams).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}

	return streams, nil
}

func (r *queryResolver) User(ctx context.Context, id *int, name string) (*model.User, error) {
	User := model.User{ID: *id}

	if err := r.DB.Select(&User); err != nil {
		return nil, err
	}
	// if err := r.DB.Model(&User).Column("user.event_id").Relation("Events").Select(); err != nil {
	// 	return nil, err
	// }

	return &User, nil
}

func (r *queryResolver) Users(ctx context.Context, limit *int) ([]*model.User, error) {
	// gc, CtxErr := InternalMiddleware.GinContextFromContext(ctx)

	// if CtxErr != nil {
	// 	fmt.Println(gc, "context resolver")
	// 	fmt.Println(CtxErr, "context resolver")
	// } else {
	// 	fmt.Println(CtxErr, "context")
	// }

	var Users []*model.User

	if limit != nil {
		QueryErr = r.DB.Model(&Users).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&Users).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}

	return Users, nil
}

func (r *queryResolver) EventSettings(ctx context.Context, eventID int) (*model.EventSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserFile(ctx context.Context, id *int, name string) (*model.UserFile, error) {
	File := model.UserFile{ID: *id, FileURI: name}

	if err := r.DB.Select(&File); err != nil {
		return nil, err
	}

	return &File, nil
}

func (r *queryResolver) UserFiles(ctx context.Context) ([]*model.UserFile, error) {
	var Files []*model.UserFile

	err := r.DB.Model(&Files).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Files, nil
}

func (r *queryResolver) EventFile(ctx context.Context, id *int, name string) (*model.EventFile, error) {
	File := model.EventFile{ID: *id, FileURI: name}

	if err := r.DB.Select(&File); err != nil {
		return nil, err
	}

	return &File, nil
}

func (r *queryResolver) EventFiles(ctx context.Context) ([]*model.EventFile, error) {
	var Files []*model.EventFile

	err := r.DB.Model(&Files).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Files, nil
}

func (r *queryResolver) Team(ctx context.Context, id *int, name string) (*model.Team, error) {
	Team := model.Team{ID: *id, Name: name}

	if err := r.DB.Select(&Team); err != nil {
		return nil, err
	}

	return &Team, nil
}

func (r *queryResolver) Teams(ctx context.Context, limit *int) ([]*model.Team, error) {
	var Teams []*model.Team

	if limit != nil {
		QueryErr = r.DB.Model(&Teams).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&Teams).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}

	return Teams, nil
}

func (r *queryResolver) Sponsor(ctx context.Context, id *int, name *string) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sponsors(ctx context.Context, limit *int) ([]*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Task(ctx context.Context, id *int) (*model.Tasks, error) {
	Task := model.Tasks{ID: *id}

	if err := r.DB.Select(&Task); err != nil {
		return nil, err
	}

	return &Task, nil
}

func (r *queryResolver) Tasks(ctx context.Context, limit *int) ([]*model.Tasks, error) {
	var Tasks []*model.Tasks

	err := r.DB.Model(&Tasks).Select()

	if err != nil {
		fmt.Println("some err here")
	}

	return Tasks, nil
}

func (r *queryResolver) Talk(ctx context.Context, id int) (*model.Talk, error) {
	talk := model.Talk{ID: id}

	if err := r.DB.Select(&talk); err != nil {
		return nil, err
	}

	return &talk, nil
}

func (r *queryResolver) Talks(ctx context.Context, limit *int) ([]*model.Talk, error) {
	var talk []*model.Talk

	if limit != nil {
		QueryErr = r.DB.Model(&talk).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&talk).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}

	return talk, nil
}

func (r *queryResolver) Track(ctx context.Context, id int) (*model.Tracks, error) {
	track := model.Tracks{ID: id}

	if err := r.DB.Select(&track); err != nil {
		return nil, err
	}

	return &track, nil
}

func (r *queryResolver) Tracks(ctx context.Context, limit *int) ([]*model.Tracks, error) {
	var track []*model.Tracks

	if limit != nil {
		QueryErr = r.DB.Model(&track).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&track).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}
	return track, nil
}

func (r *queryResolver) Volunteer(ctx context.Context, id int) (*model.Volunteer, error) {
	volunteer := model.Volunteer{ID: id}

	if err := r.DB.Select(&volunteer); err != nil {
		return nil, err
	}

	return &volunteer, nil
}

func (r *queryResolver) Volunteers(ctx context.Context, limit *int, eventID int) ([]*model.Volunteer, error) {
	var Volunteer []*model.Volunteer

	if limit != nil {
		QueryErr = r.DB.Model(&Volunteer).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&Volunteer).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}
	return Volunteer, nil
}

func (r *queryResolver) CartItems(ctx context.Context, categoryID int, limit *int) ([]*model.CartItem, error) {
	// var CartItem []*model.CartItem{ID: categoryID}
	var CartItem []*model.CartItem

	if limit != nil {
		QueryErr = r.DB.Model(&CartItem).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&CartItem).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}
	return CartItem, nil
}

func (r *queryResolver) AllCartItems(ctx context.Context, limit *int) ([]*model.CartItem, error) {
	var CartItem []*model.CartItem

	if limit != nil {
		QueryErr = r.DB.Model(&CartItem).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&CartItem).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}
	return CartItem, nil
}

func (r *queryResolver) Purchases(ctx context.Context, eventID int, limit *int) ([]*model.Purchases, error) {
	// var purchases []*model.Purchases
	//
	// if limit != nil {
	// 	QueryErr = r.DB.Model(&purchases).Limit(*limit).Select()
	// } else {
	// 	QueryErr = r.DB.Model(&purchases).Select()
	// }
	//
	// if QueryErr != nil {
	// 	return nil, CustomResponse.QueryError
	// }
	// return purchases, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllPurchases(ctx context.Context, limit *int) ([]*model.Purchases, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Category(ctx context.Context, id int, limit *int) ([]*model.Category, error) {
	var Category []*model.Category

	if limit != nil {
		QueryErr = r.DB.Model(&Category).Limit(*limit).Where("v = ?", id).Select()
	} else {
		QueryErr = r.DB.Model(&Category).Where("id = ?", id).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}
	return Category, nil
}

func (r *queryResolver) TaskComment(ctx context.Context, taskID int) ([]*model.TaskComments, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BugReports(ctx context.Context, limit *int) (*model.BugReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FeatureRequests(ctx context.Context, limit *int) (*model.BugReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Reminder(ctx context.Context, userID int) (*model.Reminder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var (
	QueryErr interface{} = nil
)
