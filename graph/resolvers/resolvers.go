package resolvers

import (
	"fmt"
	"github.com/go-pg/pg/v9"

	"github.com/vickywane/event-server/graph/model"
)

type Resolver struct {
	DB *pg.DB
}

// my custom func
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

func (r *mutationResolver) UpdateCurrentUser(user *model.User) (*model.User, error) {
	_, err := r.DB.Model(user).Where("id = ?", user.ID).Update()
	return user, err
}

func (r *mutationResolver) DeleteCurrentUser(user *model.User) error {
	_, err := r.DB.Model(user).Where("id =?", user.ID).Delete()
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

