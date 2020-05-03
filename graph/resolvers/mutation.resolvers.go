package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
	"github.com/vickywane/event-server/graph/validators"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEvent, userID int) (*model.Event, error) {
	mockbucketLink := string(rand.Int())
	Time := time.Now()

	event := model.Event{
		ID:          rand.Int(),
		Name:        input.Name,
		Description: input.Description,
		Alias:       input.Alias,
		Summary:     input.Summary,
		Email:       input.Email,
		EventType:   input.EventType,
		Website:     input.Website,
		Venue:       input.Venue,
		Date:        input.Date,
		AuthorID:    userID,
		BucketLink:  mockbucketLink,
		IsArchived:  false,
		IsLocked:    false,
		CreatedAt:   Time,
		UpdatedAt:   Time,
	}

	if err := r.DB.Insert(&event); err != nil {
		fmt.Println(err)
	}

	return &event, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id *int, input model.UpdateEvent) (*model.Event, error) {
	event, err := r.GetEventById(*id)

	if event != nil && err != nil {
		return nil, validators.NotFound
	}

	event, err = r.UpdateCurrentEvent(event)

	validators.LengthChecker(event.Name, 5)

	if len(event.Description) < 3 {
		return nil, validators.ShortInput
	} else {
		event.Description = *input.Description
	}

	if len(event.Website) < 3 {
		return nil, validators.ShortInput
	} else {
		event.Website = *input.Website
	}

	if len(event.Alias) < 1 {
		return nil, validators.ShortInput
	} else {
		event.Alias = *input.Alias
	}

	if len(event.Summary) < 6 {
		return nil, validators.ShortInput
	} else {
		event.Alias = *input.Alias
	}

	if len(event.Venue) < 4 {
		return nil, validators.ShortInput
	} else {
		event.Venue = *input.Venue
	}

	if len(event.Website) < 6 {
		return nil, validators.ShortInput
	} else {
		event.Website = *input.Website
	}

	if err != nil {
		return nil, validators.ErrorUpdating
	}

	return event, nil
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id int) (bool, error) {
	event, err := r.GetEventById(id)
	if event != nil && err != nil {
		return false, validators.NotFound
	}
	err = r.DeleteCurrentEvent(event)
	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	hashedPassword := HashPassword(input.Password)
	mockbucketLink := string(rand.Int())
	EventID := rand.Int()
	// 1976235410884491574
	user := model.User{
		ID:         rand.Int(),
		Name:       input.Name,
		Email:      input.Email,
		CreatedAt:  time.Now(),
		Role:       input.Role,
		Password:   hashedPassword,
		BucketLink: mockbucketLink,
		EventID:    EventID, // my event FK
	}

	if err := r.DB.Insert(&user); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *int, input model.UpdateUser) (*model.User, error) {
	user, err := r.GetUserById(*id)
	if user != nil && err != nil {
		return nil, validators.NotFound
	}

	// Todo: i need a better way to check if d field is nil
	// Todo: i need to be able to check if an input is in a field or skip the field
	// Todo: try REGEX
	if len(*input.Name) < 3 {
		return nil, validators.ShortInput
	} else {
		user.Name = *input.Name
	}

	if len(*input.Email) < 3 {
		return nil, validators.ShortInput
	} else {
		user.Email = *input.Email
	}

	if len(*input.Password) < 3 {
		return nil, validators.ShortInput
	} else {
		Password := HashPassword(*input.Password)
		user.Password = Password
	}

	if len(*input.Role) > 3 {
		user.Role = input.Role
	}

	user, err = r.UpdateCurrentUser(user)

	if err != nil {
		return nil, validators.ErrorUpdating
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	user, err := r.GetUserById(id)
	if user != nil && err != nil {
		return false, validators.NotFound
	}

	err = r.DeleteCurrentUser(user)

	return true, nil
}

func (r *mutationResolver) CreatePreference(ctx context.Context, input model.CreatePreference) (*model.Preference, error) {
	// testing GIN context here

	preference := model.Preference{
		ID:        rand.Int(),
		Name:      input.Name,
		CreatedAt: time.Now(),
	}

	if err := r.DB.Insert(&preference); err != nil {
		return nil, validators.ErrorInserting
	}

	return &preference, nil
}

func (r *mutationResolver) UpdatePreference(ctx context.Context, id *int, input model.UpdatePreference) (*model.Preference, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePreference(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFile(ctx context.Context, input model.CreateFile) (*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFile(ctx context.Context, id *int, input model.DeleteFile) (*model.File, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.CreateTeam, eventID int) (*model.Team, error) {
	team := model.Team{
		ID:        rand.Int(),
		Name:      input.Name,
		Goal:      input.Goal,
		EventID:   eventID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.DB.Insert(&team); err != nil {
		return nil, err
	}

	return &team, nil
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, id *int, input model.UpdateTeam) (*model.Team, error) {
	team, err := r.GetTeamById(*id)
	if team != nil && err != nil {
		return nil, validators.NotFound
	}

	if len(input.Name) < 4 {
		return nil, validators.ShortInput
	} else {
		team.Name = input.Name
	}

	if len(input.Goal) < 10 {
		return nil, validators.ShortInput
	} else {
		team.Goal = input.Goal
	}

	return team, nil
}

func (r *mutationResolver) DeleteTeam(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSponsor(ctx context.Context, input model.CreateSponsor) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSponsor(ctx context.Context, id *int, input model.UpdateSponsor) (*model.Sponsor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSponsor(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthResponse, error) {
	if len(input.Email) < 5 {
		panic(fmt.Errorf("email too short"))
	}

	user, err := r.GetUserByEmail(input.Email)

	if err != nil {
		fmt.Println(user)
		fmt.Println(err)
	}

	var auth *model.AuthResponse

	return auth, err
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTasks) (*model.Tasks, error) {
	task := model.Tasks{
		ID:        rand.Int(),
		Name:      input.Name,
		Type:      input.Type,
		CreatedAt: time.Now(),
	}

	if err := r.DB.Insert(&task); err != nil {
		fmt.Println(err)
	}

	return &task, nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id int, input model.UpdateTask) (*model.Tasks, error) {
	task, err := r.GetTaskById(id)

	if task != nil && err != nil {
		return nil, validators.NotFound
	}

	if len(input.Name) < 5 {
		task.Name = input.Name
	}

	if len(input.Type) < 2 {
		task.Type = input.Type
	}

	task.Type = input.Type
	task, err = r.UpdateCurrentTask(task)

	if err != nil {
		return nil, validators.ErrorUpdating
	}

	return task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id int) (bool, error) {
	task, err := r.GetTaskById(id)
	if task != nil && err != nil {
		return false, validators.NotFound
	}
	err = r.DeleteCurrentTask(task)
	return true, nil
}

func (r *mutationResolver) CreateTalk(ctx context.Context, input model.CreateTalk, userID int) (*model.Talk, error) {
	Talk := model.Talk{
		ID:           rand.Int(),
		Title:        input.Title,
		SpeakerID:    userID,
		TalkCoverURI: nil,
		Summary:      input.Summary,
		Description:  input.Description,
		Reviewers:    nil,
		Archived:     false,
		Duration:     input.Duration,
		Tags:         nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := r.DB.Insert(&Talk); err != nil {
		return nil, validators.ErrorInserting
	}

	return &Talk, nil
}

func (r *mutationResolver) UpdateTalk(ctx context.Context, id int, input model.UpdateTalk) (*model.Talk, error) {
	talk, err := r.GetTalkById(id)

	if talk != nil && err != nil {
		return nil, validators.NotFound
	}

	talk, err = r.UpdateCurrentTalk(talk)

	if len(input.Title) < 3 {
		return nil, validators.ShortInput
	} else {
		talk.Title = input.Title
	}

	if len(input.Summary) < 10 {
		return nil, validators.ShortInput
	} else {
		talk.Summary = input.Summary
	}

	if len(input.Description) < 20 {
		return nil, validators.ShortInput
	} else {
		talk.Description = input.Description
	}

	if input.Duration < 1 {
		return nil, validators.ShortInput
	} else {
		talk.Description = input.Description
	}

	talk.Archived = input.Archived

	return talk, nil
}

func (r *mutationResolver) DeleteTalk(ctx context.Context, id int) (bool, error) {
	talk, err := r.GetTalkById(id)

	if talk != nil && err != nil {
		return false, validators.NotFound
	}

	err = r.DeleteCurrentTalk(talk)

	return true, nil
}

func (r *mutationResolver) CreateTrack(ctx context.Context, input model.CreateTrack, eventID int) (*model.Track, error) {
	track := model.Track{
		ID:          rand.Int(),
		Name:        input.Name,
		Duration:    input.Duration,
		TotalTalks:  input.TotalTalks,
		IsCompleted: false,
		Archived:    false,
		EventID:     eventID,
	}

	if err := r.DB.Insert(track); err != nil {
		return nil, err
	}

	return &track, nil
}

func (r *mutationResolver) UpdateTrack(ctx context.Context, id int, input model.UpdateTrack) (*model.Track, error) {
	track, err := r.GetTrackById(id)

	if track != nil && err != nil {
		return nil, validators.NotFound
	}

	if len(input.Name) < 10 {
		return nil, validators.ShortInput
	} else {
		track.Name = input.Name
	}

	if len(input.Duration) < 5 {
		return nil, validators.ShortInput
	} else {
		track.Duration = input.Duration
	}

	track, err = r.UpdateCurrentTrack(track)

	if err != nil {
		return nil, validators.ErrorUpdating
	}

	return track, nil
}

func (r *mutationResolver) DeleteTrack(ctx context.Context, id int) (bool, error) {
	track, err := r.GetTrackById(id)

	if track != nil && err != nil {
		return false, validators.NotFound
	}

	err = r.DeleteCurrentTrack(track)

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
