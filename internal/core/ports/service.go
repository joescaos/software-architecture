package ports

import (
	"software-architecture/internal/dto"
)

type ProductService interface {
	CreateProduct(request *dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetProductByID(id uint) (*dto.ProductResponse, error)
	UpdateProduct(id uint, request *dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(id uint) error
	ListProducts(page, pageSize int) (dto.PaginatedResponse, error)
}
