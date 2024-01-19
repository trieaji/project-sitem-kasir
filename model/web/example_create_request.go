package web

type ExampleCreateRequest struct {
	Name string `validate:"required,min=1" json:"name"`
	Jurusan string `validate:"required,min=1" json:"jurusan"`
}