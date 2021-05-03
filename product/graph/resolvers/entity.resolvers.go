package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"product/graph/generated"
	"product/graph/model"
)

func (r *entityResolver) FindProductByUpc(ctx context.Context, upc string) (*model.Product, error) {

	hats := []*model.Product{
		{
			Upc:   "top-1",
			Name:  "Trilby",
			Price: 11,
		},
		{
			Upc:   "top-2",
			Name:  "Fedora",
			Price: 22,
		},
		{
			Upc:   "top-3",
			Name:  "Boater",
			Price: 33,
		},
	}

	for _, h := range hats {
		if h.Upc == upc {
			return h, nil
		}
	}
	return nil, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
