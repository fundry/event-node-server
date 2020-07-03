package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *subscriptionResolver) VolunteerCreated(ctx context.Context, role *string) (<-chan *model.Volunteer, error) {
	volunteer := make(chan *model.Volunteer, 1)

	fmt.Println("\n \n mutation ran  ")
	fmt.Println("\n \n mutation here  ")
	go func() {
		<-ctx.Done()
	}()

	return volunteer, nil
}

func (r *subscriptionResolver) NewTeam(ctx context.Context) (<-chan *model.Team, error) {
	// r.mutex.Lock()
	//
	// team := make(chan *model.Team )
	//
	// go func() {
	//     <- ctx.Done()
	// }()
	//
	// team = TeamChan
	//
	// return team, nil

	panic("not implemented")
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
