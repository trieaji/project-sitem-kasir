package service

import (
	"context"
	"projectsistemkasir/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.LoginUser) web.UserLoginResponse //kayaknya nanti di "web.UserCreateRequest" diganti dengan file yg baru misal "web.LoginUserCreate" dan berisi data baru yg dibutuhkan
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
	Logut(ctx context.Context, request web.LogoutUser) //kayaknya nanti di "web.UserUpdateRequest" diganti dengan file yg baru misal "web.LooutUserCreate" dan berisi data baru yg dibutuhkan
}