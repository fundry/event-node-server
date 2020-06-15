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

	eventId := time.Now().Nanosecond()

	BucketName, err := CreateBucket(eventId)

	fmt.Printf("Bucket %v \n", BucketName)
	fmt.Printf("Error from Bucket %v \n", err)

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
		AuthorID:       userID,
		DateCreated:    time.Now().Format("01-02-2006"),
		TotalAttendees: 0,
		BucketLink:     "",
		MeetupGroups: append([]*model.MeetupGroups{}, &model.MeetupGroups{
			ID: time.Now().Nanosecond(),
		}),
		BucketName:            BucketName,
		IsAcceptingTalks:      false,
		IsAcceptingVolunteers: false,
		ConfirmedEmail:        false,
		IsArchived:            false,
		Actions:               []string{fmt.Sprintf("Event was created by USER on %v", time.Now().Format("01-02-2006"))},
		IsLocked:              false,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}

	if !r.CheckEventFieldExists("name", input.Name) {
		return nil, validators.FieldTaken("name")
	}

	if !r.CheckEventFieldExists("email", input.Email) {
		return nil, validators.FieldTaken("email")
	}

	if !r.CheckEventFieldExists("description", input.Description) {
		return nil, validators.FieldTaken("description")
	}

	if len(input.Summary) > 150 {
		return nil, validators.ValueExceeded(150)
	}

	if len(input.Description) > 1500 {
		return nil, validators.ValueExceeded(1500)
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
	if validators.CheckMail(*input.Email) == false {
		return nil, validators.InvalidEmail
	}

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
		IsAccepted:    false,
		DateSubmitted: time.Now().Format("01-02-2006"),
		DraftID:       talkID,
		EventID:       eventID,
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

	if err := r.DB.Update(talk); err != nil {
		return nil, validators.ErrorUpdating
	}

	return talk, nil
}

func (r *mutationResolver) CreateMeetupGroup(ctx context.Context, eventID int, leadID int, input *model.CreateMeetupGroup) (*model.MeetupGroups, error) {
	meetupGroup := &model.MeetupGroups{
		ID:        time.Now().Nanosecond(),
		Name:      input.Name,
		EventID:   eventID,
		LeadID:    leadID,
		Location:  input.Location,
		Alias:     input.Alias,
		CreatedAt: time.Now().Format("01-02-2006"),
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

	return meetupGroup, nil
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
		ID:        time.Now().Nanosecond(),
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

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.CreateTeam, eventID int) (*model.Team, error) {
	event, err := r.GetEventById(eventID)

	if err != nil {
		return nil, validators.NotFound
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

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTasks, teamID int, userID int) (*model.Tasks, error) {
	task := model.Tasks{
		ID:        time.Now().Nanosecond(),
		Name:      input.Name,
		Category:  input.Category,
		CreatedAt: time.Now().Format("01-02-2006"),
		TeamID:    teamID,
		// id: userID,
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

	panic("likely err here")
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
