package product

type CreateProductRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Image       []string `json:"image"`
}

type UpdateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       []string `json:"image"`
}
