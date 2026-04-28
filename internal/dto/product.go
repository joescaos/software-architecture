package dto

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type PaginatedResponse struct {
	Data       []ProductResponse `json:"data"`
	Page       int               `json:"page"`
	PageSize   int               `json:"pageSize"`
	Total      int64             `json:"total"`
	TotalPages int               `json:"totalPages"`
}
