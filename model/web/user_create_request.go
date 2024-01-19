package web

type UserCreateRequest struct {
	Name string `validate:"required,min=1" json:"name"`
	Email string `validate:"required,min=1" json:"email"`
	Password string `validate:"required,min=1" json:"password"`
	Role string `validate:"required,min=1" json:"role"`
	// Token string `json:"token"`
}

type LoginUser struct {
	Email string `validate:"required,min=1" json:"email"`
	Password string `validate:"required,min=1" json:"password"`
}

type LogoutUser struct {
	Email string `validate:"required,min=1" json:"email"`
	// Password string `validate:"required,min=1" json:"password"`
}