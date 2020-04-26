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
func (r *mutationResolver) GetUser(field, value string) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()

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

func (r *mutationResolver) GetUserByEmail(email string) (*model.User , error) {
 	return r.GetUser("email", email)
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

func (r *mutationResolver) GetTeamById(id int) (*model.Team, error) {
	team := model.Team{}

	err := r.DB.Model(&team).Where("id = ?", id).First()

	return &team, err
}


