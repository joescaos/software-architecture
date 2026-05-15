package graph

import (
	"context"
	"strconv"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"software-architecture/graph/model"
	"software-architecture/internal/dto"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	desc := ""
	if input.Description != nil {
		desc = *input.Description
	}
	res, err := r.ProductService.CreateProduct(dto.CreateProductRequest{
		Name:        input.Name,
		Description: desc,
		Price:       input.Price,
	})
	if err != nil {
		return nil, gqlerror.Wrap(err)
	}
	return productToGraph(res), nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProductInput) (*model.Product, error) {
	pid, err := parseProductID(id)
	if err != nil {
		return nil, gqlerror.Errorf("invalid product id")
	}
	req := dto.UpdateProductRequest{}
	if input.Name != nil {
		req.Name = *input.Name
	}
	if input.Description != nil {
		req.Description = *input.Description
	}
	if input.Price != nil {
		req.Price = *input.Price
	}
	res, err := r.ProductService.UpdateProduct(pid, req)
	if err != nil {
		return nil, gqlerror.Wrap(err)
	}
	return productToGraph(res), nil
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	pid, err := parseProductID(id)
	if err != nil {
		return false, gqlerror.Errorf("invalid product id")
	}
	if err := r.ProductService.DeleteProduct(pid); err != nil {
		return false, gqlerror.Wrap(err)
	}
	return true, nil
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	pid, err := parseProductID(id)
	if err != nil {
		return nil, gqlerror.Errorf("invalid product id")
	}
	res, err := r.ProductService.GetProductByID(pid)
	if err != nil {
		return nil, gqlerror.Wrap(err)
	}
	return productToGraph(res), nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, page *int32, pageSize *int32) (*model.ProductConnection, error) {
	p := 1
	ps := 10
	if page != nil && *page > 0 {
		p = int(*page)
	}
	if pageSize != nil && *pageSize > 0 {
		ps = int(*pageSize)
	}
	out, err := r.ProductService.ListProducts(p, ps)
	if err != nil {
		return nil, gqlerror.Wrap(err)
	}
	data := make([]*model.Product, len(out.Data))
	for i := range out.Data {
		data[i] = productToGraph(&out.Data[i])
	}
	return &model.ProductConnection{
		Data:       data,
		Page:       int32(out.Page),
		PageSize:   int32(out.PageSize),
		Total:      int32(out.Total),
		TotalPages: int32(out.TotalPages),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func parseProductID(id string) (uint, error) {
	u, err := strconv.ParseUint(id, 10, 32)
	if err != nil || u == 0 {
		return 0, err
	}
	return uint(u), nil
}

func productToGraph(p *dto.ProductResponse) *model.Product {
	if p == nil {
		return nil
	}
	return &model.Product{
		ID:          strconv.FormatUint(uint64(p.ID), 10),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}
