package web

type PaymentUpdateRequest struct {
	Id int `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
	Type string `validate:"required,max=200,min=1" json:"type"`
}