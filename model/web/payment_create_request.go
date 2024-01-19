package web

type PaymentCreateRequest struct {
	Name string `validate:"required,min=1" json:"name"`
	Type string `validate:"required,min=1" json:"type"`
}