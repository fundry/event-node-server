package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	InternalMiddleware "github.com/vickywane/event-server/graph/middlewares"
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

func (r *queryResolver) User(ctx context.Context, id *int, name string) (*model.User, error) {
	User := model.User{ID: *id, Name: name}

	if err := r.DB.Select(&User); err != nil {
		return nil, err
	}
	// if err := r.DB.Model(&User).Column("user.event_id").Relation("Events").Select(); err != nil {
	// 	return nil, err
	// }

	return &User, nil
}

func (r *queryResolver) Users(ctx context.Context, limit *int) ([]*model.User, error) {
	gc, CtxErr := InternalMiddleware.GinContextFromContext(ctx)

	if CtxErr != nil {
		fmt.Println(gc, "context resolver")
		fmt.Println(CtxErr, "context resolver")
	} else {
		fmt.Println(CtxErr, "context")
	}

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

func (r *queryResolver) Preference(ctx context.Context, id *int, name string) (*model.Preference, error) {
	Preference := model.Preference{ID: *id, Name: name}

	if err := r.DB.Select(&Preference); err != nil {
		return nil, err
	}

	return &Preference, nil
}

func (r *queryResolver) Preferences(ctx context.Context, limit *int) ([]*model.Preference, error) {
	var Preferences []*model.Preference

	if limit != nil {
		QueryErr = r.DB.Model(&Preferences).Limit(*limit).Select()
	} else {
		QueryErr = r.DB.Model(&Preferences).Select()
	}

	if QueryErr != nil {
		return nil, CustomResponse.QueryError
	}

	return Preferences, nil
}

func (r *queryResolver) File(ctx context.Context, id *int, name string) (*model.File, error) {
	File := model.File{ID: *id, FileURI: name}

	if err := r.DB.Select(&File); err != nil {
		return nil, err
	}

	return &File, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*model.File, error) {
	var Files []*model.File

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
