package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
    "context"
    "fmt"
    "sync"

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
    m := sync.Mutex{}
    m.Lock()
    m.Unlock()

    team := make(chan *model.Team)

    fmt.Println(team)

    go func() {
        ctx.Done()
        // team <- interface {}
    }()
    m.Lock()

    m.Unlock()
    return team, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct {
    *Resolver
}
