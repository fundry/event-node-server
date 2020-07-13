package resolvers

import (
    "fmt"
    "github.com/go-pg/pg/v9"
    "sync"

    "github.com/vickywane/event-server/graph/model"
)

type Resolver struct {
    DB *pg.DB

    // teamChan map[interface{}]chan *model.Team // remove * if it doesn work

    mutex sync.Mutex
}

// var TeamChan = make(chan *model.Team, 1)


// my custom func

func (r *mutationResolver) CheckTaskFieldExists(field string , fieldValue string) bool {
    taskField := model.Tasks{}
    if err := r.DB.Model(&taskField).Where(fmt.Sprintf("%v = ?", field), fieldValue).First(); err != nil {
        return true
    }
    return false
}

func (r *mutationResolver) CheckEventTeamFieldExists(field string , fieldValue string) bool {
    teamField := model.Team{}
    if err := r.DB.Model(&teamField).Where(fmt.Sprintf("%v = ?", field), fieldValue).First(); err != nil {
        return true
    }
    return false
}

func (r *mutationResolver) CheckEventTalkFieldExists(field string , fieldValue int) bool {
    talkField := model.EventTalk{}
    if err := r.DB.Model(&talkField).Where(fmt.Sprintf("%v = ?", field), fieldValue).First(); err != nil {
        return true
    }
    return false
}

func (r *mutationResolver) CheckEventFieldExists(field, fieldValue string) bool {
    eventField := model.Event{}
    if err := r.DB.Model(&eventField).Where(fmt.Sprintf("%v = ?", field), fieldValue).First(); err != nil {
        return true
    }
    return false
}

func (r *mutationResolver) CheckMeetupFieldExists(field, fieldValue string) bool {
    meetupField := model.MeetupGroups{}
    if err := r.DB.Model(&meetupField).Where(fmt.Sprintf("%v = ?", field), fieldValue).First(); err != nil {
        return true
    }
    return false
}


func (r *mutationResolver) GetSubmittedTalkById(id int) (*model.EventTalk, error) {
    event := model.EventTalk{}
    err := r.DB.Model(&event).Where("id = ?", id).First()
    return &event, err
}


func (r *mutationResolver) CheckAttendeeFieldExists(field string, fieldValue int) bool {
    attendeeField := model.Attendee{}

    err := r.DB.Model(&attendeeField).Where(fmt.Sprintf("%v = ?", field), fieldValue).First()

    fmt.Println(err, "check exists")
    if err != nil {
        return true
    }

    return false
}

// Todo compress funcs here
func (r *mutationResolver) GetUserField(field, value string) (*model.User, error) {
    user := model.User{}

    err := r.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()

    return &user, err
}

func (r *mutationResolver) GetUserByEmail(email string) (*model.User, error) {
    user := model.User{}
    err := r.DB.Model(&user).Where("email = ?", email).First()
    return &user, err
}

func (r *mutationResolver) GetUserById(id int) (*model.User, error) {
    user := model.User{}
    err := r.DB.Model(&user).Where("id = ?", id).First()
    return &user, err
}

func (r *mutationResolver) GetNoteById(id int) (*model.Notes, error) {
    Notes := model.Notes{}
    err := r.DB.Model(&Notes).Where("id = ?", id).First()
    return &Notes, err
}

func (r *mutationResolver) UpdateCurrentUser(user *model.User) (*model.User, error) {
    _, err := r.DB.Model(user).Where("id = ?", user.ID).Update()
    return user, err
}

func (r *mutationResolver) UpdateCurrentEventSetting(setting *model.EventSettings) (*model.EventSettings, error) {
    _, err := r.DB.Model(setting).Where("id = ?", setting.ID).Update()
    return setting, err
}

func (r *mutationResolver) UpdateCurrentNote(note *model.Notes) (*model.Notes, error) {
    _, err := r.DB.Model(note).Where("id = ?", note.ID).Update()
    return note, err
}

func (r *mutationResolver) DeleteCurrentUser(user *model.User) error {
    _, err := r.DB.Model(user).Where("id =?", user.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) DeleteCurrentNote(note *model.Notes) error {
    _, err := r.DB.Model(note).Where("id =?", note.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) DeleteCurrentTeam(team *model.Team) error {
    _, err := r.DB.Model(team).Where("id =?", team.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}


func (r *mutationResolver) GetEventById(id int) (*model.Event, error) {
    event := model.Event{}
    err := r.DB.Model(&event).Where("id = ?", id).First()
    return &event, err
}

func (r *mutationResolver) GetSettingById(id int) (*model.EventSettings, error) {
    setting :=  model.EventSettings{}
    err := r.DB.Model(&setting).Where("id = ?", id).First()
    return &setting, err
}


func (r *mutationResolver) GetSponsorById(id int) (*model.Sponsor, error) {
    sponsor := model.Sponsor{}
    err := r.DB.Model(&sponsor).Where("id = ?", id).First()
    return &sponsor, err
}

func (r *mutationResolver) DeleteCurrentSponsor(sponsor *model.Sponsor) error {
    _, err := r.DB.Model(sponsor).Where("id =?", sponsor.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) UpdateCurrentEvent(event *model.Event) (*model.Event, error) {
    _, err := r.DB.Model(event).Where("id = ?", event.ID).Update()
    return event, err
}

func (r *mutationResolver) DeleteCurrentEvent(Event *model.Event) error {
    _, err := r.DB.Model(Event).Where("id =?", Event.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) GetTeamById(id int) (*model.Team, error) {
    team := model.Team{}
    err := r.DB.Model(&team).Where("id = ?", id).First()
    return &team, err
}

func (r *mutationResolver) GetTaskById(id int) (*model.Tasks, error) {
    task := model.Tasks{}
    err := r.DB.Model(&task).Where("id = ?", id).First()
    return &task, err
}

func (r *mutationResolver) UpdateCurrentTask(task *model.Tasks) (*model.Tasks, error) {
    _, err := r.DB.Model(task).Where("id = ?", task.ID).Update()
    return task, err
}

func (r *mutationResolver) DeleteCurrentTask(task *model.Tasks) error {
    _, err := r.DB.Model(task).Where("id =?", task.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) GetTalkById(id int) (*model.Talk, error) {
    Talk := model.Talk{}
    err := r.DB.Model(&Talk).Where("id = ?", id).First()
    return &Talk, err
}

func (r *mutationResolver) UpdateCurrentTalk(Talk *model.Talk) (*model.Talk, error) {
    _, err := r.DB.Model(Talk).Where("id = ?", Talk.ID).Update()
    return Talk, err
}

func (r *mutationResolver) DeleteCurrentTalk(Talk *model.Talk) error {
    _, err := r.DB.Model(Talk).Where("id =?", Talk.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) GetTrackById(id int) (*model.Tracks, error) {
    Track := model.Tracks{}
    err := r.DB.Model(&Track).Where("id = ?", id).First()
    return &Track, err
}

func (r *mutationResolver) UpdateCurrentTrack(Track *model.Tracks) (*model.Tracks, error) {
    _, err := r.DB.Model(Track).Where("id = ?", Track.ID).Update()
    return Track, err
}

func (r *mutationResolver) DeleteCurrentTrack(Track *model.Tracks) error {
    _, err := r.DB.Model(Track).Where("id =?", Track.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) DeleteCurrentReminder(Reminder *model.Reminder) error {
    _, err := r.DB.Model(Reminder).Where("id =?", Reminder.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

func (r *mutationResolver) GetReminderById(id int) (*model.Reminder, error) {
    reminder := model.Reminder{}
    err := r.DB.Model(&reminder).Where("id = ?", id).First()
    return &reminder, err
}

func (r *mutationResolver) GetVolunteerById(id int) (*model.Volunteer, error) {
    Volunteer := model.Volunteer{}
    err := r.DB.Model(&Volunteer).Where("id = ?", id).First()
    return &Volunteer, err
}

func (r *mutationResolver) UpdateCurrentVolunteer(Volunteer *model.Volunteer) (*model.Volunteer, error) {
    _, err := r.DB.Model(Volunteer).Where("id = ?", Volunteer.ID).Update()
    return Volunteer, err
}

func (r *mutationResolver) DeleteCurrentVolunteer(Volunteer *model.Volunteer) error {
    _, err := r.DB.Model(Volunteer).Where("id =?", Volunteer.ID).Delete()
    if err != nil {
        return nil
    }
    return nil
}

// Checks if file exists
func (r *mutationResolver) CheckEventFile(field, fieldValue string) bool {
    err := r.DB.Model(&model.EventFile{}).Where(fmt.Sprintf("%v = ?", field), fieldValue).First()

    if err != nil {
        fmt.Printf("Error %v", err)
        return true
    }

    println(",skgfsfbkefblu")

    return false
}
