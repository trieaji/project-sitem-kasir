package web

type ProductResponse struct {
	Id   int    `json:"id"`
	Sku string `json:"sku"`
	Name string `json:"name"`
	Stock int `json:"stock"`
	Price int `json:"price"`
	CategoryId int `json:"category_id"`
}