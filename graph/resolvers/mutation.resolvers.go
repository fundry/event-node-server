package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
	"github.com/vickywane/event-server/graph/validators"
	"google.golang.org/api/googleapi"
)

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := r.GetUserByEmail(input.Email)

	userPassword := user.Password // existing user password
	if user != nil && err != nil {
		fmt.Println(err)
		return nil, validators.LoginError
	}

	if compareErr := ComparePassword(userPassword, input.Password); compareErr != nil {
		return nil, validators.LoginError
	}

	token, err := GenToken(string(user.ID))
	return &model.AuthResponse{
		ID:        time.Now().Nanosecond(),
		Token:     token,
		ExpiredAt: time.Now(),
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEvent, userID int) (*model.Event, error) {
	// currentUser, err := middlewares.ExtractCurrentUserFromContext(ctx)
	// if err != nil {
	//     fmt.Println(err)
	//     return nil, validators.Unauthorized
	// }
	// fmt.Println(currentUser, err, "current user")

	// if _, err := SendEmail("vickywane@gmail.com", "concatenate", "Create Event"); err != nil {
	//     return nil, errors.Errorf("Error from email: %v", err)
	// }
	// }

	eventId := time.Now().Nanosecond()

	BucketName, err := CreateBucket(eventId)
	// fmt.Printf("Bucket %v \n", BucketName)
	fmt.Printf("Error from Bucket %v \n", err)
	if valid, err := validators.DataLength(9, input.Email, "Email"); !valid && err != nil {
		return nil, err
	}

	if validators.CheckMail(input.Email) == false {
		return nil, validators.InvalidEmail
	}

	event := model.Event{
		ID:             eventId,
		Name:           input.Name,
		Description:    input.Description,
		Alias:          input.Alias,
		Summary:        input.Summary,
		Email:          input.Email,
		EventType:      input.EventType,
		Website:        input.Website,
		Venue:          input.Venue,
		EventDate:      input.EventDate,
		MediaLinks:     input.MediaLinks,
		AuthorID:       userID,
		DateCreated:    time.Now().Format("01-02-2006"),
		TotalAttendees: 0,
		MeetupGroups: append([]*model.MeetupGroups{}, &model.MeetupGroups{
			ID: time.Now().Nanosecond(),
		}),
		BucketLink:            "",
		BucketName:            BucketName,
		IsVirtual:             input.IsVirtual,
		IsAcceptingAttendees:  false,
		IsAcceptingTalks:      false,
		IsAcceptingVolunteers: false,
		ConfirmedEmail:        false,
		IsArchived:            false,
		Actions:               []string{fmt.Sprintf("Event was created by USER on %v", time.Now().Format("01-02-2006"))},
		IsLocked:              false,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),

		MobileOnboarding:      false,
		InvitationsOnboarding: false,
		TeamsOnboarding:       false,
		ScheduleOnboarding:    false,
		MarketplaceOnboarding: false,
	}

	settings := &model.EventSettings{
		ID:                          time.Now().Nanosecond(),
		EventID:                     eventId,
		ShowWelcomeMeetupGroup:      true,
		ShowTeamInstruction:         true,
		ShowInvitationInstruction:   true,
		ShowWelcomeEventInstruction: true,
		EventThemeColour:            nil,
	}

	if !r.CheckEventFieldExists("name", input.Name) {
		return nil, validators.FieldTaken("name")
	}

	if !r.CheckEventFieldExists("email", input.Email) {
		return nil, validators.FieldTaken("email")
	}

	if err := r.DB.Insert(&event); err != nil {
		return nil, validators.ErrorInserting
	}

	if err := r.DB.Insert(settings); err != nil {
		fmt.Printf("Error : %v", err)
		return nil, validators.ErrorInserting
	}

	//if !r.CheckEventFieldExists("description", *input.Description) {
	//	return nil, validators.FieldTaken("description")
	//}
	//
	//if len(*input.Summary) > 150 {
	//	return nil, validators.ValueExceeded(150)
	//}
	//
	//if len(*input.Description) > 1500 {
	//	return nil, validators.ValueExceeded(1500)
	//}

	return &event, nil
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id int, input model.UpdateEvent) (*model.Event, error) {
	event, err := r.GetEventById(id)
	if event != nil && err != nil {
		return nil, validators.NotFound
	}

	if validators.CheckMail(*input.Email) == false {
		return nil, validators.InvalidEmail
	}

	event.UpdatedAt = time.Now()
	event.MediaLinks = input.MediaLinks

	if valid, err := validators.BoolRequired(*input.IsVirtual, "isVirtual"); valid && err == nil {
		event.IsVirtual = *input.IsVirtual
	} else {
		return nil, err
	}

	if valid, err := validators.BoolRequired(*input.IsAcceptingTalks, "isAcceptingTalks"); valid && err == nil {
		event.IsAcceptingTalks = *input.IsAcceptingTalks
	} else {
		return nil, err
	}

	if valid, err := validators.BoolRequired(*input.IsAcceptingVolunteers, "isAcceptingVolunteers"); valid && err == nil {
		event.IsAcceptingVolunteers = *input.IsAcceptingVolunteers
	} else {
		return nil, err
	}

	if valid, err := validators.BoolRequired(*input.IsLocked, "isLocked"); valid && err == nil {
		event.IsLocked = *input.IsLocked
	} else {
		return nil, err
	}

	if valid, err := validators.BoolRequired(*input.IsArchived, "isArchived"); valid && err == nil {
		event.IsArchived = *input.IsArchived
	} else {
		return nil, err
	}

	if valid, err := validators.DataLength(9, *input.Email, "Email"); valid && err == nil {
		event.Email = *input.Email
	}

	if valid, err := validators.DataLength(9, *input.Name, "Email"); valid && err == nil {
		event.Name = *input.Name
	}

	event.SpeakerConduct = input.SpeakerConduct

	if valid, err := validators.DataLength(9, *input.Description, "Description"); valid && err == nil {
		event.Description = input.Description
	} else {
		return nil, err
	}

	if valid, err := validators.DataLength(9, *input.Website, "Website"); valid && err == nil {
		event.Website = *input.Website
	} else {
		return nil, err
	}

	if valid, err := validators.DataLength(2, *input.Alias, "Alias"); valid && err == nil {
		event.Alias = *input.Alias
	} else {
		return nil, err
	}

	if valid, err := validators.DataLength(9, *input.Summary, "Summary"); valid && err == nil {
		event.Summary = input.Summary
	} else {
		return nil, err
	}

	if valid, err := validators.DataLength(6, *input.Venue, "Venue"); valid && err == nil {
		event.Venue = input.Venue
	} else {
		return nil, err
	}

	date := time.Now().Format("01-02-2006")
	action := fmt.Sprintf("This event was updated on %v", date)
	event.Actions = append(event.Actions, action)
	event.Actions = append(event.Actions, action)
	if err := r.DB.Update(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	if event, err = r.UpdateCurrentEvent(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	return event, nil
}

func (r *mutationResolver) AttendEvent(ctx context.Context, eventID int, userID int) (*model.Attendee, error) {
	attendee := &model.Attendee{
		ID:         time.Now().Nanosecond(),
		DateJoined: time.Now().Format("01-02-2006"),
		UserID:     userID,
		EventID:    eventID,
	}

	event, err := r.GetEventById(eventID)

	// validate user && event exists
	if event != nil && err != nil {
		return nil, validators.ValueNotFound("event")
	}

	user, UserErr := r.GetUserById(userID)
	if user != nil && UserErr != nil {
		return nil, validators.ValueNotFound("user")
	}
	// ==============================>

	if !r.CheckAttendeeFieldExists("user_id", userID) {
		return nil, validators.FieldTaken("user")
	}

	if err := r.DB.Insert(attendee); err != nil {
		return nil, err
	}

	return attendee, nil
}

func (r *mutationResolver) UpdateEventAttendee(ctx context.Context, eventID int, userID int) (*model.Attendee, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id int) (bool, error) {
	event, err := r.GetEventById(id)
	if event != nil && err != nil {
		return false, validators.NotFound
	}
	err = r.DeleteCurrentEvent(event)
	return true, nil
}

func (r *mutationResolver) SubmitEventTalk(ctx context.Context, talkID int, eventID int, input *model.SubmitEventTalk) (*model.EventTalk, error) {
	talk := &model.EventTalk{
		ID:            time.Now().Nanosecond(),
		IsAccepted:    input.IsAccepted,
		DateSubmitted: time.Now().Format("01-02-2006"),
		DraftID:       talkID,
		EventID:       eventID,
	}

	if exist := r.CheckEventTalkFieldExists("draft_id", talkID); !exist {
		return nil, validators.FieldTaken("draft")
	}

	if err := r.DB.Insert(talk); err != nil {
		return nil, validators.ErrorInserting
	}

	return talk, nil
}

func (r *mutationResolver) UpdateSubmittedTalk(ctx context.Context, talkID int, reviewerID *int, input model.UpdateSubmittedTalk) (*model.EventTalk, error) {
	talk, err := r.GetSubmittedTalkById(talkID)

	if talk != nil && err != nil {
		return nil, validators.NotFound
	}

	talk.ReviewerID = reviewerID
	talk.IsAccepted = input.IsAccepted
	talk.DateAccepted = time.Now().Format("01-02-2006")

	if input.Comment != nil {
		if valid, err := validators.DataLength(15, *input.Comment, "Commnet"); valid && err == nil {
			talk.Comment = input.Comment
		} else {
			return nil, err
		}
	}

	if err := r.DB.Update(talk); err != nil {
		return nil, validators.ErrorUpdating
	}

	return talk, nil
}

func (r *mutationResolver) CreateMeetupGroup(ctx context.Context, eventID int, leadID int, input *model.CreateMeetupGroup) (*model.MeetupGroups, error) {
	event, err := r.GetEventById(eventID)
	if event != nil && err != nil {
		return nil, validators.NotFound
	}

	meetupGroup := &model.MeetupGroups{
		ID:          time.Now().Nanosecond(),
		Name:        input.Name,
		EventID:     eventID,
		LeadID:      leadID,
		Description: input.Description,
		MediaLinks:  input.MediaLinks,
		Summary:     fmt.Sprintf("%v meetup group for %v", event.Name, input.Location),
		Location:    input.Location,
		Alias:       input.Alias,
		CreatedAt:   time.Now().Format("01-02-2006"),
	}

	if !r.CheckMeetupFieldExists("name", input.Name) {
		return nil, validators.FieldTaken("name")
	}

	if !r.CheckMeetupFieldExists("alias", input.Alias) {
		return nil, validators.FieldTaken("alias")
	}

	if err := r.DB.Insert(meetupGroup); err != nil {
		return nil, validators.ErrorInserting
	}

	date := time.Now().Format("01-02-2006")
	action := fmt.Sprintf("A new Meetup Group was launched on %v", date)
	event.Actions = append(event.Actions, action)

	if err := r.DB.Update(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	if event, err = r.UpdateCurrentEvent(event); err != nil {
		return nil, validators.ErrorUpdating

	}

	return meetupGroup, nil
}

func (r *mutationResolver) UpdateMeetupGroup(ctx context.Context, id int) (*model.MeetupGroups, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateStream(ctx context.Context, input model.CreateStream, userID int) (*model.Stream, error) {
	if user, err := r.GetUserById(userID); user != nil && err != nil {
		return nil, validators.ValueNotFound("user field")
	}

	stream := &model.Stream{
		ID:          time.Now().Nanosecond(),
		Title:       input.Title,
		UserID:      userID,
		Summary:     input.Summary,
		Duration:    input.Duration,
		StreamLinks: input.StreamLinks,
		CreatedAt:   time.Now().Format("01-02-2006"),
	}

	if err := r.DB.Insert(stream); err != nil {
		return nil, validators.ErrorInserting
	}

	return stream, nil
}

func (r *mutationResolver) UpdateStream(ctx context.Context, id int, input model.UpdateStream) (*model.Stream, error) {
	stream, err := r.GetStreamById(id)

	if stream != nil && err != nil {
		return nil, validators.ValueNotFound("stream")
	}

	stream.Title = input.Title
	stream.Summary = input.Summary
	stream.NotesID = input.NotesID
	stream.Duration = input.Duration
	stream.StreamLinks = input.StreamLinks

	if stream, err = r.UpdateCurrentStream(stream); err != nil {
		return nil, validators.ErrorUpdating
	}

	return stream, nil
}

func (r *mutationResolver) DeleteStream(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEventModals(ctx context.Context, id int, eventID int, input model.UpdateEventModals) (*model.EventSettings, error) {
	if event, err := r.GetEventById(eventID); event != nil && err != nil {
		return nil, validators.ValueNotFound("event")
	}

	eventSetting, err := r.GetSettingById(id) // set

	if eventSetting != nil && err != nil {
		return nil, validators.ErrorUpdating
	}

	eventSetting.EventThemeColour = input.EventThemeColour
	eventSetting.ShowInvitationInstruction = *input.ShowInvitationInstruction
	eventSetting.ShowWelcomeEventInstruction = *input.ShowWelcomeEventInstruction
	eventSetting.ShowWelcomeMeetupGroup = *input.ShowWelcomeMeetupGroup
	eventSetting.ShowTeamInstruction = *input.ShowTeamInstruction

	if _, err := r.UpdateCurrentEventSetting(eventSetting); err != nil {
		return nil, validators.ErrorUpdating
	}

	return eventSetting, nil
}

func (r *mutationResolver) UpdateEventSettings(ctx context.Context, eventID int, input model.UpdateEventSettings) (*model.Event, error) {
	event, err := r.GetEventById(eventID)

	if event != nil && err != nil {
		return nil, validators.ValueNotFound("event")
	}

	event.IsLocked = input.IsLocked
	event.IsArchived = input.IsArchived

	if _, err := r.UpdateCurrentEvent(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	return event, nil
}

func (r *mutationResolver) CreateSponsor(ctx context.Context, input model.CreateSponsor, eventID int) (*model.Sponsor, error) {
	event, err := r.GetEventById(eventID)

	if event != nil && err != nil {
		return nil, validators.NotFound
	}

	sponsor := &model.Sponsor{
		ID:                time.Now().Nanosecond(),
		Name:              input.Name,
		Type:              input.Type,
		ImageURL:          input.ImageURL,
		EventID:           eventID,
		SponsorshipStatus: "Awaiting Confirmation",
		IsOrganization:    input.IsOrganization,
	}

	if err := r.DB.Insert(sponsor); err != nil {
		return nil, validators.InsertError(err)
	}

	date := time.Now().Format("01-02-2006")
	action := fmt.Sprintf("%v began sponsoring this event %v", input.Name, date)
	event.Actions = append(event.Actions, action)
	if err := r.DB.Update(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	return sponsor, nil
}

func (r *mutationResolver) UpdateSponsor(ctx context.Context, id *int, input model.UpdateSponsor) (*model.Sponsor, error) {
	sponsor, err := r.GetSponsorById(*id)

	if err != nil {
		return nil, validators.NotFound
	}

	sponsor.Name = input.Name
	sponsor.SponsorshipStatus = input.SponsorhipStatus
	sponsor.ImageURL = input.ImageURL
	sponsor.Type = input.Type

	if err := r.DB.Update(sponsor); err != nil {
		return nil, validators.ErrorUpdating
	}

	return sponsor, nil
}

func (r *mutationResolver) DeleteSponsor(ctx context.Context, id int) (bool, error) {
	sponsor, err := r.GetSponsorById(id)

	if err != nil && sponsor != nil {
		return false, validators.NotFound
	}

	err = r.DeleteCurrentSponsor(sponsor)

	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.AuthResponse, error) {
	hashedPassword := HashPassword(input.Password)
	EventID := time.Now().Nanosecond()
	token, tokenGenErr := GenToken(input.Name)

	if validators.CheckMail(input.Email) == false {
		return nil, validators.InvalidEmail
	}
	if len(input.Name) < 7 {
		return nil, validators.ShortInput
	}
	if _, err := r.GetUserByEmail(input.Email); err == nil {
		return nil, validators.EmailTaken
	}

	// this doesnt return non-nil. BUG!
	if tokenGenErr == nil {
		fmt.Println(token, "User Token")
		return nil, validators.TokenGenerationError
	}

	UserID := time.Now().Nanosecond()
	BucketName, err := CreateBucket(UserID)
	user := model.User{
		ID:         UserID,
		Name:       input.Name,
		Email:      input.Email,
		CreatedAt:  time.Now(),
		Password:   hashedPassword,
		BucketName: BucketName,
		FileID:     nil,
		EventID:    EventID, // my event FK
	}

	transaction, err := r.DB.Begin()
	if err != nil {
		panic("transaction wasn't opened")
	}

	defer transaction.Rollback()
	// Todo: Fix - this transaction statement doesnt work!
	//     if _, err := transaction.Model(&user).Returning("*").Insert(&user); err != nil {
	//         fmt.Println("Error happened here")
	//         return nil, errb
	//     }

	if err := r.DB.Insert(&user); err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		ID:        UserID,
		Token:     token,
		ExpiredAt: time.Now(),
		User:      &user,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *int, input model.UpdateUser) (*model.User, error) {
	user, err := r.GetUserById(*id)
	if user != nil && err != nil {
		return nil, validators.NotFound
	}

	if validators.CheckMail(*input.Email) == false {
		return nil, validators.InvalidEmail
	}

	if input.Name != nil {
		if len(*input.Name) < 7 {
			return nil, validators.ShortInput
		} else {
			user.Name = *input.Name
		}
	}

	if input.Email != nil {
		if len(*input.Email) < 7 {
			return nil, validators.ShortInput
		} else {
			user.Email = *input.Email
		}
	}

	if input.Password != nil {
		if len(*input.Password) < 6 {
			return nil, validators.ShortInput
		} else {
			Password := HashPassword(*input.Password)
			user.Password = Password
		}
	}

	if input.ImgURI != nil {
		if len(*input.ImgURI) < 6 {
			return nil, validators.ShortInput
		} else {
			user.ImgURI = input.ImgURI
		}
	}
	user.UpdatedAt = time.Now().UTC()

	if user, err = r.UpdateCurrentUser(user); err != nil {
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

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.CreateTeam, eventID int) (*model.Team, error) {
	event, err := r.GetEventById(eventID)
	if err != nil {
		return nil, validators.NotFound
	}

	if !r.CheckEventTeamFieldExists("name", input.Name) {
		return nil, validators.FieldTaken("team name")
	}

	team := model.Team{
		ID:        time.Now().Nanosecond(),
		Name:      input.Name,
		Goal:      input.Goal,
		EventID:   eventID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.DB.Insert(&team); err != nil {
		return nil, err
	}

	date := time.Now().Format("01-02-2006")
	action := fmt.Sprintf("A new team was created on %v", date)
	event.Actions = append(event.Actions, action)
	if err := r.DB.Update(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	// teamss := []*model.Team{ &team }
	//
	// for _, team := range teamss {
	// 	TeamChan  <- team
	// }
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
	team, err := r.GetTeamById(id)

	if team != nil && err != nil {
		return false, validators.ValueNotFound("team data")
	}

	event, err := r.GetEventById(team.EventID)

	date := time.Now().Format("01-02-2006")
	action := fmt.Sprintf("%v was deleted on %v", team.Name, date)

	event.Actions = append(event.Actions, action)

	if err = r.DeleteCurrentTeam(team); err != nil {
		return false, validators.ErrorUpdating
	}

	// teamss := []*model.Team{ team }
	//
	// for _, team := range teamss {
	// 	TeamChan  <- team
	// }

	return true, nil
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTasks, teamID int, userID int) (*model.Tasks, error) {
	if team, err := r.GetTeamById(teamID); team != nil && err != nil {
		return nil, validators.ValueNotFound("team")
	}

	if !r.CheckTaskFieldExists("name", input.Name) {
		return nil, validators.FieldTaken("task name")
	}

	task := model.Tasks{
		ID:        time.Now().Nanosecond(),
		Name:      input.Name,
		Category:  input.Category,
		CreatedAt: time.Now().Format("01-02-2006"),
		TeamID:    teamID,
		Status:    input.Status,
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

	if len(input.Category) < 2 {
		task.Category = input.Category
	}

	task.Category = input.Category
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
		ID:           time.Now().Nanosecond(),
		Title:        input.Title,
		SpeakerID:    userID,
		TalkCoverURI: input.TalkCoverURI,
		Summary:      input.Summary,
		Description:  input.Description,
		Archived:     false,
		Duration:     input.Duration,
		Tags:         input.Tags,
		CreatedAt:    time.Now().Format("01-02-2006"),
		UpdatedAt:    "",
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

	if len(*input.Summary) < 10 {
		return nil, validators.ShortInput
	} else {
		talk.Summary = *input.Summary
	}

	if len(*input.Description) < 20 {
		return nil, validators.ShortInput
	} else {
		talk.Description = *input.Description
	}

	if len(*input.Duration) < 20 {
		return nil, validators.ShortInput
	} else {
		talk.Duration = *input.Duration
	}

	talk.Archived = *input.Archived

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

func (r *mutationResolver) CreateTrack(ctx context.Context, input model.CreateTrack, eventID int) (*model.Tracks, error) {
	event, err := r.GetEventById(eventID)
	if err != nil {
		return nil, validators.NotFound
	}
	track := model.Tracks{
		ID:          time.Now().Nanosecond(),
		Name:        input.Name,
		Duration:    input.Duration,
		TotalTalks:  input.TotalTalks,
		IsCompleted: false,
		Archived:    false,
		Summary:     input.Summary,
		EventID:     eventID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	date := time.Now().Format("01-02-2006")
	action := fmt.Sprintf("A new track was added on %v", date)
	event.Actions = append(event.Actions, action)
	if err := r.DB.Update(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	if err := r.DB.Insert(&track); err != nil {
		return nil, err
	}

	return &track, nil
}

func (r *mutationResolver) UpdateTrack(ctx context.Context, id int, input model.UpdateTrack) (*model.Tracks, error) {
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
	panic("likely err here")
}

func (r *mutationResolver) DeleteTrack(ctx context.Context, id int) (bool, error) {
	track, err := r.GetTrackById(id)

	if track != nil && err != nil {
		return false, validators.NotFound
	}

	err = r.DeleteCurrentTrack(track)

	return true, nil
}

func (r *mutationResolver) UploadSingleUserFile(ctx context.Context, req model.UploadFile, bucketName string) (*model.UserFile, error) {
	upload, err := UploadFileToBucket(bucketName, req.File, req.File.Filename)

	fmt.Printf("UPLOAD %v \n", upload)
	fmt.Printf("ERROR %v \n", err)
	location := fmt.Sprintf("https://storage.googleapis.com/%v/%v", bucketName, req.File.Filename)

	file := model.UserFile{
		ID:        time.Now().Nanosecond(),
		FileURI:   location,
		File:      req.File, // to be removed when i get FILE_URI
		Type:      req.Type,
		UserID:    *req.UserID,
		Timestamp: time.Now(),
	}

	if err := r.DB.Insert(&file); err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *mutationResolver) UploadMultipleUserFiles(ctx context.Context, req []*model.UploadFile) ([]*model.UserFile, error) {
	files := model.UserFile{
		ID:        time.Now().Nanosecond(),
		Timestamp: time.Now(),
	}
	fmt.Println(files)

	return nil, nil
}

func (r *mutationResolver) UploadSingleEventFile(ctx context.Context, req model.UploadFile, bucketName string) (*model.EventFile, error) {
	// Todo: this function is faulty
	if existsErr := r.CheckEventFile("file ->> 'Filename'", req.File.Filename); !existsErr {
		return nil, validators.FieldTaken("file")
	}

	upload, err := UploadFileToBucket(bucketName, req.File, req.File.Filename)
	fmt.Printf("UPLOAD %v \n", upload)
	fmt.Printf("ERROR %v \n", err)

	// way to handle err using gcloud
	if e, ok := err.(*googleapi.Error); ok {
		fmt.Println(e.Code, "comma ok")
		fmt.Println(e.Message, "comma ok")
		fmt.Println(e.Body, "comma ok")

		if e.Code == 400 {
			fmt.Println(e, "trying comma ok")
		}
	}
	// Todo Ensure filenames dont contain whitespaces && follow object Naming conv
	location := fmt.Sprintf("https://storage.googleapis.com/%v/%v", bucketName, req.File.Filename)
	fmt.Println(location)
	file := model.EventFile{
		ID:        time.Now().Nanosecond(),
		FileURI:   location,
		Type:      req.Type,
		File:      req.File,
		EventID:   req.EventID,
		UserID:    *req.UserID,
		Timestamp: time.Now(),
	}

	if err := r.DB.Insert(&file); err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *mutationResolver) UploadMultipleEventFiles(ctx context.Context, req []*model.UploadFile) ([]*model.EventFile, error) {
	files := model.EventFile{
		ID:        time.Now().Nanosecond(),
		Timestamp: time.Now(),
	}
	fmt.Println(files)

	return nil, nil
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateVolunteer(ctx context.Context, input model.CreateVolunteer, userID int, eventID int) (*model.Volunteer, error) {
	VolunteerTableId := time.Now().Nanosecond()

	volunteer := model.Volunteer{
		ID:                VolunteerTableId,
		Role:              input.Role,
		Duration:          *input.Duration,
		VolunteerProposal: input.VolunteerProposal,
		DateApplied:       time.Now().Format("01-02-2006"),
		ApprovalStatus:    "Pending",
		EventID:           eventID,
		UserID:            userID,
	}

	// update user with VolunteerId field
	user, err := r.GetUserById(userID)

	if user != nil && err != nil {
		return nil, validators.NotFound
	}

	user.VolunteerID = VolunteerTableId
	user, err = r.UpdateCurrentUser(user)
	// =================================>

	if err := r.DB.Insert(&volunteer); err != nil {
		return nil, err
	}

	return &volunteer, nil
}

func (r *mutationResolver) UpdateVolunteer(ctx context.Context, id int, input model.UpdateVolunteer) (*model.Volunteer, error) {
	volunteer, err := r.GetVolunteerById(id)

	if volunteer != nil && err != nil {
		return nil, validators.NotFound
	}

	if volunteer, err = r.UpdateCurrentVolunteer(volunteer); err != nil {
		return nil, validators.ErrorUpdating
	}

	return volunteer, nil
}

func (r *mutationResolver) DeleteVolunteer(ctx context.Context, id int) (bool, error) {
	// volunteer, err := r.GetVolunteerById(id)
	// if volunteer != nil && err != nil {
	// 	return false, validators.NotFound
	// }
	//
	// volunteer = r.DeleteCurrentVolunteer(volunteer)
	// return true, nil

	panic("HAVING SLIGHT ERRS")
}

func (r *mutationResolver) PurchaseItem(ctx context.Context, input model.MakePurchases, itemID int, userID int, eventID int) (*model.Purchases, error) {
	purchase := model.Purchases{
		ID:            time.Now().Nanosecond(),
		ItemID:        itemID,
		UserID:        userID,
		EventID:       eventID,
		DatePurchased: time.Now().Format("01-02-2006 Mon"),
		Quantity:      0,
		Price:         input.Price,
	}

	// checks if user && events exists
	if _, err := r.GetUserById(userID); err != nil {
		return nil, validators.NotFound
	}

	if _, err := r.GetEventById(eventID); err != nil {
		return nil, validators.NotFound
	}

	if err := r.DB.Insert(&purchase); err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (r *mutationResolver) DeletePurchase(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CreateCategory, eventID int) (*model.Category, error) {
	category := model.Category{
		ID:      time.Now().Nanosecond(),
		Name:    input.Name,
		EventID: eventID,
		// ItemID: 0,
	}

	if _, err := r.GetEventById(eventID); err != nil {
		return nil, validators.NotFound
	}

	if err := r.DB.Insert(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCartItem(ctx context.Context, input model.CreateCartItem, categoryID int) (*model.CartItem, error) {
	event, err := r.GetEventById(categoryID)

	if event != nil && err != nil {
		fmt.Println("Event not found")
	}

	item := model.CartItem{
		ID:          time.Now().Nanosecond(),
		Name:        input.Name,
		Description: input.Description,
		CreatedAt:   time.Now().Format("01-02-2006"),
		CategoryID:  categoryID,
		Quantity:    0,
		Price:       input.Price,
		IsFree:      input.IsFree,
	}

	if err := r.DB.Insert(&item); err != nil {
		return nil, err
	}

	date := time.Now().Format("01-02-2006")
	action := fmt.Sprintf("%v was added to this event store %v", input.Name, date)
	event.Actions = append(event.Actions, action)
	if err := r.DB.Update(event); err != nil {
		return nil, validators.ErrorUpdating
	}

	return &item, nil
}

func (r *mutationResolver) UpdateCartItem(ctx context.Context, input model.UpdateCartItem) (*model.CartItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCartItem(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.CreateTaskComment, userID int, taskID int) (*model.TaskComments, error) {
	comment := model.TaskComments{
		ID:        time.Now().Nanosecond(),
		Content:   input.Content,
		UserID:    userID,
		TaskID:    taskID,
		WrittenAt: time.Now().Format("01-02-2006"),
	}

	if err := r.DB.Insert(&comment); err != nil {
		return nil, validators.ErrorInserting
	}

	return &comment, nil
}

func (r *mutationResolver) CreateBugReport(ctx context.Context, input *model.CreateBugReport, userID int, eventID int) (*model.BugReport, error) {
	if user, err := r.GetUserById(userID); user != nil && err != nil {
		return nil, validators.ValueNotFound("user")
	}

	if event, err := r.GetEventById(eventID); event != nil && err != nil {
		return nil, validators.ValueNotFound("user")
	}

	bugReport := &model.BugReport{
		ID:          time.Now().Nanosecond(),
		Title:       input.Title,
		Description: input.Description,
		UserID:      userID,
		EventID:     eventID,
		Status:      input.Status,
		CreatedAt:   time.Now().Format("01-02-2006"),
		UpdatedAt:   "",
	}

	if err := r.DB.Insert(bugReport); err != nil {
		return nil, validators.ErrorInserting
	}

	return bugReport, nil
}

func (r *mutationResolver) UpdateBugReport(ctx context.Context, input *model.UpdateBugReport, userID int, eventID int) (*model.BugReport, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBugReport(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateFeatureRequest(ctx context.Context, input *model.CreateFeatureRequest, userID int, eventID int) (*model.FeatureRequest, error) {
	if user, err := r.GetUserById(userID); user != nil && err != nil {
		return nil, validators.ValueNotFound("user")
	}

	if event, err := r.GetEventById(eventID); event != nil && err != nil {
		return nil, validators.ValueNotFound("user")
	}

	feature := &model.FeatureRequest{
		ID:          time.Now().Nanosecond(),
		Title:       input.Title,
		Description: input.Description,
		UserID:      userID,
		EventID:     eventID,
		Status:      input.Status,
		CreatedAt:   time.Now().Format("01-02-2006"),
		UpdatedAt:   "",
	}

	if err := r.DB.Insert(feature); err != nil {
		return nil, validators.ErrorInserting
	}

	return feature, nil
}

func (r *mutationResolver) UpdateFeatureRequest(ctx context.Context, input *model.UpdateFeatureRequest, userID int, eventID int) (*model.FeatureRequest, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteFeatureRequest(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateReminder(ctx context.Context, input *model.CreateReminder, userID int) (*model.Reminder, error) {
	if user, err := r.GetUserById(userID); user != nil && err != nil {
		return nil, validators.NotFound
	}

	reminder := &model.Reminder{
		ID:     time.Now().Nanosecond(),
		Name:   input.Name,
		UserID: userID,
		From:   input.From,
		Due:    input.Due,
	}

	if err := r.DB.Insert(reminder); err != nil {
		return nil, validators.ErrorInserting
	}

	return reminder, nil
}

func (r *mutationResolver) DeleteReminder(ctx context.Context, id *int) (bool, error) {
	reminder, err := r.GetReminderById(*id)

	if reminder != nil && err != nil {
		return false, validators.NotFound
	}

	if err := r.DeleteCurrentReminder(reminder); err != nil {
		return false, nil
	}

	return true, nil
}

func (r *mutationResolver) CreateNote(ctx context.Context, input *model.CreateNote, talkID int) (*model.Notes, error) {
	if talk, err := r.GetTalkById(talkID); talk != nil && err != nil {
		return nil, validators.NotFound
	}

	note := &model.Notes{
		ID:      time.Now().Nanosecond(),
		Title:   input.Title,
		Content: input.Content,
		TalkID:  talkID,
	}

	if err := r.DB.Insert(note); err != nil {
		return nil, validators.ErrorInserting
	}

	return note, nil
}

func (r *mutationResolver) UpdateNote(ctx context.Context, input *model.UpdateNote, talkID int) (*model.Notes, error) {
	note, err :=
		r.GetNoteById(talkID)

	if err != nil {
		return nil, validators.NotFound
	}

	note.Title = input.Title
	note.Content = input.Content

	if note, err = r.UpdateCurrentNote(note); err != nil {
		return nil, validators.ErrorUpdating
	}

	return note, nil
}

func (r *mutationResolver) DeleteNote(ctx context.Context, id int) (bool, error) {
	note, err := r.GetNoteById(id)

	if note != nil && err != nil {
		return false, validators.NotFound
	}

	err = r.DeleteCurrentNote(note)

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
