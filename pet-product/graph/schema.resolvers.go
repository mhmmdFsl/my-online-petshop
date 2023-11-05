package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"github.com/mhmmdFsl/my-online-petshop/pet-product/graph/model"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	p, err := r.ProductService.Create(&input)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (string, error) {
	s, err := r.ProductService.Delete(id)
	if err != nil {
		return "", err
	}
	return s, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, input *model.UpdateProduct) (*model.Product, error) {
	p, err := r.ProductService.UpdateProduct(input)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	p, err := r.ProductService.GetAll()
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
