package web

type UserUpdateRequest struct {
	Id int `validate:"required"`
	Name string `validate:"required,min=1" json:"name"`
	Email string `validate:"required,min=1" json:"email"`
	Password string `validate:"required,min=1" json:"password"`
	Role string `validate:"required,min=1" json:"role"`
	// Token string `json:"token"`
}