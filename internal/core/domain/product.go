package domain

import "errors"

type Product struct {
	ID          uint
	Name        string
	Description string
	Price       float64
}

func NewProduct(name, description string, price float64) (*Product, error) {
	if name == "" {
		return nil, errors.New("product name cannot be empty")
	}
	if price < 0 {
		return nil, errors.New("product price cannot be negative or zero")
	}
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
	}, nil
}
