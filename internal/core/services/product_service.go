package services

import (
	"errors"
	"software-architecture/internal/core/domain"
	"software-architecture/internal/core/ports"
	"software-architecture/internal/dto"
)

type ProductService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(request dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product, err := domain.NewProduct(request.Name, request.Description, request.Price)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *ProductService) GetProductByID(id uint) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *ProductService) UpdateProduct(id uint, request dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Price > 0 {
		product.Price = request.Price
	}

	_, err = s.repo.Update(product)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *ProductService) ListProducts(page, pageSize int) (dto.PaginatedResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	results, err := s.repo.List(page, pageSize)
	if err != nil {
		return dto.PaginatedResponse{}, err
	}
	if results == nil {
		return dto.PaginatedResponse{}, errors.New("no results returned")
	}

	responses := make([]dto.ProductResponse, len(results.Items))
	for i, product := range results.Items {
		responses[i] = dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}
	}

	return dto.PaginatedResponse{
		Data:       responses,
		Total:      results.Total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((results.Total + int64(pageSize) - 1) / int64(pageSize)),
	}, nil
}
