package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vickywane/event-server/graph/generated"
	"github.com/vickywane/event-server/graph/model"
)

func (r *categoryResolver) Event(ctx context.Context, obj *model.Category) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *categoryResolver) Items(ctx context.Context, obj *model.Category) ([]*model.CartItem, error) {
	var cartItem []*model.CartItem

	if err := r.DB.Model(&cartItem).Where("category_id = ?", obj.ID).Order("id").Select(); err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (r *purchasesResolver) Item(ctx context.Context, obj *model.Purchases) ([]*model.CartItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *purchasesResolver) User(ctx context.Context, obj *model.Purchases) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *purchasesResolver) Event(ctx context.Context, obj *model.Purchases) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Purchases returns generated.PurchasesResolver implementation.
func (r *Resolver) Purchases() generated.PurchasesResolver { return &purchasesResolver{r} }

type categoryResolver struct{ *Resolver }
type purchasesResolver struct{ *Resolver }
