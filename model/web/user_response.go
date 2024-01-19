package web

type UserResponse struct {
	Id   int    `json:"id"`
	Name string	`json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
	// Role string `json:"role"`
}