package repository

import (
	"software-architecture/internal/core/domain"
	"software-architecture/internal/core/ports"

	"gorm.io/gorm"
)

type gormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ports.ProductRepository {
	return &gormProductRepository{db: db}
}

type ProductModel struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:255;not null"`
	Description string  `gorm:"size:500"`
	Price       float64 `gorm:"not null"`
}

func (r *gormProductRepository) Create(product *domain.Product) (*domain.Product, error) {
	model := ProductModel{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return nil, err
	}
	product.ID = model.ID
	return product, nil
}

func (r *gormProductRepository) GetByID(id uint) (*domain.Product, error) {
	var model ProductModel
	if err := r.db.First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &domain.Product{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
	}, nil
}

func (r *gormProductRepository) Update(product *domain.Product) (*domain.Product, error) {
	model := ProductModel{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
	if err := r.db.Save(&model).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *gormProductRepository) Delete(id uint) error {
	return r.db.Delete(&ProductModel{}, id).Error
}

func (r *gormProductRepository) List(page, pageSize int) (*ports.PageResult, error) {
	var models []ProductModel
	var total int64

	if err := r.db.Model(&ProductModel{}).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&models).Error; err != nil {
		return nil, err
	}

	products := make([]domain.Product, len(models))
	for i, model := range models {
		products[i] = domain.Product{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			Price:       model.Price,
		}
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	return &ports.PageResult{
		Items:      products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}
