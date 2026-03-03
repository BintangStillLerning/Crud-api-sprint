package web

type ProductCreateRequest struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Image string `json:"image" validate:"required"`
}