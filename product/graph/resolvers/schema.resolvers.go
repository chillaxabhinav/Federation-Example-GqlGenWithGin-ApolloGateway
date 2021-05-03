package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"product/graph/generated"
	"product/graph/model"
)

var hats = []*model.Product{
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

func (r *queryResolver) TopProducts(ctx context.Context) ([]*model.Product, error) {
	return hats, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
