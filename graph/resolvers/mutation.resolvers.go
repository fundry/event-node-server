package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
    "context"
    "fmt"
    "math/rand"
    "time"

    "github.com/vickywane/event-server/graph/generated"
    "github.com/vickywane/event-server/graph/middlewares"
    "github.com/vickywane/event-server/graph/model"
    "github.com/vickywane/event-server/graph/validators"
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
        ID:        rand.Int(),
        Token:     token,
        ExpiredAt: time.Now(),
        User:      user,
    }, nil
}

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.CreateEvent, userID int) (*model.Event, error) {
    currentUser , err := middlewares.ExtractCurrentUserFromContext(ctx)

    // returns an err output. Haven't finished work with AUTH middleware
    if err != nil {
        fmt.Println(err)
        return nil, validators.Unauthorized
    }

    fmt.Println(currentUser, "current user")
    mockbucketLink := string(rand.Int())
    Time := time.Now()

    if validators.CheckMail(input.Email) == false {
        return nil, validators.InvalidEmail
    }

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

func (r *mutationResolver) DeleteEvent(ctx context.Context, id int) (bool, error) {
    event, err := r.GetEventById(id)
    if event != nil && err != nil {
        return false, validators.NotFound
    }
    err = r.DeleteCurrentEvent(event)
    return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.AuthResponse, error) {
    hashedPassword := HashPassword(input.Password)
    mockbucketLink := string(rand.Int())
    EventID := rand.Int()
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
    fmt.Println(token, "User Token")

    UserID := rand.Int()
    user := model.User{
        ID:         UserID,
        Name:       input.Name,
        Email:      input.Email,
        CreatedAt:  time.Now(),
        Role:       input.Role,
        Password:   hashedPassword,
        BucketLink: mockbucketLink,
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
    track := model.Tracks{
        ID:          rand.Int(),
        Name:        input.Name,
        Duration:    input.Duration,
        TotalTalks:  input.TotalTalks,
        IsCompleted: false,
        Archived:    false,
        EventID:     eventID,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    // i  forgot to dereference this and it caused nightmares
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
