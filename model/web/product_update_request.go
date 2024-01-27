package web

type ProductUpdateRequest struct {
	Id int `validate:"required"`
	Sku string `validate:"required,min=1" json:"sku"`
	Name string `validate:"required,min=1" json:"name"`
	Stock int `validate:"required,min=1" json:"stock"`
	Price int `validate:"required,min=1" json:"price"`
	CategoryId int `validate:"required,min=1" json:"category_id"`
}