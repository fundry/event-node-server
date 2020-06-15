package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *taskCommentsResolver) WrittenBy(ctx context.Context, obj *model.TaskComments) ([]*model.User, error) {
	var createdBy []*model.User
	err := r.DB.Model(&createdBy).Where("id = ?", obj.WrittenBy).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return createdBy, nil
}

func (r *taskCommentsResolver) Task(ctx context.Context, obj *model.TaskComments) ([]*model.Tasks, error) {
	var tasks []*model.Tasks
	err := r.DB.Model(&tasks).Where("user_id = ?", obj.Task).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *tasksResolver) Assignees(ctx context.Context, obj *model.Tasks) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *tasksResolver) CreatedBy(ctx context.Context, obj *model.Tasks) ([]*model.User, error) {
	var createdBy []*model.User
	err := r.DB.Model(&createdBy).Where("id = ?", obj.AuthorID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return createdBy, nil
}

func (r *tasksResolver) Event(ctx context.Context, obj *model.Tasks) ([]*model.Event, error) {
	var createdBy []*model.Event
	err := r.DB.Model(&createdBy).Where("event_id = ?", obj.Event).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return createdBy, nil
}

func (r *tasksResolver) Comments(ctx context.Context, obj *model.Tasks) ([]*model.TaskComments, error) {
	var comments []*model.TaskComments
	err := r.DB.Model(&comments).Where("id = ?", obj.CommentID).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return comments, nil
}

// TaskComments returns generated.TaskCommentsResolver implementation.
func (r *Resolver) TaskComments() generated.TaskCommentsResolver { return &taskCommentsResolver{r} }

// Tasks returns generated.TasksResolver implementation.
func (r *Resolver) Tasks() generated.TasksResolver { return &tasksResolver{r} }

type taskCommentsResolver struct{ *Resolver }
type tasksResolver struct{ *Resolver }
