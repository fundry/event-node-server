package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *talkResolver) Speaker(ctx context.Context, obj *model.Talk) ([]*model.User, error) {
	var Speaker []*model.User

	if err := r.DB.Model(&Speaker).Order("id").Where("id = ?", obj.SpeakerID).Select(); err != nil {
		return nil, err
	}

	return Speaker, nil
}

func (r *talkResolver) Reviewers(ctx context.Context, obj *model.Talk) ([]*model.User, error) {
	var User []*model.User

	err := r.DB.Model(&User).Order("id").Select()

	if err != nil {
		return nil, err
	}
	return User, nil
}

// Talk returns generated.TalkResolver implementation.
func (r *Resolver) Talk() generated.TalkResolver { return &talkResolver{r} }

type talkResolver struct{ *Resolver }
