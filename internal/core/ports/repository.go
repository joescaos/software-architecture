package ports

import (
	"software-architecture/internal/core/domain"
)

type PageResult struct {
	Items      []domain.Product
	Total      int64
	Page       int
	PageSize   int
	TotalPages int
}

type ProductRepository interface {
	Create(product *domain.Product) (*domain.Product, error)
	GetByID(id uint) (*domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
	Delete(id uint) error
	List(page, pageSize int) (*PageResult, error)
}
