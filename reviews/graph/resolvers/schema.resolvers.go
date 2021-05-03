package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"reviews/graph/generated"
	"reviews/graph/model"
)

var reviews = []*model.Review{
	{
		Body:    "A highly effective form of birth control.",
		Product: &model.Product{Upc: "top-1"},
		Author:  &model.User{ID: "1234"},
	},
	{
		Body:    "Fedoras are one of the most fashionable hats around and can look great with a variety of outfits.",
		Product: &model.Product{Upc: "top-1"},
		Author:  &model.User{ID: "1234"},
	},
	{
		Body:    "This is the last straw. Hat you will wear. 11/10",
		Product: &model.Product{Upc: "top-1"},
		Author:  &model.User{ID: "7777"},
	},
}

func (r *productResolver) Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error) {
	var res []*model.Review

	for _, review := range reviews {
		if review.Product.Upc == obj.Upc {
			res = append(res, review)
		}
	}

	return res, nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	var res []*model.Review

	for _, review := range reviews {
		if review.Author.ID == obj.ID {
			res = append(res, review)
		}
	}

	return res, nil
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type productResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
